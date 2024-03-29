package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
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
	err = compile.MacAMD()
	if err != nil {
		return err
	}
	err = compile.MacARM()
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

func (b Build) MacAMD() error {
	err := b.setupMac()
	if err != nil {
		return err
	}
	nowtime := time.Now()
	defer func() {
		fmt.Println("[build:MacAMD]Elapsed time: ", time.Since(nowtime).String())
	}()
	err = sh.RunV("sudo", "fyne-cross", "darwin",
		"-icon", "./assets/cacodemon.png", "-arch=amd64", "--macosx-sdk-path", MacosSDKPath)
	if err != nil {
		return err
	}
	err = sh.Rm("./builds/" + MacZipNameAmd64)
	if err != nil {
		return err
	}
	err = os.Chdir("./fyne-cross/dist/darwin-amd64/")
	if err != nil {
		return err
	}
	err = sh.RunV("zip", "-r", "../../../builds/"+MacZipNameAmd64, ".")
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

func (b Build) MacARM() error {
	err := b.setupMac()
	if err != nil {
		return err
	}
	nowtime := time.Now()
	defer func() {
		fmt.Println("[build:MacARM]Elapsed time: ", time.Since(nowtime).String())
	}()
	err = sh.RunV("sudo", "fyne-cross", "darwin",
		"-icon", "./assets/cacodemon.png", "-arch=arm64", "--macosx-sdk-path", MacosSDKPath)
	if err != nil {
		return err
	}
	err = sh.Rm("./builds/" + MacZipNameArm64)
	if err != nil {
		return err
	}
	err = os.Chdir("./fyne-cross/dist/darwin-arm64/")
	if err != nil {
		return err
	}
	err = sh.RunV("zip", "-r", "../../../builds/"+MacZipNameArm64, ".")
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

func (b Build) DefaultMac() error {
	if runtime.GOARCH == "arm64" {
		return b.MacARM()
	}
	if runtime.GOARCH == "amd64" {
		return b.MacAMD()
	}
	return nil
}

func (Build) nativeMac() error {
	nowtime := time.Now()
	defer func() {
		fmt.Println("[build:macARM]Elapsed time: ", time.Since(nowtime).String())
	}()
	fmt.Println("Compiling for YOUR mac...")
	err := sh.RunV("fyne", "package", "--os", "darwin", "--release", "--src", MainDir)
	if err != nil {
		return err
	}
	return nil
}

func (b Build) Mac(mode string) error {
	switch mode {
	case "ARM":
		return b.MacARM()
	case "AMD":
		return b.MacAMD()
	default:
		return b.DefaultMac()
	}
}
