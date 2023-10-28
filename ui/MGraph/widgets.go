/*
 * Copyright Tom5521(c) - All Rights Reserved.
 *
 * This project is licenced under the MIT License.
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

var (
	Stundetlist  *widget.List
	RegisterList *widget.List
)

func CreateStudentList(vars basicVars, students *[]data.Student) fyne.Widget {
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
	Stundetlist.OnSelected = func(id widget.ListItemID) {
		d := *students
		Stundetlist.UnselectAll()
		LoadStudentInfo(vars, &d[id])
		Stundetlist.Refresh()
	}

	return Stundetlist
}

func GetForm(app fyne.App, d *formReturn) *fyne.Container {
	imageButton := widget.NewButton("Select Image", func() {
		ImagePicker(app, d.ImagePath)
	})
	deleteImgBtn := widget.NewButton("Delete Current Image", func() {
		*d.ImagePath = ""
	})
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

	content := container.NewVBox(container.NewHBox(imageButton, deleteImgBtn), form)
	return content
}

type formReturn struct {
	ExecFunc   func()
	NameEntry  *widget.Entry
	AgeEntry   *widget.Entry
	IDEntry    *widget.Entry
	PhoneEntry *widget.Entry
	ImagePath  *string
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func TemplateUser() *fyne.Container {
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

func Menu(a fyne.App) *fyne.MainMenu {
	menu := fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("Load a config file", func() {
				data.LoadConf(LoadConf(a))
			}),
			fyne.NewMenuItem("Add Student", func() {
				AddStudentForm(a)
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
				AboutWin(a)
			}),
		),
	)
	return menu
}

func LoadConf(app fyne.App) string {
	ret := recieveFile(app)
	fmt.Println(ret)
	return ret
}

func recieveFile(app fyne.App) string {
	resultChannel := make(chan string)
	go FilePicker(app, resultChannel)
	selectedFilePath := <-resultChannel
	return selectedFilePath
}
