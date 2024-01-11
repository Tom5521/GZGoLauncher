package gui

import (
	"runtime"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/GZGoLauncher/internal/download"
	"github.com/Tom5521/GZGoLauncher/internal/tools"
)

type configUI struct {
	mainUI *ui
	gzdoom struct {
		Label  *widget.Label
		Entry  *widget.Entry
		Button *widget.Button
	}
	zdoom struct {
		Label  *widget.Label
		Entry  *widget.Entry
		Button *widget.Button
	}
	download struct {
		gzdoom *widget.Button
		zdoom  *widget.Button
	}
	lang struct {
		currentLabel *widget.Label
		selecter     *widget.Select
	}
}

var configuration configUI

func (ui *configUI) Container(mainUI *ui) *fyne.Container {
	ui.mainUI = mainUI
	gzdoom := ui.Gzdoom()
	zdoom := ui.Zdoom()
	download := ui.Download()
	lang := ui.Lang()

	content := container.NewVBox(
		gzdoom,
		zdoom,
		download,
		lang,
	)
	return content
}

func (ui *configUI) Gzdoom() *fyne.Container {
	gzdoom := &ui.gzdoom
	gzdoom.Label = widget.NewLabel(po.Get("GZDoom Path"))
	gzdoom.Entry = widget.NewEntry()
	gzdoom.Entry.OnChanged = func(s string) {
		settings.GZDoomDir = s
		err := settings.Write()
		if err != nil {
			ErrWin(err)
		}
		ui.mainUI.ZRunnerSelect.ClearSelected()
	}
	gzdoom.Entry.SetText(settings.GZDoomDir)
	gzdoom.Button = widget.NewButton(po.Get("Select path"), func() {
		newDir := tools.ExeFilePicker()
		if newDir == "" {
			return
		}
		settings.GZDoomDir = newDir
		err := settings.Write()
		if err != nil {
			ErrWin(err)
		}
		gzdoom.Entry.SetText(newDir)
	})
	gzdirCont := container.NewBorder(nil, nil, gzdoom.Button, nil, gzdoom.Entry)
	content := container.NewVBox(
		gzdoom.Label,
		gzdirCont,
	)
	return content
}

func (ui *configUI) Zdoom() *fyne.Container {
	zdoom := &ui.zdoom
	zdoom.Label = widget.NewLabel(po.Get("ZDoom path"))
	zdoom.Entry = widget.NewEntry()
	zdoom.Entry.SetText(settings.ZDoomDir)
	zdoom.Entry.OnChanged = func(s string) {
		settings.ZDoomDir = s
		err := settings.Write()
		if err != nil {
			ErrWin(err)
		}
		ui.mainUI.ZRunnerSelect.ClearSelected()
	}
	zdoom.Button = widget.NewButton(po.Get("Select path"), func() {
		newDir := tools.ExeFilePicker()
		if newDir == "" {
			return
		}
		settings.ZDoomDir = newDir
		err := settings.Write()
		if err != nil {
			ErrWin(err)
		}
	})
	zdirCont := container.NewBorder(nil, nil, zdoom.Button, nil, zdoom.Entry)
	content := container.NewVBox(
		zdoom.Label,
		zdirCont,
	)
	return content
}

func (ui *configUI) Download() *fyne.Container {
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
		ui.gzdoom.Entry.SetText(settings.GZDoomDir)
		mainUI.ZRunnerSelect.ClearSelected()
		down.gzdoom.SetText(po.Get("Downloaded!"))
		time.Sleep(time.Second * 2)
		down.gzdoom.SetText(po.Get("Download GZDoom"))
	}

	down.zdoom = &widget.Button{Text: po.Get("Download ZDoom")}
	down.zdoom.OnTapped = func() {
		down.zdoom.SetText(po.Get("Downloading..."))
		err := download.ZDoom()
		if err != nil {
			ErrWin(err)
			down.zdoom.SetText(po.Get("Retry"))
			return
		}
		ui.zdoom.Entry.SetText(settings.ZDoomDir)
		mainUI.ZRunnerSelect.ClearSelected()
		down.zdoom.SetText(po.Get("Downloaded!"))
		time.Sleep(time.Second * 2)
		down.zdoom.SetText(po.Get("Download ZDoom"))
	}
	if runtime.GOOS == "linux" {
		down.zdoom.Disable()
	}

	downloadCont := container.NewAdaptiveGrid(2, down.gzdoom, down.zdoom)
	return downloadCont
}

func (ui *configUI) Lang() *fyne.Container {
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
		case "Spanish":
			settings.Lang = "es"
		case "English":
			settings.Lang = "en"
		case "Portuguese":
			settings.Lang = "pt"
		default:
			return
		}
		err := settings.Write()
		if err != nil {
			ErrWin(err)
			return
		}
	})
	lang.selecter.SetSelected(currentLang())
	content := container.NewBorder(nil, nil, lang.currentLabel, nil, lang.selecter)
	return content
}
