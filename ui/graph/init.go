/*
 * Copyright (c) 2023 Tom5521- All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package graph

import (
	"EduTrack/data"
	"EduTrack/ui/sizes"
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
	DB                         = &data.Db
)

// MainWindow is the main entry point of the application.
func MainWindow() {
	app.Settings().SetTheme(xtheme.AdwaitaTheme())
	mainWin := app.NewWindow("EduTrack")
	// MainWin.SetFullScreen(true)
	mainWin.SetMaster()
	mainWin.SetMainMenu(Menu())
	wintools.MaximizeWin(mainWin)

	searchButton := widget.NewButton("Search", func() {
		Search()
	})
	addButton := widget.NewButton("Add a student", func() {
		AddStudentForm()
	})
	addGradeButton := widget.NewButton("Show grades", func() {
		window := app.NewWindow("Grades")
		window.Resize(sizes.ListSize)
		content := container.NewStack(GradesList)

		window.SetContent(content)
		window.Show()
	})
	testButton := widget.NewButton("Add a grade", func() {
		AddGrade()
	})

	list := CreateStudentList(&DB.Students)
	const gridNumber int = 2
	buttonsGrid := container.NewAdaptiveGrid(gridNumber, searchButton, addButton, addGradeButton, testButton)
	vbox := container.NewVBox(buttonsGrid, widget.NewSeparator(), StudentTab)
	mainbox := container.NewHSplit(vbox, list)
	mainbox.SetOffset(0)
	mainWin.SetContent(mainbox)
	mainWin.ShowAndRun()
}
