package download

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/Tom5521/GZGoLauncher/internal/config"
)

func checkGZDoomDir() error {
	if _, err := os.Stat("gzdoom"); os.IsNotExist(err) {
		err = os.Mkdir("gzdoom", os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func linuxGZDoom() error {
	pkgpath := config.Path + "/gzdoom.tar.xz"
	url := LinuxGZDoomURL
	err := Download(url, pkgpath)
	if err != nil {
		return err
	}
	err = os.Chdir(config.Path)
	if err != nil {
		return err
	}
	err = checkGZDoomDir()
	if err != nil {
		return err
	}
	cmd := exec.Command("tar", "-xf", "gzdoom.tar.xz", "--strip-components=1", "-C", "gzdoom")
	err = cmd.Run()
	if err != nil {
		return err
	}
	settings.GZDoomDir = config.Path + "/gzdoom/gzdoom"
	err = os.RemoveAll(pkgpath)
	if err != nil {
		return err
	}
	return nil
}

func windowsGZDoom() error {
	pkgpath := config.Path + `\gzdoom.zip`
	url := WinGZDoomURL
	err := Download(url, pkgpath)
	if err != nil {
		return err
	}
	err = os.Chdir(config.Path)
	if err != nil {
		return err
	}
	err = Unzip(pkgpath, "gzdoom")
	if err != nil {
		return err
	}
	settings.GZDoomDir = config.Path + `\gzdoom\gzdoom.exe`
	err = os.RemoveAll(pkgpath)
	if err != nil {
		return err
	}
	return nil
}

func macGZDoom() error {
	pkgpath := config.Path + "/gzdoom.zip"
	tmpDir := "tmp-gzdoom"
	url := MacGZDoomURL
	err := Download(url, pkgpath)
	if err != nil {
		return err
	}
	err = os.Chdir(config.Path)
	if err != nil {
		return err
	}
	err = checkGZDoomDir()
	if err != nil {
		return err
	}
	err = Unzip(pkgpath, tmpDir)
	if err != nil {
		return err
	}
	command := fmt.Sprintf("cp -rf %s/GZDoom.app/Contents/MacOS/* gzdoom/", tmpDir)
	cmd := exec.Command("sh", "-c", command)
	err = cmd.Run()
	if err != nil {
		return err
	}
	settings.GZDoomDir = config.Path + "/gzdoom/gzdoom"

	toRemove := []string{tmpDir, pkgpath}
	for _, f := range toRemove {
		err = os.RemoveAll(f)
		if err != nil {
			return err
		}
	}
	return nil
}
