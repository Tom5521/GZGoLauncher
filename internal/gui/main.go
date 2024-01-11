package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (ui *ui) StartMainWindow() {
	ui.MainWindow = ui.App.NewWindow(ui.App.Metadata().Name)
	ui.MainWindow.SetIcon(ui.App.Metadata().Icon)
	appTabs := container.NewAppTabs(
		container.NewTabItem(po.Get("Run"), ui.MainContent()),
		container.NewTabItem(po.Get("Settings"), configuration.Container(ui)),
	)
	ui.MainWindow.Resize(fyne.NewSize(1100, 500))
	ui.MainWindow.SetContent(appTabs)
	ui.MainWindow.ShowAndRun()
}

func (ui *ui) MainContent() *fyne.Container {
	selectConts := ui.SelectCont()
	runButton := &widget.Button{
		Text:       po.Get("Run"),
		Importance: widget.HighImportance,
		OnTapped: func() {
			ui.RunDoom()
		}}
	rightContent := container.NewHBox(widget.NewSeparator(), RightCont())
	downContent := container.NewBorder(nil, nil, nil, func() *fyne.Container {
		ui.ZRunnerSelect = widget.NewSelect([]string{"GZDoom", "ZDoom"}, func(s string) {
			switch s {
			case "GZDoom":
				settings.GZDir = settings.GZDoomDir
			case "ZDoom":
				settings.GZDir = settings.ZDoomDir
			default:
				return
			}
			err := settings.Write()
			if err != nil {
				ErrWin(err)
			}
		})
		switch settings.GZDir {
		case settings.GZDoomDir:
			ui.ZRunnerSelect.SetSelected("GZDoom")
		case settings.ZDoomDir:
			ui.ZRunnerSelect.SetSelected("ZDoom")
		default:
			ui.ZRunnerSelect.ClearSelected()
		}
		ui.ZRunnerSelect.PlaceHolder = po.Get("Select a Runner")
		c := container.NewHBox(ui.ZRunnerSelect, runButton)
		return c
	}())
	content := container.NewBorder(nil, downContent, nil, rightContent, selectConts)
	return content
}
