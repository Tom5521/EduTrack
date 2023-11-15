/*
 * Copyright (c) 2023 Tom5521- All Rights Reserved.
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

// AddRecord opens a window to add a register for a student.
func AddRecord(student *data.Student) {
	getTimeNow := func() string {
		time := time.Now().Format("02/01/2006")
		return time
	}
	var tmpDate string = getTimeNow()

	window := app.NewWindow("Add a register")
	window.Resize(sizes.RecSize)

	regNameLabel := widget.NewLabel("Record name:")
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

func ShowRecords(student *data.Student) {
	GetRegisterList(student)
	var content *fyne.Container
	window := app.NewWindow(student.Name + " records")
	window.Resize(sizes.RecsListSize)
	if len(student.Registers) == 0 {
		noRegistersLabel := widget.NewLabel("No registers found")
		noRegistersLabel.Alignment = fyne.TextAlignCenter
		AddRegisterButton := widget.NewButton("Add Register", func() {
			AddRecord(student)
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
func EditRecordsData(student *data.Student, index int) {
	var tmpDate string
	window := app.NewWindow("Edit Record")
	window.Resize(sizes.RecSize)
	reg := &student.Registers[index]

	RecNameEntry := widget.NewEntry()
	RecNameEntry.SetText(reg.Name)

	recDate := widget.NewLabel("Date: " + reg.Date)
	DateButton := widget.NewButton("Select Date", func() {
		ret, err := zenity.Calendar("Select a date from below:",
			zenity.DefaultDate(2023, time.December, 1))
		if err != nil {
			wins.ErrWin(app, err.Error())
		}
		tmpDate = ret.Format("02/01/2006")
		recDate.SetText(tmpDate)
	})

	DetailsLabel := widget.NewLabel("Details")
	recDetails := widget.NewMultiLineEntry()
	recDetails.SetText(reg.Data)

	// FormItems
	RecNameForm := widget.NewFormItem("Record name:", RecNameEntry)
	RecDateForm := widget.NewFormItem("Record date:", container.NewAdaptiveGrid(2, recDate, DateButton))

	submitFunc := func() {
		reg.EditName(RecNameEntry.Text)
		reg.EditData(recDetails.Text)
		reg.EditDate(tmpDate)
		window.Close()
	}

	Form := widget.NewForm(
		RecNameForm,
		RecDateForm,
	)
	Form.OnSubmit = submitFunc

	vbox := container.NewVBox(
		DetailsLabel,
		Form,
	)

	box := container.NewVSplit(recDetails, vbox)
	box.SetOffset(1)
	window.SetContent(box)

	window.Show()
}
