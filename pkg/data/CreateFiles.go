/*
 * Copyright (c) 2023 Tom5521- All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package data

import (
	"log"
	"os"
	"runtime"

	"gorm.io/gorm"

	"github.com/Tom5521/EduTrack/pkg/conf"
	"github.com/Tom5521/EduTrack/pkg/files"
	"github.com/glebarez/sqlite"
)

var (
	Config  conf.Config
	ConfDir = conf.GetConfDir()
)

func CreateDatabase() error {
	db, err := gorm.Open(sqlite.Open(Config.DatabaseFile), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return err
	}

	err = db.AutoMigrate(&Student{})
	printErr(err)
	err = db.AutoMigrate(&Course{})
	printErr(err)
	err = db.AutoMigrate(&StudentCourse{})
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

	_, err = files.CopyFile("database.db", Config.DatabaseFile)
	if err != nil {
		log.Println(err)
	}
	return err
}
