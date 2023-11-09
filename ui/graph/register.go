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
		student.Register = append(student.Register, NewReg)
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
		content = container.NewStack(RegisterList)
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
			wins.ErrWin(app, err.Error())
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
		data.SaveStudentsData()
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
