package gui

import (
	"fmt"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (ui *ui) ShowOutWin(text string) {
	w := ui.App.NewWindow(po.Get("Output"))
	text = fmt.Sprintf("```\n%s\n```", text)
	richText := widget.NewRichTextFromMarkdown(text)
	w.SetContent(container.NewVScroll(richText))
	w.Show()
}
