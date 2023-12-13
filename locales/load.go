package locales

import (
	"embed"

	"github.com/leonelquinteros/gotext"
	"github.com/ncruces/zenity"
)

//go:embed po
var PoFiles embed.FS

func read(file string) []byte {
	data, err := PoFiles.ReadFile(file)
	if err != nil {
		zenity.Error(err.Error())
		return read("po/en.po")
	}
	return data
}

func GetPo(lang string) *gotext.Po {
	po := gotext.NewPo()
	po.Parse(read("po/" + lang + ".po"))
	return po
}
