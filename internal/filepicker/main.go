package filepicker

import (
	v "github.com/Tom5521/GZGoLauncher/pkg/values"
	"github.com/Tom5521/GoNotes/pkg/messages"
	"github.com/ncruces/zenity"
)

var latestPath string

type Picker struct {
	Filters []string
	Msg     string
	Path    string
}

func (p Picker) Start() string {
	ret, msg := zenity.SelectFile(
		zenity.Filename(p.Path),
		zenity.FileFilters{
			{p.Msg, p.Filters, true},
		})
	if msg != nil {
		messages.Error(msg)
	}
	return ret
}

func Wad() string {
	p := Picker{
		Filters: []string{"*.wad"},
		Msg:     "Wad files",
		Path:    latestPath,
	}
	return p.Start()
}

func PK3() string {
	p := Picker{
		Filters: []string{"*.pk3", "*.wad"},
		Msg:     "pk3 or wad files",
		Path:    latestPath,
	}
	return p.Start()
}

func Image() string {
	p := Picker{
		Filters: []string{"*.png", "*.gif", "*.ico", "*.jpg", "*.webp"},
		Msg:     "Images",
		Path:    latestPath,
	}
	return p.Start()
}

func Exe() string {
	p := Picker{
		Msg:  "Executable files",
		Path: latestPath,
	}
	if v.IsWindows {
		p.Filters = []string{"*.exe"}
	}
	return p.Start()
}
