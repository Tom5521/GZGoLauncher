package gui

import (
	"slices"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/GZGoLauncher/internal/config"
	"github.com/Tom5521/GZGoLauncher/internal/tools"
	"github.com/Tom5521/GoNotes/pkg/messages"
)

func (ui *ui) StartMainWindow() {
	ui.MainWindow = ui.App.NewWindow(ui.App.Metadata().Name)
	ui.MainWindow.SetContent(ui.MainContent())
	ui.MainWindow.ShowAndRun()
}

func (ui *ui) MainContent() *fyne.Container {
	selectConts := ui.SelectCont()
	runButton := &widget.Button{Text: "Run", Importance: widget.HighImportance, OnTapped: func() { RunDoom() }}
	rightContent := container.NewVBox(widget.NewLabel("TEST"))
	downContent := container.NewBorder(nil, nil, nil, runButton)
	content := container.NewBorder(nil, downContent, nil, rightContent, selectConts)
	return content
}

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
		SelectedWad = string(settings.Wads[id])
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
		config.Settings.Wads = slices.Delete(settings.Wads, selected, selected+1)
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
	var selected = -1
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
			ctr := co.(*fyne.Container)
			l := ctr.Objects[0].(*widget.Label)
			c := ctr.Objects[1].(*widget.Check)
			l.SetText(settings.Mods[lii].Path)
			c.OnChanged = func(b bool) {
				settings.Mods[lii].Enabled = b
			}
		},
	)
	ui.ModsList.OnSelected = func(id widget.ListItemID) {
		selected = id
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
		if selected == -1 {
			return
		}
		settings.Mods = slices.Delete(settings.Mods, selected, selected+1)
		err := settings.Write()
		if err != nil {
			messages.Error(err)
		}
		ui.ModsList.UnselectAll()
		selected = -1
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

func toolbar(leftItem fyne.CanvasObject, plus, minus func()) *fyne.Container {
	toolbar := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), minus),
		widget.NewToolbarAction(theme.ContentAddIcon(), plus),
	)
	content := container.NewBorder(nil, nil, leftItem, nil, toolbar)
	return content
}
