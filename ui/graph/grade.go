package graph

import (
	"EduTrack/data"
	"EduTrack/pkg/wins"
	"EduTrack/ui/sizes"
	"strings"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func AddGrade() {
	window := app.NewWindow("Add a grade")
	window.Resize(sizes.FormSize)
	gradeEntry := widget.NewEntry()
	priceEntry := widget.NewEntry()
	InfoEntry := widget.NewMultiLineEntry()
	StartEntry := widget.NewEntry()
	EndEntry := widget.NewEntry()

	gradeFormInput := widget.NewFormItem("Grade name:", gradeEntry)
	priceFormInput := widget.NewFormItem("Price per moth:", priceEntry)
	infoFormInput := widget.NewFormItem("Grade Info:", InfoEntry)
	StartFormInput := widget.NewFormItem("Start date:", StartEntry)
	EndFormInput := widget.NewFormItem("End date:", EndEntry)

	form := widget.NewForm(
		gradeFormInput,
		priceFormInput,
		StartFormInput,
		EndFormInput,
		infoFormInput,
	)
	form.OnSubmit = func() {
		if gradeEntry.Text == "" {
			wins.ErrWin(app, "Grade name entry is empty")
			return
		}
		if priceEntry.Text == "" {
			wins.ErrWin(app, "Info entry is empty")
			return
		}
		if StartEntry.Text == "" {
			wins.ErrWin(app, "Start date entry is empty")
			return
		}
		if EndEntry.Text == "" {
			wins.ErrWin(app, "End date entry is emty")
			return
		}
		if strings.Contains(strings.Join(data.GetGradesNames(), " "), gradeEntry.Text) {
			wins.ErrWin(app, "This grade already exists!")
			return
		}

		data.Grades = append(data.Grades, data.Grade{
			Name:  gradeEntry.Text,
			Info:  InfoEntry.Text,
			Start: StartEntry.Text,
			End:   EndEntry.Text,
			Price: priceEntry.Text,
		})

		data.SaveGradesData()
		window.Close()

	}

	content := container.NewVBox(form)
	window.SetContent(content)
	window.Show()
}
