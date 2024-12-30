package locales

import (
	"embed"
	"fmt"

	"github.com/Tom5521/GZGoLauncher/internal/config"
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

var locales = map[string]string{
	"es": "Español",
	"en": "English",
	"pt": "Português",
}

func LocaleShort(long string) string {
	for k, v := range locales {
		if v == long {
			return k
		}
	}
	return ""
}

func LocaleLong(short string) string {
	return locales[short]
}

func LocaleNames() (names []string) {
	for _, v := range locales {
		names = append(names, v)
	}

	return
}

func ShortLocaleExists(short string) bool {
	_, ok := locales[short]
	return ok
}

func LongLocaleExists(long string) bool {
	for _, v := range locales {
		if v == long {
			return true
		}
	}

	return false
}

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

var Current = Po(config.Settings.Lang)
