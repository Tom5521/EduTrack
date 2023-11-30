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
	"EduTrack/ui/wintools"
	"fmt"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// LoadStudentInfo loads information for a specific student.
func LoadStudentInfo(student *data.Student) {
	var nlb = widget.NewLabel
	var nhbx = container.NewHBox

	image := wintools.LoadProfileImg(student.ImageFilePath)

	// Name Label
	nameLabel := nlb("Name: ")
	nameLabel.TextStyle.Bold = true
	nameCont := nhbx(nameLabel, nlb(student.Name))
	// Age Label
	ageLabel := nlb("Age: ")
	ageLabel.TextStyle.Bold = true
	ageCont := nhbx(ageLabel, nlb(strconv.Itoa(int(student.Age))))
	// ID Label
	idLabel := nlb("DNI: ")
	idLabel.TextStyle.Bold = true
	idCont := nhbx(idLabel, nlb(student.DNI))
	// Phone Label
	phoneLabel := nlb("Phone Number: ")
	phoneLabel.TextStyle.Bold = true
	phoneCont := nhbx(phoneLabel, nlb(student.PhoneNumber))

	gradesLabel := nlb("Grades:")
	gradesLabel.TextStyle.Bold = true
	getGrades := func() string {
		d := student.GetGradesNames()
		p := strings.Join(d, ",")
		return p
	}
	gradesCont := nhbx(gradesLabel, nlb(getGrades()))

	dataContainer := container.NewVBox(nameCont, ageCont, idCont, phoneCont, gradesCont)
	editButton := widget.NewButton("Edit", func() {
		EditFormWindow(student)
	})

	deleteButton := widget.NewButton("Delete", func() {
		DeleteForm(student)
	})
	recordButton := widget.NewButton("Add a record", func() {
		AddRecord(student)
	})
	showRecordsBt := widget.NewButton("Show records", func() {
		ShowRecords(student)
	})

	const gridNumber int = 2
	content := container.NewVBox(image,
		dataContainer,
		container.NewAdaptiveGrid(gridNumber,
			editButton,
			deleteButton,
			recordButton,
			showRecordsBt,
		),
	)
	StudentTab.Objects = []fyne.CanvasObject{content}
}

// EditFormWindow opens a window to edit a student's information.
func EditFormWindow(student *data.Student) {
	window := app.NewWindow("Edit " + student.Name)
	window.Resize(sizes.FormSize)

	// Initialize form fields
	var imagePath = student.ImageFilePath
	nameEntry := widget.NewEntry()
	nameEntry.SetText(student.Name)
	ageEntry := widget.NewEntry()
	ageEntry.SetText(strconv.Itoa(int(student.Age)))
	dniEntry := widget.NewEntry()
	dniEntry.SetText(student.DNI)
	phoneEntry := widget.NewEntry()
	phoneEntry.SetText(student.PhoneNumber)

	imageLabel := widget.NewLabel(imagePath)

	nameForm := widget.NewFormItem("Name:", nameEntry)
	ageForm := widget.NewFormItem("Age:", ageEntry)
	dniForm := widget.NewFormItem("DNI:", dniEntry)

	phoneForm := widget.NewFormItem("Phone:", phoneEntry)

	deleteImgButton := widget.NewButton("Delete selected image", func() {
		imagePath = ""
		imageLabel.SetText(imagePath)
	})
	selectImgButton := widget.NewButton("Select student image", func() {
		wins.ImagePicker(app, &imagePath)
		imageLabel.SetText(imagePath)
	})

	const gridNumber int = 2
	imgCont := container.NewAdaptiveGrid(gridNumber, deleteImgButton, selectImgButton)
	imgForm := widget.NewFormItem("", imgCont)
	imgInfoForm := widget.NewFormItem("Image path:", imageLabel)

	form := widget.NewForm(
		nameForm,
		ageForm,
		dniForm,
		phoneForm,
		imgForm,
		imgInfoForm,
	)
	form.OnSubmit = func() {
		// Validate form fields
		if !checkValues(data.Student{
			Age:         uint(atoi(ageEntry.Text)),
			DNI:         dniEntry.Text,
			PhoneNumber: phoneEntry.Text,
			Name:        nameEntry.Text,
		}) {
			wins.ErrWin(app, "Some value in the form is empty")
			return
		}
		if dniEntry.Text != student.DNI {
			if existsID(dniEntry.Text, data.GetStudentDNIs()) {
				wins.ErrWin(app, "The DNI already exists")
				return
			}
		}
		err := student.Edit(data.Student{
			Name:          nameEntry.Text,
			Age:           uint(atoi(ageEntry.Text)),
			DNI:           dniEntry.Text,
			PhoneNumber:   phoneEntry.Text,
			ImageFilePath: imagePath,
		})
		if err != nil {
			wins.ErrWin(app, err.Error())
		}
		StundentList.Refresh()
		LoadStudentInfo(&data.Students[data.FindStudentIndexByID(student.ID)])
		window.Close()
	}

	window.SetContent(form)
	window.Show()
}

