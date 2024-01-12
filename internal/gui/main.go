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
	rightContent := container.NewHBox(widget.NewSeparator(), container.NewVScroll(RightCont()))
	downContent := container.NewBorder(nil, nil, nil, func() *fyne.Container {
		ui.ZRunnerSelect = &widget.Select{
			Selected: func() string {
				switch settings.GZDir {
				case settings.GZDoomDir:
					return "GZDoom"
				case settings.ZDoomDir:
					return "ZDoom"
				default:
					return "GZDoom"
				}
			}(),
			Options: []string{"GZDoom", "ZDoom"},
			OnChanged: func(s string) {
				switch s {
				case "GZDoom":
					settings.GZDir = settings.GZDoomDir
				case "ZDoom":
					settings.GZDir = settings.ZDoomDir
				default:
					return
				}
			},
			PlaceHolder: po.Get("Select a Runner"),
		}
		c := container.NewHBox(ui.ZRunnerSelect, runButton)
		return c
	}(), func() *fyne.Container {
		label := widget.NewLabel(po.Get("Custom arguments:"))
		ui.CustomArgs = &widget.Entry{Text: settings.CustomArgs}
		ui.CustomArgs.OnChanged = func(s string) {
			settings.CustomArgs = s
		}
		ui.CustomArgs.SetPlaceHolder(po.Get("Example: %s", "-fast"))
		return container.NewBorder(nil, nil, label, nil, ui.CustomArgs)
	}(),
	)
	content := container.NewBorder(nil, downContent, nil, rightContent, selectConts)
	return content
}
