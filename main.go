/*
 * Copyright Tom5521(c) - All Rights Reserved.
 *
 * This project is licenced under the MIT License.
 */

package main

import (
	"EduTrack/data"
	"EduTrack/iconloader"
	mgraph "EduTrack/ui/MGraph"
)

func main() {
	data.GetYamlData()
	iconloader.SetThemeIcons()
	mgraph.MainWindow()
}
