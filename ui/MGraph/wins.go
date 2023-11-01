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
	"net/url"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

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
			ErrWin(app, "Some value in the form is empty")
			return
		}

		if existsId(idEntry.Text, data.GetIDs()) && idEntry.Text != student.ID {
			ErrWin(app, "The ID already exists")
			return
		}

		// Update student information
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

	content := GetForm(&retForm)

	window.SetContent(content)
	window.Show()
}

// ImagePicker opens a file picker window to select an image file.
func ImagePicker(app fyne.App, imageFilePath *string) {
	window := app.NewWindow("Pick an image!")
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

// FilePicker opens a file picker window to select a configuration file.
func FilePicker(app fyne.App, resultChan chan string) {
	window := app.NewWindow("Select a config file!")
	window.Resize(sizes.PickSize)
	window.SetFixedSize(true)

	dialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err == nil && reader != nil {
			t := strings.TrimPrefix(reader.URI().String(), "file://")
			resultChan <- t
			window.Close()
		}
	}, window)
	dialog.SetFilter(storage.NewExtensionFileFilter([]string{".yaml", ".yml"}))
	dialog.Show()
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
						data.SaveData()
						data.GetYamlData()
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
func AddStudentForm(app fyne.App) {
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
			ErrWin(app, "Some value in the form is empty")
			return
		}
		if existsId(idEntry.Text, data.GetIDs()) {
			ErrWin(app, "The ID already exists")
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
		data.SaveData()
		data.GetYamlData()
		Stundetlist.Refresh()
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

// ErrWin opens an error window with a message.
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

// Search opens a window to search for a student by ID.
func Search() {
	w := app.NewWindow("Search Student")
	w.Resize(sizes.SearchSize)
	entry := widget.NewEntry()
	searchButton := widget.NewButton("Search", func() {
		studentID := entry.Text
		student := data.FindStudentByID(studentID)
		if student != nil {
			LoadStudentInfo(student)
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

// existsId checks if a given string exists in a list of strings.
func existsId(check string, list []string) bool {
	var contains bool
	for _, v := range list {
		if v == check {
			contains = true
			break
		}
	}
	return contains
}

// checkValues checks if all required form fields are not empty.
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

// AddRegister opens a window to add a register for a student.
func AddRegister(student *data.Student) {
	getTimeNow := func() string {
		time := time.Now().Format("02/01/2006")
		return time
	}

	window := app.NewWindow("Add a register")
	window.Resize(sizes.RegSize)

	regNameLabel := widget.NewLabel("Register name:")
	regnameEntry := widget.NewEntry()
	regnameEntry.SetPlaceHolder(getTimeNow())
	DetailsLabel := widget.NewLabel("Details")
	regDetails := widget.NewMultiLineEntry()
	regDetails.SetPlaceHolder("E.g., The student has not attended")

	var Rname string
	if regnameEntry.Text == "" {
		Rname = getTimeNow()
	} else {
		Rname = regnameEntry.Text
	}

	submitButton := widget.NewButton("Submit", func() {
		NewReg := struct {
			Date string
			Name string
			Data string
		}{
			Date: getTimeNow(),
			Name: Rname,
			Data: regDetails.Text,
		}
		student.Register = append(student.Register, NewReg)
		data.SaveData()
		data.GetYamlData()
		RegisterList.Refresh()
		window.Close()
	})

	vbox := container.NewVBox(
		DetailsLabel,
		regNameLabel,
		regnameEntry,
		submitButton,
	)
	box := container.NewVSplit(regDetails, vbox)
	box.SetOffset(1)
	window.SetContent(box)
	window.Show()
}

func ShowRegisters(student *data.Student) {
	GetRegisterList(student)
	var content *fyne.Container
	window := app.NewWindow(student.Name + " registers")
	window.Resize(sizes.RegsListSize)
	if len(student.Register) == 0 {
		noRegistersLabel := widget.NewLabel("No registers found")
		noRegistersLabel.Alignment = fyne.TextAlignCenter
		AddRegisterButton := widget.NewButton("Add Register", func() {
			AddRegister(student)
			window.Close()
		})
		content = container.NewVBox(noRegistersLabel, AddRegisterButton)
	} else {
		GetRegisterList(student)
		content = container.NewMax(RegisterList)
	}
	window.SetContent(content)
	window.Show()
}
func EditRegisterData(student *data.Student, index int) {
	/*
		getTimeNow := func() string {
			time := time.Now().Format("02/01/2006")
			return time
		}
	*/
	window := app.NewWindow("Edit Register")
	window.Resize(sizes.RegSize)
	reg := &student.Register[index]

	regNameLabel := widget.NewLabel("Register name:")
	regnameEntry := widget.NewEntry()
	regnameEntry.SetText(reg.Name)

	regDate := widget.NewLabel("Date: " + reg.Date)

	DetailsLabel := widget.NewLabel("Details")

	regDetails := widget.NewMultiLineEntry()
	regDetails.SetText(reg.Data)

	submitButton := widget.NewButton("Submit", func() {

		reg.Name = regnameEntry.Text
		reg.Data = regDetails.Text

		data.SaveData()

		window.Close()
	})

	vbox := container.NewVBox(
		DetailsLabel,
		regNameLabel,
		regnameEntry,
		regDate,
		submitButton,
	)

	box := container.NewVSplit(regDetails, vbox)
	box.SetOffset(1)
	window.SetContent(box)

	window.Show()
}

// AboutWin opens an "About" window to display information about the app.
func AboutWin(app fyne.App) {
	window := app.NewWindow("About")
	label1 := widget.NewLabel("Created by:")
	link, _ := url.Parse("https://github.com/Tom5521")
	gitLabel := widget.NewHyperlink("Tom5521", link)
	LicenceLabel := widget.NewLabel("Under the MIT license")
	AuthorCont := container.NewHBox(label1, gitLabel)
	logo := canvas.NewImageFromResource(iconloader.AppICON)
	logo.SetMinSize(fyne.NewSize(300, 300))
	vbox1 := container.NewVBox(
		AuthorCont,
		LicenceLabel,
		logo,
	)
	window.SetContent(vbox1)
	window.Show()
}
