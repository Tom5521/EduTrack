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
	"EduTrack/ui/wintools"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// LoadStudentInfo loads information for a specific student.
func LoadStudentInfo(student *data.Student) {
	var Nlb = widget.NewLabel
	var Nhbx = container.NewHBox

	image := wintools.LoadProfileImg(student.ImageFilePath)

	// Name Label
	NameLabel := Nlb("Name: ")
	NameLabel.TextStyle.Bold = true
	NameCont := Nhbx(NameLabel, Nlb(student.Name))
	// Age Label
	AgeLabel := Nlb("Age: ")
	AgeLabel.TextStyle.Bold = true
	AgeCont := Nhbx(AgeLabel, Nlb(fmt.Sprintf("%v", student.Age)))
	// ID Label
	IDLabel := Nlb("ID: ")
	IDLabel.TextStyle.Bold = true
	IDCont := Nhbx(IDLabel, Nlb(student.ID))
	// Phone Label
	PhoneLabel := Nlb("Phone Number: ")
	PhoneLabel.TextStyle.Bold = true
	PhoneCont := Nhbx(PhoneLabel, Nlb(student.Phone_number))

	GradesLabel := Nlb("Grades:")
	GradesLabel.TextStyle.Bold = true
	GradesCont := Nhbx(GradesLabel, Nlb(fmt.Sprint(student.Grades)))

	dataContainer := container.NewVBox(NameCont, AgeCont, IDCont, PhoneCont, GradesCont)
	editButton := widget.NewButton("Edit", func() {
		EditFormWindow(student)
	})

	deleteButton := widget.NewButton("Delete", func() {
		DeleteForm(student)
	})
	registerButton := widget.NewButton("Add register", func() {
		AddRegister(student)
	})
	ShowRegistersBt := widget.NewButton("Show Registers", func() {
		ShowRegisters(student)
	})

	content := container.NewVBox(image,
		dataContainer,
		container.NewAdaptiveGrid(2,
			editButton,
			deleteButton,
			registerButton,
			ShowRegistersBt,
		),
	)
	StudentTab.Objects = []fyne.CanvasObject{content}
}

// EditFormWindow opens a window to edit a student's information.
func EditFormWindow(student *data.Student) {
	window := app.NewWindow("Edit " + student.Name)
	window.Resize(sizes.FormSize)

	// Initialize form fields
	var imagePath string = student.ImageFilePath
	nameEntry := widget.NewEntry()
	nameEntry.SetText(student.Name)
	ageEntry := widget.NewEntry()
	ageEntry.SetText(fmt.Sprintf("%v", student.Age))
	idEntry := widget.NewEntry()
	idEntry.SetText(student.ID)
	phoneEntry := widget.NewEntry()
	phoneEntry.SetText(student.Phone_number)

	saveEdited := func() {
		// Validate form fields
		if !checkValues(formReturn{NameEntry: nameEntry, AgeEntry: ageEntry, IDEntry: idEntry, PhoneEntry: phoneEntry}) {
			wins.ErrWin(app, "Some value in the form is empty")
			return
		}

		if existsId(idEntry.Text, data.GetStudentIDs()) && idEntry.Text != student.ID {
			wins.ErrWin(app, "The ID already exists")
			return
		}

		// Update student information
		student.Age = atoi(ageEntry.Text)
		student.Name = nameEntry.Text
		student.Phone_number = phoneEntry.Text
		student.ID = idEntry.Text
		student.ImageFilePath = imagePath
		data.SaveStudentsData()
		data.GetStundentData()
		StundentList.Refresh()
		LoadStudentInfo(student)
		window.Close()
	}

	retForm := formReturn{
		NameEntry:  nameEntry,
		AgeEntry:   ageEntry,
		IDEntry:    idEntry,
		PhoneEntry: phoneEntry,
		ExecFunc:   saveEdited,
		ImagePath:  &imagePath,
	}

	content := GetForm(&retForm)

	window.SetContent(content)
	window.Show()
}

