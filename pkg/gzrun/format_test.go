package gzrun_test

import (
	"testing"

	"github.com/Tom5521/GZGoLauncher/pkg/gzrun"
)

func TestFormat(t *testing.T) {
	gzrun.GZDir = "gzdoom"
	newRun := gzrun.Pars{
		IWad: "/run/media/tom/Archivos/Juegos PC/Doom/DOOM.WAD",
	}
	newRun.Mods.Enabled = true
	newRun.Mods.List = []string{
		"/run/media/tom/Archivos/Juegos PC/Doom/brutalv21.pk3",
		"/run/media/tom/Archivos/Juegos PC/Doom/mapsofchaos.wad",
	}
	err := newRun.Run()
	if err != nil {
		t.Fail()
	}
}
