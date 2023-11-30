/*
 * Copyright (c) 2023 Tom5521- All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package data

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/user"
	"runtime"

	"gopkg.in/yaml.v3"
	"gorm.io/gorm"

	"github.com/glebarez/sqlite"
)

var Config ConfigStr
var databaseFile, configFile = getOSConfFile()

type ConfigStr struct {
	DatabaseFile string `yaml:"database"`
	Lang         string `yaml:"lang"` // TODO: Add multilanguage support
}

func CopyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func CreateDatabase() error {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return err
	}

	err = db.AutoMigrate(&Student{})
	printErr(err)
	err = db.AutoMigrate(&Grade{})
	printErr(err)
	err = db.AutoMigrate(&StudentGrade{})
	printErr(err)
	err = db.AutoMigrate(&Record{})
	printErr(err)

	defer func() { // Delete temporal database file
		if (runtime.GOOS == "linux" || runtime.GOOS == "unix") && Config.DatabaseFile != "database.db" {
			err = os.Remove("database.db")
			if err != nil {
				log.Println(err)
			}
		}
	}()

	_, err = CopyFile("database.db", Config.DatabaseFile)
	if err != nil {
		log.Println(err)
	}
	return err
}

func GetConfData() ConfigStr {
	conf := ConfigStr{}
	data, err := os.ReadFile(configFile)
	if err != nil {
		log.Println("Error reading config file!", err)
	}
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		log.Println(err)
	}
	return conf
}

func getOSConfFile() (string, string) {
	cOS := runtime.GOOS
	if cOS == "linux" || cOS == "unix" {
		currentUser, err := user.Current()
		if err != nil {
			fmt.Println(err)
		}
		confDir := fmt.Sprintf("%v/.config/EduTrack", currentUser.HomeDir)
		if _, err = os.Stat(confDir); os.IsNotExist(err) {
			err = os.Mkdir(confDir, os.ModePerm)
			if err != nil {
				log.Println("Error creating ~/.config/EduTrack/", err)
			}
		}
		return confDir + "/database.db", confDir + "/config.yml"
	}
	return "database.db", "config.yml"
}

func NewConfigurationFile() {
	var err error
	ymlData, err := yaml.Marshal(ConfigStr{DatabaseFile: databaseFile})
	if err != nil {
		log.Println("Error marshalling new configuration file", err)
	}
	err = os.WriteFile(configFile, ymlData, os.ModePerm)
	if err != nil {
		log.Println("Error writing config file", err)
	}
}
