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
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func GetRegisterList(student *data.Student) {
	list := widget.NewList(
		func() int {
			return len(student.Registers)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(student.Registers[i].Name)
		},
	)
	list.OnSelected = func(id widget.ListItemID) {
		list.UnselectAll()
		EditRegisterData(student, id)
	}
	RegisterList = list
}

func GetGradesList(grades *[]*data.Grade) *widget.List {
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
		g := *grades
		GradeDetailsWin(g[id])
	}

	return list
}

func EditGrade() {}

func GradeDetailsWin(g *data.Grade) {

}

func StudentGradeDetailsWin(g *data.StudentGrade) {
	window := app.NewWindow("Details for " + g.Name)

	gradeNameLabel := widget.NewLabel(g.Name)
	gradePricePMLabel := widget.NewLabel(g.Price)
	gradeStartLabel := widget.NewLabel(g.Start)
	gradeEndLabel := widget.NewLabel(g.End)
	gradeInfoLabel := widget.NewMultiLineEntry()
	gradeInfoLabel.SetText(g.Info)
	gradeInfoLabel.Disable()

	editGradeButton := widget.NewButton("Edit Grade", func() {})
	editStudentButton := widget.NewButton("Edit Student", func() {})

	NameForm := widget.NewFormItem("Name:", gradeNameLabel)
	PriceForm := widget.NewFormItem("Price:", gradePricePMLabel)
	StartForm := widget.NewFormItem("Start:", gradeStartLabel)
	EndForm := widget.NewFormItem("End:", gradeEndLabel)
	InfoForm := widget.NewFormItem("Info:", gradeInfoLabel)
	EditForm := widget.NewFormItem("", container.NewAdaptiveGrid(2, editGradeButton, editStudentButton))

	Form := widget.NewForm(
		NameForm,
		PriceForm,
		StartForm,
		EditForm,
		EndForm,
		InfoForm,
		EditForm,
	)

	window.SetContent(Form)
	window.Show()
}

func AddGrade() {
	window := app.NewWindow("Add a grade")
	window.Resize(sizes.FormSize)
	gradeEntry := widget.NewEntry()
	priceEntry := widget.NewEntry()
	InfoEntry := widget.NewMultiLineEntry()

	gradeFormInput := widget.NewFormItem("Grade name:", gradeEntry)
	priceFormInput := widget.NewFormItem("Price per moth:", priceEntry)
	infoFormInput := widget.NewFormItem("Grade Info:", InfoEntry)

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
		if strings.Contains(strings.Join(Db.GetGradesNames(), " "), gradeEntry.Text) {
			wins.ErrWin(app, "This grade already exists!")
			return
		}

		Db.AddGrade(data.Grade{
			Name:  gradeEntry.Text,
			Info:  InfoEntry.Text,
			Price: priceEntry.Text,
		})
		window.Close()
	}
	content := container.NewVBox(form)
	window.SetContent(content)
	window.Show()
}
