package gui

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/EduTrack/assets"
	"github.com/Tom5521/EduTrack/internal/pkg/sizes"
	"github.com/Tom5521/EduTrack/pkg/data"
	"github.com/Tom5521/EduTrack/pkg/wins"
	"github.com/ncruces/zenity"
)

func (ui *ui) MakeRecList(l *[]data.Record) *widget.List {
	err := data.LoadRecords()
	if err != nil {
		wins.ErrWin(ui.App, err.Error())
	}
	list := widget.NewList(
		func() int {
			return len(*l)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			l := *l
			o.(*widget.Label).SetText(l[i].Name)
		},
	)
	return list
}

func (ui *ui) GetStudentRecordsList(student *data.Student) *widget.List {
	student.GetRecords()
	list := ui.MakeRecList(&student.Records)
	return list
}

func (ui *ui) GetRecordsList() *widget.List {
	list := ui.MakeRecList(&data.Records)
	return list
}

func (ui *ui) AddRecord(studentID uint) {
	getTimeNow := func() string {
		return time.Now().Format("02/01/2006 12:14")
	}
	var tmpDate = getTimeNow()

	window := ui.App.NewWindow("Add a record")
	window.Resize(sizes.RecSize)

	recNameLabel := widget.NewLabel("Record name:")
	recnameEntry := widget.NewEntry()
	recnameEntry.SetPlaceHolder(getTimeNow())

	recDateButton := widget.NewButton("Select Date", func() {
		const year int = 2023
		ret, err := zenity.Calendar("Select a date from below:",
			zenity.DefaultDate(year, time.December, 1))
		if err != nil {
			wins.ErrWin(ui.App, err.Error())
		}
		tmpDate = ret.Format("02/01/2006")
	})
	detailsLabel := widget.NewLabel("Details")
	recDetails := widget.NewMultiLineEntry()
	recDetails.SetPlaceHolder("E.g., The student has not attended")
	submitButton := widget.NewButton("Submit", func() {
		err := data.AddRecord(&data.Record{
			StudentID: studentID,
			Date:      tmpDate,
			Name:      recnameEntry.Text,
			Info:      recDetails.Text,
		})
		if err != nil {
			wins.ErrWin(ui.App, err.Error())
		}
		i := data.FindStudentIndexByID(studentID)

		data.Students[i].GetRecords()
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

func (ui *ui) EditRecordData(recordID uint) {
	var tmpDate string
	window := ui.App.NewWindow("Edit Record")
	window.Resize(sizes.RecSize)
	i := data.FindRecordIndexByID(recordID)
	if i == -1 {
		wins.ErrWin(ui.App, fmt.Sprintf("Record ID<%v> not found!", recordID))
		window.Close()
		return
	}
	rec := &data.Records[i]

	i = data.FindStudentIndexByID(rec.StudentID)
	if i == -1 {
		wins.ErrWin(ui.App, fmt.Sprintf("Student ID<%v> not found!", rec.StudentID))
		window.Close()
		return
	}
	student := &data.Students[i]
	student.GetRecords()

	recNameEntry := widget.NewEntry()
	recNameEntry.SetText(rec.Name)

	recDate := widget.NewLabel("Date: " + rec.Date)
	dateButton := widget.NewButton("Select Date", func() {
		const year, day int = 2023, 1
		ret, err := zenity.Calendar("Select a date from below:",
			zenity.DefaultDate(year, time.December, day))
		if err != nil {
			wins.ErrWin(ui.App, err.Error())
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
		err := rec.Edit(data.Record{Name: recNameEntry.Text, Info: recDetails.Text, Date: tmpDate, StudentID: student.ID})
		if err != nil {
			wins.ErrWin(ui.App, err.Error())
		}

		student.GetRecords()
		if err != nil {
			wins.ErrWin(ui.App, err.Error())
		}
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

func (ui *ui) StudentRecordsMainWin(student *data.Student) {
	w := ui.App.NewWindow(student.Name + " Records")
	var selected int
	list := ui.GetStudentRecordsList(student)
	bar := widget.NewToolbar(
		widget.NewToolbarAction(assets.Plus, func() {
			ui.AddRecord(student.ID)
			student.GetRecords()
			list.UnselectAll()
		}),
		widget.NewToolbarAction(assets.Cross, func() {
			if selected == -1 {
				return
			}
			err := student.DeleteRecord(student.Records[selected].ID)
			if err != nil {
				wins.ErrWin(ui.App, err.Error())
			}
			list.UnselectAll()
			selected = -1
		}),
		widget.NewToolbarAction(assets.Edit, func() {
			if selected == -1 {
				return
			}
			ui.EditRecordData(data.Records[selected].ID)
			student.GetRecords()
			list.UnselectAll()
		}),
	)

	list.OnSelected = func(id widget.ListItemID) {
		selected = id
	}
	content := container.NewBorder(bar, nil, nil, nil, list)

	if len(student.Records) == 0 {
		content = container.NewStack(widget.NewButton("Add a Record", func() {
			ui.AddRecord(student.ID)
			w.Close()
		}))
	}

	w.SetContent(content)
	w.Show()
}
