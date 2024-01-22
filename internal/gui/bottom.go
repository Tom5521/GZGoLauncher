package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	boxes "fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (ui *ui) BottomBox() *fyne.Container {
	runButton := &widget.Button{
		Text:       po.Get("Run"),
		Importance: widget.HighImportance,
		OnTapped: func() {
			ui.RunDoom()
		}}
	ui.ZRunnerSelect = &widget.Select{
		Options: []string{"GZDoom", "ZDoom"},
		OnChanged: func(s string) {
			switch s {
			case "GZDoom":
				settings.ExecDir = settings.GZDoomDir
			case "ZDoom":
				settings.ExecDir = settings.ZDoomDir
			default:
				return
			}
		},
		PlaceHolder: po.Get("Select a Runner"),
	}
	switch settings.ExecDir {
	case "":
		ui.ZRunnerSelect.ClearSelected()
	case settings.GZDoomDir:
		ui.ZRunnerSelect.SetSelected("GZDoom")
	case settings.ZDoomDir:
		ui.ZRunnerSelect.SetSelected("GZDoom")
	default:
		ui.ZRunnerSelect.ClearSelected()
	}
	cArgsLabel := widget.NewRichTextFromMarkdown(fmt.Sprintf(
		"[%s](https://zdoom.org/wiki/Command_line_parameters)",
		po.Get("Custom arguments:"),
	))
	ui.CustomArgs = &widget.Entry{Text: settings.CustomArgs}
	ui.CustomArgs.OnChanged = func(s string) {
		settings.CustomArgs = s
	}
	ui.CustomArgs.SetPlaceHolder(po.Get("Example: %s", "-fast"))

	customArgsBox := boxes.NewBorder(nil, nil, cArgsLabel, nil, ui.CustomArgs)
	rightBox := boxes.NewHBox(ui.ZRunnerSelect, runButton)
	content := boxes.NewBorder(nil, nil, nil, rightBox, customArgsBox)
	return content
}
