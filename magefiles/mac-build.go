package main

import "runtime"

func (b Build) setupMac() error {
	err := checkdir()
	if err != nil {
		return err
	}
	if runtime.GOOS == "darwin" {
		return b.nativeMac()
	}
	err = copyfile("./cmd/GZGoLauncher/main.go", "./main.go")
	if err != nil {
		return err
	}
	err = copyfile("./cmd/GZGoLauncher/FyneApp.toml", "FyneApp.toml")
	if err != nil {
		return err
	}
	return nil
}
