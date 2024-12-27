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

func linuxGZDoom() (path string, err error) {
	pkgpath := config.Path + "/gzdoom.tar.xz"
	url := LinuxGZDoomURL
	path = config.Path + "/gzdoom/gzdoom"
	err = Download(url, pkgpath)
	if err != nil {
		return
	}
	err = os.Chdir(config.Path)
	if err != nil {
		return
	}
	err = checkGZDoomDir()
	if err != nil {
		return
	}
	cmd := exec.Command("tar", "-xf", "gzdoom.tar.xz", "--strip-components=1", "-C", "gzdoom")
	err = cmd.Run()
	if err != nil {
		return
	}

	err = os.RemoveAll(pkgpath)
	if err != nil {
		return
	}
	return
}

func windowsGZDoom() (path string, err error) {
	pkgpath := config.Path + `\gzdoom.zip`
	url := WinGZDoomURL
	path = config.Path + `\gzdoom\gzdoom.exe`
	err = Download(url, pkgpath)
	if err != nil {
		return
	}
	err = os.Chdir(config.Path)
	if err != nil {
		return
	}
	err = Unzip(pkgpath, "gzdoom")
	if err != nil {
		return
	}

	err = os.RemoveAll(pkgpath)
	if err != nil {
		return
	}
	return
}

func macGZDoom() (path string, err error) {
	pkgpath := config.Path + "/gzdoom.zip"
	tmpDir := "tmp-gzdoom"
	path = config.Path + "/gzdoom/gzdoom"
	url := MacGZDoomURL
	err = Download(url, pkgpath)
	if err != nil {
		return
	}
	err = os.Chdir(config.Path)
	if err != nil {
		return
	}
	err = checkGZDoomDir()
	if err != nil {
		return
	}
	err = Unzip(pkgpath, tmpDir)
	if err != nil {
		return
	}
	command := fmt.Sprintf("cp -rf %s/GZDoom.app/Contents/MacOS/* gzdoom/", tmpDir)
	cmd := exec.Command("sh", "-c", command)
	err = cmd.Run()
	if err != nil {
		return
	}

	toRemove := []string{tmpDir, pkgpath}
	for _, f := range toRemove {
		err = os.RemoveAll(f)
		if err != nil {
			return
		}
	}
	return
}
