package filepicker

import (
	"fmt"

	v "github.com/Tom5521/GZGoLauncher/pkg/values"
	"github.com/ncruces/zenity"
)

var latestPath string

func Picker(filters []string, msg, path string) string {
	latestPath = path
	ret, err := zenity.SelectFile(
		zenity.Filename(path),
		zenity.FileFilters{
			{msg, filters, true},
		})
	if err != nil {
		fmt.Println(err)
	}
	return ret
}

func Wad() string {
	return Picker([]string{"*.wad"}, "wad files", latestPath)
}

func PK3() string {
	return Picker([]string{"*.pk3", "*.wad"}, "pk3 or wad files", latestPath)
}

func Image() string {
	return Picker([]string{"*.png", "*.gif", "*.ico", "*.jpg", "*.webp"}, "Images", latestPath)
}

func Exe() string {
	if v.IsWindows {
		return Picker([]string{"*.exe"}, ".exe files", latestPath)
	}
	return Picker([]string{}, "Executable files", latestPath)
}
