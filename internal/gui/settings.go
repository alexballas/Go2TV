package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func settingsWindow(s *NewScreen) fyne.CanvasObject {
	themeText := canvas.NewText("Theme", nil)
	dropdown := widget.NewSelect([]string{"Light", "Dark", "Default"}, parseTheme(s))
	dropdown.PlaceHolder = "yo"
	settings := container.NewVBox(themeText, dropdown)
	return settings
}

func parseTheme(s *NewScreen) func(string) {
	return func(t string) {
		switch t {
		case "Light":
			s.config.Theme = "Light"
		case "Dark":
			s.config.Theme = "Dark"
		case "Default":
			s.config.Theme = "Default"
		}
		s.config.ApplyAppConfig()
	}
}
