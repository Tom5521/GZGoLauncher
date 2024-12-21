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
	cmd := widget.NewRichTextFromMarkdown(fmt.Sprintf("**Command:** `%s`", Runner.MakeCmd().String()))
	topBox.Add(cmd)
	richText := widget.NewRichTextFromMarkdown(fmt.Sprintf("## Output:\n%s", text))
	if Runner.Error != nil {
		errLb := widget.NewRichTextFromMarkdown(fmt.Sprintf("## Error:\n`%s`", Runner.Error))
		topBox.Add(errLb)
		richText.ParseMarkdown(fmt.Sprintf("## Error Output:\n%s", text))
	}
	richText.Wrapping = fyne.TextWrapBreak
	content := boxes.NewBorder(topBox, nil, nil, nil, richText)
	w.SetContent(
		boxes.NewBorder(
			nil,
			widget.NewButton("Ok", func() { w.Close() }),
			nil,
			nil,
			boxes.NewScroll(content),
		),
	)
	w.Show()
}
