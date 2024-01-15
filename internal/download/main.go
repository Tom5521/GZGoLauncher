package download

import (
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/Tom5521/GZGoLauncher/internal/config"
	v "github.com/Tom5521/GZGoLauncher/pkg/values"
)

var (
	ErrIncompatiblePlattaform = errors.New("incompatible plattaform")
	ErrOnlyForWindows         = errors.New("only for windows")
	ErrZDoomOnMac             = errors.New("zdoom is not available on mac")
)

var settings = &config.Settings

const (
	WinGZDoomURL = "https://github.com/ZDoom/gzdoom/releases/download/g4.11.3/gzdoom-4-11-3.a.-Windows-64bit.zip"
	WinZDoomURL  = "https://zdoom.org/files/zdoom/2.8/zdoom-2.8.1.zip"

	LinuxGZDoomURL = "https://github.com/ZDoom/gzdoom/releases/download/g4.11.3/gzdoom-g4.11.3-linux-portable.tar.xz"
	LinuxZDoomURL  = "https://zdoom.org/files/zdoom/2.8/zdoom_2.8.1_amd64.deb"

	MacGZDoomURL = "https://github.com/ZDoom/gzdoom/releases/download/g4.11.3/gzdoom-4-11-3-macOS.zip"
	MacZDoomURL  = "https://zdoom.org/files/zdoom/2.8/zdoom-2.8.1.dmg"
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
	if v.IsMac {
		return macGZDoom()
	}
	if v.IsLinux {
		return linuxGZDoom()
	}
	if v.IsWindows {
		return windowsGZDoom()
	}
	return ErrIncompatiblePlattaform
}

// Only with windows.
func ZDoom() error {
	if !v.IsMac {
		return macZDoom()
	}
	if v.IsLinux {
		return linuxZdoom()
	}
	if v.IsWindows {
		return windowsZdoom()
	}
	return ErrIncompatiblePlattaform
}
