/*
 * Copyright (c) 2023 Tom5521- All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package graph

import (
	"EduTrack/iconloader"
	"EduTrack/ui/sizes"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Define global variables
var (
	StundentList *widget.List
	GradesList   *widget.List = GetGradesList(Db.Grades)
)

// atoi converts a string to an integer, handling errors.
func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

// TemplateUser returns a container with user data.
func TemplateUser() *fyne.Container {
	// Create user data elements
	iconloader.SetThemeIcons(app.Settings().ThemeVariant())
	image := canvas.NewImageFromResource(iconloader.UserTemplateICON)
	image.SetMinSize(sizes.ProfileSize)
	dataLabel := widget.NewLabel(
		"Name: " + "--" + "\n" +
			"Age: " + "--" + "\n" +
			"ID: " + "--" + "\n" +
			"Phone number: " + "--",
	)
	content := container.NewVBox(image, dataLabel)
	return content
}

// Menu returns the main application menu.
func Menu() *fyne.MainMenu {
	// Create the main menu
	menu := fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("Load a config file", func() {
				//data.LoadConf(wins.FilePicker(app))
			}),
			fyne.NewMenuItem("Add Student", func() {
				AddStudentForm()
			}),
			fyne.NewMenuItem("Re-Save Changes", func() {
				//data.SaveStudentsData()
			})),
		fyne.NewMenu("Edit",
			fyne.NewMenuItem("Reload data", func() {
				//data.GetStundentData()
			}),
		),
		fyne.NewMenu("Help",
			fyne.NewMenuItem("About", func() {
				AboutWin()
			}),
		),
	)
	return menu
}
