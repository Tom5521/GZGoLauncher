package main

const (
	// Mesa32Url            = "https://downloads.fdossena.com/geth.php?r=mesa-latest"
	Mesa64Url            = "https://downloads.fdossena.com/geth.php?r=mesa64-latest"
	GoInstallURL         = "github.com/Tom5521/GZGoLauncher/cmd/GZGoLauncher@latest"
	ProyectName          = "GZGoLauncher"
	TmpDir               = "tmp/"
	MainDir              = "./cmd/GZGoLauncher/"
	MainFile             = MainDir + "main.go"
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
)
