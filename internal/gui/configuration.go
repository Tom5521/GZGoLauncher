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
}

var configuration configUI

func (ui *configUI) Container(mainUI *ui) *fyne.Container {
	ui.mainUI = mainUI
	gzdoom := ui.Gzdoom()
	zdoom := ui.Zdoom()
	download := ui.Download()

	content := container.NewVBox(
		gzdoom,
		zdoom,
		download,
	)
	return content
}

func (ui *configUI) Gzdoom() *fyne.Container {
	gzdoom := &ui.gzdoom
	gzdoom.Label = widget.NewLabel("GZDoom path")
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
	gzdoom.Button = widget.NewButton("Select path", func() {
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
	zdoom.Label = widget.NewLabel("ZDoom path")
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
	zdoom.Button = widget.NewButton("Select path", func() {
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
	down.gzdoom = &widget.Button{Text: "Download GZDoom"}
	down.gzdoom.OnTapped = func() {
		down.gzdoom.SetText("Downloading...")
		err := download.GZDoom()
		if err != nil {
			ErrWin(err)
			return
		}
		ui.gzdoom.Entry.SetText(settings.GZDoomDir)
		mainUI.ZRunnerSelect.ClearSelected()
		down.gzdoom.SetText("Downloaded!")
		time.Sleep(time.Second * 2)
		down.gzdoom.SetText("DownloadGZDoom")
	}

	down.zdoom = &widget.Button{Text: "Download ZDoom"}
	down.zdoom.OnTapped = func() {
		down.zdoom.SetText("Downloading...")
		err := download.ZDoom()
		if err != nil {
			ErrWin(err)
			down.zdoom.SetText("Retry")
			return
		}
		ui.zdoom.Entry.SetText(settings.ZDoomDir)
		mainUI.ZRunnerSelect.ClearSelected()
		down.zdoom.SetText("Downloaded!")
		time.Sleep(time.Second * 2)
		down.zdoom.SetText("Download ZDoom")
	}
	if runtime.GOOS == "linux" {
		down.zdoom.Disable()
	}

	downloadCont := container.NewAdaptiveGrid(2, down.gzdoom, down.zdoom)
	return downloadCont
}
