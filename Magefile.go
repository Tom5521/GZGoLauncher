//go:build mage
// +build mage

package main

import (
	"fmt"
	"go/build"
	"io"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/Tom5521/GoNotes/pkg/messages"
	"github.com/artdarek/go-unzip"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/yi-ge/unxz"
)

const (
	// Mesa32Url            = "https://downloads.fdossena.com/geth.php?r=mesa-latest"
	Mesa64Url            = "https://downloads.fdossena.com/geth.php?r=mesa64-latest"
	GoInstallURL         = "github.com/Tom5521/GZGoLauncher/cmd/GZGoLauncher@latest"
	ProyectName          = "GZGoLauncher"
	TmpDir               = "tmp/"
	MainDir              = "./cmd/GZGoLauncher/"
	WindowsExeName       = ProyectName + "-win64.exe"
	MakeWindowsZipTmpDir = "windows-tmp"
	WindowsZipName       = ProyectName + "-win64.zip"
	LinuxTarName         = ProyectName + "-linux64.tar.xz"
)

var (
	FilesToClean = []string{
		TmpDir,
		"./builds",
		"./cmd/GZGoLauncher/builds/",
	}
	compile    = Build{}
	extract    = ex{}
	WindowsEnv = windowsEnv()
)

func windowsEnv() map[string]string {
	if runtime.GOOS != "linux" {
		return map[string]string{}
	}
	env := map[string]string{
		"CC":          "/usr/bin/x86_64-w64-mingw32-gcc",
		"CXX":         "/usr/bin/x86_64-w64-mingw32-c++",
		"CGO_ENABLED": "1",
	}
	return env
}

type ex struct{}
type Build mg.Namespace
type Install mg.Namespace
type Uninstall mg.Namespace

func copyfile(src, dest string) error {
	fmt.Printf("Copying %s file to %s\n", src, dest)
	nowtime := time.Now()
	source, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	err = os.WriteFile(dest, source, os.ModePerm)
	if err != nil {
		return err
	}
	fmt.Println("[Copy]Elapsed time: ", time.Since(nowtime).String())
	return nil
}

func movefile(src, dest string) error {
	fmt.Printf("Moving %s file to %s\n", src, dest)
	nowtime := time.Now()
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
	fmt.Println("[Move]Elapsed time: ", time.Since(nowtime).String())
	return nil
}

func mkdir(dir string) error {
	fmt.Printf("Making \"%s\" directory...\n", dir)
	return os.Mkdir(dir, os.ModePerm)
}

func download(url, file string) error {
	fmt.Printf("Downloading %s", url)
	nowtime := time.Now()
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	outputFile, err := os.Create(file)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, response.Body)
	fmt.Println("Downloaded in ", file)
	fmt.Println("[Download]Elapsed time: ", time.Since(nowtime).String())
	return err
}

func (ex) tarXz(src, destDir string) error {
	if _, err := os.Stat(destDir); os.IsNotExist(err) {
		err := mkdir(destDir)
		if err != nil {
			return err
		}
	}
	u := unxz.New(src, destDir)
	return u.Extract()
}

func (ex) zip(src, dest string) error {
	uz := unzip.New(src, dest)
	return uz.Extract()
}

func downloadWinFiles() error {
	if _, err := os.Stat(TmpDir); os.IsNotExist(err) {
		err = mkdir(TmpDir)
		if err != nil {
			return err
		}
	}
	if _, err := os.Stat(TmpDir + "/opengl32.7z"); os.IsNotExist(err) {
		fmt.Println("Downloading opengl32.dll...")
		err = download(Mesa64Url, TmpDir+"/opengl32.7z")
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
		err = mkdir("builds")
		if err != nil {
			return err
		}
	}
	return nil
}

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
		fmt.Sprintf("./builds/%s", WindowsExeName),
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

func setupLinuxMake() error {
	if _, err := os.Stat("builds/" + LinuxTarName); os.IsNotExist(err) {
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
		err = extract.tarXz(LinuxTarName, ".")
		if err != nil {
			return err
		}
	}
	return nil
}

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
		fmt.Println("Making temporal dir...")
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
	err = os.RemoveAll(zipDir)
	if err != nil {
		return err
	}
	fmt.Println("[MakeWindowsZip]Elapsed time: ", time.Since(nowtime).String())
	return nil
}

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
