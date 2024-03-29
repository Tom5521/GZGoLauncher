package gui

import (
	"time"

	"fyne.io/fyne/v2"
	boxes "fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
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
	credits *widget.Button
}

var configuration configUI

func (ui *configUI) MainBox(mainUI *ui) *fyne.Container {
	ui.mainUI = mainUI
	th := ui.themeBox()
	gzdoom := ui.gzBox()
	zdoom := ui.zBox()
	download := ui.downloadBox()
	lang := ui.langBox()
	credits := ui.creditsBox()

	content := boxes.NewVBox(
		th,
		gzdoom,
		zdoom,
		download,
		lang,
		credits,
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

func (ui *configUI) gzBox() *fyne.Container {
	gzdoom := &ui.gzdoom
	gzdoom.Label = &widget.Label{
		Text:      po.Get("GZDoom Path"),
		Alignment: fyne.TextAlignCenter,
		TextStyle: fyne.TextStyle{
			Bold: true,
		},
	}
	gzdoom.Entry = &widget.Entry{Text: settings.GZDoomDir}
	gzdoom.Entry.OnChanged = func(s string) {
		settings.GZDoomDir = s
		ui.mainUI.ZRunnerSelect.ClearSelected()
	}
	gzdoom.Button = widget.NewButton(po.Get("Select path"), func() {
		newDir := filepicker.Executable.Start()
		if newDir == "" {
			return
		}
		settings.GZDoomDir = newDir
		gzdoom.Entry.SetText(newDir)
	})
	gzdirCont := boxes.NewBorder(nil, nil, gzdoom.Button, nil, gzdoom.Entry)
	content := boxes.NewVBox(
		gzdoom.Label,
		gzdirCont,
	)
	return content
}

func (ui *configUI) zBox() *fyne.Container {
	zdoom := &ui.zdoom
	zdoom.Label = &widget.Label{
		Text:      po.Get("ZDoom path"),
		Alignment: fyne.TextAlignCenter,
		TextStyle: fyne.TextStyle{
			Bold: true,
		},
	}
	zdoom.Entry = &widget.Entry{Text: settings.ZDoomDir}
	zdoom.Entry.OnChanged = func(s string) {
		settings.ZDoomDir = s
		ui.mainUI.ZRunnerSelect.ClearSelected()
	}
	zdoom.Button = widget.NewButton(po.Get("Select path"), func() {
		newDir := filepicker.Executable.Start()
		if newDir == "" {
			return
		}
		settings.ZDoomDir = newDir
	})
	zdirCont := boxes.NewBorder(nil, nil, zdoom.Button, nil, zdoom.Entry)
	content := boxes.NewVBox(
		zdoom.Label,
		zdirCont,
	)
	return content
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
		if err == download.ErrZDoomOnMac {
			ErrWin(po.Get("The automatic download of zdoom on mac is not supported, I recommend you to download the binary by yourself and select its path in the corresponding zdoom path."))
			down.zdoom.SetText(po.Get("Download ZDoom"))
			return
		}
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
