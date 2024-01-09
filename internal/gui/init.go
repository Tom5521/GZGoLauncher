package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/GZGoLauncher/internal/config"
)

type Gui struct {
	ui
}

type ui struct {
	App           fyne.App
	MainWindow    fyne.Window
	WadList       *widget.List
	ModsList      *widget.List
	ZRunnerSelect *widget.Select
}

var (
	settings     = &config.Settings
	CloseOnStart bool
)

func Init() *Gui {
	app := app.New()
	ui := &Gui{}
	ui.App = app
	return ui
}
