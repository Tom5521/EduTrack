/*
 * Copyright (c) 2023 Tom5521- All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package assets

import "fyne.io/fyne/v2"

// Themed

var (
	App           fyne.Resource
	AddUser       fyne.Resource
	Dev           fyne.Resource
	Download      fyne.Resource
	Error         fyne.Resource
	Info          fyne.Resource
	Install       fyne.Resource
	Restart       fyne.Resource
	Save          fyne.Resource
	Uninstall     fyne.Resource
	UserTemplate  fyne.Resource
	Lens1         fyne.Resource
	ShowGrades    fyne.Resource
	DeleteGrade   fyne.Resource
	Plus          fyne.Resource
	Edit          fyne.Resource
	DeleteStudent fyne.Resource
	Cross         fyne.Resource
	Refresh       fyne.Resource
)

// Shared

var (
	Placeholder fyne.Resource
)

func Load() {
	ShowGrades = getResource("showGrades.png")
	Lens1 = getResource("lens1.png")
	UserTemplate = getResource("UserTemplate.png")
	Uninstall = getResource("Uninstall.png")
	Save = getResource("Save.png")
	Restart = getResource("Restart.png")
	Install = getResource("Install.png")
	Info = getResource("Info.png")
	Error = getResource("Error.png")
	Download = getResource("Download.png")
	Dev = getResource("Dev.png")
	AddUser = getResource("AddUser.png")
	App = getResource("Icon.png")
	DeleteGrade = getResource("DeleteGrade.png")
	Plus = getResource("plus.png")
	Edit = getResource("Edit.png")
	DeleteStudent = getResource("deleteStudent.png")
	Cross = getResource("cross.png")
	Refresh = getResource("refresh.png")

	Placeholder = getCmResource("Placeholder.png")
}
