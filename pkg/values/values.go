package values

import (
	"os/user"
	"runtime"
)

const (
	IsWindows = runtime.GOOS == "windows"
	IsLinux   = runtime.GOOS == "linux"
	IsMac     = runtime.GOOS == "darwin"
)

var (
	HomeDir = func() string {
		usr, err := user.Current()
		if err != nil {
			panic(err)
		}
		return usr.HomeDir
	}()
)
