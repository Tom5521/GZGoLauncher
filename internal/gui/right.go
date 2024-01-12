package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func RightCont() *fyne.Container {
	// Launch Options.
	launchLabel := &widget.Label{Text: po.Get("Launch Options"), Alignment: fyne.TextAlignCenter}
	closeOnStart := widget.NewCheck(po.Get("Close launcher on start"), func(b bool) {
		CloseOnStart = b
	})
	nostartup := widget.NewCheck(po.Get("No startup"), func(b bool) {
		Runner.NoStartup = b
	})

	// Audio options.
	audioLabel := widget.NewLabel(po.Get("Audio Options"))

	nosound := widget.NewCheck(po.Get("No sound"), func(b bool) {
		Runner.NoSound = b
	})
	nomusic := widget.NewCheck(po.Get("No music"), func(b bool) {
		Runner.NoMusic = b
	})
	nosfx := widget.NewCheck(po.Get("No SFX"), func(b bool) {
		Runner.NoSFX = b
	})

	// Gameplay options.
	gameplayLabel := &widget.Label{Text: po.Get("Gameplay Options"), Alignment: fyne.TextAlignCenter}
	skillLabel := widget.NewLabel(po.Get("Select skill:"))
	skillList := []string{
		po.Get("Cancel"),
		po.Get("I'm too young to die."),
		po.Get("Hey, not too rough."),
		po.Get("Hurt me plenty."),
		po.Get("Ultra-Violence."),
		po.Get("Nightmare!"),
	}
	selectSkill := widget.NewSelect(skillList, func(s string) {})
	selectSkill.PlaceHolder = po.Get("Select a skill")
	selectSkill.OnChanged = func(s string) {
		setSkill := func(level int) {
			Runner.Skill.Level = level
		}
		if s == po.Get("Cancel") {
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

	warpLabel := widget.NewLabel(po.Get("Select warp"))
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

	fastMonsters := widget.NewCheck(po.Get("Fast monsters"), func(b bool) {
		Runner.FastMonsters = b
	})

	noMonsters := widget.NewCheck(po.Get("No monsters"), func(b bool) {
		Runner.NoMonsters = b
	})

	respawnMonsters := widget.NewCheck(po.Get("Respawn monsters"), func(b bool) {
		Runner.RespawnMonsters = b
	})

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
	gameplayChecks := container.NewAdaptiveGrid(2,
		fastMonsters,
		noMonsters,
		respawnMonsters,
	)
	gameplayCont := container.NewVBox(
		gameplayLabel,
		skillCont,
		warpCont,
		gameplayChecks,
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
