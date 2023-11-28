/*
 * Copyright (c) 2023 Tom5521- All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package data

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"os/user"
	"runtime"

	"gopkg.in/yaml.v3"

	_ "github.com/glebarez/go-sqlite"
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
	db, err := sql.Open("sqlite", "database.db")
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()
	defer func() { // Delete temporal database file
		if (runtime.GOOS == "linux" || runtime.GOOS == "unix") && Config.DatabaseFile != "database.db" {
			err = os.Remove("database.db")
			if err != nil {
				log.Println(err)
			}
		}
	}()
	const query string = `
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

	_, err = db.Exec(query)
	if err != nil {
		log.Println(err)
		return err
	}

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
		NotifyError("Error reading config file!", err)
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
				NotifyError("Error creating ~/.config/EduTrack/", err)
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
		NotifyError("Error marshalling new configuration file", err)
	}
	err = os.WriteFile(configFile, ymlData, os.ModePerm)
	if err != nil {
		NotifyError("Error writing config file", err)
	}
}
