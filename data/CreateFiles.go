/*
 * Copyright (c) 2023 Tom5521- All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/user"
	"runtime"

	"gopkg.in/yaml.v3"

	_ "github.com/glebarez/go-sqlite"
)

var Config ConfigStr
var DatabaseFile, ConfigFile = getOSConfFile()

type ConfigStr struct {
	DatabaseFile string
	Lang         string // TODO: Add multilanguage support
}

//goasd:embed database.db
//var SqlTemplate []byte

func CreateDatabase() error {
	/*
		file, err := os.Create(Config.DatabaseFile)
		if err != nil {
			log.Println(err)
			return err
		}
		defer file.Close()
		if err != nil {
			log.Println(err)
			return err
		}
		defer file.Close()
		_, err = file.Write(SqlTemplate)
		if err != nil {
			log.Println(err)
			return err
		}*/
	db, err := sql.Open("sqlite", Config.DatabaseFile)
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()
	const Query string = `
CREATE TABLE IF NOT EXISTS "Grades" (
	"grade_id"	INTEGER,
	"Name"	TEXT,
	"Info"	TEXT,
	"Price"	INTEGER,
	PRIMARY KEY("grade_id" AUTOINCREMENT)
);

CREATE TABLE IF NOT EXISTS "Records" (
	"record_id"	INTEGER,
	"student_id"	INTEGER,
	"Name"	TEXT,
	"Date"	TEXT,
	"Info"	TEXT,
	PRIMARY KEY("record_id" AUTOINCREMENT)
);

CREATE TABLE IF NOT EXISTS "Student_grades" (
	"student_id"	INTEGER,
	"grade_id"	INTEGER,
	"start"	TEXT,
	"end"	BLOB,
	"student_grade_id"	INTEGER,
	PRIMARY KEY("student_grade_id" AUTOINCREMENT)
);

CREATE TABLE IF NOT EXISTS "Students" (
	"student_id"	INTEGER,
	"Name"	TEXT,
	"DNI"	TEXT,
	"Age"	INTEGER,
	"Phone_Number"	TEXT,
	"ImagePath"	TEXT,
	PRIMARY KEY("student_id" AUTOINCREMENT)
);
`

	_, err = db.Exec(Query)
	if err != nil {
		log.Println(err)
		return err
	}
	return err

	/*currentUser, _ := user.Current()

	atoi := func(s string) int {
		res, _ := strconv.Atoi(s)
		return res
	}
	err = os.Chown(Config.DatabaseFile, atoi(currentUser.Uid), atoi(currentUser.Gid))
	if err != nil {
		log.Println(err)
	}
	_, err = file.Write(SqlTemplate)
	if err != nil {
		log.Println(err)
	}
	return err*/
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
