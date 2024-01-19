package FilePicker

import (
	"fmt"

	v "github.com/Tom5521/GZGoLauncher/pkg/values"
	"github.com/ncruces/zenity"
)

func Picker(filters []string, msg string) string {
	const defaultPath string = ""
	ret, err := zenity.SelectFile(
		zenity.Filename(defaultPath),
		zenity.FileFilters{
			{msg, filters, true},
		})
	if err != nil {
		fmt.Println(err)
	}
	return ret
}

func Wad() string {
	return Picker([]string{"*.wad"}, "wad files")
}

func PK3() string {
	return Picker([]string{"*.pk3", "*.wad"}, "pk3 or wad files")
}

func Image() string {
	return Picker([]string{"*.png", "*.gif", "*.ico", "*.jpg", "*.webp"}, "Images")
}

func Exe() string {
	if v.IsWindows {
		return Picker([]string{"*.exe"}, ".exe files")
	}
	return Picker([]string{}, "Executable files")
}
