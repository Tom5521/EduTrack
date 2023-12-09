package locales

import (
	"embed"
	"fmt"

	"gopkg.in/yaml.v3"
)

//go:embed files
var LocaleFiles embed.FS

func LoadFiles(lang string) Locale {
	english, err := LocaleFiles.ReadFile("files/english.yml")
	if err != nil {
		fmt.Println(err)
	}
	spanish, err := LocaleFiles.ReadFile("files/spanish.yml")
	if err != nil {
		fmt.Println(err)
	}
	var (
		locale     Locale
		fileToRead []byte
	)
	if lang == "es" {
		fileToRead = spanish
	}
	if lang == "en" {
		fileToRead = english
	}
	err = yaml.Unmarshal(fileToRead, &locale)
	if err != nil {
		fmt.Println(err)
	}
	return locale
}
