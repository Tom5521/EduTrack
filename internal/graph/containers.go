package graph

import (
	"EduTrack/cmd/data"
	icon "EduTrack/pkg/icons"
	"fmt"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/MyGolangTools/file"
)

var (
	FormSize = fyne.NewSize(300, 200)
	PickSize = fyne.NewSize(600, 400)
	ErrSize  = fyne.NewSize(400, 80)
)

func AddStudentForm(a fyne.App) {
	w := a.NewWindow("Students")
	w.Resize(FormSize)
	nameEntry := widget.NewEntry()
	ageEntry := widget.NewEntry()
	idEntry := widget.NewEntry()
	phoneNumberEntry := widget.NewEntry()
	imageFilePath := ""

	imageButton := widget.NewButton("Select Image", func() {
		w.Resize(PickSize)
		defer w.Resize(FormSize)
		dialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err == nil && reader != nil {
				imageFilePath = reader.URI().String()
			}
		}, w)
		dialog.SetFilter(storage.NewExtensionFileFilter([]string{".png", ".jpg", ".jpeg"}))
		dialog.Show()
	})

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Name", Widget: nameEntry},
			{Text: "Age", Widget: ageEntry},
			{Text: "ID", Widget: idEntry},
			{Text: "Phone Number", Widget: phoneNumberEntry},
		},
		OnSubmit: func() {
			age, err := strconv.Atoi(ageEntry.Text)
			if err != nil {
				ErrWin(a, err, nil)
			}

			student := data.Student{
				Name:          nameEntry.Text,
				Age:           age,
				ID:            idEntry.Text,
				Phone_number:  phoneNumberEntry.Text,
				ImageFilePath: strings.TrimPrefix(imageFilePath, "file://"),
			}
			data.Students = append(data.Students, student)

			nameEntry.SetText("")
			ageEntry.SetText("")
			idEntry.SetText("")
			phoneNumberEntry.SetText("")
			imageFilePath = "" // Restablece la variable de la ruta de la imagen de perfil
			w.Close()
		},
	}

	w.SetContent(container.NewVBox(form, imageButton))
	w.Show()
}
func StudentList(app fyne.App) *widget.List {
	list := widget.NewList(
		func() int {
			return len(data.Students)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(data.Students[id].Name)
		},
	)
	list.OnSelected = func(id widget.ListItemID) {
		student := data.Students[id]
		ShowStudentInfoWindow(app, student)
	}

	return list
}
func ShowStudentInfoWindow(a fyne.App, student data.Student) {
	infoWindow := a.NewWindow("Student Information")
	res, err := fyne.LoadResourceFromPath(student.ImageFilePath)
	if err != nil {
		ErrWin(a, err, nil)
	}
	image := canvas.NewImageFromResource(res)
	image.FillMode = canvas.ImageFillOriginal

	dataLabel := widget.NewLabel(
		"Name: " + student.Name + "\n" +
			"Age: " + fmt.Sprintf("%d", student.Age) + "\n" +
			"ID: " + student.ID + "\n" +
			"Phone number: " + student.Phone_number,
	)

	editButton := widget.NewButton("Edit", func() {
		EditForm(a, student)
	})

	deleteButton := widget.NewButton("Delete", func() {
		DeleteForm(a, student)
	})

	content := container.NewVBox(image, dataLabel, container.NewHBox(editButton, deleteButton))
	infoWindow.SetContent(content)
	infoWindow.Resize(FormSize)
	infoWindow.Show()
}
func EditForm(a fyne.App, student data.Student) {
	w := a.NewWindow("Students")
	w.Resize(FormSize)
	nameEntry := widget.NewEntry()
	ageEntry := widget.NewEntry()
	idEntry := widget.NewEntry()
	phoneNumberEntry := widget.NewEntry()
	imageFilePath := student.ImageFilePath
	fmt.Println(file.CheckFile(student.ImageFilePath))

	imageButton := widget.NewButton("Select Image", func() {
		w.Resize(PickSize)

		dialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err == nil && reader != nil {
				imageFilePath = reader.URI().String()
			}
		}, w)
		dialog.SetFilter(storage.NewExtensionFileFilter([]string{".png", ".jpg", ".jpeg"}))
		dialog.Show()
	})

	nameEntry.SetText(student.Name)
	ageEntry.SetText(fmt.Sprintf("%d", student.Age))
	idEntry.SetText(student.ID)
	phoneNumberEntry.SetText(student.Phone_number)

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Name", Widget: nameEntry},
			{Text: "Age", Widget: ageEntry},
			{Text: "ID", Widget: idEntry},
			{Text: "Phone Number", Widget: phoneNumberEntry},
		},
		OnSubmit: func() {
			age, _ := strconv.Atoi(ageEntry.Text)

			for i, s := range data.Students {
				if s == student {
					data.Students[i] = data.Student{
						Name:          nameEntry.Text,
						Age:           age,
						ID:            idEntry.Text,
						Phone_number:  phoneNumberEntry.Text,
						ImageFilePath: strings.TrimPrefix(imageFilePath, "file://"),
					}
					break
				}
			}

			w.Close()
		},
	}

	w.SetContent(container.NewVBox(imageButton, form))
	w.Show()
}
func DeleteForm(a fyne.App, student data.Student) {
	w := a.NewWindow("Delete Student")

	content := container.NewVBox(
		widget.NewLabel("Are you sure you want to eliminate the student?"),
		widget.NewButton("Yes", func() {
			for i, s := range data.Students {
				if s == student {
					data.Students = append(data.Students[:i], data.Students[i+1:]...)
					break
				}
			}
			w.Close()
		}),
		widget.NewButton("No", func() {
			w.Close()
		}),
	)

	w.SetContent(content)
	w.Show()
}

func Menu(a fyne.App) *fyne.MainMenu {
	menu := fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("Add", func() {
				AddStudentForm(a)
			}),
			fyne.NewMenuItem("Save", func() {
				data.SaveData()
			}),
			fyne.NewMenuItem("Exit", func() {
				a.Quit()
			}),
		),
		fyne.NewMenu("Edit",
			fyne.NewMenuItem("Undo Changes", func() {
				data.GetYamlData()
			}),
		),
	)
	return menu
}

func ErrWin(app fyne.App, err error, clWindow fyne.Window) {
	window := app.NewWindow("Error")
	window.Resize(ErrSize)
	//window.SetFixedSize(true)
	window.SetIcon(icon.ErrorICON)
	errlabel := widget.NewLabel(err.Error())
	errlabel.TextStyle.Bold = true
	errlabel.Alignment = fyne.TextAlignCenter
	acceptButton := widget.NewButton("Accept", func() {
		window.Close()
		if clWindow != nil {
			clWindow.Close()
		}
	})

	content := container.NewVBox(
		errlabel,
		acceptButton,
	)
	window.SetContent(content)
	window.SetMainMenu(window.MainMenu())
	window.Show()
}
