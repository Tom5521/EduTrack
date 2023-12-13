package locales

import (
	"embed"
	"fmt"

	"github.com/leonelquinteros/gotext"
)

//go:embed po
var PoFiles embed.FS

func read(file string, fs embed.FS) []byte {
	data, err := fs.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func GetPo(lang string) *gotext.Po {
	po := gotext.NewPo()
	po.Parse(read("po/"+lang+".po", PoFiles))
	return po
}
