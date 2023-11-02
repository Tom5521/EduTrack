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
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/ncruces/zenity"
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
func ImagePicker(imageFilePath *string) {
	filter := []string{"*.png", "*.gif", "*.ico", "*.jpg", "*.webp"}
	const defaultPath string = ""
	ret, err := zenity.SelectFile(
		zenity.Filename(defaultPath),
		zenity.FileFilters{
			{"Image files", filter, true},
		})
	if err != nil {
		ErrWin(app, err.Error())
	} else {
		*imageFilePath = ret
	}
}

// FilePicker opens a file picker window to select a configuration file.
func FilePicker() string {
	filter := []string{"*.yml", "*.yaml"}
	const defaultPath string = ""
	ret, err := zenity.SelectFile(
		zenity.Filename(defaultPath),
		zenity.FileFilters{
			{"Yaml files", filter, true},
		})
	if err != nil {
		ErrWin(app, err.Error())
	}
	return ret
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
	window.RequestFocus()
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
	var tmpDate string = getTimeNow()

	window := app.NewWindow("Add a register")
	window.Resize(sizes.RegSize)

	regNameLabel := widget.NewLabel("Register name:")
	regnameEntry := widget.NewEntry()
	regnameEntry.SetPlaceHolder(getTimeNow())

	regDateButton := widget.NewButton("Select Date", func() {
		ret, err := zenity.Calendar("Select a date from below:",
			zenity.DefaultDate(2023, time.December, 1))
		if err != nil {
			ErrWin(app, err.Error())
		}
		tmpDate = ret.Format("02/01/2006")
	})
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
			Date: tmpDate,
			Name: Rname,
			Data: regDetails.Text,
		}
		student.Register = append(student.Register, NewReg)
		data.SaveData()
		data.GetYamlData()
		RegisterList.Refresh()
		window.Close()
	})

	endBox := container.NewAdaptiveGrid(2, regDateButton, submitButton)

	vbox := container.NewVBox(
		DetailsLabel,
		regNameLabel,
		regnameEntry,
		endBox,
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
	var tmpDate string

	window := app.NewWindow("Edit Register")
	window.Resize(sizes.RegSize)
	reg := &student.Register[index]
	regNameLabel := widget.NewLabel("Register name:")
	regNameLabel.Alignment = fyne.TextAlignCenter
	regnameEntry := widget.NewEntry()
	regnameEntry.SetText(reg.Name)
	regnameBox := container.NewAdaptiveGrid(2, regNameLabel, regnameEntry)
	regDate := widget.NewLabel("Date: " + reg.Date)
	regDate.Alignment = fyne.TextAlignCenter
	DateButton := widget.NewButton("Select Date", func() {
		ret, err := zenity.Calendar("Select a date from below:",
			zenity.DefaultDate(2023, time.December, 1))
		if err != nil {
			ErrWin(app, err.Error())
		}
		tmpDate = ret.Format("02/01/2006")
	})
	datebox := container.NewAdaptiveGrid(2, regDate, DateButton)
	DetailsLabel := widget.NewLabel("Details")
	regDetails := widget.NewMultiLineEntry()
	regDetails.SetText(reg.Data)

	submitButton := widget.NewButton("Submit", func() {
		reg.Name = regnameEntry.Text
		reg.Data = regDetails.Text
		reg.Date = tmpDate
		data.SaveData()
		window.Close()
	})

	vbox := container.NewVBox(
		DetailsLabel,
		regnameBox,
		datebox,
		submitButton,
	)

	box := container.NewVSplit(regDetails, vbox)
	box.SetOffset(1)
	window.SetContent(box)

	window.Show()
}

// AboutWin opens an "About" window to display information about the app.
func AboutWin() {
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
