package main

import (
	"fmt"
	"os"
	"time"

	"github.com/magefile/mage/sh"
)

func (Install) Go() error {
	nowtime := time.Now()
	fmt.Println("Running go install...")
	err := sh.RunV("go", "install", "-v", GoInstallURL)
	if err != nil {
		return err
	}
	fmt.Println("[install:go]Elapsed time: ", time.Since(nowtime).String())
	return err
}

func (Install) Root() error {
	nowtime := time.Now()
	err := setupLinuxMake()
	if err != nil {
		return err
	}
	fmt.Println("Running make...")
	err = sh.RunV("sudo", "make", "install")
	if err != nil {
		return err
	}
	err = os.Chdir("..")
	if err != nil {
		return err
	}
	fmt.Println("[install:root]Elapsed time: ", time.Since(nowtime).String())
	return nil
}

func (Install) User() error {
	nowtime := time.Now()
	err := setupLinuxMake()
	if err != nil {
		return err
	}
	fmt.Println("Running make...")
	err = sh.RunV("make", "user-install")
	if err != nil {
		return err
	}
	err = os.Chdir("..")
	if err != nil {
		return err
	}
	fmt.Println("[install:user]Elapsed time: ", time.Since(nowtime).String())
	return nil
}
