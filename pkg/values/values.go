package values

import (
	"os/user"
	"runtime"
)

const (
	IsWindows = runtime.GOOS == "windows"
	IsLinux   = runtime.GOOS == "linux"
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
