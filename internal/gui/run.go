package gui

import (
	"runtime"

	"github.com/Tom5521/GZGoLauncher/pkg/gzrun"
	"github.com/Tom5521/GoNotes/pkg/messages"
	"github.com/ncruces/zenity"
)

var Runner gzrun.Pars

func (ui *ui) RunDoom() {
	if Runner.IWad == "" {
		zenity.Error("Select a wad first!")
		return
	}
	if settings.GZDir == "" {
		gzdir := GZDir()
		settings.GZDir = gzdir
		gzrun.GZDir = gzdir
		err := settings.Write()
		if err != nil {
			messages.Error(err)
		}
	}
	gzrun.GZDir = settings.GZDir
	mods := enabledPaths()
	if len(mods) > 0 {
		Runner.Mods.Enabled = true
		Runner.Mods.List = mods
	}
	if CloseOnStart {
		ui.MainWindow.Hide()
		err := Runner.Run()
		if err != nil {
			zenity.Error(err.Error())
		}
		ui.MainWindow.Show()
		return
	}
	err := Runner.Start()
	if err != nil {
		zenity.Error(err.Error())
		return
	}
}

func GZDir() string {
	if runtime.GOOS == "windows" {
		return "gzdoom.exe"
	}
	return "gzdoom"
}
