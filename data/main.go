/*
 * Copyright (c) 2023 Tom5521- All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package data

import (
	"log"
	"os"

	"github.com/ncruces/zenity"
)

var Db DbStr

func NotifyError(text string, err error) {
	zenity.Notify(text + "::" + err.Error())
}

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
	Db = InitDB()
}
