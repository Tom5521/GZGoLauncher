package gui

import (
	"errors"
	"slices"
	"time"

	"fyne.io/fyne/v2"
	boxes "fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/GZGoLauncher/internal/config"
	"github.com/Tom5521/GZGoLauncher/internal/download"
	"github.com/Tom5521/GZGoLauncher/internal/filepicker"
	"github.com/Tom5521/GZGoLauncher/internal/gui/credits"
	"github.com/Tom5521/GZGoLauncher/locales"
)

type configUI struct {
	mainUI *ui
	theme  struct {
		label *widget.Label
		mode  *widget.Select
	}
	sourcePorts struct {
		list *widget.List
	}
	download struct {
		gzdoom *widget.Button
		zdoom  *widget.Button
	}
	lang struct {
		currentLabel *widget.Label
		selecter     *widget.Select
	}
	credits *widget.Button
}

func cast[T any](a any) T {
	return a.(T)
}

var configuration configUI

func (ui *configUI) MainBox(mainUI *ui) *fyne.Container {
	ui.mainUI = mainUI

	var objs []fyne.CanvasObject
	objs = append(
		objs,
		ui.creditsBox(),
		ui.langBox(),
		ui.downloadBox(),
	)

	content := boxes.NewVBox()
	for i, obj := range objs {
		content.Add(obj)
		if i != len(objs)-1 {
			content.Add(widget.NewSeparator())
		}
	}

	content = boxes.NewBorder(
		boxes.NewVBox(content, widget.NewSeparator()),
		nil,
		nil,
		nil,
		ui.sourcePortsBox(),
	)

	return content
}

func (ui *configUI) sourcePortsBox() *fyne.Container {
	sp := &ui.sourcePorts

	var curLii widget.ListItemID = -1

	sp.list = widget.NewList(
		func() int { return len(settings.SourcePorts) },
		func() fyne.CanvasObject {
			return boxes.NewBorder(nil, nil, &widget.Label{
				Alignment: fyne.TextAlignLeading,
				TextStyle: fyne.TextStyle{
					Bold: true,
				},
			},
				nil,
				boxes.NewScroll(&widget.Label{Alignment: fyne.TextAlignTrailing}),
			)
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			container := co.(*fyne.Container)
			pathLb := cast[*widget.Label](cast[*boxes.Scroll](container.Objects[0]).Content)
			nameLb := cast[*widget.Label](container.Objects[1])

			nameLb.SetText(settings.SourcePorts[lii].Name)
			pathLb.SetText(settings.SourcePorts[lii].ExecutablePath)
		},
	)

	sp.list.OnSelected = func(id widget.ListItemID) {
		curLii = id
	}
	sp.list.OnUnselected = func(id widget.ListItemID) {
		curLii = -1
	}

	nameExists := func(name string) bool {
		return slices.ContainsFunc(settings.SourcePorts, func(i config.SourcePort) bool {
			return i.Name == name
		})
	}

	displayForm := func(
		title, confirm, dimiss string,
		nameText, pathText string,
		callback func(b bool, name, path *widget.Entry),
	) {
		nameEntry := widget.NewEntry()
		nameEntry.SetText(nameText)
		pathEntry := widget.NewEntry()
		pathEntry.SetText(pathText)

		dialog.ShowForm(
			po.Get(title),
			po.Get(confirm),
			po.Get(dimiss),
			[]*widget.FormItem{
				widget.NewFormItem(po.Get("Name:"), nameEntry),
				widget.NewFormItem(
					po.Get("Path"),
					boxes.NewVBox(
						pathEntry,
						widget.NewButton(po.Get("Select path"), func() {
							pathEntry.SetText(filepicker.Executable.Start())
						}),
					),
				),
			},
			func(b bool) {
				callback(b, nameEntry, pathEntry)
			},
			ui.mainUI.MainWindow,
		)
	}

	addButton := &widget.Button{
		Text: po.Get("Add"),
		OnTapped: func() {
			displayForm("Add", "Add", "Cancel", "", "",
				func(b bool, name, path *widget.Entry) {
					if b {
						if nameExists(name.Text) {
							dialog.ShowError(
								errors.New(po.Get("Name already exists!")),
								ui.mainUI.MainWindow,
							)
							return
						}
						settings.SourcePorts = append(
							settings.SourcePorts,
							config.SourcePort{Name: name.Text, ExecutablePath: path.Text},
						)
						sp.list.Refresh()
						ui.mainUI.refreshZRunnerSelection()

					}
				},
			)
		},
	}
	removeButton := &widget.Button{
		Text: po.Get("Remove"),
		OnTapped: func() {
			if curLii == -1 {
				return
			}
			settings.SourcePorts = slices.Delete(settings.SourcePorts, curLii, curLii+1)
			settings.CurrentSourcePort = -1
			ui.mainUI.refreshZRunnerSelection()
			sp.list.UnselectAll()
		},
	}
	editButton := &widget.Button{
		Text: po.Get("Edit"),
		OnTapped: func() {
			if curLii == -1 {
				return
			}

			displayForm("Edit", "Confirm", "Cancel",
				settings.SourcePorts[curLii].Name,
				settings.SourcePorts[curLii].ExecutablePath,
				func(b bool, name, path *widget.Entry) {
					if b {
						settings.SourcePorts[curLii] = config.SourcePort{
							Name:           name.Text,
							ExecutablePath: path.Text,
						}
						sp.list.RefreshItem(curLii)
						settings.CurrentSourcePort = -1
						ui.mainUI.refreshZRunnerSelection()
					}
				},
			)
		},
	}

	container := boxes.NewBorder(
		boxes.NewVBox(
			&widget.Label{
				Text:      po.Get("Source Ports"),
				Alignment: fyne.TextAlignCenter,
				TextStyle: fyne.TextStyle{
					Bold: true,
				},
			},
			boxes.NewAdaptiveGrid(2,
				&widget.Label{
					Text:      po.Get("Name"),
					Alignment: fyne.TextAlignLeading,
					TextStyle: fyne.TextStyle{
						Bold: true,
					},
				},
				&widget.Label{
					Text:      po.Get("Path"),
					Alignment: fyne.TextAlignCenter,
					TextStyle: fyne.TextStyle{
						Bold: true,
					},
				},
			),
		),
		boxes.NewAdaptiveGrid(3, addButton, removeButton, editButton),
		nil,
		nil,
		sp.list,
	)

	return container
}

