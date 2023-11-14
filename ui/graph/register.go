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
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/ncruces/zenity"
)

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
			wins.ErrWin(app, err.Error())
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
		student.Registers = append(student.Registers, NewReg)
		data.SaveStudentsData()
		data.GetStundentData()
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
	if len(student.Registers) == 0 {
		noRegistersLabel := widget.NewLabel("No registers found")
		noRegistersLabel.Alignment = fyne.TextAlignCenter
		AddRegisterButton := widget.NewButton("Add Register", func() {
			AddRegister(student)
			window.Close()
		})
		content = container.NewVBox(noRegistersLabel, AddRegisterButton)
	} else {
		GetRegisterList(student)
		content = container.NewStack(RegisterList)
	}
	window.SetContent(content)
	window.Show()
}
func EditRegisterData(student *data.Student, index int) {
	var tmpDate string
	window := app.NewWindow("Edit Register")
	window.Resize(sizes.RegSize)
	reg := &student.Registers[index]

	regnameEntry := widget.NewEntry()
	regnameEntry.SetText(reg.Name)

	regDate := widget.NewLabel("Date: " + reg.Date)
	DateButton := widget.NewButton("Select Date", func() {
		ret, err := zenity.Calendar("Select a date from below:",
			zenity.DefaultDate(2023, time.December, 1))
		if err != nil {
			wins.ErrWin(app, err.Error())
		}
		tmpDate = ret.Format("02/01/2006")
		regDate.SetText(tmpDate)
	})

	DetailsLabel := widget.NewLabel("Details")
	regDetails := widget.NewMultiLineEntry()
	regDetails.SetText(reg.Data)

	// FormItems
	RegNameForm := widget.NewFormItem("Register name:", regnameEntry)
	RegDateForm := widget.NewFormItem("Register date:", container.NewAdaptiveGrid(2, regDate, DateButton))

	submitFunc := func() {
		reg.EditName(regnameEntry.Text)
		reg.EditData(regDetails.Text)
		reg.EditDate(tmpDate)
		data.SaveStudentsData()
		window.Close()
	}

	Form := widget.NewForm(
		RegNameForm,
		RegDateForm,
	)
	Form.OnSubmit = submitFunc

	vbox := container.NewVBox(
		DetailsLabel,
		Form,
	)

	box := container.NewVSplit(regDetails, vbox)
	box.SetOffset(1)
	window.SetContent(box)

	window.Show()
}
