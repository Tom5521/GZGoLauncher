package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/GZGoLauncher/internal/config"
	"github.com/Tom5521/GZGoLauncher/internal/tools"
	"github.com/Tom5521/GoNotes/pkg/messages"
)

func (ui *ui) SelectCont() *container.Split {
	content := container.NewVSplit(ui.IwadsCont(), ui.ModsCont())
	return content
}

func (ui *ui) IwadsCont() *fyne.Container {
	var selected = -1
	selectLabel := widget.NewRichTextFromMarkdown("**Select wad**")
	ui.WadList = widget.NewList(
		func() int {
			return len(settings.Wads)
		},
		func() fyne.CanvasObject {
			return &widget.Label{}
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(string(settings.Wads[lii]))
		},
	)
	ui.WadList.OnSelected = func(id widget.ListItemID) {
		Runner.IWad = string(settings.Wads[id])
		selected = id
	}

	add := func() {
		file := tools.WadFilePicker()
		if file == "" {
			return
		}
		newWad := config.Wad(file)
		if !newWad.IsValid() {
			return
		}
		settings.Wads = append(settings.Wads, newWad)
		err := settings.Write()
		if err != nil {
			messages.Error(err)
		}
		ui.WadList.Refresh()
	}
	remove := func() {
		if selected == -1 {
			return
		}
		settings.Wads = deleteSlice(settings.Wads, selected)
		err := config.Settings.Write()
		if err != nil {
			messages.Error(err)
		}
		ui.WadList.UnselectAll()
		selected = -1
	}
	toolbar := toolbar(selectLabel, add, remove)

	content := container.NewBorder(toolbar, nil, nil, nil, ui.WadList)
	return content
}

func (ui *ui) ModsCont() *fyne.Container {
	selectModLabel := widget.NewRichTextFromMarkdown("**Select mods to use**")
	ui.ModsList = widget.NewList(
		func() int {
			return len(settings.Mods)
		},
		func() fyne.CanvasObject {
			return container.NewBorder(
				nil, nil, nil,
				&widget.Check{},
				&widget.Label{},
			)
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			mod := &settings.Mods[lii]
			ctr := co.(*fyne.Container)
			l := ctr.Objects[0].(*widget.Label)
			c := ctr.Objects[1].(*widget.Check)
			l.SetText(settings.Mods[lii].Path)
			c.SetChecked(mod.Enabled)
			c.OnChanged = func(b bool) {
				mod.Enabled = b
				settings.Write()
			}
		},
	)
	ui.ModsList.OnSelected = func(id widget.ListItemID) {
		ui.ModsList.UnselectAll()
	}
	add := func() {
		newMod := tools.PK3FilePicker()
		if newMod == "" {
			return
		}
		settings.Mods = append(settings.Mods, config.Mod{Path: newMod})
		err := settings.Write()
		if err != nil {
			messages.Error(err)
		}
		ui.ModsList.Refresh()
	}
	remove := func() {
		for index, i := range settings.Mods {
			if i.Enabled {
				settings.Mods = deleteSlice(settings.Mods, index)
			}
		}
		err := settings.Write()
		if err != nil {
			messages.Error(err)
		}
		ui.ModsList.Refresh()
	}
	bar := toolbar(selectModLabel, add, remove)
	content := container.NewBorder(bar, nil, nil, nil, ui.ModsList)
	return content
}

func enabledMods() []config.Mod {
	var enableds []config.Mod
	for _, i := range settings.Mods {
		if i.Enabled {
			enableds = append(enableds, i)
		}
	}
	return enableds
}

func enabledPaths() []string {
	mods := enabledMods()
	var paths []string
	for _, i := range mods {
		paths = append(paths, i.Path)
	}
	return paths
}

func deleteSlice[S ~[]E, E any](slice S, index int) S {
	if index < 0 || index >= len(slice) {
		return slice
	}

	newSlice := append(slice[:index], slice[index+1:]...)

	return newSlice
}

func toolbar(leftItem fyne.CanvasObject, plus, minus func()) *fyne.Container {
	toolbar := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), minus),
		widget.NewToolbarAction(theme.ContentAddIcon(), plus),
	)
	content := container.NewBorder(nil, nil, leftItem, nil, toolbar)
	return content
}
