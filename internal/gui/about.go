package gui

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func aboutWindow(s *NewScreen) fyne.CanvasObject {
	richhead := widget.NewRichTextFromMarkdown(`
# Go2TV

Cast your media files to UPnP/DLNA Media Renderers and Smart TVs

---

## Author
Alex Ballas - alex@ballas.org

## License
MIT

## Version

` + s.version)

	for i := range richhead.Segments {
		if seg, ok := richhead.Segments[i].(*widget.TextSegment); ok {
			seg.Style.Alignment = fyne.TextAlignCenter
		}
		if seg, ok := richhead.Segments[i].(*widget.HyperlinkSegment); ok {
			seg.Alignment = fyne.TextAlignCenter
		}
	}
	githubbutton := widget.NewButton("Github page", func() {
		go func() {
			u, _ := url.Parse("https://github.com/alexballas/go2tv")
			_ = fyne.CurrentApp().OpenURL(u)
		}()
	})
	checkversion := widget.NewButton("Check version", func() {
		go checkVersion(s)
	})

	s.CheckVersion = checkversion

	return container.NewVBox(richhead, container.NewCenter(container.NewHBox(githubbutton, checkversion)))

}

func checkVersion(s *NewScreen) {
	s.CheckVersion.Disable()
	defer s.CheckVersion.Enable()
	errRedirectChecker := fmt.Errorf("redirect")
	errVersioncomp := fmt.Errorf("cant compare versions - on develop or non-compiled version")
	errVersionGet := fmt.Errorf("failed to get version info - check your internet connection")

	extractVer := strings.Split(s.version, "/")
	str := strings.ReplaceAll(extractVer[0], ".", "")
	str = strings.TrimSpace(str)
	currversion, err := strconv.Atoi(str)
	if err != nil {
		dialog.ShowError(errVersioncomp, s.Current)
		return
	}

	req, err := http.NewRequest("GET", "https://github.com/alexballas/Go2TV/releases/latest", nil)
	if err != nil {
		dialog.ShowError(errVersionGet, s.Current)
	}

	client := new(http.Client)
	client.Timeout = time.Duration(3 * time.Second)
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return errRedirectChecker
	}

	response, err := client.Do(req)
	if errors.Is(err, errRedirectChecker) {
		url, err := response.Location()
		if err != nil {
			dialog.ShowError(errVersionGet, s.Current)
			return
		}
		str := strings.Trim(filepath.Base(url.Path), "v")
		str = strings.ReplaceAll(str, ".", "")
		chversion, err := strconv.Atoi(str)
		if err != nil {
			dialog.ShowError(errVersionGet, s.Current)
			return
		}

		switch {
		case chversion > currversion:
			dialog.ShowInformation("Version checker", "New version: "+filepath.Base(url.Path), s.Current)
			return
		default:
			dialog.ShowInformation("Version checker", "No new version", s.Current)
			return
		}
	}

	dialog.ShowError(errVersionGet, s.Current)
}