/*
 * Copyright (c) 2023 Tom5521- All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package main

import (
	"github.com/Tom5521/EduTrack/internal/gui"
	"github.com/Tom5521/EduTrack/pkg/data"
)

func main() {
	// Load the configuration and data files
	data.LoadFiles()
	// Init windows
	InitGUI()
}

func InitGUI() {
	ui := gui.Init()
	ui.MainWin()
}
