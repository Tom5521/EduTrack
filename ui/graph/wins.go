/*
 * Copyright Tom5521(c) - All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package graph

import (
	"EduTrack/data"
	"EduTrack/iconloader"
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
		studentID := entry.Text
		student := data.FindStudentByID(studentID)
		if student != nil {
			LoadStudentInfo(student)
			w.Close()
		} else {
			wins.ErrWin(app, "Student not found!")
		}
	})

	content := container.NewVBox(
		widget.NewLabel("Enter Student ID:"),
		entry,
		searchButton,
	)

	w.SetContent(content)
	w.Show()
}

// existsId checks if a given string exists in a list of strings.
func existsId(check string, list []string) bool {
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
func checkValues(d formReturn) bool {
	if d.AgeEntry.Text == "" {
		return false
	}
	if d.IDEntry.Text == "" {
		return false
	}
	if d.PhoneEntry.Text == "" {
		return false
	}
	if d.NameEntry.Text == "" {
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
	LicenceLabel := widget.NewLabel("Under the MIT license")
	AuthorCont := container.NewHBox(label1, gitLabel)
	logo := canvas.NewImageFromResource(iconloader.AppICON)
	logo.SetMinSize(fyne.NewSize(300, 300))
	vbox1 := container.NewVBox(
		AuthorCont,
		LicenceLabel,
		logo,
	)
	window.SetContent(vbox1)
	window.Show()
}
