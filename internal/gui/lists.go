package gui

import (
	"fmt"
	"slices"

	"github.com/Tom5521/GZLauncher-gtk/internal/config"
	"github.com/Tom5521/GZLauncher-gtk/internal/filepicker"
	"github.com/Tom5521/GZLauncher-gtk/internal/tools"
	t "github.com/Tom5521/GZLauncher-gtk/internal/tools"
	"github.com/Tom5521/gtk4tools/pkg/boxes"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func (ui *ui) WadList() *gtk.ListView {
	ui.WadsListModel = gtk.NewStringList(settings.Wads)

	factory := gtk.NewSignalListItemFactory()
	factory.ConnectSetup(func(listitem *gtk.ListItem) {
		listitem.SetChild(gtk.NewLabel(""))
	})
	factory.ConnectBind(func(listitem *gtk.ListItem) {
		label := listitem.Child().(*gtk.Label)
		obj := listitem.Item().Cast().(*gtk.StringObject)
		label.SetLabel(obj.String())
	})

	ui.WadsListSModel = gtk.NewSingleSelection(ui.WadsListModel)

	if runner.IWad != "" {
		i := slices.Index(settings.Wads, runner.IWad)
		if i != -1 {
			ui.WadsListSModel.SetSelected(uint(i))
		}
	}

	ui.WadsListSModel.ConnectSelectionChanged(func(_, _ uint) {
		if ui.WadsListSModel.SelectedItem() == nil {
			runner.IWad = ""
			return
		}
		index := ui.WadsListSModel.Selected()
		runner.IWad = settings.Wads[index]
	})
	ui.WadsListView = gtk.NewListView(ui.WadsListSModel, &factory.ListItemFactory)
	return ui.WadsListView
}

func (ui *ui) ModsList() *gtk.ListView {
	var mods []string
	for _, m := range settings.Mods {
		mods = append(mods, m.Path)
	}
	ui.ModsListModel = gtk.NewStringList(mods)
	ui.ModsListSModel = gtk.NewNoSelection(ui.ModsListModel)

	factory := gtk.NewSignalListItemFactory()
	factory.ConnectSetup(func(listitem *gtk.ListItem) {
		checkBox := gtk.NewCheckButton()
		label := gtk.NewLabel("")
		label.SetHExpand(true)

		box := boxes.NewHbox(label, checkBox)

		listitem.SetChild(box)
	})
	factory.ConnectBind(func(listitem *gtk.ListItem) {
		obj := listitem.Item().Cast().(*gtk.StringObject)
		box := listitem.Child().(*gtk.Box)
		boxItems := box.ObserveChildren()
		label := boxItems.Item(0).Cast().(*gtk.Label)
		check := boxItems.Item(1).Cast().(*gtk.CheckButton)

		label.SetLabel(obj.String())

		var modIndex int

		for i, m := range settings.Mods {
			if m.Path == obj.String() {
				modIndex = i
				if m.Enabled {
					check.Activate()
					break
				}
			}
		}

		check.ConnectToggled(func() {
			settings.Mods[modIndex].Enabled = check.Active()
			err := settings.Write()
			if err != nil {
				t.ErrWin(err)
			}
		})

	})

	ui.ModsListView = gtk.NewListView(ui.ModsListSModel, &factory.ListItemFactory)
	return ui.ModsListView
}

func (ui *ui) WadListBox() gtk.Widgetter {
	label := gtk.NewLabel(po.Get("Select wad"))
	label.SetMarkup(fmt.Sprintf("<b>%s</b>", label.Label()))
	label.SetHExpand(true)

	addButton := gtk.NewButtonFromIconName("list-add")
	addButton.ConnectClicked(func() {
		wads := filepicker.Wad.MultiStart()
		for _, w := range wads {
			i := slices.Index(settings.Wads, w)
			if i != -1 {
				tools.ErrWin(po.Get("The file already exists"))
				return
			}
			ui.WadsListModel.Append(w)
			settings.Wads = append(settings.Wads, w)
		}
	})
	removeButton := gtk.NewButtonFromIconName("list-remove")
	removeButton.ConnectClicked(func() {
		pos := ui.WadsListSModel.Selected()
		item := ui.WadsListModel.Item(pos)
		if item == nil {
			return
		}
		if settings.Wads[pos] == runner.IWad {
			runner.IWad = ""
		}
		ui.WadsListModel.Remove(pos)
		settings.Wads = slices.Delete(settings.Wads, int(pos), int(pos)+1)
	})

	list := ui.WadList()
	list.SetVExpand(true)

	labelBox := boxes.NewHbox(
		label,
		addButton,
		removeButton,
	)

	listBox := boxes.NewVbox(
		labelBox,
		boxes.NewScrolledVbox(
			list,
		),
	)
	listBox.SetHExpand(true)

	return listBox
}

func (ui *ui) ModsListBox() gtk.Widgetter {
	label := gtk.NewLabel(po.Get("Select mods to use"))
	label.SetMarkup(fmt.Sprintf("<b>%s</b>", label.Label()))
	label.SetHExpand(true)

	addButton := gtk.NewButtonFromIconName("list-add")
	addButton.ConnectClicked(func() {
		mods := filepicker.Pk3.MultiStart()
		for _, m := range mods {
			i := slices.IndexFunc(settings.Mods, func(mod config.Mod) bool {
				return m == mod.Path
			})
			if i != -1 {
				tools.ErrWin(po.Get("The file already exists"))
				return
			}
			settings.Mods = append(settings.Mods, config.Mod{Path: m})
			ui.ModsListModel.Append(m)
		}
	})
	removeButton := gtk.NewButtonFromIconName("list-remove")
	removeButton.ConnectClicked(func() {
	next:
		var toDelete []int
		for i, m := range settings.Mods {
			if m.Enabled {
				toDelete = append(toDelete, i)
			}
		}
		for _, f := range toDelete {
			settings.Mods = slices.Delete(settings.Mods, f, f+1)
			goto next
		}
		var newMods []string
		for _, m := range settings.Mods {
			newMods = append(newMods, m.Path)
		}
		ui.ModsListModel.Splice(0, ui.ModsListModel.NItems(), newMods)
	})

	list := ui.ModsList()
	list.SetVExpand(true)

	labelBox := boxes.NewHbox(
		label,
		addButton,
		removeButton,
	)
	listBox := boxes.NewVbox(
		labelBox,
		boxes.NewScrolledVbox(
			list,
		),
	)
	listBox.SetHExpand(true)

	return listBox
}
