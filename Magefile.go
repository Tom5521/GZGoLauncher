//go:build mage && (windows || linux)

package main

import (
	"fmt"
	"go/build"
	"os"
	"runtime"

	"github.com/Tom5521/GoNotes/pkg/messages"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

const (
	Mesa64Url = "https://downloads.fdossena.com/geth.php?r=mesa64-latest"
)

var (
	TmpDir     = "./tmp"
	MainDir    = "./cmd/GZGoLauncher/"
	WindowsEnv = windowsEnv()
	compile    = Build{}
)

type Build mg.Namespace
type Install mg.Namespace
type Uninstall mg.Namespace

func windowsEnv() map[string]string {
	var env map[string]string
	if runtime.GOOS == "linux" {
		env = map[string]string{
			"CC":          "/usr/bin/x86_64-w64-mingw32-gcc",
			"CXX":         "/usr/bin/x86_64-w64-mingw32-c++",
			"CGO_ENABLED": "1",
		}
	}
	return env
}

func copyfile(src, dest string) error {
	fmt.Printf("Copying %s file to %s\n", src, dest)
	source, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	err = os.WriteFile(dest, source, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func movefile(src, dest string) error {
	fmt.Printf("Moving %s file to %s\n", src, dest)
	source, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	err = os.WriteFile(dest, source, os.ModePerm)
	if err != nil {
		return err
	}
	err = os.Remove(src)
	if err != nil {
		return err
	}
	return nil
}

func downloadWinFiles() error {
	if _, err := os.Stat(TmpDir); os.IsNotExist(err) {
		err = os.Mkdir(TmpDir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	if _, err := os.Stat(TmpDir + "/opengl32.7z"); os.IsNotExist(err) {
		fmt.Println("Downloading opengl32.dll...")
		err = sh.RunV("wget", "-O", "tmp/opengl32.7z", Mesa64Url)
		if err != nil {
			return err
		}
	}
	if _, err := os.Stat(TmpDir + "/opengl32.dll"); os.IsNotExist(err) {
		fmt.Println("Extracting...")
		err = os.Chdir(TmpDir)
		if err != nil {
			return err
		}
		if err = sh.RunV("7z", "e", "opengl32.7z"); err != nil {
			return err
		}
		err = os.Chdir("..")
		if err != nil {
			return err
		}
	}
	return nil
}

func checkdir() error {
	if _, err := os.Stat("builds"); os.IsNotExist(err) {
		fmt.Println("Making \"builds\" directory...")
		err = os.Mkdir("builds", os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func (Build) All() error {
	err := compile.Linux()
	if err != nil {
		return err
	}
	err = compile.Windows()
	if err != nil {
		return err
	}
	return nil
}

// Compile the program to be distributed on windows, NOTE: This will only return an .exe of the program, the installation in windows can only be done through the installer.
func (Build) Windows() error {
	if err := checkdir(); err != nil {
		return err
	}
	fmt.Println("Compiling for windows...")
	err := sh.RunWithV(WindowsEnv, "fyne", "package", "--os", "windows", "--release",
		"--tags", "windows", "--src", MainDir, "--exe", "builds/GZGoLauncher.exe")
	if err != nil {
		return err
	}
	err = movefile(MainDir+"/builds/GZGoLauncher.exe", "./builds/GZGoLauncher.exe")
	if err != nil {
		return err
	}
	return nil
}

// Compile the program to be distributed on linux.
func (Build) Linux() error {
	if err := checkdir(); err != nil {
		return err
	}
	fmt.Println("Compiling for linux...")
	err := sh.RunV("fyne", "package", "--os", "linux", "--release", "--tags", "linux", "--src", MainDir)
	if err != nil {
		return err
	}
	err = movefile("GZGoLauncher.tar.xz", "builds/GZGoLauncher-linux64.tar.xz")
	if err != nil {
		return err
	}
	return nil
}

func setupLinuxMake() error {
	if _, err := os.Stat("builds/GZGoLauncher-linux64.tar.xz"); os.IsNotExist(err) {
		err = compile.Linux()
		if err != nil {
			return err
		}
	}
	err := os.Chdir("builds")
	if err != nil {
		return err
	}
	if _, err = os.Stat("Makefile"); os.IsNotExist(err) {
		fmt.Println("Extracting tarfile...")
		err = sh.RunV("tar", "-xvf", "GZGoLauncher-linux64.tar.xz")
		if err != nil {
			return err
		}
	}

	return nil
}

// Delete temporary directories, compilation files, etc, It leaves it as if it had just been cloned.
func Clean() {
	toRemove := []string{
		"./tmp",
		"./builds",
		"./cmd/GZGoLauncher/builds/",
	}

	for _, f := range toRemove {
		if _, err := os.Stat(f); os.IsNotExist(err) {
			continue
		}
		fmt.Printf("Deleting %s...\n", f)
		err := sh.Rm(f)
		if err != nil {
			messages.Error(err)
		}
	}
}

func MakeWindowsZip() error {
	var zipDir = "windows-tmp"
	if _, err := os.Stat(zipDir); os.IsNotExist(err) {
		fmt.Println("Making temporal dir...")
		err = os.Mkdir(zipDir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	if _, err := os.Stat("tmp/opengl32.dll"); os.IsNotExist(err) {
		err = downloadWinFiles()
		if err != nil {
			return err
		}
	}
	err := copyfile("tmp/opengl32.dll", zipDir+"/opengl32.dll")
	if err != nil {
		return err
	}
	if _, err = os.Stat("builds/"); os.IsNotExist(err) {
		err = compile.Windows()
		if err != nil {
			return err
		}
	}
	err = copyfile("builds/GZGoLauncher.exe", zipDir+"/GZGoLauncher.exe")
	if err != nil {
		return err
	}
	err = copyfile("README.md", zipDir+"/README.md")
	if err != nil {
		return err
	}
	if _, err = os.Stat("builds/GZGoLauncher-win64.zip"); os.IsExist(err) {
		err = os.Remove("builds/GZGoLauncher-win64.zip")
		if err != nil {
			return err
		}
	}
	err = os.Chdir(zipDir)
	if err != nil {
		return err
	}

	fmt.Println("Zipping content...")
	err = sh.RunV("zip", "-r", "../builds/GZGoLauncher-win64.zip", ".")
	if err != nil {
		return err
	}
	err = os.Chdir("..")
	if err != nil {
		return err
	}
	fmt.Println("Cleaning...")
	err = os.RemoveAll(zipDir)
	if err != nil {
		return err
	}
	return nil
}

func (Install) Go() error {
	fmt.Println("Running go install...")
	err := sh.RunV("go", "install", "-v", "github.com/Tom5521/GZGoLauncher/cmd/GZGoLauncher@latest")
	if err != nil {
		return err
	}
	return err
}

// NOTE: Only works in linux, in windows you will have to use the installer.
func (Install) Root() error {
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
	return nil
}

// NOTE: Only works in linux, in windows you will have to use the installer.
func (Install) User() error {
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
	return nil
}

func (Uninstall) Go() error {
	path := build.Default.GOPATH + "/bin/GZGoLauncher"
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
	return nil
}

func (Uninstall) User() error {
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
	return nil
}

func (Uninstall) Root() error {
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
	return nil
}
