/*
 * Copyright (c) 2023 Tom5521- All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package data

import (
	"log"

	"gorm.io/gorm"

	"github.com/Tom5521/EduTrack/pkg/conf"
	"github.com/glebarez/sqlite"
)

var (
	ConfDir = conf.GetConfDir()
)

func CreateDatabase() error {
	driver := sqlite.Open(conf.Config.DatabaseFile)
	db, err := gorm.Open(driver, &gorm.Config{})
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

	return err
}
