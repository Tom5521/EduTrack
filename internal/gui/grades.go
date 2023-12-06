package gui

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/EduTrack/assets"
	"github.com/Tom5521/EduTrack/internal/pkg/sizes"
	"github.com/Tom5521/EduTrack/pkg/data"
	"github.com/Tom5521/EduTrack/pkg/wins"
)

const editString string = "Edit "

func (ui ui) GetGradesList(grades *[]data.Grade) *widget.List {
	list := widget.NewList(
		func() int {
			return len(*grades)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			mod := *grades
			o.(*widget.Label).SetText(mod[i].Name)
		},
	)

	return list
}

func (ui *ui) EditGrade(g *data.Grade) {
	window := ui.App.NewWindow(editString + g.Name)
	window.Resize(sizes.FormSize)

	nameEntry := widget.NewEntry()
	nameEntry.SetText(g.Name)

	priceEntry := widget.NewEntry()
	priceEntry.SetText(g.Price)

	infoEntry := widget.NewMultiLineEntry()
	infoEntry.SetText(g.Info)

	form := widget.NewForm(
		widget.NewFormItem("Name:", nameEntry),
		widget.NewFormItem("Price:", priceEntry),
		widget.NewFormItem("Info:", infoEntry),
	)

	form.OnSubmit = func() {
		newGrade := data.Grade{
			Name:  nameEntry.Text,
			Price: priceEntry.Text,
			Info:  infoEntry.Text,
		}
		err := g.Edit(&newGrade)
		if err != nil {
			log.Println(err)
			wins.ErrWin(ui.App, err.Error())
		}
		window.Close()
	}

	window.SetContent(form)

	window.Show()
}

func (ui *ui) GradeDetailsWin(g *data.Grade) {
	window := ui.App.NewWindow(g.Name)
	window.Resize(sizes.FormSize)

	editButton := widget.NewButton("Edit", func() {
		ui.EditGrade(g)
		window.Close()
	})
	deleteButton := widget.NewButton("Delete", func() {
		err := data.Delete(g)
		if err != nil {
			wins.ErrWin(ui.App, err.Error())
		}
		ui.GradesList = ui.GetGradesList(&data.Grades)
		ui.GradesList.Refresh()
		window.Close()
	})
	const gridNumber int = 2

	form := widget.NewForm(
		widget.NewFormItem("Name:", widget.NewLabel(g.Name)),
		widget.NewFormItem("Price:", widget.NewLabel(g.Price)),
		widget.NewFormItem("Info", widget.NewLabel(g.Info)),
		widget.NewFormItem("", container.NewAdaptiveGrid(gridNumber, deleteButton, editButton)),
	)

	window.SetContent(form)
	window.Show()
}

func (ui *ui) AddGrade() {
	window := ui.App.NewWindow("Add a grade")
	window.Resize(sizes.FormSize)
	gradeEntry := widget.NewEntry()
	priceEntry := widget.NewEntry()
	infoEntry := widget.NewMultiLineEntry()

	gradeFormInput := widget.NewFormItem("Grade name:", gradeEntry)
	priceFormInput := widget.NewFormItem("Price per moth:", priceEntry)
	infoFormInput := widget.NewFormItem("Grade Info:", infoEntry)

	form := widget.NewForm(
		gradeFormInput,
		priceFormInput,
		infoFormInput,
	)
	form.OnSubmit = func() {
		if gradeEntry.Text == "" {
			wins.ErrWin(ui.App, "Grade name entry is empty")
			return
		}
		if priceEntry.Text == "" {
			wins.ErrWin(ui.App, "Info entry is empty")
			return
		}
		if func() bool {
			for _, grade := range data.Grades {
				if grade.Name == gradeEntry.Text {
					return true
				}
			}
			return false
		}() {
			wins.ErrWin(ui.App, "This grade already exists!")
			return
		}
		newGrade := data.Grade{
			Name:  gradeEntry.Text,
			Info:  infoEntry.Text,
			Price: priceEntry.Text,
		}
		err := data.AddGrade(&newGrade)
		if err != nil {
			wins.ErrWin(ui.App, err.Error())
		}
		ui.GradesList = ui.GetGradesList(&data.Grades)
		window.Close()
	}
	content := container.NewVBox(form)
	window.SetContent(content)
	window.Show()
}

func (ui *ui) GradesMainWin() {
	w := ui.App.NewWindow("Grades")
	w.Resize(sizes.ListSize)

	selected := -1

	list := ui.GetGradesList(&data.Grades)

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(assets.Plus, func() { ui.AddGrade() }),
		widget.NewToolbarAction(assets.DeleteGrade, func() {
			if selected == -1 {
				return
			}
			err := data.Delete(data.Grades[selected])
			if err != nil {
				wins.ErrWin(ui.App, err.Error())
				return
			}
			list.Refresh()
			selected = -1
		}),
		widget.NewToolbarAction(assets.ShowGrades, func() {
			if selected == -1 {
				return
			}
			ui.GradeDetailsWin(&data.Grades[selected])
		}),
		widget.NewToolbarAction(assets.Edit, func() {
			if selected == -1 {
				return
			}
			ui.EditGrade(&data.Grades[selected])
			list.Refresh()
		}),
	)

	list.OnSelected = func(id widget.ListItemID) {
		selected = id
	}

	content := container.NewBorder(toolbar, nil, nil, nil, list)
	w.SetIcon(assets.ShowGrades)
	w.SetContent(content)

	w.Show()
}
