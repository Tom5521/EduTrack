/*
 * Copyright Tom5521(c) - All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package mgraph

import (
	"EduTrack/data"
	"EduTrack/ui/wintools"
	"fmt"

	"fyne.io/fyne/v2"
	fyne_app "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	xtheme "fyne.io/x/fyne/theme"
)

// basicVars is a struct to hold basic variables used in the application.

var (
	app        fyne.App        = fyne_app.New()
	StudentTab *fyne.Container = TemplateUser()
)

// MainWindow is the main entry point of the application.
func MainWindow() {
	app.Settings().SetTheme(xtheme.AdwaitaTheme())
	window := app.NewWindow("EduTrack")
	window.SetMaster()
	window.SetMainMenu(Menu(app))
	wintools.MaximizeWin(window)

	searchButton := widget.NewButton("Search", func() {
		Search()
	})
	addButton := widget.NewButton("Add a student", func() {
		AddStudentForm(app)
	})

	list := CreateStudentList(&data.Students)
	vbox := container.NewVBox(searchButton, addButton, widget.NewSeparator(), StudentTab)
	mainbox := container.NewHSplit(vbox, list)
	mainbox.SetOffset(0)
	window.SetContent(mainbox)
	window.ShowAndRun()
}

// LoadStudentInfo loads information for a specific student.
func LoadStudentInfo(student *data.Student) {
	var Nlb = widget.NewLabel
	var Nhbx = container.NewHBox
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
	image := wintools.LoadProfileImg(student.ImageFilePath)

	dataContainer := container.NewVBox(NameCont, AgeCont, IDCont, PhoneCont)

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
