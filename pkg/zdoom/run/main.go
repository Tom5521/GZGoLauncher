package run

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/Tom5521/GZGoLauncher/locales"
)

var (
	po = locales.Current

	ErrMissingGZDoom = errors.New(po.Get("gzdoom/zdoom not found"))
	ErrBadIwad       = errors.New(po.Get("the iwad is incorrect or does not exist"))
)

type Pars struct {
	Output bytes.Buffer `json:"-"`

	Error  error  `json:"-"`
	ErrOut string `json:"-"`
	IWad   string
	Skill  struct {
		Enabled bool
		Level   int
	}
	Warp struct {
		Enabled bool
		Level   string
	}
	Mods struct {
		Enabled bool
		List    []string
	}
	CustomArgs struct {
		Enabled bool
		Args    []string
	}

	Multiplayer struct {
		Enabled    bool
		NetMode    uint // 0: Peer to peer | 1: Packet server
		Host       int
		Port       int // Default:5029
		Deathmatch bool
		IP         string
	}

	FastMonsters    bool
	NoMonsters      bool
	RespawnMonsters bool

	NoSFX     bool
	NoMusic   bool
	NoSound   bool
	NoStartup bool
}

var GZDir string

func ExistsGZInPath() bool {
	_, err := exec.LookPath(GZDir)
	return err == nil
}

func (p *Pars) MakeCmd() *exec.Cmd {
	cmd := exec.Command(GZDir, "-iwad", p.IWad)
	if p.Mods.Enabled {
		cmd.Args = append(cmd.Args, "-file")
		cmd.Args = append(cmd.Args, p.Mods.List...)
	}
	if p.Skill.Enabled {
		cmd.Args = append(cmd.Args, "-skill", strconv.Itoa(p.Skill.Level))
	}
	if p.Warp.Enabled {
		cmd.Args = append(cmd.Args, "-warp", p.Warp.Level)
	}
	if p.NoMusic {
		cmd.Args = append(cmd.Args, "-nomusic")
	}
	if p.NoSound {
		cmd.Args = append(cmd.Args, "-nosound")
	}
	if p.NoSFX {
		cmd.Args = append(cmd.Args, "-nosfx")
	}
	if p.NoStartup {
		cmd.Args = append(cmd.Args, "-nostartup")
	}
	if p.RespawnMonsters {
		cmd.Args = append(cmd.Args, "-respawn")
	}
	if p.NoMonsters {
		cmd.Args = append(cmd.Args, "-nomonsters")
	}
	if p.FastMonsters {
		cmd.Args = append(cmd.Args, "-fast")
	}
	if p.Multiplayer.Enabled {
		cmd.Args = append(cmd.Args, "join", fmt.Sprintf("%s:%v", p.Multiplayer.IP, "5029"))
		if p.Multiplayer.Port != 0 {
			cmd.Args = append(cmd.Args, "-port", strconv.Itoa(p.Multiplayer.Port))
		}
		if p.Multiplayer.Host != 0 {
			cmd.Args = append(cmd.Args, "-host", strconv.Itoa(p.Multiplayer.Host))
		}
		if p.Multiplayer.Deathmatch {
			cmd.Args = append(cmd.Args, "-deathmatch")
		}
		cmd.Args = append(cmd.Args, "-netmode", strconv.Itoa(int(p.Multiplayer.NetMode)))
	}
	if p.CustomArgs.Enabled {
		cmd.Args = append(cmd.Args, p.CustomArgs.Args...)
	}

	cmd.Stderr = &p.Output
	cmd.Stdout = &p.Output

	return cmd
}

func (p *Pars) check() error {
	if p.IWad == "" {
		p.Error = ErrBadIwad
		return ErrBadIwad
	}
	if _, err := os.Stat(p.IWad); os.IsNotExist(err) {
		p.Error = ErrBadIwad
		return ErrBadIwad
	}
	if !ExistsGZInPath() {
		p.Error = ErrMissingGZDoom
		return ErrMissingGZDoom
	}
	return nil
}

func (p *Pars) Run() error {
	err := p.check()
	if err != nil {
		return err
	}
	cmd := p.MakeCmd()
	err = cmd.Run()
	p.Error = err

	return err
}

func (p *Pars) Start() error {
	err := p.check()
	if err != nil {
		return err
	}
	cmd := p.MakeCmd()
	err = cmd.Start()
	p.Error = err

	return err
}