func (ui *configUI) downloadBox() *fyne.Container {
	mainUI := ui.mainUI
	down := &ui.download
	down.gzdoom = &widget.Button{Text: po.Get("Download GZDoom")}
	down.gzdoom.OnTapped = func() {
		down.gzdoom.SetText(po.Get("Downloading..."))
		err := download.GZDoom()
		if err != nil {
			ErrWin(err)
			return
		}
		ui.sourcePorts.list.Refresh()
		mainUI.refreshZRunnerSelection()
		down.gzdoom.SetText(po.Get("Downloaded!"))
		time.Sleep(time.Second * 2)
		down.gzdoom.SetText(po.Get("Download GZDoom"))
	}

	down.zdoom = &widget.Button{Text: po.Get("Download ZDoom")}
	down.zdoom.OnTapped = func() {
		down.zdoom.SetText(po.Get("Downloading..."))
		err := download.ZDoom()
		if err == download.ErrZDoomOnMac {
			ErrWin(
				po.Get(
					"The automatic download of zdoom on mac is not supported, I recommend you to download the binary by yourself and select its path in the corresponding zdoom path.",
				),
			)
			down.zdoom.SetText(po.Get("Download ZDoom"))
			return
		}
		if err != nil {
			ErrWin(err)
			down.zdoom.SetText(po.Get("Retry"))
			return
		}
		ui.sourcePorts.list.Refresh()
		mainUI.refreshZRunnerSelection()
		down.zdoom.SetText(po.Get("Downloaded!"))
		time.Sleep(time.Second * 2)
		down.zdoom.SetText(po.Get("Download ZDoom"))
	}

	downloadCont := boxes.NewAdaptiveGrid(2, down.gzdoom, down.zdoom)
	return downloadCont
}

func (ui *configUI) langBox() *fyne.Container {
	lang := &ui.lang
	lang.currentLabel = &widget.Label{Text: po.Get("Current language:")}
	lang.selecter = widget.NewSelect(locales.LocaleNames(), func(s string) {
		short := locales.LocaleShort(s)
		if settings.Lang == short {
			return
		}
		if !locales.ShortLocaleExists(short) {
			return
		}

		settings.Lang = short
		po.Parse(locales.Parser(settings.Lang))
		dialog.ShowInformation(
			po.Get("Info"),
			po.Get("You will be able to see the changes after restarting"),
			ui.mainUI.MainWindow,
		)
	})
	lang.selecter.SetSelected(locales.LocaleLong(settings.Lang))
	content := boxes.NewBorder(nil, nil, lang.currentLabel, nil, lang.selecter)
	return content
}

func (ui *configUI) creditsBox() *fyne.Container {
	ui.credits = widget.NewButton(po.Get("Show credits"), func() {
		credits.CreditsWindow(ui.mainUI.App, fyne.NewSize(400, 800)).Show()
	})
	content := boxes.NewVBox(
		ui.credits,
	)
	return content
}
