package gzrun

import (
	"errors"
	"os"
	"os/exec"
	"strconv"

	"github.com/Tom5521/GZGoLauncher/internal/config"
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
	NoSFX     bool
	NoMusic   bool
	NoSound   bool
	NoStartup bool
}

var (
	GZDir   string
	LogFile = config.CurrentPath + "/out.log"
)

func ExistsGZInPath() bool {
	_, err := exec.LookPath(GZDir)
	return err == nil
}

func (p Pars) FormatCmd() *exec.Cmd {
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
	cmd := p.FormatCmd()

	var out []byte
	out, err = cmd.CombinedOutput()
	errWrite := writeOut(out)
	if err != nil {
		return err
	}
	if errWrite != nil {
		return errWrite
	}

	return nil
}

func (p Pars) Start() error {
	err := p.check()
	if err != nil {
		return err
	}
	cmd := p.FormatCmd()
	go func() {
		var out []byte
		out, err = cmd.CombinedOutput()
		errWrite := writeOut(out)
		if err != nil {
			return
		}
		if errWrite != nil {
			err = errWrite
			return
		}
	}()
	return err
}

func writeOut(b []byte) error {
	return os.WriteFile(LogFile, b, os.ModePerm)
}
