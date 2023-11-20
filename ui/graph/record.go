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
		time := time.Now().Format("02/01/2006 12:14")
		return time
	}
	var tmpDate string = getTimeNow()

	window := app.NewWindow("Add a record")
	window.Resize(sizes.RecSize)

	recNameLabel := widget.NewLabel("Record name:")
	recnameEntry := widget.NewEntry()
	recnameEntry.SetPlaceHolder(getTimeNow())

	recDateButton := widget.NewButton("Select Date", func() {
		ret, err := zenity.Calendar("Select a date from below:",
			zenity.DefaultDate(2023, time.December, 1))
		if err != nil {
			wins.ErrWin(app, err.Error())
		}
		tmpDate = ret.Format("02/01/2006")
	})
	DetailsLabel := widget.NewLabel("Details")
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
		student.AddRecord(data.Record{
			Date: tmpDate,
			Name: recnameEntry.Text,
			Info: recDetails.Text,
		})
		student.LoadRecords()
		UpdateRecordsList(student)
		window.Close()
	})

	endBox := container.NewAdaptiveGrid(2, recDateButton, submitButton)

	vbox := container.NewVBox(
		DetailsLabel,
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
	student.LoadRecords()
	UpdateRecordsList(student)
	var content *fyne.Container
	window := app.NewWindow(student.Name + " records")
	window.Resize(sizes.RecsListSize)
	if len(student.Records) == 0 {
		noRegistersLabel := widget.NewLabel("No records found")
		noRegistersLabel.Alignment = fyne.TextAlignCenter
		AddRegisterButton := widget.NewButton("Add record", func() {
			AddRecord(student)
			window.Close()
		})
		content = container.NewVBox(noRegistersLabel, AddRegisterButton)
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

	RecNameEntry := widget.NewEntry()
	RecNameEntry.SetText(rec.Name)

	recDate := widget.NewLabel("Date: " + rec.Date)
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
	recDetails.SetText(rec.Info)

	// FormItems
	RecNameForm := widget.NewFormItem("Record name:", RecNameEntry)
	RecDateForm := widget.NewFormItem("Record date:", container.NewAdaptiveGrid(2, recDate, DateButton))

	submitFunc := func() {
		rec.Edit(data.Record{Name: RecNameEntry.Text, Info: recDetails.Text, Date: tmpDate, StudentId: student.ID})
		student.LoadRecords()
		UpdateRecordsList(student)
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