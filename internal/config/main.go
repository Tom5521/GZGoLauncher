package config

import (
	"encoding/json"
	"os"

	t "github.com/Tom5521/GZGoLauncher/internal/tools"
	v "github.com/Tom5521/GZGoLauncher/pkg/values"
)

var (
	FilePath = Path + "/config.json"
	Path     = func() string {
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

type Mod struct {
	Path    string
	Enabled bool
}

type Wad string

type SourcePort struct {
	ExecutablePath string `json:"executable"`
	Name           string
}

func (w Wad) IsValid() bool {
	stat, err := os.Stat(string(w))
	return !os.IsNotExist(err) || !stat.IsDir()
}

type Config struct {
	ThemeMode         bool         `json:"theme-mode"` // 1 = light, 0 = dark
	ShowOutOnClose    bool         `json:"show-out-on-close"`
	CloseOnStart      bool         `json:"close-on-start"`
	CustomArgs        string       `json:"custom-args"`
	Lang              string       `json:"lang"`
	Mods              []Mod        `json:"mods"`
	Wads              []Wad        `json:"wads"`
	CurrentSourcePort int          `json:"current-source-port"`
	SourcePorts       []SourcePort `json:"source-ports"`
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

func Read() Config {
	var err error
	if _, err = os.Stat(Path); os.IsNotExist(err) {
		err = os.Mkdir(Path, os.ModePerm)
		if err != nil {
			t.FatalErrWin(err)
		}
	}
	if _, err = os.Stat(FilePath); os.IsNotExist(err) {
		s := Config{Lang: "en"}
		var bytedata []byte
		bytedata, err = json.Marshal(s)
		if err != nil {
			t.FatalErrWin(err)
		}
		err = os.WriteFile(FilePath, bytedata, os.ModePerm)
		if err != nil {
			t.FatalErrWin(err)
		}
	}
	bytedata, err := os.ReadFile(FilePath)
	if err != nil {
		t.FatalErrWin(err)
	}
	ret := Config{}
	err = json.Unmarshal(bytedata, &ret)
	if err != nil {
		t.FatalErrWin(err)
	}
	return ret
}
