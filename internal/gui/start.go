package gui

import (
	"github.com/Tom5521/gtk4tools/pkg/boxes"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func (ui *ui) StartTab() gtk.Widgetter {
	listBox := boxes.NewVbox(
		ui.WadListBox(),
		gtk.NewSeparator(gtk.OrientationHorizontal),
		ui.ModsListBox(),
	)
	listBox.SetHExpand(true)

	right := ui.Right()

	mainBox := boxes.NewVbox(
		boxes.NewHbox(
			listBox,
			gtk.NewSeparator(gtk.OrientationVertical),
			right,
		),
		gtk.NewSeparator(gtk.OrientationHorizontal),
		ui.BottomBox(),
	)

	return mainBox
}
