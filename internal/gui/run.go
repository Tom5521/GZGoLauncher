package gui

import (
	"fmt"
	"strings"

	"github.com/Tom5521/GZGoLauncher/pkg/gzrun"
)

var Runner gzrun.Pars

func (ui *ui) RunDoom() {
	if Runner.IWad == "" {
		ErrWin("Select a wad first!")
		return
	}
	if ui.ZRunnerSelect.Selected == "" {
		ErrWin("Select a runner first!")
		return
	}
	gzrun.GZDir = settings.GZDir
	mods := enabledPaths()

	Runner.Mods.Enabled = len(mods) > 0
	Runner.Mods.List = mods
	Runner.CustomArgs.Enabled = len(settings.CustomArgs) > 0
	Runner.CustomArgs.Args = strings.Fields(settings.CustomArgs)

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
	fmt.Println(Runner.MakeCmd())
	err := Runner.Start()
	if err != nil {
		ErrWin(err)
		return
	}
}
