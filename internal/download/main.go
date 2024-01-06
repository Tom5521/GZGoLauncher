package download

import (
	"archive/zip"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/Tom5521/GZGoLauncher/internal/config"
)

var ErrOnlyForWindows = errors.New("only for windows!")

var settings = &config.Settings

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

func Unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		fpath := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, f.Mode())
		} else {
			var fdir string
			if lastIndex := strings.LastIndex(fpath, string(os.PathSeparator)); lastIndex > -1 {
				fdir = fpath[:lastIndex]
			}

			err = os.MkdirAll(fdir, f.Mode())
			if err != nil {
				log.Fatal(err)
				return err
			}
			f, err := os.OpenFile(
				fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func GZDoom() error {
	path := func() string {
		if runtime.GOOS == "windows" {
			return config.CurrentPath + "/gzdoom.zip"
		}
		return config.CurrentPath + "/gzdoom.tar.xz"
	}()
	url := func() string {
		if runtime.GOOS == "windows" {
			return WinGZDoomURL
		}
		if runtime.GOOS == "linux" {
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

	if runtime.GOOS == "windows" {
		err = Unzip("gzdoom.zip", "gzdoom")
	}
	if runtime.GOOS == "linux" {
		if _, err = os.Stat("gzdoom"); os.IsNotExist(err) {
			err = os.Mkdir("gzdoom", os.ModePerm)
			if err != nil {
				return err
			}
		}
		cmd := exec.Command("tar", "-xf", "gzdoom.tar.xz", "--strip-components=1", "-C", "gzdoom")
		err = cmd.Run()
		if err != nil {
			return err
		}
	}
	if runtime.GOOS == "linux" {
		settings.GZDoomDir = config.CurrentPath + "/gzdoom/gzdoom"
	}
	if runtime.GOOS == "windows" {
		settings.GZDoomDir = config.CurrentPath + "/gzdoom/gzdoom.exe"
	}
	err = settings.Write()
	if err != nil {
		return err
	}
	return nil
}

// Only with windows.
func ZDoom() error {
	if runtime.GOOS == "linux" {
		return ErrOnlyForWindows
	}
	path := config.CurrentFilePath + "/zdoom.zip"
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
	settings.ZDoomDir = "zdoom/zdoom.exe"
	err = settings.Write()
	if err != nil {
		return err
	}
	return nil
}
