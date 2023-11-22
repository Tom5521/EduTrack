/*
 * Copyright (c) 2023 Tom5521- All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package main

import (
	"EduTrack/data"
	"EduTrack/ui/graph"
)

func main() {
	// Load the configuration and data files
	data.LoadFiles()
	// Init windows
	graph.MainWindow()
}
