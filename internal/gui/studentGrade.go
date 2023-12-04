package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/EduTrack/assets"
	"github.com/Tom5521/EduTrack/pkg/data"
	"github.com/Tom5521/EduTrack/pkg/wins"
)

func (_ ui) GetStudentGradesList(g *[]data.StudentGrade) *widget.List {
	list := widget.NewList(
		func() int {
			return len(*g)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			// mod := *g
			name := data.Grades[data.FindGradeIndexByID(data.StudentGrades[i].GradeID)].Name
			o.(*widget.Label).SetText(name)
		},
	)
	return list
}

func (ui *ui) StudentGradeDetailsWin(sg *data.StudentGrade) {
	getGrade := func() *data.Grade {
		i := data.FindGradeIndexByID(sg.GradeID)
		return &data.Grades[i]
	}
	g := getGrade()

	window := ui.App.NewWindow("Details for " + g.Name)

	gradeNameLabel := widget.NewLabel(g.Name)
	gradePricePMLabel := widget.NewLabel(g.Price)
	gradeStartLabel := widget.NewLabel(sg.Start)
	gradeEndLabel := widget.NewLabel(sg.End)
	gradeInfoLabel := widget.NewMultiLineEntry()
	gradeInfoLabel.SetText(g.Info)
	gradeInfoLabel.Disable()

	editGradeButton := widget.NewButton("Edit Grade", func() {})
	editStudentButton := widget.NewButton("Edit Student", func() {})

	nameForm := widget.NewFormItem("Name:", gradeNameLabel)
	priceForm := widget.NewFormItem("Price:", gradePricePMLabel)
	startForm := widget.NewFormItem("Start:", gradeStartLabel)
	endForm := widget.NewFormItem("End:", gradeEndLabel)
	infoForm := widget.NewFormItem("Info:", gradeInfoLabel)
	const gridNumber int = 2
	editForm := widget.NewFormItem("",
		container.NewAdaptiveGrid(gridNumber,
			editGradeButton,
			editStudentButton,
		),
	)

	form := widget.NewForm(
		nameForm,
		priceForm,
		startForm,
		editForm,
		endForm,
		infoForm,
		editForm,
	)

	window.SetContent(form)
	window.Show()
}

func (ui *ui) AddStudentGradeWin(student *data.Student) {
	w := ui.App.NewWindow("Add Grade")

	currrentGradesLabel := widget.NewLabel("Current labels")
	currrentGradesLabel.Alignment = fyne.TextAlignCenter
	currentGradesList := ui.GetStudentGradesList(&student.Grades)
	bar := widget.NewToolbar(
		widget.NewToolbarAction(assets.AddUser, func() {}),
	)

	barCont := container.NewVBox(bar, currrentGradesLabel)

	mainBox := container.NewBorder(barCont, nil, nil, nil, currentGradesList)

	w.SetContent(mainBox)
	w.Show()
}

func (ui *ui) StudentGradesMainWin(s *data.Student) {
	w := ui.App.NewWindow(s.Name + " Grades")
	s.GetGrades()
	var selected = -1
	list := ui.GetStudentGradesList(&s.Grades)
	list.OnSelected = func(id widget.ListItemID) {
		selected = id
	}

	bar := widget.NewToolbar(
		widget.NewToolbarAction(assets.Plus, func() {
			ui.AddStudentGradeWin(s)
			list.Refresh()
		}),
		widget.NewToolbarAction(assets.Cross, func() {
			if selected == -1 {
				return
			}
			tg := s.Grades[selected]
			i := data.FindGradeIndexByID(tg.GradeID)
			err := data.Delete(data.Grades[i])
			if err != nil {
				wins.ErrWin(ui.App, err.Error())
			}
			s.GetGrades()
			list.Refresh()
		}),
	)

	content := container.NewBorder(bar, nil, nil, nil, list)

	w.SetContent(content)

	w.Show()
}
