package gui

import (
	"fmt"
	"runtime"
	"strconv"

	"github.com/Tom5521/GZLauncher-gtk/pkg/gtk/boxes"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

type SetSensitiver interface {
	SetSensitive(bool)
}

func (ui *ui) Right() gtk.Widgetter {
	newLabel := func(text string) *gtk.Label {
		lb := gtk.NewLabel(po.Get(text))
		lb.SetMarkup(fmt.Sprintf("<b>%s</b>", lb.Label()))
		lb.SetHAlign(gtk.AlignCenter)
		return lb
	}
	newCheck := func(text string) *gtk.CheckButton {
		check := gtk.NewCheckButtonWithLabel(po.Get(text))
		return check
	}

	launchLb := newLabel("Launch Options")

	closeOnStartCheck := newCheck("Close launcher on start")
	closeOnStartCheck.ConnectToggled(func() {
		settings.CloseOnStart = closeOnStartCheck.Active()
	})
	noStartupCheck := newCheck("No startup")
	noStartupCheck.ConnectToggled(func() {
		runner.NoStartup = noStartupCheck.Active()
	})

	showOutputOnCloseCheck := newCheck("Show output on close")
	showOutputOnCloseCheck.ConnectToggled(func() {
		settings.ShowOutOnClose = showOutputOnCloseCheck.Active()
	})

	audioLb := newLabel("Audio Options")

	noMusicCheck := newCheck("No music")
	noMusicCheck.ConnectToggled(func() {
		runner.NoMusic = noMusicCheck.Active()
	})
	noSoundCheck := newCheck("No sound")
	noSoundCheck.ConnectToggled(func() {
		runner.NoSound = noSoundCheck.Active()
	})
	noSFXCheck := newCheck("No SFX")
	noSFXCheck.ConnectToggled(func() {
		runner.NoSFX = noSFXCheck.Active()
	})

	gameplayLb := newLabel("Gameplay Options")

	skillLb := newLabel("Select skill:")
	skills := []string{
		po.Get("None."),
		po.Get("I'm too young to die."),
		po.Get("Hey, not too rough."),
		po.Get("Hurt me plenty."),
		po.Get("Ultra-Violence."),
		po.Get("Nightmare!"),
	}
	skillDropDown := gtk.NewDropDownFromStrings(skills)
	skillDropDown.SetHExpand(true)
	skillDropDown.ConnectAfter("notify::selected", func() {
		switch skillDropDown.Selected() {
		case 0:
			runner.Skill.Enabled = false
		default:
			runner.Skill.Enabled = true
			runner.Skill.Level = int(skillDropDown.Selected()) - 1
		}
	})

	warpLb := newLabel("Select warp")
	warpEntry := gtk.NewEntry()
	warpEntry.SetPlaceholderText("E1M1")
	warpEntry.ConnectChanged(func() {
		if warpEntry.Text() == "" {
			runner.Warp.Enabled = false
		}
		runner.Warp.Level = warpEntry.Text()
	})
	warpEntry.SetText(runner.Warp.Level)
	warpEntry.SetHExpand(true)

	fastMonstersCheck := newCheck("Fast monsters")
	fastMonstersCheck.ConnectToggled(func() {
		runner.FastMonsters = fastMonstersCheck.Active()
	})
	noMonstersCheck := newCheck("No monsters")
	noMonstersCheck.ConnectToggled(func() {
		runner.NoMonsters = noMonstersCheck.Active()
	})
	respawnMonstersCheck := newCheck("Respawn monsters")
	respawnMonstersCheck.ConnectToggled(func() {
		runner.RespawnMonsters = respawnMonstersCheck.Active()
	})

	multiplayerLb := newLabel("Multiplayer")

	hostLb := newLabel("Host:")
	hostEntry := gtk.NewEntry()
	hostEntry.SetHExpand(true)
	hostEntry.SetText(strconv.Itoa(runner.Multiplayer.Host))
	hostEntry.SetPlaceholderText("0")
	hostEntry.ConnectChanged(func() {
		txt := hostEntry.Text()
		for i, l := range txt {
			_, err := strconv.Atoi(string(l))
			if err != nil {
				hostEntry.SetText(txt[:i])
			}
		}
		host, err := strconv.Atoi(txt)
		if err != nil {
			hostEntry.SetText("0")
			runner.Multiplayer.Host = 0
			return
		}
		runner.Multiplayer.Host = host
	})

	deathmatchCheck := newCheck("Deathmatch")
	packetServerCheck := newCheck("Packet server")

	portLb := newLabel("Port:")
	portEntry := gtk.NewEntry()
	portEntry.SetHExpand(true)
	portEntry.SetText(strconv.Itoa(runner.Multiplayer.Port))
	portEntry.ConnectChanged(func() {
		txt := portEntry.Text()
		if txt == "" {
			portEntry.SetText("5029")
			runner.Multiplayer.Port = 5029
			return
		}
		for i, l := range txt {
			_, err := strconv.Atoi(string(l))
			if err != nil {
				portEntry.SetText(txt[:i])
			}
		}
		port, err := strconv.Atoi(txt)
		if err != nil {
			portEntry.SetText("5029")
			runner.Multiplayer.Port = 5029
			return
		}
		runner.Multiplayer.Port = port
	})

	connectToLb := newLabel("Connect to")
	connectToEntry := gtk.NewEntry()
	connectToEntry.SetHExpand(true)
	connectToEntry.SetText(runner.Multiplayer.IP)
	connectToEntry.ConnectChanged(func() {
		runner.Multiplayer.IP = connectToEntry.Text()
	})

	toggleMultiplayer := func(mode bool) {
		widgets := []SetSensitiver{
			hostEntry,
			deathmatchCheck,
			packetServerCheck,
			portEntry,
			connectToEntry,
		}
		for _, w := range widgets {
			w.SetSensitive(mode)
		}
	}
	multiplayerEnabledCheck := newCheck("Enabled")
	multiplayerEnabledCheck.ConnectToggled(func() {
		runner.Multiplayer.Enabled = multiplayerEnabledCheck.Active()
		toggleMultiplayer(multiplayerEnabledCheck.Active())
	})

	closeOnStartCheck.SetActive(settings.CloseOnStart)
	if runtime.GOOS != "windows" {
		noStartupCheck.Hide()
	}
	noStartupCheck.SetActive(runner.NoStartup)
	showOutputOnCloseCheck.SetActive(settings.ShowOutOnClose)
	noMonstersCheck.SetActive(runner.NoMusic)
	noSoundCheck.SetActive(runner.NoSound)
	noSFXCheck.SetActive(runner.NoSFX)

	switch {
	case !runner.Skill.Enabled:
		skillDropDown.SetSelected(0)
	default:
		skillDropDown.SetSelected(uint(runner.Skill.Level + 1))
	}

	fastMonstersCheck.SetActive(runner.FastMonsters)
	noMonstersCheck.SetActive(runner.NoMonsters)
	respawnMonstersCheck.SetActive(runner.RespawnMonsters)

	multiplayerEnabledCheck.SetActive(runner.Multiplayer.Enabled)
	toggleMultiplayer(runner.Multiplayer.Enabled)
	deathmatchCheck.SetActive(runner.Multiplayer.Deathmatch)

	if runner.Multiplayer.NetMode == 1 {
		packetServerCheck.SetActive(true)
	}

	launchBox := boxes.NewVbox(
		launchLb,
		boxes.NewAdaptativeGrid(2,
			closeOnStartCheck,
			noStartupCheck,
			showOutputOnCloseCheck,
		),
	)
	launchLb.SetHAlign(gtk.AlignCenter)
	audioBox := boxes.NewVbox(
		audioLb,
		boxes.NewAdaptativeGrid(2,
			noMusicCheck,
			noSoundCheck,
			noSFXCheck,
		),
	)
	audioBox.SetHAlign(gtk.AlignCenter)
	gameplayBox := boxes.NewVbox(
		gameplayLb,
		boxes.NewHbox(
			skillLb,
			skillDropDown,
		),
		boxes.NewHbox(
			warpLb,
			warpEntry,
		),
		boxes.NewAdaptativeGrid(2,
			fastMonstersCheck,
			noMonstersCheck,
			respawnMonstersCheck,
		),
	)
	multiplayerBox := boxes.NewVbox(
		multiplayerLb,
		multiplayerEnabledCheck,
		boxes.NewHbox(hostLb, hostEntry),
		boxes.NewAdaptativeGrid(2,
			deathmatchCheck,
			packetServerCheck,
		),
		boxes.NewHbox(portLb, portEntry),
		boxes.NewHbox(connectToLb, connectToEntry),
	)

	mainBox := boxes.NewScrolledVbox(
		launchBox,
		audioBox,
		gameplayBox,
		multiplayerBox,
	)
	return mainBox
}
