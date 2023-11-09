/*
 * Copyright Tom5521(c) - All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package graph

import (
	"EduTrack/data"
	"EduTrack/pkg/wins"
	"EduTrack/ui/sizes"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func GetRegisterList(student *data.Student) {
	list := widget.NewList(
		func() int {
			return len(student.Register)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(student.Register[i].Name)
		},
	)
	list.OnSelected = func(id widget.ListItemID) {
		list.UnselectAll()
		EditRegisterData(student, id)
	}
	RegisterList = list
}

func GetGradesList(grades *[]data.Grade) *widget.List {
	list := widget.NewList(
		func() int {
			return len(*grades)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			mod := *grades
			o.(*widget.Label).SetText(mod[i].Name)
		},
	)

	list.OnSelected = func(id widget.ListItemID) {
		list.UnselectAll()
	}

	return list
}

func AddGrade() {
	window := app.NewWindow("Add a grade")
	window.Resize(sizes.FormSize)
	gradeEntry := widget.NewEntry()
	priceEntry := widget.NewEntry()
	InfoEntry := widget.NewMultiLineEntry()
	StartEntry := widget.NewEntry()
	EndEntry := widget.NewEntry()

	gradeFormInput := widget.NewFormItem("Grade name:", gradeEntry)
	priceFormInput := widget.NewFormItem("Price per moth:", priceEntry)
	infoFormInput := widget.NewFormItem("Grade Info:", InfoEntry)
	StartFormInput := widget.NewFormItem("Start date:", StartEntry)
	EndFormInput := widget.NewFormItem("End date:", EndEntry)

	form := widget.NewForm(
		gradeFormInput,
		priceFormInput,
		StartFormInput,
		EndFormInput,
		infoFormInput,
	)
	form.OnSubmit = func() {
		if gradeEntry.Text == "" {
			wins.ErrWin(app, "Grade name entry is empty")
			return
		}
		if priceEntry.Text == "" {
			wins.ErrWin(app, "Info entry is empty")
			return
		}
		if StartEntry.Text == "" {
			wins.ErrWin(app, "Start date entry is empty")
			return
		}
		if EndEntry.Text == "" {
			wins.ErrWin(app, "End date entry is emty")
			return
		}
		if strings.Contains(strings.Join(data.GetGradesNames(), " "), gradeEntry.Text) {
			wins.ErrWin(app, "This grade already exists!")
			return
		}

		data.Grades = append(data.Grades, data.Grade{
			Name:  gradeEntry.Text,
			Info:  InfoEntry.Text,
			Start: StartEntry.Text,
			End:   EndEntry.Text,
			Price: priceEntry.Text,
		})

		data.SaveGradesData()
		window.Close()

	}

	content := container.NewVBox(form)
	window.SetContent(content)
	window.Show()
}
