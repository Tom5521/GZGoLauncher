package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/artdarek/go-unzip"
	"github.com/magefile/mage/sh"
	"github.com/yi-ge/unxz"
)

var (
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
		"GOOS":        "windows",
		"CGO_ENABLED": "1",
	}
	return env
}

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

type ex struct{}

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
