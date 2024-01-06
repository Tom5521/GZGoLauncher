package gui

import (
	"runtime"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/GZGoLauncher/internal/download"
	"github.com/Tom5521/GZGoLauncher/internal/tools"
	"github.com/ncruces/zenity"
)

func (ui *ui) Configuration() *fyne.Container {
	gzlabel := widget.NewLabel("GZDoom path")
	gzdirEntry := widget.NewEntry()
	gzdirEntry.OnChanged = func(s string) {
		settings.GZDoomDir = s
		settings.Write()
	}
	gzdirEntry.SetText(settings.GZDoomDir)
	gzdirBt := widget.NewButton("Select path", func() {
		newDir := tools.ExeFilePicker()
		if newDir == "" {
			return
		}
		settings.GZDoomDir = newDir
		err := settings.Write()
		if err != nil {
			zenity.Error(err.Error())
		}
		gzdirEntry.SetText(newDir)
	})
	gzdirCont := container.NewBorder(nil, nil, gzdirBt, nil, gzdirEntry)

	zdirLabel := widget.NewLabel("ZDoom path")
	zdirEntry := widget.NewEntry()
	zdirEntry.SetText(settings.ZDoomDir)
	zdirEntry.OnChanged = func(s string) {
		settings.ZDoomDir = s
		settings.Write()
	}
	zdirBt := widget.NewButton("Select path", func() {
		newDir := tools.ExeFilePicker()
		if newDir == "" {
			return
		}
		settings.ZDoomDir = newDir
		err := settings.Write()
		if err != nil {
			zenity.Error(err.Error())
		}
	})

	downGZDoomBt := &widget.Button{Text: "Download GZDoom"}
	downGZDoomBt.OnTapped = func() {
		downGZDoomBt.SetText("Downloading...")
		err := download.GZDoom()
		if err != nil {
			zenity.Error(err.Error())
			return
		}
		gzdirEntry.SetText(settings.GZDoomDir)
		downGZDoomBt.SetText("Downloaded!")
		time.Sleep(time.Second * 2)
		downGZDoomBt.SetText("DownloadGZDoom")
	}
	downZDoomBt := &widget.Button{Text: "Download ZDoom (only for windows)"}
	downZDoomBt.OnTapped = func() {
		downZDoomBt.SetText("Downloading...")
		err := download.ZDoom()
		if err != nil {
			zenity.Error(err.Error())
			if runtime.GOOS == "linux" {
				downZDoomBt.SetText("Only for windows!")
				time.Sleep(time.Second * 2)
			}
			downZDoomBt.SetText("Download ZDoom (only for windows)")
			return
		}
		zdirEntry.SetText(settings.ZDoomDir)
		downZDoomBt.SetText("Downloaded!")
		time.Sleep(time.Second * 2)
		downZDoomBt.SetText("Download ZDoom (only for windows)")
	}
	downloadCont := container.NewAdaptiveGrid(2, downGZDoomBt, downZDoomBt)
	zdirCont := container.NewBorder(nil, nil, zdirBt, nil, zdirEntry)
	content := container.NewVBox(
		gzlabel,
		gzdirCont,
		zdirLabel,
		zdirCont,
		downloadCont,
	)
	return content
}
