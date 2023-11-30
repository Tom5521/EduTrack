package data

import (
	"log"
	"os"
)

func CheckFiles() {
	check := func(file string) bool {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			return false
		}
		return true
	}
	db, conf := getOSConfFile()
	if !check(conf) {
		NewConfigurationFile()
		Config = GetConfData()
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
	Config = GetConfData()
	CheckFiles()
	DB = GetDB()
	LoadEverything()
}
