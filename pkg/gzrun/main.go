package gzrun

import (
	"errors"
	"os"
	"os/exec"
	"strconv"
)

var (
	ErrMissingGZDoom = errors.New("gzdoom/zdoom not found")
	ErrBadIwad       = errors.New("the iwad is incorrect or does not exist")
)

type Pars struct {
	IWad  string
	Skill struct {
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
	FastMonsters    bool
	NoMonsters      bool
	RespawnMonsters bool

	NoSFX     bool
	NoMusic   bool
	NoSound   bool
	NoStartup bool
}

var (
	GZDir string
)

func ExistsGZInPath() bool {
	_, err := exec.LookPath(GZDir)
	return err == nil
}

func (p Pars) MakeCmd() *exec.Cmd {
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
	if p.CustomArgs.Enabled {
		cmd.Args = append(cmd.Args, p.CustomArgs.Args...)
	}
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd
}

func (p Pars) check() error {
	if p.IWad == "" {
		return ErrBadIwad
	}
	if _, err := os.Stat(p.IWad); os.IsNotExist(err) {
		return ErrBadIwad
	}
	if !ExistsGZInPath() {
		return ErrMissingGZDoom
	}
	return nil
}

func (p Pars) Run() error {
	err := p.check()
	if err != nil {
		return err
	}
	cmd := p.MakeCmd()
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func (p Pars) Start() error {
	err := p.check()
	if err != nil {
		return err
	}
	cmd := p.MakeCmd()
	err = cmd.Start()
	if err != nil {
		return err
	}
	return nil
}
