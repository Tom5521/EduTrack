package mgraph

import (
	"EduTrack/data"
	"EduTrack/ui/wintools"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	xtheme "fyne.io/x/fyne/theme"
)

type basicVars struct {
	app  *fyne.App
	cont *fyne.Container
}

func MainWindow() {
	app := app.New()
	app.Settings().SetTheme(xtheme.AdwaitaTheme())
	var StudentTab *fyne.Container = TemplateUser()
	window := app.NewWindow("EduTrack")
	window.SetMaster()
	window.SetMainMenu(Menu(app))
	wintools.MaximizeWin(window)

	listBasic := basicVars{
		app:  &app,
		cont: StudentTab,
	}
	searchButton := widget.NewButton("Search", func() {
		Search(listBasic)
	})
	addButton := widget.NewButton("Add a student", func() {
		AddStudentForm(app)
	})

	list := CreateStudentList(listBasic, &data.Students)
	vbox := container.NewVBox(searchButton, addButton, widget.NewSeparator(), StudentTab)
	mainbox := container.NewHSplit(vbox, list)
	mainbox.SetOffset(0)
	window.SetContent(mainbox)
	window.ShowAndRun()
}

func LoadStudentInfo(vars basicVars, student *data.Student) {
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
		EditFormWindow(*vars.app, student)
	})

	deleteButton := widget.NewButton("Delete", func() {
		DeleteForm(*vars.app, student)
	})

	content := container.NewVBox(image, dataContainer, container.NewHBox(editButton, deleteButton))
	vars.cont.Objects = []fyne.CanvasObject{content}
}
