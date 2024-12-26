package gui

import (
	"strings"
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
	mods := enabledPaths()

	Runner.Mods.Enabled = len(mods) > 0
	Runner.Mods.List = mods
	Runner.CustomArgs.Enabled = len(settings.CustomArgs) > 0
	Runner.CustomArgs.Args = strings.Fields(settings.CustomArgs)

	showWindow := make(chan struct{}, 1)
	defer close(showWindow)

	go func() {
		if settings.CloseOnStart {
			ui.MainWindow.Hide()
		}
		err := Runner.Run()

		if settings.ShowOutOnClose || err != nil {
			ui.ShowOutWin(Runner.Output.String())
		}

		if settings.CloseOnStart {
			showWindow <- struct{}{}
		}
	}()

	if settings.CloseOnStart {
		<-showWindow
		ui.MainWindow.Show()
	}
}
