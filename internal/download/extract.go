package download

import (
	"os"
	"os/exec"

	"github.com/artdarek/go-unzip"
	"github.com/yi-ge/unxz"
)

func Unzip(src, dest string) error {
	uz := unzip.New(src, dest)
	return uz.Extract()
}

func ExtractTarXz(src, destDir string) error {
	if _, err := os.Stat(destDir); os.IsNotExist(err) {
		err = os.Mkdir(destDir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	u := unxz.New(src, destDir)
	return u.Extract()
}

func ExtractDeb(file, dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	cmd := exec.Command("ar", "x", file, "--output="+dir)
	return cmd.Run()
}
