package locales

import (
	"embed"

	"github.com/leonelquinteros/gotext"
	"github.com/ncruces/zenity"
)

//go:embed po
var PoFiles embed.FS

const readError string = `A language is not available/does not exist in the configuration.
The available ones are:
- Spanish
- English
- Portuguese
`

func read(file string) []byte {
	data, err := PoFiles.ReadFile(file)
	if err != nil {
		zenity.Error(readError)
		return read("po/en.pot")
	}
	return data
}

func GetPo(lang string) *gotext.Po {
	var bytedata []byte
	if lang == "en" {
		bytedata = read("po/en.pot")
		po := gotext.NewPo()
		po.Parse(bytedata)
		return po
	}
	bytedata = read("po/" + lang + ".po")
	po := gotext.NewPo()
	po.Parse(bytedata)
	return po
}
