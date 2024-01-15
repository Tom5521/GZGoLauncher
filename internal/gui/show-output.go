package gui

import (
	"fmt"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (ui *ui) ShowOutWin(text string) {
	w := ui.App.NewWindow(po.Get("Output"))
	text = fmt.Sprintf("```\n%s\n```", text)
	cmd := widget.NewRichTextFromMarkdown(fmt.Sprintf("**Command:** `%s`", Runner.MakeCmd()))
	richText := widget.NewRichTextFromMarkdown(text)
	content := container.NewBorder(cmd, nil, nil, nil, richText)
	w.SetContent(container.NewVScroll(content))
	w.Show()
}
