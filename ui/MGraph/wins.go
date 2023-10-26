package mgraph

import (
	"EduTrack/data"
	"EduTrack/iconloader"
	"EduTrack/ui/sizes"
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

func EditFormWindow(app fyne.App, student *data.Student) {
	window := app.NewWindow("Edit " + student.Name)
	window.Resize(sizes.FormSize)
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
		if !checkValues(formReturn{NameEntry: nameEntry, AgeEntry: ageEntry, IDEntry: idEntry, PhoneEntry: phoneEntry}) {
			ErrWin(app, "Some value in the form is empty")
			return
		}
		if existsId(idEntry.Text, data.GetIDs()) && idEntry.Text != student.ID {
			ErrWin(app, "The id already exists")
			return
		}
		student.Age = atoi(ageEntry.Text)
		student.Name = nameEntry.Text
		student.Phone_number = phoneEntry.Text
		student.ID = idEntry.Text
		student.ImageFilePath = imagePath
		data.SaveData()
		data.GetYamlData()
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

	content := GetForm(app, &retForm)

	window.SetContent(content)
	window.Show()

}
func ImagePicker(app fyne.App, imageFilePath *string) {
	window := app.NewWindow("Pick a image!")
	window.Resize(sizes.PickSize)
	dialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err == nil && reader != nil {
			*imageFilePath = strings.TrimPrefix(reader.URI().String(), "file://")
			window.Close()
		}
	}, window)
	dialog.SetFilter(storage.NewExtensionFileFilter([]string{".png", ".jpg", ".jpeg"}))
	dialog.Show()

	window.Show()
}

func FilePicker(app fyne.App, resultChan chan string) {
	window := app.NewWindow("Select a config file!")
	window.Resize(sizes.PickSize)
	window.SetFixedSize(true)
	dialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err == nil && reader != nil {
			t := strings.TrimPrefix(reader.URI().String(), "file://")
			resultChan <- t // Envía el valor de t a través del canal
			window.Close()
		}
	}, window)
	dialog.SetFilter(storage.NewExtensionFileFilter([]string{".yaml", ".yml"}))
	dialog.Show()
	window.Show()
}
func recieveFile(app fyne.App) string {
	resultChannel := make(chan string)
	go FilePicker(app, resultChannel)
	selectedFilePath := <-resultChannel
	return selectedFilePath
}

func DeleteForm(a fyne.App, student *data.Student) {
	w := a.NewWindow("Delete Student")
	content := container.NewVBox(
		widget.NewLabel("Are you sure you want to eliminate the student?"),
		widget.NewButton("Yes", func() {
			for i, s := range data.Students {
				if s == *student {
					data.Students = append(data.Students[:i], data.Students[i+1:]...)
					data.SaveData()
					data.GetYamlData()
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

func AddStudentForm(app fyne.App) {
	var imagePath string
	window := app.NewWindow("Add a student")
	window.Resize(sizes.FormSize)
	nameEntry := widget.NewEntry()
	ageEntry := widget.NewEntry()
	idEntry := widget.NewEntry()
	phoneEntry := widget.NewEntry()

	submitFunc := func() {
		if !checkValues(formReturn{NameEntry: nameEntry, AgeEntry: ageEntry, IDEntry: idEntry, PhoneEntry: phoneEntry}) {
			ErrWin(app, "Some value in the form is empty")
			return
		}
		if existsId(idEntry.Text, data.GetIDs()) {
			ErrWin(app, "The id already exists")
			return
		}
		data.Students = append(data.Students, data.Student{
			Name:          nameEntry.Text,
			Age:           atoi(ageEntry.Text),
			ID:            idEntry.Text,
			Phone_number:  phoneEntry.Text,
			ImageFilePath: imagePath,
		})
		data.SaveData()
		data.GetYamlData()
		list.Refresh()
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
	content := GetForm(app, &formRet)
	window.SetContent(content)
	window.Show()

}

func ErrWin(app fyne.App, err string, clWindow ...fyne.Window) {
	window := app.NewWindow("Error")
	window.Resize(sizes.ErrSize)
	window.SetIcon(iconloader.ErrorICON)
	errlabel := widget.NewLabel(err)
	errlabel.TextStyle.Bold = true
	errlabel.Alignment = fyne.TextAlignCenter
	acceptButton := widget.NewButton("Accept", func() {
		window.Close()
		if len(clWindow) > 0 {
			clWindow[0].Close()
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
func Search(vars basicVars) {
	app := *vars.app
	w := app.NewWindow("Search Student")
	w.Resize(sizes.SearchSize)
	entry := widget.NewEntry()
	searchButton := widget.NewButton("Search", func() {
		studentID := entry.Text
		student := data.FindStudentByID(studentID)
		if student != nil {
			LoadStudentInfo(vars, student)
			w.Close()
		} else {
			ErrWin(app, "Student not found!")
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

func existsId(check string, list []string) bool {
	var conains bool
	for _, v := range list {
		if v == check {
			conains = true
			break
		}
	}
	return conains
}
func checkValues(d formReturn) bool {
	if d.AgeEntry.Text == "" {
		return false
	}
	if d.IDEntry.Text == "" {
		return false
	}
	if d.PhoneEntry.Text == "" {
		return false
	}
	if d.NameEntry.Text == "" {
		return false
	}
	return true
}
