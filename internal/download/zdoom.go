package download

import (
	"os"
	"os/exec"

	"github.com/Tom5521/GZGoLauncher/internal/config"
)

func windowsZdoom() (path string, err error) {
	pkgpath := config.Path + `\zdoom.zip`
	url := WinZDoomURL
	path = config.Path + `\zdoom\zdoom.exe`
	err = Download(url, pkgpath)
	if err != nil {
		return
	}
	err = os.Chdir(config.Path)
	if err != nil {
		return
	}
	err = Unzip("zdoom.zip", "zdoom")
	if err != nil {
		return
	}

	return
}

func linuxZdoom() (path string, err error) {
	var (
		debName = "zdoom.deb"
		pkgpath = config.Path + "/" + debName
		url     = LinuxZDoomURL
		tmpDir  = "tmp-zdoom"
	)
	path = config.Path + "/zdoom/zdoom"
	err = Download(url, pkgpath)
	if err != nil {
		return
	}
	err = os.Chdir(config.Path)
	if err != nil {
		return
	}
	if _, err = os.Stat(tmpDir); os.IsNotExist(err) {
		err = os.Mkdir(tmpDir, os.ModePerm)
		if err != nil {
			return
		}
	}
	err = ExtractDeb(debName, tmpDir)
	if err != nil {
		return
	}
	err = os.Chdir(tmpDir)
	if err != nil {
		return
	}
	if _, err = os.Stat("zdoom"); os.IsNotExist(err) {
		err = os.Mkdir("zdoom", os.ModePerm)
		if err != nil {
			return
		}
	}
	cmd := exec.Command("tar", "-xf", "data.tar.xz", "-C", "zdoom")
	err = cmd.Run()
	if err != nil {
		return
	}
	cmd = exec.Command("cp", "-rf", "zdoom/opt/zdoom", config.Path)
	err = cmd.Run()
	if err != nil {
		return
	}
	err = os.Chdir(config.Path)
	if err != nil {
		return
	}
	toRemove := []string{tmpDir, debName}
	for _, f := range toRemove {
		err = os.RemoveAll(f)
		if err != nil {
			return
		}
	}

	return
}

func macZDoom() (path string, err error) {
	return "", ErrZDoomOnMac
}
