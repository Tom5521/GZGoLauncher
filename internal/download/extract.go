package download

import (
	"os"
	"os/exec"

	"github.com/artdarek/go-unzip"
	"github.com/yi-ge/unxz"
)

func Unzip(src, dest string) error {
	if _, err := os.Stat(dest); os.IsNotExist(err) {
		err = os.Mkdir(dest, os.ModePerm)
		if err != nil {
			return err
		}
	}
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
	_, err := exec.LookPath("ar")
	if err != nil {
		return ErrArIsNotInPath
	}
	cmd := exec.Command("ar", "x", file, "--output="+dir)
	return cmd.Run()
}