// DeleteForm opens a confirmation window to delete a student.
func DeleteForm(student *data.Student) {
	window := app.NewWindow("Delete Student")
	const gridNumber int = 2
	content := container.NewVBox(
		widget.NewLabel("Are you sure you want to delete the student?"),
		container.NewAdaptiveGrid(gridNumber,
			widget.NewButton("Yes", func() {
				err := data.Delete(student)
				if err != nil {
					wins.ErrWin(app, err.Error())
				}
				window.Close()
			}),
			widget.NewButton("No", func() {
				window.Close()
			}),
		))
	window.SetContent(content)
	window.Show()
}

// AddStudentForm opens a window to add a new student.
func AddStudentForm() {
	var imagePath string
	var gradesStr []data.Grade
	window := app.NewWindow("Add a student")
	window.Resize(sizes.FormSize)

	// Initialize form fields
	nameEntry := widget.NewEntry()
	ageEntry := widget.NewEntry()
	dniEntry := widget.NewEntry()
	phoneEntry := widget.NewEntry()

	imageButton := widget.NewButton("Select Image", func() {
		wins.ImagePicker(app, &imagePath)
	})
	deleteImgBtn := widget.NewButton("Delete Current Image", func() {
		imagePath = ""
	})

	getStGrade := func() string {
		var grades string
		for _, g := range gradesStr {
			grades += g.Name + ","
		}
		return grades
	}

	studentGradesLabel := widget.NewLabel("")
	studentGradesLabel.SetText(getStGrade())
	gradeSelect := widget.NewSelect(data.GetGradesNames(), func(s string) {
		gradesStr = append(gradesStr, data.FindGradeByName(s))

		studentGradesLabel.SetText(getStGrade())
	})

	nameForm := widget.NewFormItem("Name:", nameEntry)
	idForm := widget.NewFormItem("DNI:", dniEntry)
	ageForm := widget.NewFormItem("Age:", ageEntry)
	phoneForm := widget.NewFormItem("Phone:", phoneEntry)
	gradeForm := widget.NewFormItem("Select Grades:", gradeSelect)
	gradesShowForm := widget.NewFormItem("Grades:", studentGradesLabel)
	const gridNumber int = 2
	imageForm := widget.NewFormItem("Image:", container.NewAdaptiveGrid(gridNumber, imageButton, deleteImgBtn))

	form := widget.NewForm(
		nameForm,
		idForm,
		ageForm,
		phoneForm,
		gradeForm,
		gradesShowForm,
		imageForm,
	)
	form.OnSubmit = func() {
		// Validate form fields
		if !checkValues(data.Student{
			Age:         uint(atoi(ageEntry.Text)),
			DNI:         dniEntry.Text,
			PhoneNumber: phoneEntry.Text,
			Name:        nameEntry.Text,
		}) {
			wins.ErrWin(app, "Some value in the form is empty")
			return
		}
		if existsID(dniEntry.Text, data.GetStudentDNIs()) {
			wins.ErrWin(app, "The DNI already exists")
			return
		}
		StGrades := func() []data.StudentGrade {
			var stgrades []data.StudentGrade
			for _, grade := range gradesStr {
				tmpGrade := data.StudentGrade{GradeID: grade.ID}
				stgrades = append(stgrades, tmpGrade)
			}
			return stgrades
		}()

		// Add a new student
		newStudent := data.Student{
			Name:          nameEntry.Text,
			Age:           uint(atoi(ageEntry.Text)),
			DNI:           dniEntry.Text,
			PhoneNumber:   phoneEntry.Text,
			ImageFilePath: imagePath,
			Grades:        StGrades,
		}
		err := data.AddStudent(&newStudent)
		if err != nil {
			wins.ErrWin(app, err.Error())
		}
		fmt.Println(gradesStr)
		fmt.Println(data.Students[len(data.Students)-1])
		StundentList.Refresh()
		s := data.Students[data.FindStudentIndexByID(newStudent.ID)]
		LoadStudentInfo(&s)
		window.Close()
	}
	window.SetContent(form)
	window.Show()
}

// CreateStudentList creates a list of students and their names.
func CreateStudentList(students *[]data.Student) fyne.Widget {
	// Initialize the student list widget
	StundentList = widget.NewList(
		func() int {
			return len(*students)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("---")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			student2 := *students
			o.(*widget.Label).SetText(student2[i].Name)
		},
	)

	// Handle item selection
	StundentList.OnSelected = func(id widget.ListItemID) {
		d := *students
		StundentList.UnselectAll()
		LoadStudentInfo(&d[id])
		StundentList.Refresh()
	}

	return StundentList
}
