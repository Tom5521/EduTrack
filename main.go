/*
 * Copyright Tom5521(c) - All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package main

import (
	"EduTrack/data"
	mgraph "EduTrack/ui/MGraph"
)

// The main function is the entry point for the application.
func main() {
	// Load data from YAML files.
	data.GetYamlData()

	// Set theme-specific icons for the application.

	// Open the main application window provided by the "MGraph" package.
	mgraph.MainWindow()
}
