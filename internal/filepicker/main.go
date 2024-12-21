package filepicker

import (
	"errors"
	"path/filepath"

	t "github.com/Tom5521/GZGoLauncher/internal/tools"
	v "github.com/Tom5521/GZGoLauncher/pkg/values"
	"github.com/ncruces/zenity"
)

var latestPath string

var (
	Wad = Picker{
		Filters: []string{"*.wad"},
		Msg:     "Wad files",
		Path:    latestPath,
	}
	Pk3 = Picker{
		Filters: []string{"*.pk3", "*.wad"},
		Msg:     "pk3 or wad files",
		Path:    latestPath,
	}
	Image = Picker{
		Filters: []string{"*.png", "*.gif", "*.ico", "*.jpg", "*.webp"},
		Msg:     "Images",
		Path:    latestPath,
	}
	Executable = func() Picker {
		p := Picker{
			Msg:  "Executable files",
			Path: latestPath,
		}
		if v.IsWindows {
			p.Filters = []string{"*.exe"}
		} else {
			p.Filters = []string{"*"}
		}
		return p
	}()
)

type Picker struct {
	Filters []string
	Msg     string
	Path    string
}

func (p Picker) Start() string {
	f, err := zenity.SelectFile(
		zenity.Filename(p.Path),
		zenity.FileFilters{
			zenity.FileFilter{
				Name:     p.Msg,
				Patterns: p.Filters,
				CaseFold: true,
			},
		})
	if err != nil && !errors.Is(err, zenity.ErrCanceled) {
		t.ErrWin(err)
	}
	latestPath = filepath.Dir(f)
	return f
}

func (p Picker) MultiStart() []string {
	f, err := zenity.SelectFileMultiple(
		zenity.Filename(p.Path),
		zenity.FileFilters{
			zenity.FileFilter{
				Name:     p.Msg,
				Patterns: p.Filters,
				CaseFold: true,
			},
		})
	if err != nil && !errors.Is(err, zenity.ErrCanceled) {
		t.ErrWin(err)
	}
	return f
}
