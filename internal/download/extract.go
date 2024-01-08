package download

import (
	"os"

	"github.com/artdarek/go-unzip"
	"github.com/yi-ge/unxz"
)

func Unzip(src, dest string) error {
	uz := unzip.New(src, dest)
	return uz.Extract()
}

func ExtractTarXz(src, destDir string) error {
	if _, err := os.Stat(destDir); os.IsNotExist(err) {
		err := os.Mkdir(destDir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	u := unxz.New(src, destDir)
	return u.Extract()
}
