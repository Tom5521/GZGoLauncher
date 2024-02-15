package gui

import (
	"fmt"
	"strings"

	t "github.com/Tom5521/GZLauncher-gtk/internal/tools"
	"github.com/Tom5521/GZLauncher-gtk/pkg/zdoom/run"
)

func (ui *ui) RunDoom() {
	if runner.IWad == "" {
		t.ErrWin(po.Get("Select a wad first!"))
		return
	}
	mods := func() []string {
		var mods []string
		for _, m := range settings.Mods {
			if m.Enabled {
				mods = append(mods, m.Path)
			}
		}
		return mods
	}()

	runner.Mods.Enabled = len(mods) > 0
	runner.Mods.List = mods
	runner.CustomArgs.Enabled = len(settings.CustomArgs) > 0
	runner.CustomArgs.Args = strings.Fields(settings.CustomArgs)

	run.GZDir = settings.ExecDir

	fmt.Println(runner.MakeCmd())
	if settings.CloseOnStart {
		ui.MainWindow.Hide()
		err := runner.Run()
		if err != nil {
			t.ErrWin(err)
		}
		ui.MainWindow.Show()
		if settings.ShowOutOnClose {

		}
		return
	}

	if settings.ShowOutOnClose {
		go func() {
			err := runner.Run()
			if err != nil {
				t.ErrWin(err)
			}
		}()
		// TODO: Set the window to display the output.
		return
	}

	err := runner.Start()
	if err != nil {
		t.ErrWin(err)
	}
}
