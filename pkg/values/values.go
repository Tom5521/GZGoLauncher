package values

import (
	"os/user"
	"runtime"

	"github.com/Tom5521/GoNotes/pkg/messages"
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
			messages.FatalError(err)
		}
		return usr.HomeDir
	}()
)
