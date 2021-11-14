package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"time"

	"github.com/alexballas/go2tv/internal/devices"
	"github.com/alexballas/go2tv/internal/gui"
	"github.com/alexballas/go2tv/internal/httphandlers"
	"github.com/alexballas/go2tv/internal/interactive"
	"github.com/alexballas/go2tv/internal/soapcalls"
	"github.com/alexballas/go2tv/internal/urlstreamer"
	"github.com/alexballas/go2tv/internal/utils"
	"github.com/pkg/errors"
)

var (
	version    string
	build      string
	dmrURL     string
	mediaArg   = flag.String("v", "", "Path to the video/audio file. (Triggers the CLI mode)")
	urlArg     = flag.String("u", "", "Path to the URL media file. URL streaming does not support seek operations. (Triggers the CLI mode)")
	subsArg    = flag.String("s", "", "Path to the subtitles file.")
	listPtr    = flag.Bool("l", false, "List all available UPnP/DLNA Media Renderer models and URLs.")
	targetPtr  = flag.String("t", "", "Cast to a specific UPnP/DLNA Media Renderer URL.")
	versionPtr = flag.Bool("version", false, "Print version.")
)

func main() {
	guiEnabled := true
	var mediaFile interface{}
	flag.Parse()

	exit, err := checkflags()
	check(err)
	if exit {
		os.Exit(0)
	}
	if *mediaArg != "" || *urlArg != "" {
		guiEnabled = false
	}

	if *mediaArg != "" {
		mediaFile = *mediaArg
	}

	if *mediaArg == "" && *urlArg != "" {
		mediaFile, err = urlstreamer.StreamURL(context.Background(), *urlArg)
		check(err)
	}

	if guiEnabled {
		scr := gui.InitFyneNewScreen(version + " / " + build)
		gui.Start(scr)
	}

	var absMediaFile string
	var mediaType string

	switch t := mediaFile.(type) {
	case string:
		absMediaFile, err = filepath.Abs(t)
		check(err)
		mediaFile = absMediaFile

		mediaType, err = utils.GetMimeDetailsFromFile(absMediaFile)
		check(err)
	case *io.PipeReader:
		absMediaFile = *urlArg
	}

	absSubtitlesFile, err := filepath.Abs(*subsArg)
	check(err)

	upnpServicesURLs, err := soapcalls.DMRextractor(dmrURL)
	check(err)

	whereToListen, err := utils.URLtoListenIPandPort(dmrURL)
	check(err)

	scr, err := interactive.InitTcellNewScreen()
	check(err)

	callbackPath, err := utils.RandomString()
	check(err)

	tvdata := &soapcalls.TVPayload{
		ControlURL:          upnpServicesURLs.AvtransportControlURL,
		EventURL:            upnpServicesURLs.AvtransportEventSubURL,
		RenderingControlURL: upnpServicesURLs.RenderingControlURL,
		CallbackURL:         "http://" + whereToListen + "/" + callbackPath,
		MediaURL:            "http://" + whereToListen + "/" + utils.ConvertFilename(absMediaFile),
		SubtitlesURL:        "http://" + whereToListen + "/" + utils.ConvertFilename(absSubtitlesFile),
		MediaType:           mediaType,
		CurrentTimers:       make(map[string]*time.Timer),
	}

	s := httphandlers.NewServer(whereToListen)
	serverStarted := make(chan struct{})

	// We pass the tvdata here as we need the callback handlers to be able to react
	// to the different media renderer states.
	go func() {
		err := s.ServeFiles(serverStarted, mediaFile, absSubtitlesFile, tvdata, scr)
		check(err)
	}()
	// Wait for HTTP server to properly initialize
	<-serverStarted

	err = tvdata.SendtoTV("Play1")
	check(err)

	scr.InterInit(tvdata)
}

func check(err error) {
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Encountered error(s): %s\n", err)
		os.Exit(1)
	}
}

func listFlagFunction() error {
	if len(devices.Devices) == 0 {
		return errors.New("-l and -t can't be used together")
	}
	fmt.Println()

	// We loop through this map twice as we need to maintain
	// the correct order.
	keys := make([]string, 0)
	for k := range devices.Devices {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for q, k := range keys {
		boldStart := ""
		boldEnd := ""

		if runtime.GOOS == "linux" {
			boldStart = "\033[1m"
			boldEnd = "\033[0m"
		}
		fmt.Printf("%sDevice %v%s\n", boldStart, q+1, boldEnd)
		fmt.Printf("%s--------%s\n", boldStart, boldEnd)
		fmt.Printf("%sModel:%s %s\n", boldStart, boldEnd, k)
		fmt.Printf("%sURL:%s   %s\n", boldStart, boldEnd, devices.Devices[k])
		fmt.Println()
	}

	return nil
}

func checkflags() (exit bool, err error) {
	checkVerflag()

	if checkGUI() {
		return false, nil
	}

	if err := checkTflag(); err != nil {
		return false, fmt.Errorf("checkflags error: %w", err)
	}

	list, err := checkLflag()
	if err != nil {
		return false, fmt.Errorf("checkflags error: %w", err)
	}

	if err := checkVflag(); err != nil {
		return false, fmt.Errorf("checkflags error: %w", err)
	}

	if list {
		return true, nil
	}

	if err := checkSflag(); err != nil {
		return false, fmt.Errorf("checkflags error: %w", err)
	}

	return false, nil
}

func checkVflag() error {
	if !*listPtr && *urlArg == "" {
		if _, err := os.Stat(*mediaArg); os.IsNotExist(err) {
			return fmt.Errorf("checkVflags error: %w", err)
		}
	}

	return nil
}

func checkSflag() error {
	if *subsArg != "" {
		if _, err := os.Stat(*subsArg); os.IsNotExist(err) {
			return fmt.Errorf("checkSflags error: %w", err)
		}
	} else {
		// The checkVflag should happen before
		// checkSflag so we're safe to call *mediaArg
		// here. If *subsArg is empty, try to
		// automatically find the srt from the
		// media file filename.
		*subsArg = (*mediaArg)[0:len(*mediaArg)-
			len(filepath.Ext(*mediaArg))] + ".srt"
	}

	return nil
}

func checkTflag() error {
	if *targetPtr == "" {
		err := devices.LoadSSDPservices(1)
		if err != nil {
			return fmt.Errorf("checkTflag service loading error: %w", err)
		}

		dmrURL, err = devices.DevicePicker(1)
		if err != nil {
			return fmt.Errorf("checkTflag device picker error: %w", err)
		}
	} else {
		// Validate URL before proceeding.
		_, err := url.ParseRequestURI(*targetPtr)
		if err != nil {
			return fmt.Errorf("checkTflag parse error: %w", err)
		}
		dmrURL = *targetPtr
	}

	return nil
}

func checkLflag() (bool, error) {
	if *listPtr {
		if err := listFlagFunction(); err != nil {
			return false, fmt.Errorf("checkLflag error: %w", err)
		}
		return true, nil
	}

	return false, nil
}

func checkVerflag() {
	if *versionPtr {
		fmt.Printf("Go2TV Version: %s, ", version)
		fmt.Printf("Build: %s\n", build)
		os.Exit(0)
	}
}

func checkGUI() bool {
	return *mediaArg == "" && !*listPtr && *urlArg == ""
}
