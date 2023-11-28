/*
 * Copyright (c) 2023 Tom5521- All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package graph

import (
	"EduTrack/assets"
	"EduTrack/data"
	"EduTrack/ui/sizes"
	"EduTrack/ui/wintools"
	"fmt"

	"fyne.io/fyne/v2"
	fyne_app "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	xtheme "fyne.io/x/fyne/theme"
)

// basicVars is a struct to hold basic variables used in the application.

var (
	app        fyne.App = fyne_app.New()
	StudentTab *fyne.Container
	DB         *data.DBStr = &data.DB
)

// MainWindow is the main entry point of the application.
func MainWindow() {
	app.Settings().SetTheme(xtheme.AdwaitaTheme())
	assets.Load()

	StudentTab = TemplateUser()

	mainWin := app.NewWindow("EduTrack")
	// MainWin.SetFullScreen(true)
	mainWin.SetMaster()
	mainWin.SetMainMenu(Menu())
	wintools.MaximizeWin(mainWin)

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

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(assets.AddUser, func() {
			AddStudentForm()
		}),
		widget.NewToolbarAction(assets.Lens1, func() {
			Search()
		}),
	)

	list := CreateStudentList(&DB.Students)
	const gridNumber int = 4
	buttonsGrid := container.NewAdaptiveGrid(gridNumber, addGradeButton, testButton)
	fmt.Println(buttonsGrid)
	vbox := container.NewVBox( /*buttonsGrid, widget.NewSeparator(),*/ StudentTab)

	listContainer := container.NewVSplit(toolbar, list)
	listContainer.SetOffset(0)
	downbox := container.NewHSplit(vbox, listContainer)
	downbox.SetOffset(0)

	// mainbox := container.NewVSplit(toolbar, downbox)
	// mainbox.SetOffset(0)

	mainWin.SetContent(downbox)
	mainWin.ShowAndRun()
}
