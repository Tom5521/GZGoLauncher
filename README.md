# GZGoLauncher

A cross-platform launcher for *ZDoom

## Features

- Light/dark mode toggle
- Cross-platform
- Assisted download of GZDoom/ZDoom

## Installation

Install GZGoLauncher with Go or by executing the binary in its own folder.

### All compatible systems(linux,windows,macos)

Requirements:

- Go compiler
- C compiler

```bash
go install -v github.com/Tom5521/GZGoLauncher/cmd/GZGoLauncher@latest
```

### Windows

1. Download the package from [releases](https://github.com/Tom5521/GZGoLauncher/releases/latest)
for your architecture.

2. Unzip the package

3. Run `GZGoLauncher.exe`; the binary is completely portable.

### MacOS

1. Download the package from [releases](https://github.com/Tom5521/GZGoLauncher/releases/latest)
for your architecture.

2. Unzip the package and move `GZGoLauncher.app` to your Applications folder,
or you can execute the binary in `GZGoLauncher.app/Contents/MacOS/GZGoLauncher`.

### Linux

1. Download the package from [releases](https://github.com/Tom5521/GZGoLauncher/releases/latest)
for your architecture.

2. Untar the package and cd into the folder.

3. Run `make install` or `make user-install` for user local installation.

## Build

Advance Notes:

- For all these methods you have to clone the repository and open a terminal
inside it.
- Compilation binaries/packages will always be in `builds/` and/or `cmd/GZGoLauncher/builds`
unless otherwise stated in the documentation.

### To Linux

#### Linux systems

Requirements:

- C compiler
- Go compiler
- Mage
- Fyne Package

Instructions:

- Run `mage build:linux`

#### Windows systems

Requirements:

- Same as linux
- WSL

Instructions:

- Open cmd or powershell and enter on WSL
- Run `mage build:linux`

### To Windows

#### Windows systems

Requirements:

- C compiler
- Go compiler
- Zip (to release the package)
- 7z (to unzip the opengl32.dll file)
- Mage
- Fyne Package

Instructions:

- Run mage `build:windows`

#### Non Windows systems

Requirements:

- Same as windows
- x86_64-w64-mingw32-gcc

Instructions:

- Run `mage build:windows`

### To macOS

#### Mac systems

Requirements:

- Fyne Package
- C compiler
- Go compiler
- Mage
- Zip (To release the package)

#### Non mac systems

Requirements:

- Same as mac

Instructions:

- [Download Command line tools for xcode](https://developer.apple.com/download/all/?q=Command%20Line%20Tools)>= 12.4 (macOS SDK 11.x)
- Extract it following the [instructions](https://github.com/fyne-io/fyne-cross?tab=readme-ov-file#extract-the-macos-sdk-for-osxdarwinapple-cross-compiling)
- Create the SDKs folder and inside paste or create a symlink to MacOSX11.sdk
- Run `mage build:macARM` or `mage build:macAMD` depending on your architecture

## Screenshots

![screenshot](./screenshots/Screenshot1.png)
![screenshot](./screenshots/Screenshot2.png)

<https://github.com/Tom5521/GZGoLauncher/assets/88908582/d9692c54-9dfa-4f68-b89e-4828cc160daa>

## License

[MIT](https://choosealicense.com/licenses/mit/)
