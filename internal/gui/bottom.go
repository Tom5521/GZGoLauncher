package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (ui *ui) Bottom() *fyne.Container {
	runButton := &widget.Button{
		Text:       po.Get("Run"),
		Importance: widget.HighImportance,
		OnTapped: func() {
			ui.RunDoom()
		}}
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
	cArgsLabel := widget.NewRichTextFromMarkdown(fmt.Sprintf(
		"[%s](https://zdoom.org/wiki/Command_line_parameters)",
		po.Get("Custom arguments:"),
	))
	ui.CustomArgs = &widget.Entry{Text: settings.CustomArgs}
	ui.CustomArgs.OnChanged = func(s string) {
		settings.CustomArgs = s
	}
	ui.CustomArgs.SetPlaceHolder(po.Get("Example: %s", "-fast"))

	customArgsBox := container.NewBorder(nil, nil, cArgsLabel, nil, ui.CustomArgs)
	rightBox := container.NewHBox(ui.ZRunnerSelect, runButton)
	content := container.NewBorder(nil, nil, nil, rightBox, customArgsBox)
	return content
}
