package gui

import (
	"runtime"

	"github.com/Tom5521/GZGoLauncher/pkg/gzrun"
	"github.com/Tom5521/GoNotes/pkg/messages"
	"github.com/ncruces/zenity"
)

func RunDoom() {
	if SelectedWad == "" {
		zenity.Error("Select a wad first!")
		return
	}
	newRunner := gzrun.Pars{}
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
	newRunner.IWad = SelectedWad
	mods := enabledPaths()
	if len(mods) > 0 {
		newRunner.Mods.Enabled = true
		newRunner.Mods.List = mods
	}
	err := newRunner.Start()
	if err != nil {
		zenity.Error(err.Error())
	}
}

func GZDir() string {
	if runtime.GOOS == "windows" {
		return "gzdoom.exe"
	}
	return "gzdoom"
}
