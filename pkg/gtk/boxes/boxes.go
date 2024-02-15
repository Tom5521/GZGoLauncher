package boxes

import "github.com/diamondburned/gotk4/pkg/gtk/v4"

func NewVbox(widgets ...gtk.Widgetter) *gtk.Box {
	vbox := gtk.NewBox(gtk.OrientationVertical, 6)
	for _, w := range widgets {
		vbox.Append(w)
	}

	return vbox
}

func NewHbox(widgets ...gtk.Widgetter) *gtk.Box {
	hbox := gtk.NewBox(gtk.OrientationHorizontal, 6)
	for _, w := range widgets {
		hbox.Append(w)
	}
	return hbox
}

func NewScrolledVbox(widgets ...gtk.Widgetter) *gtk.ScrolledWindow {
	vbox := NewVbox(widgets...)
	sbox := gtk.NewScrolledWindow()
	sbox.SetChild(vbox)
	return sbox
}

func NewScrolledHbox(widgets ...gtk.Widgetter) *gtk.ScrolledWindow {
	hbox := NewHbox(widgets...)
	sbox := gtk.NewScrolledWindow()
	sbox.SetChild(hbox)
	return sbox
}

func NewVPaned(top, bottom gtk.Widgetter) *gtk.Paned {
	paned := gtk.NewPaned(gtk.OrientationVertical)
	paned.SetStartChild(top)
	paned.SetEndChild(bottom)
	return paned
}
func NewHPaned(left, right gtk.Widgetter) *gtk.Paned {
	paned := gtk.NewPaned(gtk.OrientationHorizontal)
	paned.SetStartChild(left)
	paned.SetEndChild(right)
	return paned
}

func NewFrame(label string, child gtk.Widgetter) *gtk.Frame {
	frame := gtk.NewFrame(label)
	frame.SetChild(child)

	return frame
}

func NewAdaptativeGrid(size int, widgets ...gtk.Widgetter) *gtk.Grid {
	grid := gtk.NewGrid()

	var rowCount, columnCount int
	for _, w := range widgets {
		isHided := w.ObjectProperty("visible")
		if !isHided.(bool) {
			continue
		}
		grid.Attach(w, columnCount, rowCount, 1, 1)
		if columnCount >= size {
			rowCount++
			columnCount = 0
			continue
		}
		columnCount++
	}

	return grid
}
