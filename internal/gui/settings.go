package gui

import "github.com/diamondburned/gotk4/pkg/gtk/v4"

func (ui *ui) SettingsTab() gtk.Widgetter {
	box := gtk.NewBox(gtk.OrientationVertical, 6)
	return box
}
