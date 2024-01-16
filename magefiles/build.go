package main

import (
	"fmt"
	"os"
	"time"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/pelletier/go-toml/v2"
)

type Build mg.Namespace

func (Build) All() error {
	nowtime := time.Now()
	defer func() {
		fmt.Println("[build:all]Elapsed time: ", time.Since(nowtime).String())
	}()
	err := compile.Linux()
	if err != nil {
		return err
	}
	err = compile.Windows()
	if err != nil {
		return err
	}
	err = compile.Mac()
	if err != nil {
		return err
	}
	return nil
}

func (Build) Windows() error {
	nowtime := time.Now()
	if err := checkdir(); err != nil {
		return err
	}
	fmt.Println("Compiling for windows...")
	err := sh.RunWithV(WindowsEnv, "fyne", "package", "--os", "windows", "--release",
		"--tags", "windows", "--src", MainDir, "--exe", fmt.Sprintf("builds/%s", WindowsExeName))
	if err != nil {
		return err
	}
	err = movefile(
		fmt.Sprintf("%s/builds/%s", MainDir, WindowsExeName),
		"./builds/"+WindowsExeName,
	)
	if err != nil {
		return err
	}
	fmt.Println("[build:windows]Elapsed time: ", time.Since(nowtime).String())
	return nil
}

func (Build) Linux() error {
	nowtime := time.Now()
	if err := checkdir(); err != nil {
		return err
	}
	fmt.Println("Compiling for linux...")
	err := sh.RunV("fyne", "package", "--os", "linux", "--release", "--tags", "linux", "--src", MainDir)
	if err != nil {
		return err
	}
	err = movefile(ProyectName+".tar.xz", fmt.Sprintf("builds/%s", LinuxTarName))
	if err != nil {
		return err
	}
	fmt.Println("[build:linux]Elapsed time: ", time.Since(nowtime).String())
	return nil
}

func (Build) Mac() error {
	err := checkdir()
	if err != nil {
		return err
	}
	type FyneApp struct {
		Details struct {
			Icon    string `toml:"Icon"`
			Name    string `toml:"Name"`
			ID      string `toml:"ID"`
			Version string `toml:"Version"`
			Build   int    `toml:"Build"`
		} `toml:"Details"`
	}
	var file FyneApp
	nowtime := time.Now()
	defer func() {
		fmt.Println("[build:Mac]Elapsed time: ", time.Since(nowtime).String())
	}()
	err = copyfile("./cmd/GZGoLauncher/main.go", "./main.go")
	if err != nil {
		return err
	}
	tomlbytedata, err := os.ReadFile("./cmd/GZGoLauncher/FyneApp.toml")
	if err != nil {
		return err
	}
	err = toml.Unmarshal(tomlbytedata, &file)
	if err != nil {
		return err
	}
	file.Details.Icon = IconPath
	bytedata, err := toml.Marshal(file)
	if err != nil {
		return err
	}
	err = os.WriteFile("FyneApp.toml", bytedata, os.ModePerm)
	if err != nil {
		return err
	}
	err = sh.RunV("sudo", "fyne-cross", "darwin", "--macosx-sdk-path", MacosSDKPath)
	if err != nil {
		return err
	}
	err = sh.Rm("./builds/GZGoLauncher-MacOS.zip")
	if err != nil {
		return err
	}
	err = os.Chdir("./fyne-cross/dist/darwin-amd64/")
	if err != nil {
		return err
	}
	err = sh.RunV("zip", "-r", "../../../builds/GZGoLauncher-MacOS.zip", ".")
	if err != nil {
		return err
	}
	err = os.Chdir("../../../")
	if err != nil {
		return err
	}
	toRemove := []string{"main.go", "FyneApp.toml"}
	for _, f := range toRemove {
		err = sh.Rm(f)
		if err != nil {
			return err
		}
	}
	return nil
}
