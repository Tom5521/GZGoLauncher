package gui

import (
	"strconv"

	"fyne.io/fyne/v2"
	boxes "fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	v "github.com/Tom5521/GZGoLauncher/pkg/values"
)

func (ui *ui) RightCont() *fyne.Container {
	// Launch Options.
	launchLabel := &widget.Label{Text: po.Get("Launch Options"), Alignment: fyne.TextAlignCenter}
	closeOnStart := &widget.Check{
		Text:    po.Get("Close launcher on start"),
		Checked: settings.CloseOnStart,
		OnChanged: func(b bool) {
			settings.CloseOnStart = b
		},
	}
	nostartup := &widget.Check{
		Text:    po.Get("No startup"),
		Checked: Runner.NoStartup,
		OnChanged: func(b bool) {
			Runner.NoStartup = b
		},
	}
	showOutOnClose := &widget.Check{
		Text:    po.Get("Show output on close"),
		Checked: settings.ShowOutOnClose,
		OnChanged: func(b bool) {
			settings.ShowOutOnClose = b
		},
	}
	if v.IsWindows {
		showOutOnClose.Hide()
	}

	// Audio options.
	audioLabel := &widget.Label{Text: po.Get("Audio Options"), Alignment: fyne.TextAlignCenter}

	nosound := &widget.Check{
		Text:    po.Get("No sound"),
		Checked: Runner.NoSound,
		OnChanged: func(b bool) {
			Runner.NoSound = b
		},
	}
	nomusic := &widget.Check{
		Text:    po.Get("No music"),
		Checked: Runner.NoSound,
		OnChanged: func(b bool) {
			Runner.NoMusic = b
		},
	}
	nosfx := &widget.Check{
		Text:    po.Get("No SFX"),
		Checked: Runner.NoSFX,
		OnChanged: func(b bool) {
			Runner.NoSFX = b
		},
	}

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
	if Runner.Skill.Enabled {
		switch Runner.Skill.Level {
		case 0:
			selectSkill.Selected = skillList[1]
		case 1:
			selectSkill.Selected = skillList[2]
		case 2:
			selectSkill.Selected = skillList[3]
		case 3:
			selectSkill.Selected = skillList[4]
		case 4:
			selectSkill.Selected = skillList[5]
		}
	}

	warpLabel := widget.NewLabel(po.Get("Select warp"))
	warpEntry := &widget.Entry{Text: Runner.Warp.Level}
	warpEntry.OnChanged = func(s string) {
		Runner.Warp.Level = s
		Runner.Warp.Enabled = s != ""
	}

	fastMonsters := &widget.Check{
		Text:    po.Get("Fast monsters"),
		Checked: Runner.FastMonsters,
		OnChanged: func(b bool) {
			Runner.FastMonsters = b
		},
	}

	noMonsters := &widget.Check{
		Text:    po.Get("No monsters"),
		Checked: Runner.NoMonsters,
		OnChanged: func(b bool) {
			Runner.NoMonsters = b
		},
	}

	respawnMonsters := &widget.Check{
		Text:    po.Get("Respawn monsters"),
		Checked: Runner.RespawnMonsters,
		OnChanged: func(b bool) {
			Runner.RespawnMonsters = b
		},
	}

	// Multiplayer

	multiplayerLabel := &widget.Label{Text: po.Get("Multiplayer"), Alignment: fyne.TextAlignCenter}

	hostLabel := &widget.Label{Text: po.Get("Host:")}
	hostEntry := &widget.Entry{Text: strconv.Itoa(Runner.Multiplayer.Host)}
	hostEntry.OnChanged = func(s string) {
		h, err := strconv.Atoi(s)
		if err != nil {
			hostEntry.SetText("0")
			return
		}
		Runner.Multiplayer.Host = h
	}

	deathMatch := &widget.Check{
		Text:    po.Get("Deathmatch"),
		Checked: Runner.Multiplayer.Deathmatch,
		OnChanged: func(b bool) {
			Runner.Multiplayer.Deathmatch = b
		},
	}

	packetServer := &widget.Check{
		Text: po.Get("Packet server"),
		Checked: func() bool {
			if Runner.Multiplayer.NetMode == 1 {
				return true
			}
			return false
		}(),
		OnChanged: func(b bool) {
			if b {
				Runner.Multiplayer.NetMode = 1
				return
			}
			Runner.Multiplayer.NetMode = 0
		},
	}

	portLabel := &widget.Label{Text: po.Get("Port:")}
	portEntry := &widget.Entry{Text: strconv.Itoa(Runner.Multiplayer.Port)}
	portEntry.OnChanged = func(s string) {
		p, err := strconv.Atoi(s)
		if err != nil {
			portEntry.SetText("5029")
			return
		}
		Runner.Multiplayer.Port = p
	}

	connectToLb := &widget.Label{Text: po.Get("Connect to")}
	connectToEntry := &widget.Entry{Text: Runner.Multiplayer.IP}
	connectToEntry.OnChanged = func(s string) {
		Runner.Multiplayer.IP = s
	}

	disableMultiplayer := func() {
		hostEntry.Disable()
		connectToEntry.Disable()
		portEntry.Disable()
		deathMatch.Disable()
		packetServer.Disable()
		Runner.Multiplayer.Enabled = false
	}
	enableMultiplayer := func() {
		hostEntry.Enable()
		connectToEntry.Enable()
		portEntry.Enable()
		deathMatch.Enable()
		packetServer.Enable()
		Runner.Multiplayer.Enabled = true
	}
	enabledMultiplayer := &widget.Check{
		Text:    po.Get("Enabled"),
		Checked: Runner.Multiplayer.Enabled,
		OnChanged: func(b bool) {
			if !b {
				disableMultiplayer()
				return
			}
			enableMultiplayer()
		},
	}
	if !Runner.Multiplayer.Enabled {
		disableMultiplayer()
	}

	// Containers

	launchBox := boxes.NewAdaptiveGrid(2,
		closeOnStart,
		nostartup,
		showOutOnClose,
	)

	audioBox := boxes.NewAdaptiveGrid(2,
		nomusic,
		nosound,
		nosfx,
	)

	skillBox := boxes.NewBorder(nil, nil, skillLabel, nil, selectSkill)
	warpBox := boxes.NewBorder(nil, nil, warpLabel, nil, warpEntry)
	gameplayCheckBox := boxes.NewAdaptiveGrid(2,
		fastMonsters,
		noMonsters,
		respawnMonsters,
	)
	gameplayBox := boxes.NewVBox(
		gameplayLabel,
		skillBox,
		warpBox,
		gameplayCheckBox,
	)

	multiplayerBox := boxes.NewVBox(
		multiplayerLabel,
		enabledMultiplayer,
		boxes.NewBorder(nil, nil, hostLabel, nil, hostEntry),
		deathMatch,
		packetServer,
		boxes.NewBorder(nil, nil, portLabel, nil, portEntry),
		connectToLb,
		connectToEntry,
	)

	mainBox := boxes.NewVBox(
		launchLabel,
		launchBox,
		audioLabel,
		audioBox,
		gameplayBox,
		multiplayerBox,
	)
	return mainBox
}
