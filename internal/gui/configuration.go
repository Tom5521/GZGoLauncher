package gui

import (
	"errors"
	"slices"
	"time"

	"fyne.io/fyne/v2"
	boxes "fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
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
		ui.themeBox(),
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

func (ui *configUI) themeBox() *fyne.Container {
	themes := []string{
		po.Get("Dark"),
		po.Get("Light"),
	}
	ui.theme.label = widget.NewLabel(po.Get("Select theme:"))
	ui.theme.mode = widget.NewSelect(themes, func(s string) {
		var t fyne.Theme
		var mode bool
		if s == themes[0] {
			t = theme.DarkTheme()
			mode = false
		} else {
			t = theme.LightTheme()
			mode = true
		}
		ui.mainUI.App.Settings().SetTheme(t)
		settings.ThemeMode = mode
	})

	// Themes
	// 1 = Light
	// 0 = Dark
	if settings.ThemeMode {
		ui.theme.mode.SetSelected(themes[1])
	} else {
		ui.theme.mode.SetSelected(themes[0])
	}

	content := boxes.NewBorder(nil, nil, ui.theme.label, nil, ui.theme.mode)
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

	addButton := &widget.Button{
		Text: po.Get("Add"),
		OnTapped: func() {
			nameEntry := widget.NewEntry()
			pathEntry := widget.NewEntry()

			dialog.NewForm(
				po.Get("Add"),
				po.Get("Add"),
				po.Get("Cancel"),
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
					if b {
						if nameExists(nameEntry.Text) {
							dialog.ShowError(
								errors.New(po.Get("Name already exists!")),
								ui.mainUI.MainWindow,
							)
							return
						}
						settings.SourcePorts = append(
							settings.SourcePorts,
							config.SourcePort{Name: nameEntry.Text, ExecutablePath: pathEntry.Text},
						)
						sp.list.Refresh()
						ui.mainUI.refreshZRunnerSelection()

					}
				},
				ui.mainUI.MainWindow,
			).Show()
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

			nameEntry := &widget.Entry{Text: settings.SourcePorts[curLii].Name}
			pathEntry := &widget.Entry{Text: settings.SourcePorts[curLii].ExecutablePath}

			dialog.NewForm(
				po.Get("Edit"),
				po.Get("Confirm"),
				po.Get("Cancel"),
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
					if b {
						settings.SourcePorts[curLii] = config.SourcePort{
							Name:           nameEntry.Text,
							ExecutablePath: pathEntry.Text,
						}
						sp.list.RefreshItem(curLii)
						settings.CurrentSourcePort = -1
						ui.mainUI.refreshZRunnerSelection()
					}
				},
				ui.mainUI.MainWindow,
			).Show()
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
	disponibleLangs := []string{
		"Spanish",
		"English",
		"Portuguese",
	}
	currentLang := func() string {
		switch settings.Lang {
		case "es":
			return "Spanish"
		case "en":
			return "English"
		case "pt":
			return "Portuguese"
		default:
			return "English"
		}
	}
	lang.currentLabel = &widget.Label{Text: po.Get("Current language:")}
	lang.selecter = widget.NewSelect(disponibleLangs, func(s string) {
		switch s {
		case currentLang():
			return
		case "Spanish":
			settings.Lang = "es"
		case "English":
			settings.Lang = "en"
		case "Portuguese":
			settings.Lang = "pt"
		default:
			return
		}
		po.Parse(locales.Parser(settings.Lang))
		dialog.ShowInformation(
			po.Get("Info"),
			po.Get("You will be able to see the changes after restarting"),
			ui.mainUI.MainWindow,
		)
	})
	lang.selecter.SetSelected(currentLang())
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
