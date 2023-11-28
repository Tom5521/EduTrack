/*
 * Copyright (c) 2023 Tom5521- All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package graph

import (
	assets "EduTrack/Assets"
	"EduTrack/data"
	"EduTrack/pkg/wins"
	"EduTrack/ui/sizes"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Search opens a window to search for a student by ID.
func Search() {
	w := app.NewWindow("Search Student")
	w.Resize(sizes.SearchSize)
	entry := widget.NewEntry()
	searchButton := widget.NewButton("Search", func() {
		studentDNI := entry.Text
		i := DB.FindStudentIndexByDNI(studentDNI)
		if i == -1 {
			wins.ErrWin(app, "Student not found!")
			return
		}
		student := DB.Students[i]
		LoadStudentInfo(&student)
		w.Close()
	})

	content := container.NewVBox(
		widget.NewLabel("Enter Student DNI:"),
		entry,
		searchButton,
	)

	w.SetContent(content)
	w.Show()
}

// existsID checks if a given string exists in a list of strings.
func existsID(check string, list []string) bool {
	var contains bool
	for _, v := range list {
		if v == check {
			contains = true
			break
		}
	}
	return contains
}

// checkValues checks if all required form fields are not empty.
func checkValues(s data.Student) bool {
	if s.Age == 0 {
		return false
	}
	if s.DNI == "" {
		return false
	}
	if s.PhoneNumber == "" {
		return false
	}
	if s.Name == "" {
		return false
	}
	return true
}

// AboutWin opens an "About" window to display information about the app.
func AboutWin() {
	window := app.NewWindow("About")
	label1 := widget.NewLabel("Created by:")
	link, _ := url.Parse("https://github.com/Tom5521")
	gitLabel := widget.NewHyperlink("Tom5521", link)
	licenceLabel := widget.NewLabel("Under the MIT license")
	authorCont := container.NewHBox(label1, gitLabel)
	logo := canvas.NewImageFromResource(assets.App)
	const w, h float32 = 300, 300
	logo.SetMinSize(fyne.NewSize(w, h))
	vbox1 := container.NewVBox(
		authorCont,
		licenceLabel,
		logo,
	)
	window.SetContent(vbox1)
	window.Show()
}
