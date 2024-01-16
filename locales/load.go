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

func read(file string) []byte {
	data, err := PoFiles.ReadFile(file)
	if err != nil {
		err = zenity.Error(readError)
		if err != nil {
			fmt.Println(err)
		}
		return read("po/en.po")
	}
	return data
}

func Po(lang string) *gotext.Po {
	po := gotext.NewPo()
	po.Parse(Parser(lang))
	return po
}

func Parser(lang string) []byte {
	var bytedata []byte
	bytedata = read("po/" + lang + ".po")
	return bytedata
}
