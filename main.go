package main

import (
	"EduTrack/cmd/data"
	"EduTrack/internal/graph"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	data.GetYamlData()
	app := app.New()
	mainWin := app.NewWindow("EduTrack")
	mainWin.SetMaster()
	mainWin.SetMainMenu(graph.Menu(app))
	//windowtools.MaximizeWin(mainWin)
	content := container.NewMax(graph.StudentList(app))
	mainWin.SetContent(content)
	mainWin.Show()

	app.Run()
}
