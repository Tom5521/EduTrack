/*
 * Copyright (c) 2023 Tom5521- All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package data

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"os/user"
	"runtime"
	"strconv"

	"gopkg.in/yaml.v3"
)

var Config ConfigStr
var DatabaseFile, ConfigFile = getOSConfFile()

type ConfigStr struct {
	DatabaseFile string
	Lang         string // TODO: Add multilanguage support
}

//go:embed database.db
var SqlTemplate []byte

func CreateDatabase() error {
	_, err := os.Create(Config.DatabaseFile)
	if err != nil {
		log.Println(err)
		return err
	}
	file, err := os.OpenFile(Config.DatabaseFile, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = file.Write(SqlTemplate)
	if err != nil {
		log.Println(err)
		return err
	}
	currentUser, _ := user.Current()

	atoi := func(s string) int {
		res, _ := strconv.Atoi(s)
		return res
	}
	err = os.Chown(Config.DatabaseFile, atoi(currentUser.Uid), atoi(currentUser.Gid))
	return err
}

func GetConfData() ConfigStr {
	var err error
	conf := ConfigStr{}
	data, err := os.ReadFile(ConfigFile)
	if err != nil {
		NotifyError("Error reading config file!", err)
	}
	yaml.Unmarshal(data, &conf)
	return conf
}

func getOSConfFile() (dataDB string, ConfYml string) {
	if runtime.GOOS == "linux" || runtime.GOOS == "unix" {
		CurrentUser, err := user.Current()
		if err != nil {
			fmt.Println(err)
		}
		confDir := fmt.Sprintf("%v/.config/EduTrack", CurrentUser.HomeDir)
		if _, err := os.Stat(confDir); os.IsNotExist(err) {
			err := os.Mkdir(confDir, os.ModePerm)
			if err != nil {
				NotifyError("Error creating ~/.config/EduTrack/", err)
			}
		}
		return confDir + "/database.db", confDir + "/config.yml"
	} else {
		return "database.db", "config.yml"
	}
}

func NewConfigurationFile() {
	var err error
	ymlData, err := yaml.Marshal(ConfigStr{DatabaseFile: DatabaseFile})
	if err != nil {
		NotifyError("Error marshalling new configuration file", err)
	}
	err = os.WriteFile(ConfigFile, ymlData, os.ModePerm)
	if err != nil {
		NotifyError("Error writing config file", err)
	}
}
