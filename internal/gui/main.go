package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (ui *ui) StartMainWindow() {
	ui.MainWindow = ui.App.NewWindow(ui.App.Metadata().Name)
	ui.MainWindow.Resize(fyne.NewSize(500, 600))
	ui.MainWindow.SetContent(ui.MainContent())
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
	downContent := container.NewBorder(nil, nil, nil, runButton)
	content := container.NewBorder(nil, downContent, nil, rightContent, selectConts)
	return content
}
