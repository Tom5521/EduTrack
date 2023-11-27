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

var RecordsList *widget.List

func UpdateRecordsList(student *data.Student) {
	RecordsList = GetRecordsList(student)
}

// AddRecord opens a window to add a register for a student.
func AddRecord(student *data.Student) {
	getTimeNow := func() string {
		return time.Now().Format("02/01/2006 12:14")
	}
	var tmpDate = getTimeNow()

	window := app.NewWindow("Add a record")
	window.Resize(sizes.RecSize)

	recNameLabel := widget.NewLabel("Record name:")
	recnameEntry := widget.NewEntry()
	recnameEntry.SetPlaceHolder(getTimeNow())

	recDateButton := widget.NewButton("Select Date", func() {
		const year int = 2023
		ret, err := zenity.Calendar("Select a date from below:",
			zenity.DefaultDate(year, time.December, 1))
		if err != nil {
			wins.ErrWin(app, err.Error())
		}
		tmpDate = ret.Format("02/01/2006")
	})
	detailsLabel := widget.NewLabel("Details")
	recDetails := widget.NewMultiLineEntry()
	recDetails.SetPlaceHolder("E.g., The student has not attended")

	/*
		var Rname string
		if recnameEntry.Text == "" {
			Rname = getTimeNow()
		} else {
			Rname = recnameEntry.Text
		}*/

	submitButton := widget.NewButton("Submit", func() {
		_, err := student.AddRecord(data.Record{
			Date: tmpDate,
			Name: recnameEntry.Text,
			Info: recDetails.Text,
		})
		if err != nil {
			wins.ErrWin(app, err.Error())
		}
		err = student.LoadRecords()
		if err != nil {
			wins.ErrWin(app, err.Error())
		}
		UpdateRecordsList(student)
		window.Close()
	})

	const gridNumber int = 2
	endBox := container.NewAdaptiveGrid(gridNumber, recDateButton, submitButton)

	vbox := container.NewVBox(
		detailsLabel,
		recNameLabel,
		recnameEntry,
		endBox,
	)
	box := container.NewVSplit(recDetails, vbox)
	box.SetOffset(1)
	window.SetContent(box)
	window.Show()
}

func ShowRecords(student *data.Student) {
	err := student.LoadRecords()
	if err != nil {
		wins.ErrWin(app, err.Error())
	}
	UpdateRecordsList(student)
	var content *fyne.Container
	window := app.NewWindow(student.Name + " records")
	window.Resize(sizes.ListSize)
	if len(student.Records) == 0 {
		noRegistersLabel := widget.NewLabel("No records found")
		noRegistersLabel.Alignment = fyne.TextAlignCenter
		addRegisterButton := widget.NewButton("Add record", func() {
			AddRecord(student)
			window.Close()
		})
		content = container.NewVBox(noRegistersLabel, addRegisterButton)
	} else {
		GetRecordsList(student)
		content = container.NewStack(GetRecordsList(student))
	}
	window.SetContent(content)
	window.Show()
}
func EditRecordsData(student *data.Student, index int) {
	UpdateRecordsList(student)
	var tmpDate string
	window := app.NewWindow("Edit Record")
	window.Resize(sizes.RecSize)
	rec := &student.Records[index]

	recNameEntry := widget.NewEntry()
	recNameEntry.SetText(rec.Name)

	recDate := widget.NewLabel("Date: " + rec.Date)
	dateButton := widget.NewButton("Select Date", func() {
		const year, day int = 2023, 1
		ret, err := zenity.Calendar("Select a date from below:",
			zenity.DefaultDate(year, time.December, day))
		if err != nil {
			wins.ErrWin(app, err.Error())
		}
		tmpDate = ret.Format("02/01/2006")
		recDate.SetText(tmpDate)
	})

	detailsLabel := widget.NewLabel("Details")
	recDetails := widget.NewMultiLineEntry()
	recDetails.SetText(rec.Info)

	// FormItems
	const gridNumber int = 2
	recNameForm := widget.NewFormItem("Record name:", recNameEntry)
	recDateForm := widget.NewFormItem("Record date:", container.NewAdaptiveGrid(gridNumber, recDate, dateButton))

	submitFunc := func() {
		err := rec.Edit(data.Record{Name: recNameEntry.Text, Info: recDetails.Text, Date: tmpDate, StudentId: student.ID})
		if err != nil {
			wins.ErrWin(app, err.Error())
		}

		err = student.LoadRecords()
		if err != nil {
			wins.ErrWin(app, err.Error())
		}
		UpdateRecordsList(student)
		window.Close()
	}

	form := widget.NewForm(
		recNameForm,
		recDateForm,
	)
	form.OnSubmit = submitFunc

	vbox := container.NewVBox(
		detailsLabel,
		form,
	)

	box := container.NewVSplit(recDetails, vbox)
	box.SetOffset(1)
	window.SetContent(box)

	window.Show()
}
