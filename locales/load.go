package locales

import (
	"embed"
	"fmt"

	"github.com/leonelquinteros/gotext"
)

//go:embed po
var PoFiles embed.FS

func read(file string) []byte {
	data, err := PoFiles.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return read("po/en.po")
	}
	return data
}

func GetPo(lang string) *gotext.Po {
	po := gotext.NewPo()
	po.Parse(read("po/" + lang + ".po"))
	return po
}
