package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	boxes "fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (ui *ui) ShowOutWin(text string) {
	w := ui.App.NewWindow(po.Get("Output"))
	w.Resize(fyne.NewSize(550, 550))
	text = fmt.Sprintf("```\n%s\n```", text)
	topBox := boxes.NewVBox()
	cmd := widget.NewRichTextFromMarkdown(fmt.Sprintf("**Command:** `%s`", Runner.MakeCmd()))
	topBox.Add(cmd)
	richText := widget.NewRichTextFromMarkdown(text)
	if Runner.Error != nil {
		richText.ParseMarkdown(fmt.Sprintf("## Error:\n ```\n%s\n```", Runner.ErrOut))
	}
	content := boxes.NewBorder(topBox, nil, nil, nil, richText)
	w.SetContent(boxes.NewScroll(content))
	w.Show()
}