// DeleteForm opens a confirmation window to delete a student.
func DeleteForm(student *data.Student) {
	window := app.NewWindow("Delete Student")
	content := container.NewVBox(
		widget.NewLabel("Are you sure you want to delete the student?"),
		container.NewAdaptiveGrid(2,
			widget.NewButton("Yes", func() {
				for i, s := range data.Students {
					if s.ID == student.ID {
						data.Students = append(data.Students[:i], data.Students[i+1:]...)
						data.SaveStudentsData()
						data.GetStundentData()
						break
					}
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
	window := app.NewWindow("Add a student")
	window.Resize(sizes.FormSize)

	// Initialize form fields
	nameEntry := widget.NewEntry()
	ageEntry := widget.NewEntry()
	idEntry := widget.NewEntry()
	phoneEntry := widget.NewEntry()

	submitFunc := func() {
		// Validate form fields
		if !checkValues(formReturn{NameEntry: nameEntry, AgeEntry: ageEntry, IDEntry: idEntry, PhoneEntry: phoneEntry}) {
			wins.ErrWin(app, "Some value in the form is empty")
			return
		}
		if existsId(idEntry.Text, data.GetStudentIDs()) {
			wins.ErrWin(app, "The ID already exists")
			return
		}

		// Add a new student
		data.Students = append(data.Students, data.Student{
			Name:          nameEntry.Text,
			Age:           atoi(ageEntry.Text),
			ID:            idEntry.Text,
			Phone_number:  phoneEntry.Text,
			ImageFilePath: imagePath,
		})
		data.SaveStudentsData()
		data.GetStundentData()
		StundentList.Refresh()
		LoadStudentInfo(data.FindStudentByID(idEntry.Text))
		window.Close()
	}

	formRet := formReturn{
		ExecFunc:   submitFunc,
		NameEntry:  nameEntry,
		IDEntry:    idEntry,
		AgeEntry:   ageEntry,
		PhoneEntry: phoneEntry,
		ImagePath:  &imagePath,
	}
	content := GetForm(&formRet)
	window.SetContent(content)
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
			student_ := *students
			o.(*widget.Label).SetText(student_[i].Name)
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

// GetForm returns a container with form elements.
func GetForm(d *formReturn) *fyne.Container {
	// Create buttons
	imageButton := widget.NewButton("Select Image", func() {
		wins.ImagePicker(app, d.ImagePath)
	})
	deleteImgBtn := widget.NewButton("Delete Current Image", func() {
		*d.ImagePath = ""
	})

	getStGrade := func() string {
		var grades string
		for _, g := range d.StundentGrades {
			grades += g.Name + ","
		}
		return grades
	}

	studentGradesLabel := widget.NewLabel("")
	studentGradesLabel.SetText(getStGrade())

	gradeSelect := widget.NewSelect(data.GetGradesNames(), func(s string) {
		d.StundentGrades = append(d.StundentGrades, *data.FindGradeByName(s))
		studentGradesLabel.SetText(getStGrade())
	})

	NameFormEntry := widget.NewFormItem("Name", d.NameEntry)
	AgeFormEntry := widget.NewFormItem("Age", d.AgeEntry)
	IDFormEntry := widget.NewFormItem("ID", d.IDEntry)
	Phone_numberFormEntry := widget.NewFormItem("Phone number", d.PhoneEntry)
	selectGradeForm := widget.NewFormItem("Select grade", gradeSelect)
	imageSelectForm := widget.NewFormItem("Stundet Image:", container.NewAdaptiveGrid(2, imageButton, deleteImgBtn))
	GradesFormLabel := widget.NewFormItem("Grades:", studentGradesLabel)
	// Create the form
	form := widget.NewForm(
		NameFormEntry,
		AgeFormEntry,
		IDFormEntry,
		Phone_numberFormEntry,
		selectGradeForm,
		GradesFormLabel,
		imageSelectForm,
	)
	form.OnSubmit = d.ExecFunc

	// Create the content container
	content := container.NewVBox(form)
	return content
}

// formReturn represents a structure for form elements.
type formReturn struct {
	ExecFunc       func()
	NameEntry      *widget.Entry
	AgeEntry       *widget.Entry
	IDEntry        *widget.Entry
	PhoneEntry     *widget.Entry
	ImagePath      *string
	StundentGrades []data.Grade
}
