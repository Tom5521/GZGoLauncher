

latest_tag := `git describe --tags --abbrev=0`

main_dir := "./cmd/GZGoLauncher/"
main_file := main_dir+"main.go"

icon_path := "./assets/cacodemon.png"
macos_sdk := "SDKs/MacOSX.sdk/"

go_install_url := "github.com/Tom5521/GZGoLauncher/cmd/GZGoLauncher"

mesa64_url := "https://downloads.fdossena.com/geth.php?r=mesa64-latest"
mesa32_url := "https://downloads.fdossena.com/geth.php?r=mesa-latest"

proyect_name := "GZGoLauncher"


build os arch:  
  #!/bin/bash
  cp {{main_dir}}/* .
 
  build_cmd="fyne-cross {{os}} -arch {{arch}} -icon {{icon_path}}"

  if [[ {{os}} == "darwin" ]]; then
    build_cmd=$build_cmd --macosx-sdk-path={{macos_sdk}} 
  fi

  $build_cmd

  rm main.go FyneApp.toml

run:
  go run -v ./cmd/GZGoLauncher/

build-all:
  just build windows amd64
  just build linux amd64
  # fyne-cross is broken...
  # just build darwin arm64
  # just build darwin amd64

release:
  just build-all

  gh release create {{latest_tag}} ./fyne-cross/dist/*/* --generate-notes


clean:
  rm -rf builds cmd/GZGoLauncher/builds ./main.go ./FyneApp.toml fyne-cross

go-install:
  go install -v {{go_install_url}}@latest

go-uninstall:
  rm ~/go/bin/{{proyect_name}}
