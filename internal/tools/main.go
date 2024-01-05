package tools

import (
	"fmt"

	"github.com/ncruces/zenity"
)

func FilePicker(filters []string, msg string) string {
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

func WadFilePicker() string {
	return FilePicker([]string{"*.wad"}, "wad files")
}

func PK3FilePicker() string {
	return FilePicker([]string{"*.pk3", "*.wad"}, "pk3 or wad files")
}
