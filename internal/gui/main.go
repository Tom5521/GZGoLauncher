package gui

import (
	"slices"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (ui *ui) InitMainWindow() {
	ui.MainWindow = ui.App.NewWindow(ProyectName)
	ui.MainWindow.SetContent(ui.SelectContainer())
	ui.MainWindow.ShowAndRun()
}

func (ui *ui) SelectContainer() *container.Split {
	wadSelect := ui.WadSelect()
	modSelect := ui.ModSelect()
	split := container.NewVSplit(wadSelect, modSelect)
	return split
}

func (ui *ui) WadSelect() *fyne.Container {
	selectWadLabel := widget.NewRichTextFromMarkdown("**IWADs**")
	var selectedID = -1
	ui.WadList = GetList(&IWADs)
	ui.WadList.OnSelected = func(id widget.ListItemID) {
		selectedID = id
	}
	wadEditToolbar := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			newItem, err := selectWadFile()
			IWADs = append(IWADs)
			settings.SetStringList(IWADsID, IWADs)
			ui.WadList.Refresh()
		}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {
			if selectedID == -1 {
				return
			}
			IWADs = slices.Delete(IWADs, selectedID, selectedID+1)
			settings.SetStringList(IWADsID, IWADs)
			ui.WadList.UnselectAll()
			selectedID = -1
		}),
	)
	wadTopContainter := container.NewBorder(nil, nil, selectWadLabel, nil, wadEditToolbar)
	wadContainer := container.NewBorder(wadTopContainter, nil, nil, nil, ui.WadList)
	return wadContainer
}

func (ui *ui) ModSelect() *fyne.Container {
	selectModLabel := widget.NewRichTextFromMarkdown("**Select mods**")
	ui.ModList = GetList(&MODs)
	modEditToolbar := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {}),
	)
	modTopContainer := container.NewBorder(nil, nil, selectModLabel, nil, modEditToolbar)
	modContainer := container.NewBorder(modTopContainer, nil, nil, nil, ui.ModList)
	return modContainer
}

func GetList(list *[]string) *widget.List {
	l := widget.NewList(
		func() int {
			return len(*list)
		},
		func() fyne.CanvasObject { return &widget.Label{} },
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			l := *list
			co.(*widget.Label).SetText(l[lii])
		},
	)
	return l
}
