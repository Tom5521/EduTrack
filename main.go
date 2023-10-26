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
