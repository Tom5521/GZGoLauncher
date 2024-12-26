package gui

import (
	"fmt"
	"slices"

	"fyne.io/fyne/v2"
	boxes "fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func runnerNames() (runners []string) {
	for _, sp := range settings.SourcePorts {
		runners = append(runners, sp.Name)
	}
	return
}

func (ui *ui) refreshZRunnerSelection() {
	ui.ZRunnerSelect.SetOptions(runnerNames())
	ui.ZRunnerSelect.ClearSelected()
}

func (ui *ui) BottomBox() *fyne.Container {
	runButton := &widget.Button{
		Text:       po.Get("Run"),
		Importance: widget.HighImportance,
		OnTapped: func() {
			ui.RunDoom()
		},
	}

	ui.ZRunnerSelect = &widget.Select{
		Options: runnerNames(),
		OnChanged: func(s string) {
			index := slices.Index(runnerNames(), s)
			settings.CurrentSourcePort = index
			if index == -1 {
				return
			}
			sp := settings.SourcePorts[index]
			Runner.RunnerPath = sp.ExecutablePath
		},
		PlaceHolder: po.Get("Select a Runner"),
	}
	ui.ZRunnerSelect.SetSelectedIndex(settings.CurrentSourcePort)

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
