package tools

import (
	"fmt"

	"github.com/ncruces/zenity"
)

func FilePicker(filters []string) string {
	const defaultPath string = ""
	ret, err := zenity.SelectFile(
		zenity.Filename(defaultPath),
		zenity.FileFilters{
			{"Wad files", filters, true},
		})
	if err != nil {
		fmt.Println(err)
	}
	return ret
}

func WadFilePicker() string {
	return FilePicker([]string{"*.wad"})
}

func PK3FilePicker() string {
	return FilePicker([]string{"*.pkg", "*.wad"})
}
