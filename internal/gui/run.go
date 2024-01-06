package gui

import (
	"runtime"

	"github.com/Tom5521/GZGoLauncher/pkg/gzrun"
)

var Runner gzrun.Pars

func (ui *ui) RunDoom() {
	if Runner.IWad == "" {
		ErrWin("Select a wad first!")
		return
	}
	gzrun.GZDir = settings.GZDir
	mods := enabledPaths()

	Runner.Mods.Enabled = len(mods) > 0

	Runner.Mods.List = mods

	if CloseOnStart {
		ui.MainWindow.Hide()
		err := Runner.Run()
		if err != nil {
			ErrWin(err)
		}
		ui.MainWindow.Show()
		return
	}
	gzrun.GZDir = settings.GZDir
	err := Runner.Run()
	if err != nil {
		ErrWin(err)
		return
	}
}

func GZDir() string {
	if runtime.GOOS == "windows" {
		return "gzdoom.exe"
	}
	return "gzdoom"
}
