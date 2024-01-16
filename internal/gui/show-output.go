package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (ui *ui) ShowOutWin(text string) {
	w := ui.App.NewWindow(po.Get("Output"))
	w.Resize(fyne.NewSize(550, 550))
	text = fmt.Sprintf("```\n%s\n```", text)
	topBox := container.NewVBox()
	cmd := widget.NewRichTextFromMarkdown(fmt.Sprintf("**Command:** `%s`", Runner.MakeCmd()))
	topBox.Add(cmd)
	richText := widget.NewRichTextFromMarkdown(text)
	if Runner.Error != nil {
		richText.ParseMarkdown(fmt.Sprintf("## Error:\n ```\n%s\n```", Runner.ErrOut))
	}
	content := container.NewBorder(topBox, nil, nil, nil, richText)
	w.SetContent(container.NewScroll(content))
	w.Show()
}
