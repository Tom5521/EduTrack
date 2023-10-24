package graph

import (
	"EduTrack/cmd/data"
	icon "EduTrack/pkg/icons"
	windowtools "EduTrack/pkg/windowTools"
	win "EduTrack/pkg/windows"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	xtheme "fyne.io/x/fyne/theme"
	"github.com/Tom5521/MyGolangTools/file"
)

var (
	FormSize    = win.FormSize
	PickSize    = win.PickSize
	ErrSize     = win.ErrSize
	ProfileSize = win.ProfileSize
	SearchSize  = win.SearchSize
)

func Init() {
	app := app.New()
	app.Settings().SetTheme(xtheme.AdwaitaTheme())

	mainWin := app.NewWindow("EduTrack")
	mainWin.SetMaster()
	mainWin.SetMainMenu(Menu(app))
	mainWin.Resize(fyne.NewSize(800, 600))
	windowtools.MaximizeWin(mainWin)
	StudentTab := container.NewHBox(TemplateUser())
	searchButton := widget.NewButton("Search", func() {
		Search(app, StudentTab)
	})
	hbox := container.NewVBox(searchButton, StudentTab)
	list := StudentList(app, StudentTab, mainWin)
	content := container.NewHSplit(hbox, list)
	content.SetOffset(0)
	mainWin.SetContent(content)
	mainWin.Show()

	app.Run()
}

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
				win.ErrWin(a, err, nil)
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
			imageFilePath = ""
			w.Close()
		},
	}

	w.SetContent(container.NewVBox(imageButton, form))
	w.Show()
}

func StudentList(app fyne.App, infoTab *fyne.Container, w fyne.Window) *widget.List {
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
		list.UnselectAll()
		LoadStudentInfo(app, infoTab, student)
	}
	return list
}

func LoadStudentInfo(a fyne.App, cont *fyne.Container, student data.Student) {
	var (
		image *canvas.Image
		err   error
	)
	if check, _ := file.CheckFile(student.ImageFilePath); !check {
		image = canvas.NewImageFromResource(icon.UserTemplateICON)
		image.SetMinSize(win.ProfileSize)
	} else {
		image, err = win.LoadProfileImg(student.ImageFilePath)
	}
	if err != nil {
		win.ErrWin(a, err, nil)
	}

	dataLabel := widget.NewLabel(
		"Name: " + student.Name + "\n" +
			"Age: " + fmt.Sprintf("%d", student.Age) + "\n" +
			"ID: " + student.ID + "\n" +
			"Phone number: " + student.Phone_number,
	)

	editButton := widget.NewButton("Edit", func() {
		EditForm(a, cont, student)
	})

	deleteButton := widget.NewButton("Delete", func() {
		DeleteForm(a, student)
	})

	content := container.NewVBox(image, dataLabel, container.NewHBox(editButton, deleteButton))
	cont.Objects = []fyne.CanvasObject{content}
}

func EditForm(a fyne.App, cont *fyne.Container, student data.Student) {
	w := a.NewWindow("Edit " + student.Name)
	w.Resize(FormSize)
	nameEntry := widget.NewEntry()
	ageEntry := widget.NewEntry()
	idEntry := widget.NewEntry()
	phoneNumberEntry := widget.NewEntry()
	imageFilePath := student.ImageFilePath

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
			name := nameEntry.Text
			id := idEntry.Text

			// Comprobar si el nombre o el ID están en blanco
			if name == "" || id == "" {
				win.ErrWin(a, errors.New("Name and ID cannot be empty."), nil)
				return
			}

			// Comprobar si el nombre o el ID están duplicados
			for _, s := range data.Students {
				if s.Name == name && s.ID != student.ID {
					win.ErrWin(a, errors.New("Name already exists for another student."), nil)
					return
				}
				if s.ID == id && s.ID != student.ID {
					win.ErrWin(a, errors.New("ID already exists for another student."), nil)
					return
				}
			}

			age, _ := strconv.Atoi(ageEntry.Text)

			for i, s := range data.Students {
				if s == student {
					data.Students[i] = data.Student{
						Name:          name,
						Age:           age,
						ID:            id,
						Phone_number:  phoneNumberEntry.Text,
						ImageFilePath: strings.TrimPrefix(imageFilePath, "file://"),
					}
					break
				}
			}
			data.SaveData()
			LoadStudentInfo(a, cont, student)
			w.Close()
		},
	}
	content := container.NewVBox(imageButton, form)
	w.SetContent(content)
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
			})),
		fyne.NewMenu("Edit",
			fyne.NewMenuItem("Undo Changes", func() {
				data.GetYamlData()
			}),
		),
	)
	return menu
}

func TemplateUser() *fyne.Container {
	image := canvas.NewImageFromResource(icon.UserTemplateICON)
	image.SetMinSize(win.ProfileSize)
	dataLabel := widget.NewLabel(
		"Name: " + "--" + "\n" +
			"Age: " + "--" + "\n" +
			"ID: " + "--" + "\n" +
			"Phone number: " + "--",
	)
	content := container.NewVBox(image, dataLabel)
	return content
}

func Search(a fyne.App, cont *fyne.Container) {
	w := a.NewWindow("Search Student")
	w.Resize(SearchSize)
	entry := widget.NewEntry()
	searchButton := widget.NewButton("Search", func() {
		studentID := entry.Text
		student := data.FindStudentByID(studentID)
		if student != nil {
			LoadStudentInfo(a, cont, *student)
			w.Close()
		} else {
			win.ErrWin(a, errors.New("Student not found!"), nil)
		}
	})

	content := container.NewVBox(
		widget.NewLabel("Enter Student ID:"),
		entry,
		searchButton,
	)

	w.SetContent(content)
	w.Show()
}
