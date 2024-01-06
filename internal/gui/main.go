package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/ncruces/zenity"
)

func (ui *ui) StartMainWindow() {
	ui.MainWindow = ui.App.NewWindow(ui.App.Metadata().Name)
	ui.MainWindow.SetIcon(ui.App.Metadata().Icon)
	appTabs := container.NewAppTabs(
		container.NewTabItem("Run", ui.MainContent()),
		container.NewTabItem("Settings", ui.Configuration()),
	)
	ui.MainWindow.Resize(fyne.NewSize(500, 600))
	ui.MainWindow.SetContent(appTabs)
	ui.MainWindow.ShowAndRun()
}

func (ui *ui) MainContent() *fyne.Container {
	selectConts := ui.SelectCont()
	runButton := &widget.Button{
		Text:       "Run",
		Importance: widget.HighImportance,
		OnTapped: func() {
			ui.RunDoom()
		}}
	rightContent := RightCont()
	downContent := container.NewBorder(nil, nil, nil, func() *fyne.Container {
		runnerSelect := widget.NewSelect([]string{"GZDoom", "ZDoom"}, func(s string) {
			if s == "" {
				return
			}
			if s == "GZDoom" {
				settings.GZDir = settings.GZDoomDir
			}
			if s == "ZDoom" {
				settings.GZDir = settings.ZDoomDir
			}
			err := settings.Write()
			if err != nil {
				zenity.Error(err.Error())
			}
		})
		runnerSelect.SetSelected("GZDoom")
		c := container.NewHBox(runnerSelect, runButton)
		return c
	}())
	content := container.NewBorder(nil, downContent, nil, rightContent, selectConts)
	return content
}
