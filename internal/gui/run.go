package gui

import (
	"fmt"
	"strings"

	"github.com/Tom5521/GZGoLauncher/pkg/zdoom/run"
	"github.com/Tom5521/GoNotes/pkg/messages"
)

func (ui *ui) RunDoom() {
	if Runner.IWad == "" {
		ErrWin(po.Get("Select a wad first!"))
		return
	}
	if ui.ZRunnerSelect.Selected == "" {
		ErrWin(po.Get("Select a runner first!"))
		return
	}
	run.GZDir = settings.ExecDir
	mods := enabledPaths()

	Runner.Mods.Enabled = len(mods) > 0
	Runner.Mods.List = mods
	Runner.CustomArgs.Enabled = len(settings.CustomArgs) > 0
	Runner.CustomArgs.Args = strings.Fields(settings.CustomArgs)

	if settings.CloseOnStart {
		ui.MainWindow.Hide()
		err := Runner.Run()
		if err != nil {
			ErrWin(err)
		}
		ui.MainWindow.Show()
		return
	}
	run.GZDir = settings.ExecDir
	fmt.Println(Runner.MakeCmd())
	// return // NOTE: Uncomment this to view the cmd without starting *zdoom.
	if settings.ShowOutOnClose {
		out, err := Runner.Out()
		if err != nil {
			messages.Error(err)
		}
		ui.ShowOutWin(out)
		return
	}
	err := Runner.Start()
	if err != nil {
		ErrWin(err)
		return
	}
}
