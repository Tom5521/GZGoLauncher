package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (ui *ui) StartMainWindow() {
	metadata := ui.App.Metadata()
	ui.MainWindow = ui.App.NewWindow(metadata.Name)
	ui.MainWindow.SetIcon(metadata.Icon)
	mainTabs := container.NewAppTabs(
		container.NewTabItem(po.Get("Run"), ui.MainBox()),
		container.NewTabItem(po.Get("Settings"), configuration.MainBox(ui)),
	)
	ui.MainWindow.Resize(fyne.NewSize(1100, 500))
	ui.MainWindow.SetContent(mainTabs)
	ui.MainWindow.ShowAndRun()
}

func (ui *ui) MainBox() *fyne.Container {
	selectbox := ui.SelectCont()
	rightbox := container.NewHBox(widget.NewSeparator(), container.NewVScroll(ui.RightCont()))
	bottom := ui.Bottom()
	content := container.NewBorder(nil, bottom, nil, rightbox, selectbox)
	return content
}
