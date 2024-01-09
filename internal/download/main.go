package download

import (
	"errors"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"github.com/Tom5521/GZGoLauncher/internal/config"
)

var ErrOnlyForWindows = errors.New("only for windows!")

var settings = &config.Settings

var (
	IsWindows = runtime.GOOS == "windows"
	IsLinux   = runtime.GOOS == "linux"
)

const (
	WinGZDoomURL   = "https://github.com/ZDoom/gzdoom/releases/download/g4.11.3/gzdoom-4-11-3.a.-Windows-64bit.zip"
	WinZDoomURL    = "https://zdoom.org/files/zdoom/2.8/zdoom-2.8.1.zip"
	LinuxGZDoomURL = "https://github.com/ZDoom/gzdoom/releases/download/g4.11.3/gzdoom-g4.11.3-linux-portable.tar.xz"
)

func Download(url, file string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	outputFile, err := os.Create(file)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, response.Body)
	return err
}

func GZDoom() error {
	path := func() string {
		if IsWindows {
			return config.CurrentPath + "\\gzdoom.zip"
		}
		return config.CurrentPath + "/gzdoom.tar.xz"
	}()
	url := func() string {
		if IsWindows {
			return WinGZDoomURL
		}
		if IsLinux {
			return LinuxGZDoomURL
		}
		return ""
	}()
	err := Download(url, path)
	if err != nil {
		return err
	}
	err = os.Chdir(config.CurrentPath)
	if err != nil {
		return err
	}
	if _, err := os.Stat("gzdoom"); os.IsNotExist(err) {
		err = os.Mkdir("gzdoom", os.ModePerm)
		if err != nil {
			return err
		}
	}
	if IsWindows {
		err = Unzip(path, "gzdoom")
		if err != nil {
			return err
		}
	}
	if IsLinux {
		cmd := exec.Command("tar", "-xf", "gzdoom.tar.xz", "--strip-components=1", "-C", "gzdoom")
		err = cmd.Run()
		if err != nil {
			return err
		}
	}
	if IsLinux {
		settings.GZDoomDir = config.CurrentPath + "/gzdoom/gzdoom"
	}
	if IsWindows {
		settings.GZDoomDir = config.CurrentPath + "\\gzdoom\\gzdoom.exe"
	}
	err = settings.Write()
	if err != nil {
		return err
	}
	return nil
}

// Only with windows.
func ZDoom() error {
	if IsLinux {
		return ErrOnlyForWindows
	}
	path := config.CurrentPath + "\\zdoom.zip"
	url := WinZDoomURL
	err := Download(url, path)
	if err != nil {
		return err
	}
	err = os.Chdir(config.CurrentPath)
	if err != nil {
		return err
	}
	err = Unzip("zdoom.zip", "zdoom")
	if err != nil {
		return err
	}
	settings.ZDoomDir = config.CurrentPath + "\\zdoom\\zdoom.exe"
	err = settings.Write()
	if err != nil {
		return err
	}
	return nil
}
