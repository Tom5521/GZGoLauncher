package download

import (
	"os"
	"os/exec"

	"github.com/Tom5521/GZGoLauncher/internal/config"
)

func windowsZdoom() error {
	path := config.Path + `\zdoom.zip`
	url := WinZDoomURL
	err := Download(url, path)
	if err != nil {
		return err
	}
	err = os.Chdir(config.Path)
	if err != nil {
		return err
	}
	err = Unzip("zdoom.zip", "zdoom")
	if err != nil {
		return err
	}
	// settings.ZDoomDir = config.Path + `\zdoom\zdoom.exe`
	return nil
}

func linuxZdoom() error {
	var (
		debName = "zdoom.deb"
		path    = config.Path + "/" + debName
		url     = LinuxZDoomURL
		tmpDir  = "tmp-zdoom"
	)
	err := Download(url, path)
	if err != nil {
		return err
	}
	err = os.Chdir(config.Path)
	if err != nil {
		return err
	}
	if _, err = os.Stat(tmpDir); os.IsNotExist(err) {
		err = os.Mkdir(tmpDir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	err = ExtractDeb(debName, tmpDir)
	if err != nil {
		return err
	}
	err = os.Chdir(tmpDir)
	if err != nil {
		return err
	}
	if _, err = os.Stat("zdoom"); os.IsNotExist(err) {
		err = os.Mkdir("zdoom", os.ModePerm)
		if err != nil {
			return err
		}
	}
	cmd := exec.Command("tar", "-xf", "data.tar.xz", "-C", "zdoom")
	err = cmd.Run()
	if err != nil {
		return err
	}
	cmd = exec.Command("cp", "-rf", "zdoom/opt/zdoom", config.Path)
	err = cmd.Run()
	if err != nil {
		return err
	}
	err = os.Chdir(config.Path)
	if err != nil {
		return err
	}
	toRemove := []string{tmpDir, debName}
	for _, f := range toRemove {
		err = os.RemoveAll(f)
		if err != nil {
			return err
		}
	}
	// settings.ZDoomDir = config.Path + "/zdoom/zdoom"
	return nil
}

func macZDoom() error {
	return ErrZDoomOnMac
}
