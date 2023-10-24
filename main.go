package main

import (
	"EduTrack/cmd/data"
	"EduTrack/internal/graph"
	icon "EduTrack/pkg/icons"
)

func main() {
	icon.SetThemeIcons()
	data.GetYamlData()
	graph.Init()
}
