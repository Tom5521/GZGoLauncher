package gtools

import "github.com/diamondburned/gotk4/pkg/gtk/v4"

type Appender interface {
	Append(gtk.Widgetter)
}

func Appends(parent Appender, widgets ...gtk.Widgetter) {
	for _, w := range widgets {
		parent.Append(w)
	}
}
