package data

import (
	"log"
	"os"

	"github.com/Tom5521/EduTrack/pkg/conf"
)

func CheckFiles() {
	check := func(file string) bool {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			return false
		}
		return true
	}
	db, config := Config.DatabaseFile, conf.ConfFile
	if !check(config) {
		conf.NewConfigurationFile()
		Config = conf.GetConfData()
		CheckFiles()
	}
	if !check(db) {
		if !check(Config.DatabaseFile) {
			err := CreateDatabase()
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func LoadFiles() {
	CheckFiles()
	Config = conf.GetConfData()
	DB = GetDB()
	LoadEverything()
}
