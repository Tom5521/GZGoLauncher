package gui

import (
	"os/exec"
	"runtime"

	t "github.com/Tom5521/GZLauncher-gtk/internal/tools"
	"github.com/Tom5521/gtk4tools/pkg/boxes"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func (ui *ui) BottomBox() gtk.Widgetter {
	const url string = "https://zdoom.org/wiki/Command_line_parameters"
	label := gtk.NewLinkButtonWithLabel(
		url,
		po.Get("Custom arguments:"),
	)
	label.ConnectClicked(func() {
		var cmd *exec.Cmd
		switch runtime.GOOS {
		case "windows":
			cmd = exec.Command("start", url)
		default:
			cmd = exec.Command("xdg-open", url)
		}
		err := cmd.Run()
		if err != nil {
			t.ErrWin(err)
		}
	})
	entry := gtk.NewEntry()
	entry.SetText(settings.CustomArgs)
	entry.ConnectChanged(func() {
		settings.CustomArgs = entry.Text()
	})
	entry.SetHExpand(true)

	runners := []string{
		po.Get("None."),
		"GZDoom",
		"ZDoom",
	}

	runnerDropDown := gtk.NewDropDownFromStrings(runners)
	runnerDropDown.ConnectAfter("notify::selected", func() {
		switch runners[runnerDropDown.Selected()] {
		case runners[0]:
			settings.ExecDir = ""
		case runners[1]:
			settings.ExecDir = settings.GZDoomDir
		case runners[2]:
			settings.ExecDir = settings.ZDoomDir
		}
	})
	switch settings.ExecDir {
	case settings.GZDoomDir:
		runnerDropDown.SetSelected(1)
	case settings.GZDoomDir:
		runnerDropDown.SetSelected(2)
	default:
		runnerDropDown.SetSelected(0)
	}

	runButton := gtk.NewButtonWithLabel(po.Get("Run"))
	runButton.ConnectClicked(ui.RunDoom)

	box := boxes.NewHbox(
		label,
		entry,
		runnerDropDown,
		runButton,
	)
	return box
}
