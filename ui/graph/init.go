/*
 * Copyright (c) 2023 Tom5521- All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package graph

import (
	"EduTrack/assets"
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
	app        fyne.App = fyne_app.New()
	StudentTab *fyne.Container
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
	var selected = -1
	list := CreateStudentList(&data.Students)
	list.OnSelected = func(id widget.ListItemID) {
		selected = id
	}

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(assets.AddUser, AddStudentForm),
		widget.NewToolbarAction(assets.DeleteStudent, func() {
			if selected == -1 {
				return
			}
			DeleteForm(&data.Students[selected])
		}),
		widget.NewToolbarAction(assets.Lens1, Search),
		widget.NewToolbarAction(assets.ShowGrades, GradesMainWin),
	)

	listContainer := container.NewBorder(toolbar, nil, nil, nil, list)
	ncontent := container.NewBorder(nil, nil, container.NewHBox(StudentTab, widget.NewSeparator()), nil, listContainer)

	downbox := container.NewHSplit(StudentTab, listContainer)
	downbox.SetOffset(0)
	mainWin.SetContent(ncontent)
	mainWin.ShowAndRun()
}
