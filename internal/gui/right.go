package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func RightCont() *fyne.Container {
	// Launch Options.
	launchLabel := widget.NewLabel("Launch Options")
	closeOnStart := widget.NewCheck("Close launcher on start", func(b bool) {
		CloseOnStart = b
	})
	nostartup := widget.NewCheck("No startup", func(b bool) {
		Runner.NoStartup = b
	})

	// Audio options.
	audioLabel := widget.NewLabel("Audio Options")

	nosound := widget.NewCheck("No sound", func(b bool) {
		Runner.NoSound = b
	})
	nomusic := widget.NewCheck("No music", func(b bool) {
		Runner.NoMusic = b
	})
	nosfx := widget.NewCheck("No SFX", func(b bool) {
		Runner.NoSFX = b
	})

	// Gameplay options.
	gameplayLabel := widget.NewLabel("Gameplay Options")
	skillLabel := widget.NewLabel("Select skill:")
	skillList := []string{
		"Cancel",
		"I'm too young to die.",
		"Hey, not too rough.",
		"Hurt me plenty.",
		"Ultra-Violence.",
		"Nightmare!",
	}
	selectSkill := widget.NewSelect(skillList, func(s string) {})
	selectSkill.PlaceHolder = "Select a skill"
	selectSkill.OnChanged = func(s string) {
		setSkill := func(level int) {
			Runner.Skill.Level = level
		}
		if s == "Cancel" {
			Runner.Skill.Enabled = false
			selectSkill.ClearSelected()
			return
		}
		if s != "" {
			Runner.Skill.Enabled = true
		}
		switch s {
		case skillList[1]:
			setSkill(0)
		case skillList[2]:
			setSkill(1)
		case skillList[3]:
			setSkill(2)
		case skillList[4]:
			setSkill(3)
		case skillList[5]:
			setSkill(4)
		}
	}

	warpLabel := widget.NewLabel("Select warp")
	warpEntry := widget.NewEntry()
	warpEntry.OnChanged = func(s string) {
		if len(s) > 4 {
			warpEntry.SetText(s[4:])
		}
		if s == "" {
			Runner.Warp.Enabled = false
			return
		}
		Runner.Warp.Enabled = true
		Runner.Warp.Level = s
	}

	// Containers

	launchCont := container.NewAdaptiveGrid(2,
		closeOnStart,
		nostartup,
	)

	audioCont := container.NewAdaptiveGrid(2,
		nomusic,
		nosound,
		nosfx,
	)

	skillCont := container.NewBorder(nil, nil, skillLabel, nil, selectSkill)
	warpCont := container.NewBorder(nil, nil, warpLabel, nil, warpEntry)
	gameplayCont := container.NewVBox(
		gameplayLabel,
		skillCont,
		warpCont,
	)

	mainBox := container.NewVBox(
		launchLabel,
		launchCont,
		audioLabel,
		audioCont,
		gameplayCont,
	)
	return mainBox
}
