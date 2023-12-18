package gui

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/EduTrack/assets"
	"github.com/Tom5521/EduTrack/internal/pkg/sizes"
	"github.com/Tom5521/EduTrack/pkg/data"
	"github.com/Tom5521/EduTrack/pkg/widgets"
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
		now := time.Now()
		return now.Format("02/01/2006 15:04")
	}
	var tmpDate = getTimeNow()

	window := ui.App.NewWindow(po.Get("Add a Record"))
	window.Resize(sizes.RecSize)

	recNameLabel := widget.NewLabel(po.Get("Record name:"))
	recnameEntry := widget.NewEntry()
	recnameEntry.SetPlaceHolder(getTimeNow())

	recDateButton := widget.NewButton(po.Get("Select Date"), func() {
		const year int = 2023
		ret, err := zenity.Calendar(po.Get("Select Date from below"),
			zenity.DefaultDate(year, time.December, 1))
		if err != nil {
			wins.ErrWin(ui.App, err.Error())
		}
		tmpDate = ret.Format("02/01/2006")
	})
	detailsLabel := widget.NewLabel(po.Get("Details"))
	recDetails := widget.NewMultiLineEntry()
	recDetails.SetPlaceHolder(po.Get("E.g., The student has not attended"))
	submitButton := widget.NewButton(po.Get("Submit"), func() {
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
	window := ui.App.NewWindow(po.Get("Edit Record"))
	window.Resize(sizes.RecSize)
	i := data.FindRecordIndexByID(recordID)
	if i == -1 {
		wins.ErrWin(ui.App, po.Get("Record ID<%v> not found!", recordID))
		window.Close()
		return
	}
	rec := &data.Records[i]

	i = data.FindStudentIndexByID(rec.StudentID)
	if i == -1 {
		wins.ErrWin(ui.App, po.Get("Student ID<%v> not found!", rec.StudentID))
		window.Close()
		return
	}
	student := &data.Students[i]
	student.GetRecords()

	recNameEntry := widget.NewEntry()
	recNameEntry.SetText(rec.Name)

	recDate := widget.NewLabel(po.Get("Date: %s", rec.Date))
	dateButton := widget.NewButton(po.Get("Select Date"), func() {
		const year, day int = 2023, 1
		ret, err := zenity.Calendar(po.Get("Select a date from below:"),
			zenity.DefaultDate(year, time.December, day))
		if err != nil {
			wins.ErrWin(ui.App, err.Error())
		}
		tmpDate = ret.Format("02/01/2006")
		recDate.SetText(tmpDate)
	})

	detailsLabel := widget.NewLabel(po.Get("Details"))
	recDetails := widget.NewMultiLineEntry()
	recDetails.SetText(rec.Info)

	// FormItems
	const gridNumber int = 2
	recNameForm := widget.NewFormItem(po.Get("Record name:"), recNameEntry)
	recDateForm := widget.NewFormItem(po.Get("Record date:"), container.NewAdaptiveGrid(gridNumber, recDate, dateButton))

	submitFunc := func() {
		err := rec.Edit(&data.Record{Name: recNameEntry.Text, Info: recDetails.Text, Date: tmpDate, StudentID: student.ID})
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
	form.OnCancel = func() {
		window.Close()
	}

	form.SubmitText = po.Get("Submit")
	form.CancelText = po.Get("Cancel")

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
	w := ui.App.NewWindow(po.Get("%s Records", student.Name))
	w.Resize(sizes.FormSize)
	var selected int
	list := ui.GetStudentRecordsList(student)
	bar := widget.NewToolbar(
		widget.NewToolbarAction(assets.Plus, func() {
			ui.AddRecord(student.ID)
			student.GetRecords()
			list.Refresh()
		}),
		widget.NewToolbarAction(assets.Cross, func() {
			if selected == -1 {
				return
			}
			err := data.Delete(data.Records[data.FindRecordIndexByID(student.Records[selected].ID)])
			if err != nil {
				wins.ErrWin(ui.App, err.Error())
			}
			student.GetRecords()
			list.UnselectAll()
			selected = -1
		}),
		widget.NewToolbarAction(assets.Edit, func() {
			if selected == -1 {
				return
			}
			ui.EditRecordData(data.Records[selected].ID)
			student.GetRecords()
			list.Refresh()
		}),
		widget.NewToolbarAction(assets.Info, func() {
			if selected == -1 {
				return
			}
			i := data.FindRecordIndexByID(student.Records[selected].ID)
			ui.RecordDetailsWin(&data.Records[i])
		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(assets.Refresh, func() {
			student.GetRecords()
			list.Refresh()
		}),
	)

	list.OnSelected = func(id widget.ListItemID) {
		selected = id
	}
	content := container.NewBorder(bar, nil, nil, nil, list)

	if len(student.Records) == 0 {
		content = container.NewStack(widget.NewButton(po.Get("Add a Record"), func() {
			ui.AddRecord(student.ID)
			w.Close()
		}))
	}

	w.SetContent(content)
	w.Show()
}

func (ui *ui) RecordDetailsWin(r *data.Record) {
	w := ui.App.NewWindow((po.Get("Details of %s", r.Name)))
	// w.Resize(sizes.RecSize)
	infoEntry := widget.NewMultiLineEntry()
	infoEntry.SetText(r.Info)

	editButton := widget.NewButton(po.Get("Edit"), func() {
		ui.EditRecordData(r.ID)
	})
	deleteButton := widget.NewButton(po.Get("Delete"), func() {
		dialog.ShowConfirm(po.Get("Please Confirm"),
			po.Get("Do you really want to delete the record"), func(b bool) {
				if b {
					err := data.Delete(r)
					if err != nil {
						wins.ErrWin(ui.App, err.Error())
					}
					w.Close()
				}
			},
			w,
		)
	})

	const gridNumber int = 2
	student := data.Students[data.FindStudentIndexByID(r.StudentID)]
	form := widgets.NewForm(
		widget.NewFormItem(po.Get("Student Name:"), widget.NewLabel(student.Name)),
		widget.NewFormItem(po.Get("Name:"), widget.NewLabel(r.Name)),
		widget.NewFormItem(po.Get("Date:"), widget.NewLabel(r.Date)),
		widget.NewFormItem(po.Get("Info:"), infoEntry),
	)
	form.CustomItems = container.NewAdaptiveGrid(gridNumber, deleteButton, editButton)
	form.OnSubmit = func() {
		w.Close()
	}
	form.SubmitText = po.Get("Close")
	form.SubmitText = po.Get("Submit")

	w.SetContent(form)

	w.Show()
}
