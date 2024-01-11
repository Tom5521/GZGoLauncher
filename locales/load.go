package locales

import (
	"embed"
	"fmt"

	"github.com/leonelquinteros/gotext"
	"github.com/ncruces/zenity"
)

//go:embed po
var PoFiles embed.FS

const readError string = `
A language is not available/does not exist in the configuration.
The available ones are:
- Spanish
- English
- Portuguese
`

var Languages = []string{"Español", "English", "Português"}

func read(file string) []byte {
	data, err := PoFiles.ReadFile(file)
	if err != nil {
		err = zenity.Error(readError)
		if err != nil {
			fmt.Println(err)
		}
		return read("po/en.pot")
	}
	return data
}

func GetPo(lang string) *gotext.Po {
	po := gotext.NewPo()
	po.Parse(GetParser(lang))
	return po
}

func GetParser(lang string) []byte {
	var bytedata []byte
	if lang == "en" {
		bytedata = read("po/en.pot")
		return bytedata
	}
	bytedata = read("po/" + lang + ".po")
	return bytedata
}
