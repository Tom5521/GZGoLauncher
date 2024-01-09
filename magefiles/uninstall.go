package main

import (
	"fmt"
	"go/build"
	"os"
	"runtime"
	"time"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Uninstall mg.Namespace

func (Uninstall) Go() error {
	nowtime := time.Now()
	path := build.Default.GOPATH + "/bin/" + ProyectName
	if runtime.GOOS == "linux" {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			return nil
		}
		fmt.Println("Removing from GOPATH...")
		err := sh.Rm(path)
		if err != nil {
			return err
		}
	}
	if runtime.GOOS == "windows" {
		if _, err := os.Stat(path + ".exe"); os.IsNotExist(err) {
			return nil
		}
		fmt.Println("Removing from GOPATH...")
		err := sh.Rm(path + ".exe")
		if err != nil {
			return err
		}
	}
	fmt.Println("[uninstall:go]Elapsed time: ", time.Since(nowtime).String())
	return nil
}

func (Uninstall) User() error {
	nowtime := time.Now()
	err := setupLinuxMake()
	if err != nil {
		return err
	}
	fmt.Println("Running make...")
	err = sh.RunV("make", "user-uninstall")
	if err != nil {
		return err
	}
	err = os.Chdir("..")
	if err != nil {
		return err
	}
	fmt.Println("[uninstall:user]Elapsed time: ", time.Since(nowtime).String())
	return nil
}

func (Uninstall) Root() error {
	nowtime := time.Now()
	err := setupLinuxMake()
	if err != nil {
		return err
	}
	fmt.Println("Running make...")
	err = sh.RunV("sudo", "make", "uninstall")
	if err != nil {
		return err
	}
	err = os.Chdir("..")
	if err != nil {
		return err
	}
	fmt.Println("[uninstall:root]Elapsed time: ", time.Since(nowtime).String())
	return nil
}
