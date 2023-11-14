/*
 * Copyright Tom5521(c) - All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package graph

import (
	"EduTrack/data"
	"EduTrack/ui/wintools"

	"fyne.io/fyne/v2"
	fyne_app "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	xtheme "fyne.io/x/fyne/theme"
)

// basicVars is a struct to hold basic variables used in the application.

var (
	app        fyne.App        = fyne_app.New()
	StudentTab *fyne.Container = TemplateUser()
	Data                       = &data.Data
)

// MainWindow is the main entry point of the application.
func MainWindow() {
	app.Settings().SetTheme(xtheme.AdwaitaTheme())
	MainWin := app.NewWindow("EduTrack")
	//MainWin.SetFullScreen(true)
	MainWin.SetMaster()
	MainWin.SetMainMenu(Menu())
	wintools.MaximizeWin(MainWin)

	searchButton := widget.NewButton("Search", func() {
		Search()
	})
	addButton := widget.NewButton("Add a student", func() {
		AddStudentForm()
	})
	AddGradeButton := widget.NewButton("Show grades", func() {
		window := app.NewWindow("Grades")

		content := container.NewStack(GradesList)

		window.SetContent(content)
		window.Show()
	})
	testButton := widget.NewButton("Add a grade", func() {
		AddGrade()
	})

	list := CreateStudentList(&Data.Students)
	buttonsGrid := container.NewAdaptiveGrid(2, searchButton, addButton, AddGradeButton, testButton)
	vbox := container.NewVBox(buttonsGrid, widget.NewSeparator(), StudentTab)
	mainbox := container.NewHSplit(vbox, list)
	mainbox.SetOffset(0)
	MainWin.SetContent(mainbox)
	MainWin.ShowAndRun()
}
