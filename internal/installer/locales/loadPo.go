package locales

import (
	"embed"

	"github.com/leonelquinteros/gotext"
)

//go:embed po
var poFS embed.FS

func Read(file string) []byte {
	reader, _ := poFS.ReadFile("po/" + file)
	return reader
}

func GetPo(lang string) *gotext.Po {
	npo := gotext.NewPo()
	if lang != "en" {
		npo.Parse(Read(lang + ".po"))
	} else {
		npo.Parse(Read(lang + ".pot"))
	}
	return npo
}
