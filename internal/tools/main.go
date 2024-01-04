package tools

import "github.com/ncruces/zenity"

func SelectWadFile() (string, error) {
	f, err := zenity.SelectFile()
}
