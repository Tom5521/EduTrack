/*
 * Copyright (c) 2023 Tom5521- All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package graph

import (
	"EduTrack/data"
	"EduTrack/pkg/wins"
	"EduTrack/ui/sizes"
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func GetRecordsList(student *data.Student) *widget.List {
	list := widget.NewList(
		func() int {
			return len(student.Records)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(student.Records[i].Name)
		},
	)
	list.OnSelected = func(id widget.ListItemID) {
		list.UnselectAll()
		EditRecordsData(student, id)
	}
	return list
}

func GetGradesList(grades []data.Grade) *widget.List {
	list := widget.NewList(
		func() int {
			return len(grades)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			mod := grades
			o.(*widget.Label).SetText(mod[i].Name)
		},
	)

	list.OnSelected = func(id widget.ListItemID) {
		list.UnselectAll()
		g := grades[id]
		GradeDetailsWin(&g)
	}

	return list
}

func EditGrade() {}

func GradeDetailsWin(g *data.Grade) {
	fmt.Println(g)
}

func StudentGradeDetailsWin(sg *data.StudentGrade) {
	getGrade := func() *data.Grade {
		i := DB.FindGradeIndexByID(sg.GradeID)
		return &DB.Grades[i]
	}
	g := getGrade()

	window := app.NewWindow("Details for " + g.Name)

	gradeNameLabel := widget.NewLabel(g.Name)
	gradePricePMLabel := widget.NewLabel(g.Price)
	gradeStartLabel := widget.NewLabel(sg.Start)
	gradeEndLabel := widget.NewLabel(sg.End)
	gradeInfoLabel := widget.NewMultiLineEntry()
	gradeInfoLabel.SetText(g.Info)
	gradeInfoLabel.Disable()

	editGradeButton := widget.NewButton("Edit Grade", func() {})
	editStudentButton := widget.NewButton("Edit Student", func() {})

	nameForm := widget.NewFormItem("Name:", gradeNameLabel)
	priceForm := widget.NewFormItem("Price:", gradePricePMLabel)
	startForm := widget.NewFormItem("Start:", gradeStartLabel)
	endForm := widget.NewFormItem("End:", gradeEndLabel)
	infoForm := widget.NewFormItem("Info:", gradeInfoLabel)
	const gridNumber int = 2
	editForm := widget.NewFormItem("",
		container.NewAdaptiveGrid(gridNumber,
			editGradeButton,
			editStudentButton,
		),
	)

	form := widget.NewForm(
		nameForm,
		priceForm,
		startForm,
		editForm,
		endForm,
		infoForm,
		editForm,
	)

	window.SetContent(form)
	window.Show()
}

func AddGrade() {
	window := app.NewWindow("Add a grade")
	window.Resize(sizes.FormSize)
	gradeEntry := widget.NewEntry()
	priceEntry := widget.NewEntry()
	infoEntry := widget.NewMultiLineEntry()

	gradeFormInput := widget.NewFormItem("Grade name:", gradeEntry)
	priceFormInput := widget.NewFormItem("Price per moth:", priceEntry)
	infoFormInput := widget.NewFormItem("Grade Info:", infoEntry)

	form := widget.NewForm(
		gradeFormInput,
		priceFormInput,
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
		if strings.Contains(strings.Join(DB.GetGradesNames(), " "), gradeEntry.Text) {
			wins.ErrWin(app, "This grade already exists!")
			return
		}

		_, err := DB.AddGrade(data.Grade{
			Name:  gradeEntry.Text,
			Info:  infoEntry.Text,
			Price: priceEntry.Text,
		})
		if err != nil {
			wins.ErrWin(app, err.Error())
		}
		window.Close()
	}
	content := container.NewVBox(form)
	window.SetContent(content)
	window.Show()
}

func GradesMainWin() {

}
