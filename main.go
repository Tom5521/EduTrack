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

// The main function is the entry point for the application.
func main() {
	// Load the configuration and data files
	data.LoadFiles()
	// Open the main application window provided by the "MGraph" package.
	graph.MainWindow()
}
