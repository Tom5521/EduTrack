package locales

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

func LoadFiles(lang string) Locale {
	english, _ := LocaleFiles.ReadFile("english.yml")
	spanish, _ := LocaleFiles.ReadFile("spanish.yml")
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
	err := yaml.Unmarshal(fileToRead, &locale)
	if err != nil {
		fmt.Println(err)
	}
	return locale
}
