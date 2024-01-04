package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

type gui struct {
	ui
}

type ui struct {
	App        fyne.App
	MainWindow fyne.Window
	WadList    *widget.List
	ModList    *widget.List
}

const (
	IWADsID     string = "IWADs"
	MODsID      string = "MODs"
	ProyectName string = "GZGoLauncher"
)

var (
	IWADs    []string
	MODs     []string
	settings fyne.Preferences
)

func Init() *ui {
	newApp := app.NewWithID("com.github.Tom5521.GZGoLauncher.preferences")
	newUI := &ui{
		App: newApp,
	}
	settings = newApp.Preferences()
	IWADs = settings.StringList(IWADsID)

	return newUI
}
