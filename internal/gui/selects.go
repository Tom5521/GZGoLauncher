package gui

import (
	"os"
	"slices"

	"fyne.io/fyne/v2"
	boxes "fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/GZGoLauncher/internal/config"
	"github.com/Tom5521/GZGoLauncher/internal/filepicker"
)

func (ui *ui) SelectBox() *boxes.Split {
	content := boxes.NewVSplit(ui.IwadsCont(), ui.ModsCont())
	return content
}

func (ui *ui) IwadsCont() *fyne.Container {
	var selected = -1
	selectLabel := widget.NewRichTextFromMarkdown(po.Get("**Select wad**"))
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
	if Runner.IWad != "" {
		for i, w := range settings.Wads {
			if Runner.IWad == string(w) {
				ui.WadList.Select(i)
			}
		}
	}

	add := func() {
		files := filepicker.Wad.MultiStart()
		for _, file := range files {
			if file == "" {
				return
			}
			newWad := config.Wad(file)
			if !newWad.IsValid() {
				ErrWin(po.Get("The file is not valid!"))
				return
			}
			i := slices.IndexFunc(settings.Wads, func(w config.Wad) bool {
				return w == newWad
			})
			if i != -1 {
				ErrWin(po.Get("The file already exists"))
				return
			}

			settings.Wads = append(settings.Wads, newWad)
		}
		ui.WadList.Refresh()
	}
	remove := func() {
		if selected == -1 {
			return
		}
		settings.Wads = slices.Delete(settings.Wads, selected, selected+1)
		ui.WadList.UnselectAll()
		Runner.IWad = ""
		selected = -1
	}
	toolbar := toolbar(selectLabel, add, remove)

	content := boxes.NewBorder(toolbar, nil, nil, nil, ui.WadList)
	return content
}

func (ui *ui) ModsCont() *fyne.Container {
	selectModLabel := widget.NewRichTextFromMarkdown(po.Get("**Select mods to use**"))
	ui.ModsList = widget.NewList(
		func() int {
			return len(settings.Mods)
		},
		func() fyne.CanvasObject {
			return boxes.NewBorder(
				nil, nil, nil,
				&widget.Check{},
				&widget.Label{},
			)
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			mod := &settings.Mods[i]
			ctr := o.(*fyne.Container)
			l := ctr.Objects[0].(*widget.Label)
			c := ctr.Objects[1].(*widget.Check)
			l.SetText(mod.Path)
			c.SetChecked(mod.Enabled)
			c.OnChanged = func(b bool) {
				mod.Enabled = b
				err := settings.Write()
				if err != nil {
					ErrWin(err)
				}
			}
		},
	)
	ui.ModsList.OnSelected = func(id widget.ListItemID) {
		ui.ModsList.Unselect(id)
	}
	add := func() {
		newMods := filepicker.Pk3.MultiStart()
		for _, newMod := range newMods {
			if newMod == "" {
				return
			}
			stat, err := os.Stat(newMod)
			if os.IsNotExist(err) {
				ErrWin(po.Get("The file is not valid!"))
				return
			}
			if stat.IsDir() {
				ErrWin(po.Get("The file is not valid!"))
				return
			}
			i := slices.IndexFunc(settings.Mods, func(m config.Mod) bool {
				return m.Path == newMod
			})
			if i != -1 {
				ErrWin(po.Get("The file already exists"))
				return
			}
			settings.Mods = append(settings.Mods, config.Mod{Path: newMod})
		}
		ui.ModsList.Refresh()
	}
	remove := func() {
		var toDelete []int
	next:
		toDelete = []int{}
		for i, mod := range settings.Mods {
			if mod.Enabled {
				toDelete = append(toDelete, i)
			}
		}
		for _, f := range toDelete {
			settings.Mods = slices.Delete(settings.Mods, f, f+1)
			goto next
		}
		ui.ModsList.Refresh()
	}
	bar := toolbar(selectModLabel, add, remove)
	content := boxes.NewBorder(bar, nil, nil, nil, ui.ModsList)
	return content
}

func enabledPaths() []string {
	var paths []string
	for _, i := range settings.Mods {
		if i.Enabled {
			paths = append(paths, i.Path)
		}
	}
	return paths
}

func toolbar(leftItem fyne.CanvasObject, plus, minus func()) *fyne.Container {
	toolbar := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), minus),
		widget.NewToolbarAction(theme.ContentAddIcon(), plus),
	)
	content := boxes.NewBorder(nil, nil, leftItem, nil, toolbar)
	return content
}
