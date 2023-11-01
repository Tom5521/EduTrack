/*
 * Copyright Tom5521(c) - All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package mgraph

import (
	"EduTrack/data"
	"EduTrack/iconloader"
	"EduTrack/ui/sizes"
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Define global variables
var (
	Stundetlist  *widget.List
	RegisterList *widget.List
)

// CreateStudentList creates a list of students and their names.
func CreateStudentList(students *[]data.Student) fyne.Widget {
	// Initialize the student list widget
	Stundetlist = widget.NewList(
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
	Stundetlist.OnSelected = func(id widget.ListItemID) {
		d := *students
		Stundetlist.UnselectAll()
		LoadStudentInfo(&d[id])
		Stundetlist.Refresh()
	}

	return Stundetlist
}

// GetForm returns a container with form elements.
func GetForm(d *formReturn) *fyne.Container {
	// Create buttons
	imageButton := widget.NewButton("Select Image", func() {
		ImagePicker(app, d.ImagePath)
	})
	deleteImgBtn := widget.NewButton("Delete Current Image", func() {
		*d.ImagePath = ""
	})

	// Create the form
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Name", Widget: d.NameEntry},
			{Text: "Age", Widget: d.AgeEntry},
			{Text: "ID", Widget: d.IDEntry},
			{Text: "Phone Number", Widget: d.PhoneEntry},
		},
		OnSubmit: d.ExecFunc,
	}
	form.SubmitText = "Submit"

	// Create the content container
	content := container.NewVBox(container.NewHBox(imageButton, deleteImgBtn), form)
	return content
}

// formReturn represents a structure for form elements.
type formReturn struct {
	ExecFunc   func()
	NameEntry  *widget.Entry
	AgeEntry   *widget.Entry
	IDEntry    *widget.Entry
	PhoneEntry *widget.Entry
	ImagePath  *string
}

// atoi converts a string to an integer, handling errors.
func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

// TemplateUser returns a container with user data.
func TemplateUser() *fyne.Container {
	// Create user data elements
	iconloader.SetThemeIcons(app.Settings().ThemeVariant())
	image := canvas.NewImageFromResource(iconloader.UserTemplateICON)
	image.SetMinSize(sizes.ProfileSize)
	dataLabel := widget.NewLabel(
		"Name: " + "--" + "\n" +
			"Age: " + "--" + "\n" +
			"ID: " + "--" + "\n" +
			"Phone number: " + "--",
	)
	content := container.NewVBox(image, dataLabel)
	return content
}

// Menu returns the main application menu.
func Menu() *fyne.MainMenu {
	// Create the main menu
	menu := fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("Load a config file", func() {
				data.LoadConf(LoadConf())
			}),
			fyne.NewMenuItem("Add Student", func() {
				AddStudentForm()
			}),
			fyne.NewMenuItem("Re-Save Changes", func() {
				data.SaveData()
			})),
		fyne.NewMenu("Edit",
			fyne.NewMenuItem("Reload data", func() {
				data.GetYamlData()
			}),
		),
		fyne.NewMenu("Help",
			fyne.NewMenuItem("About", func() {
				AboutWin()
			}),
		),
	)
	return menu
}

// LoadConf loads a configuration file.
func LoadConf() string {
	ret := recieveFile()
	fmt.Println(ret)
	return ret
}

// recieveFile receives a file from the user.
func recieveFile() string {
	resultChannel := make(chan string)
	go FilePicker(app, resultChannel)
	selectedFilePath := <-resultChannel
	return selectedFilePath
}

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
