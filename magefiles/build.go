package main

import (
	"fmt"
	"time"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Build mg.Namespace

func (Build) All() error {
	nowtime := time.Now()
	err := compile.Linux()
	if err != nil {
		return err
	}
	err = compile.Windows()
	if err != nil {
		return err
	}
	fmt.Println("[build:all]Elapsed time: ", time.Since(nowtime).String())
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
