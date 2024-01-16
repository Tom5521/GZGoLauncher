package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Tom5521/GoNotes/pkg/messages"
	"github.com/magefile/mage/sh"
)

func Clean() {
	nowtime := time.Now()
	for _, f := range FilesToClean {
		if _, err := os.Stat(f); os.IsNotExist(err) {
			continue
		}
		fmt.Printf("Deleting %s...\n", f)
		err := sh.Rm(f)
		if err != nil {
			messages.Error(err)
		}
	}
	fmt.Println("[clean]Elapsed time: ", time.Since(nowtime).String())
}

func MakeWindowsZip() error {
	nowtime := time.Now()
	var zipDir = MakeWindowsZipTmpDir
	if _, err := os.Stat(zipDir); os.IsNotExist(err) {
		err = mkdir(zipDir)
		if err != nil {
			return err
		}
	}
	if _, err := os.Stat(TmpDir + "/opengl32.dll"); os.IsNotExist(err) {
		err = downloadWinFiles()
		if err != nil {
			return err
		}
	}
	err := copyfile(TmpDir+"/opengl32.dll", zipDir+"/opengl32.dll")
	if err != nil {
		return err
	}
	if _, err = os.Stat("builds/" + WindowsExeName); os.IsNotExist(err) {
		err = compile.Windows()
		if err != nil {
			return err
		}
	}
	err = copyfile("builds/"+WindowsExeName, fmt.Sprintf("%s/%s", zipDir, WindowsExeName))
	if err != nil {
		return err
	}
	err = copyfile("README.md", zipDir+"/README.md")
	if err != nil {
		return err
	}
	if _, err = os.Stat("builds/" + WindowsZipName); os.IsExist(err) {
		err = os.Remove("builds/" + WindowsZipName)
		if err != nil {
			return err
		}
	}
	err = os.Chdir(zipDir)
	if err != nil {
		return err
	}

	fmt.Println("Zipping content...")
	err = sh.RunV("zip", "-r", "../builds/"+WindowsZipName, ".")
	if err != nil {
		return err
	}
	err = os.Chdir("..")
	if err != nil {
		return err
	}
	fmt.Println("Cleaning...")
	err = sh.Rm(zipDir)
	if err != nil {
		return err
	}
	fmt.Println("[MakeWindowsZip]Elapsed time: ", time.Since(nowtime).String())
	return nil
}

func Release() error {
	nowtime := time.Now()
	defer func() {
		fmt.Println("[release]Elapsed time: ", time.Since(nowtime).String())
	}()
	err := compile.All()
	if err != nil {
		return err
	}
	err = MakeWindowsZip()
	if err != nil {
		return err
	}
	return nil
}
