/*
 * Copyright (c) 2023 Tom5521- All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package graph

import (
	"EduTrack/assets"
	"EduTrack/data"
	"EduTrack/pkg/wins"
	"EduTrack/ui/sizes"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

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

	return list
}

func EditGrade(g *data.Grade) {
	window := app.NewWindow("Edit " + g.Name)

	nameEntry := widget.NewEntry()
	nameEntry.SetText(g.Name)

	priceEntry := widget.NewEntry()
	priceEntry.SetText(g.Price)

	infoEntry := widget.NewMultiLineEntry()
	infoEntry.SetText(g.Info)

	form := widget.NewForm(
		widget.NewFormItem("Name:", nameEntry),
		widget.NewFormItem("Price:", priceEntry),
		widget.NewFormItem("Info:", infoEntry),
	)

	form.OnSubmit = func() {
		newGrade := data.Grade{
			Name:  nameEntry.Text,
			Price: priceEntry.Text,
			Info:  infoEntry.Text,
		}
		err := g.Edit(newGrade)
		if err != nil {
			log.Println(err)
			wins.ErrWin(app, err.Error())
		}
		window.Close()
	}

	window.SetContent(form)

	window.Show()
}

func GetGradeDetailsCont(g *data.Grade, window fyne.Window) *fyne.Container {
	editButton := widget.NewButton("Edit", func() {
		EditGrade(g)
		window.Close()
	})
	deleteButton := widget.NewButton("Delete", func() {
		err := data.Delete(g)
		if err != nil {
			wins.ErrWin(app, err.Error())
		}
		GradesList = GetGradesList(&data.Grades)
		GradesList.Refresh()
		window.Close()
	})
	const gridNumber int = 2

	form := widget.NewForm(
		widget.NewFormItem("Name:", widget.NewLabel(g.Name)),
		widget.NewFormItem("Price:", widget.NewLabel(g.Price)),
		widget.NewFormItem("Info", widget.NewLabel(g.Info)),
		widget.NewFormItem("", container.NewAdaptiveGrid(gridNumber, deleteButton, editButton)),
	)
	return container.NewStack(form)
}

func GradeDetailsWin(g *data.Grade) {
	window := app.NewWindow(g.Name)

	form := GetGradeDetailsCont(g, window)

	window.SetContent(form)
	window.Show()
}

func StudentGradeDetailsWin(sg *data.StudentGrade) {
	getGrade := func() *data.Grade {
		i := data.FindGradeIndexByID(sg.GradeID)
		return &data.Grades[i]
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
		if func() bool {
			for _, grade := range data.Grades {
				if grade.Name == gradeEntry.Text {
					return true
				}
			}
			return false
		}() {
			wins.ErrWin(app, "This grade already exists!")
			return
		}
		newGrade := data.Grade{
			Name:  gradeEntry.Text,
			Info:  infoEntry.Text,
			Price: priceEntry.Text,
		}
		err := data.AddGrade(&newGrade)
		if err != nil {
			wins.ErrWin(app, err.Error())
		}
		GradesList = GetGradesList(&data.Grades)
		window.Close()
	}
	content := container.NewVBox(form)
	window.SetContent(content)
	window.Show()
}

func GradesMainWin() {
	w := app.NewWindow("Grades")

	var selected = -1

	list := GetGradesList(&data.Grades)

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(assets.DeleteGrade, func() {
			if selected == -1 {
				return
			}
			err := data.Delete(data.Grades[selected])
			if err != nil {
				wins.ErrWin(app, err.Error())
				return
			}
			selected = -1
		}),
		widget.NewToolbarAction(assets.Plus, func() { AddGrade() }),
		widget.NewToolbarAction(assets.ShowGrades, func() { GradeDetailsWin(&data.Grades[selected]) }),
		widget.NewToolbarAction(assets.Edit, func() { EditGrade(&data.Grades[selected]) }),
	)

	list.OnSelected = func(id widget.ListItemID) {
		selected = id
	}

	ncontent := container.NewBorder(toolbar, nil, nil, nil, list)
	//content := container.NewStack(ncontent, list)
	//content.SetOffset(0)
	w.SetIcon(assets.ShowGrades)
	w.SetContent(ncontent)

	w.Show()
}
