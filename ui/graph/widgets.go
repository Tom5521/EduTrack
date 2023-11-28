/*
 * Copyright (c) 2023 Tom5521- All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package graph

import (
	assets "EduTrack/Assets"
	"EduTrack/pkg/wins"
	"EduTrack/ui/sizes"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Define global variables.
var (
	StundentList *widget.List
	GradesList   *widget.List = GetGradesList(DB.Grades)
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
	image := canvas.NewImageFromResource(assets.UserTemplate)
	image.SetMinSize(sizes.ProfileSize)
	dataForm := widget.NewForm(
		widget.NewFormItem("Name:", widget.NewLabel("--")),
		widget.NewFormItem("Age:", widget.NewLabel("--")),
		widget.NewFormItem("DNI:", widget.NewLabel("--")),
		widget.NewFormItem("Phone number:", widget.NewLabel("--")),
	)
	content := container.NewVBox(image, dataForm)
	return content
}

// Menu returns the main application menu.
func Menu() *fyne.MainMenu {
	// Create the main menu
	menu := fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("Add Student", func() {
				AddStudentForm()
			}),
		),
		fyne.NewMenu("Edit",
			fyne.NewMenuItem("Reload data", func() {
				err := DB.Update()
				if err != nil {
					wins.ErrWin(app, err.Error())
				}
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
