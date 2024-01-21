package config

import (
	"encoding/json"
	"os"

	v "github.com/Tom5521/GZGoLauncher/pkg/values"
	"github.com/Tom5521/GoNotes/pkg/messages"
)

type Mod struct {
	Path    string
	Enabled bool
}

type Wad string

func (w *Wad) IsValid() bool {
	stat, err := os.Stat(string(*w))
	if os.IsNotExist(err) {
		return false
	}
	if stat.IsDir() {
		return false
	}
	return true
}

type Config struct {
	ThemeMode      bool   `json:"theme-mode"` // 1 = light, 0 = dark
	ShowOutOnClose bool   `json:"show-out-on-close"`
	CloseOnStart   bool   `json:"close-on-start"`
	CustomArgs     string `json:"custom-args"`
	Lang           string `json:"lang"`
	GZDoomDir      string `json:"gzdoom-dir"`
	ZDoomDir       string `json:"zdoom-dir"`
	ExecDir        string `json:"runner-dir"`
	Mods           []Mod  `json:"mods"`
	Wads           []Wad  `json:"wads"`
}

func (c *Config) Write() error {
	bytedata, err := json.Marshal(c)
	if err != nil {
		return err
	}
	err = os.WriteFile(FilePath, bytedata, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

var (
	FilePath = func() string {
		if v.IsWindows {
			return v.HomeDir + WindowsPath + "/config.json"
		}
		return v.HomeDir + UnixPath + "/config.json"
	}()
	Path = func() string {
		if v.IsWindows {
			return v.HomeDir + WindowsPath
		}
		return v.HomeDir + UnixPath
	}()
	Settings = Read()
)

const (
	UnixPath    string = "/.config/GZGoLauncher"
	WindowsPath string = `\Documents\GZGoLauncher`
)

func Read() Config {
	var err error
	if _, err = os.Stat(Path); os.IsNotExist(err) {
		err = os.Mkdir(Path, os.ModePerm)
		if err != nil {
			messages.FatalError(err)
		}
	}
	if _, err = os.Stat(FilePath); os.IsNotExist(err) {
		zdoom, gzdoom := func() (string, string) {
			const z, g string = "zdoom", "gzdoom"
			if v.IsWindows {
				return z + ".exe", g + ".exe"
			}
			return z, g
		}()
		s := Config{Lang: "en", GZDoomDir: gzdoom, ZDoomDir: zdoom}
		var bytedata []byte
		bytedata, err = json.Marshal(s)
		if err != nil {
			messages.FatalError(err)
		}
		err = os.WriteFile(FilePath, bytedata, os.ModePerm)
		if err != nil {
			messages.FatalError(err)
		}
	}
	bytedata, err := os.ReadFile(FilePath)
	if err != nil {
		messages.FatalError(err)
	}
	ret := Config{}
	err = json.Unmarshal(bytedata, &ret)
	if err != nil {
		messages.FatalError(err)
	}
	return ret
}
