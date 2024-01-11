package config

import (
	"encoding/json"
	"os"
	"os/user"
	"runtime"

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
	Lang      string `json:"lang"`
	GZDoomDir string `json:"gzdoom-dir"`
	ZDoomDir  string `json:"zdoom-dir"`
	GZDir     string `json:"gzdir"`
	Mods      []Mod  `json:"mods"`
	Wads      []Wad  `json:"wads"`
}

func (c *Config) Write() error {
	bytedata, err := json.Marshal(c)
	if err != nil {
		return err
	}
	err = os.WriteFile(CurrentFilePath, bytedata, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

var (
	HomeDir = func() string {
		usr, err := user.Current()
		if err != nil {
			messages.FatalError(err)
		}
		return usr.HomeDir
	}()
	CurrentFilePath = func() string {
		usr, err := user.Current()
		if err != nil {
			messages.FatalError(err)
		}
		if runtime.GOOS == "windows" {
			return usr.HomeDir + WindowsPath + "/config.json"
		}
		return usr.HomeDir + UnixPath + "/config.json"
	}()
	CurrentPath = func() string {
		if runtime.GOOS == "windows" {
			return HomeDir + WindowsPath
		}
		return HomeDir + UnixPath
	}()
	Settings = Read()
)

const (
	UnixPath    string = "/.config/GZGoLauncher"
	WindowsPath string = "\\Documents\\GZGoLauncher"
)

func Read() Config {
	var err error
	if _, err = os.Stat(CurrentPath); os.IsNotExist(err) {
		err = os.Mkdir(CurrentPath, os.ModePerm)
		if err != nil {
			messages.FatalError(err)
		}
	}
	if _, err = os.Stat(CurrentFilePath); os.IsNotExist(err) {
		s := Config{Lang: "en", GZDoomDir: "gzdoom", ZDoomDir: "zdoom"}
		var bytedata []byte
		bytedata, err = json.Marshal(s)
		if err != nil {
			messages.FatalError(err)
		}
		err = os.WriteFile(CurrentFilePath, bytedata, os.ModePerm)
		if err != nil {
			messages.FatalError(err)
		}
	}
	bytedata, err := os.ReadFile(CurrentFilePath)
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
