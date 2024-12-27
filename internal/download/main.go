package download

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"slices"

	"github.com/Tom5521/GZGoLauncher/internal/config"
	"github.com/Tom5521/GZGoLauncher/locales"
	v "github.com/Tom5521/GZGoLauncher/pkg/values"
)

var (
	po = locales.Current

	ErrIncompatiblePlattaform = errors.New("incompatible plattaform")
	ErrOnlyForWindows         = errors.New("only for windows")
	ErrZDoomOnMac             = errors.New("zdoom is not available on mac")
	ErrArIsNotInPath          = errors.New("ar executable is not in path")
)

var settings = &config.Settings

const (
	WinGZDoomURL = "https://github.com/ZDoom/gzdoom/releases/download/g4.14.0/gzdoom-4-14-0a-windows.zip"
	WinZDoomURL  = "https://zdoom.org/files/zdoom/2.8/zdoom-2.8.1.zip"

	LinuxGZDoomURL = "https://github.com/ZDoom/gzdoom/releases/download/g4.11.3/gzdoom-g4.11.3-linux-portable.tar.xz"
	LinuxZDoomURL  = "https://zdoom.org/files/zdoom/2.8/zdoom_2.8.1_amd64.deb"

	MacGZDoomURL = "https://github.com/ZDoom/gzdoom/releases/download/g4.14.0/gzdoom-4-14-0-macos.zip"
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
	var (
		path string
		err  error
	)
	switch {
	case v.IsDarwin:
		path, err = macGZDoom()
	case v.IsLinux:
		path, err = linuxGZDoom()
	case v.IsWindows:
		path, err = windowsGZDoom()
	default:
		return ErrIncompatiblePlattaform
	}

	if err != nil {
		return err
	}

	settings.SourcePorts = append(config.Settings.SourcePorts,
		config.SourcePort{
			Name:           genName("GZDoom"),
			ExecutablePath: path,
		},
	)

	return settings.Write()
}

// Only with windows.
func ZDoom() error {
	var (
		err  error
		path string
	)
	switch {
	case v.IsDarwin:
		path, err = macZDoom()
	case v.IsLinux:
		path, err = linuxZdoom()
	case v.IsWindows:
		path, err = windowsZdoom()
	default:
		return ErrIncompatiblePlattaform
	}
	if err != nil {
		return err
	}

	settings.SourcePorts = append(settings.SourcePorts,
		config.SourcePort{
			Name:           genName("ZDoom"),
			ExecutablePath: path,
		},
	)

	return err
}

func genName(prefix string) (name string) {
	var diff int
	name = prefix
	for slices.ContainsFunc(settings.SourcePorts, func(i config.SourcePort) bool { return i.Name == name }) {
		diff++
		name = fmt.Sprintf("%s-%d", prefix, diff)
	}

	return
}
