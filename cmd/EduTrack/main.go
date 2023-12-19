/*
 * Copyright (c) 2023 Tom5521- All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package main

import (
	"fmt"

	"github.com/Tom5521/EduTrack/internal/gui"
	askpasswd "github.com/Tom5521/EduTrack/pkg/askPasswd"
	"github.com/Tom5521/EduTrack/pkg/conf"
	"github.com/Tom5521/EduTrack/pkg/data"
	"github.com/ncruces/zenity"
)

func main() {
	// Load the configuration and data files
	data.LoadFiles()
	// Ask password
	CheckPwd()
	// Init windows
	InitGUI()
}

func CheckPwd() {
	insertedPwd := askpasswd.AskPwd()
	originalHash := askpasswd.Hash(conf.Config.Password.Hash)
	err := originalHash.Compare(insertedPwd)
	if err != nil {
		if zenity.Error("Incorrect Passsword") != nil {
			fmt.Println("Error in error window")
		}
		panic(err)
	}
}

func InitGUI() {
	ui := gui.Init()
	ui.MainWin()
}
