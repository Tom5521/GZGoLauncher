package gui

import (
	"fyne.io/fyne/v2"
	boxes "fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (ui *ui) StartMainWindow() {
	metadata := ui.App.Metadata()
	ui.MainWindow = ui.App.NewWindow(metadata.Name)
	ui.MainWindow.SetMaster()
	ui.MainWindow.SetIcon(metadata.Icon)
	mainTabs := boxes.NewAppTabs(
		boxes.NewTabItem(po.Get("Run"), ui.MainBox()),
		boxes.NewTabItem(po.Get("Settings"), configuration.MainBox(ui)),
	)
	ui.MainWindow.Resize(fyne.NewSize(1100, 500))
	ui.MainWindow.SetContent(mainTabs)
	ui.MainWindow.ShowAndRun()
}

func (ui *ui) MainBox() *fyne.Container {
	selectbox := ui.SelectBox()
	rightbox := boxes.NewHBox(widget.NewSeparator(), boxes.NewVScroll(ui.RightBox()))
	bottom := ui.BottomBox()
	content := boxes.NewBorder(nil, bottom, nil, rightbox, selectbox)
	return content
}
