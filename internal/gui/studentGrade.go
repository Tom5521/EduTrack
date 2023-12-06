package gui

import (
	"slices"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/EduTrack/assets"
	"github.com/Tom5521/EduTrack/internal/pkg/sizes"
	"github.com/Tom5521/EduTrack/pkg/data"
	"github.com/Tom5521/EduTrack/pkg/wins"
)

func (_ ui) GetStudentGradesList(g *[]data.StudentGrade) *widget.List {
	list := widget.NewList(
		func() int {
			return len(*g)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			// mod := *g
			var name string
			if len(data.Grades) != 0 {
				i := data.FindGradeIndexByID(data.StudentGrades[i].GradeID)
				if i == -1 {
					return
				}
				name = data.Grades[i].Name
			}
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
	window.Resize(sizes.ListSize)

	gradeNameLabel := widget.NewLabel(g.Name)
	gradePricePMLabel := widget.NewLabel(g.Price)
	gradeStartLabel := widget.NewLabel(sg.Start)
	gradeEndLabel := widget.NewLabel(sg.End)
	gradeInfoLabel := widget.NewMultiLineEntry()
	gradeInfoLabel.SetText(g.Info)

	nameForm := widget.NewFormItem("Name:", gradeNameLabel)
	priceForm := widget.NewFormItem("Price:", gradePricePMLabel)
	startForm := widget.NewFormItem("Start:", gradeStartLabel)
	endForm := widget.NewFormItem("End:", gradeEndLabel)
	infoForm := widget.NewFormItem("Info:", gradeInfoLabel)

	form := widget.NewForm(
		nameForm,
		priceForm,
		startForm,
		endForm,
		infoForm,
	)
	form.OnSubmit = func() {
		window.Close()
	}
	form.SubmitText = "Close"

	window.SetContent(form)
	window.Show()
}

func (ui *ui) StartEndWin(submitFunc func(start, end string)) {
	w := ui.App.NewWindow("Select start and end")
	w.Resize(sizes.StartEndSize)
	startEntry := widget.NewEntry()
	endEntry := widget.NewEntry()
	form := widget.NewForm(
		widget.NewFormItem("Start:", startEntry),
		widget.NewFormItem("End:", endEntry),
	)

	form.OnSubmit = func() {
		submitFunc(startEntry.Text, endEntry.Text)
		w.Close()
	}
	w.SetContent(form)

	w.Show()
}

func (ui *ui) SelectGradeWin(s *data.Student) {
	w := ui.App.NewWindow("Select a grade!")
	w.Resize(fyne.NewSize(600, 500))

	// Selection vars
	var addedSelected, toAddSelected = -1, -1

	getNoAddedGrades := func() []data.Grade {
		var grades []data.Grade
		for _, grade := range data.Grades {
			var found bool
			for _, sg := range s.Grades {
				if grade.ID == sg.GradeID {
					found = true
					break
				}
			}
			if !found {
				grades = append(grades, grade)
			}
		}
		return grades
	}

	// Temporal lists
	tmpToAddList := getNoAddedGrades()

	// Lists (widgets)
	toAddList := ui.GetGradesList(&tmpToAddList)
	toAddList.OnSelected = func(id widget.ListItemID) {
		toAddSelected = id
	}
	addedList := ui.GetStudentGradesList(&s.Grades)
	addedList.OnSelected = func(id widget.ListItemID) {
		addedSelected = id
	}

	addGrade := func() {
		if toAddSelected == -1 {
			return
		}
		ui.StartEndWin(func(start, end string) {
			studentGrade := data.StudentGrade{
				Start:     start,
				End:       end,
				GradeID:   tmpToAddList[toAddSelected].ID,
				StudentID: s.ID,
			}
			err := data.AddStudentGrade(&studentGrade)
			if err != nil {
				wins.ErrWin(ui.App, err.Error())
			}
			tmpToAddList = slices.Delete(tmpToAddList, toAddSelected, toAddSelected+1)
			s.GetGrades()
			addedList.Refresh()
			toAddList.Refresh()
		})
	}
	quitGrade := func() {
		if addedSelected == -1 {
			return
		}
		i := data.FindStudentGradeIndexByID(s.Grades[addedSelected].ID)
		err := data.Delete(data.StudentGrades[i])
		if err != nil {
			wins.ErrWin(ui.App, err.Error())
		}
		s.GetGrades()
		addedList.Refresh()
		tmpToAddList = getNoAddedGrades()
		toAddList.Refresh()
	}

	addToButton := widget.NewButton("Add grade to student", addGrade)
	quitToButton := widget.NewButton("Delete from student", quitGrade)

	// Layout
	const gridNumber int = 1
	toAddCont := container.NewBorder(
		container.NewAdaptiveGrid(gridNumber, addToButton),
		nil, nil, nil,
		toAddList,
	)
	toQuitCont := container.NewBorder(
		container.NewAdaptiveGrid(gridNumber, quitToButton),
		nil, nil, nil,
		addedList,
	)

	split := container.NewHSplit(toAddCont, toQuitCont)

	// content := container.NewBorder(container.NewAdaptiveGrid(gridNumber, addToButton))
	w.SetContent(split)
	w.Show()
}

func (ui *ui) StudentGradesMainWin(s *data.Student) {
	w := ui.App.NewWindow(s.Name + " Grades")
	w.Resize(fyne.NewSize(300, 300))
	s.GetGrades()
	var selected = -1
	list := ui.GetStudentGradesList(&s.Grades)
	list.OnSelected = func(id widget.ListItemID) {
		selected = id
	}

	bar := widget.NewToolbar(
		widget.NewToolbarAction(assets.Plus, func() {
			ui.SelectGradeWin(s)
			list.Refresh()
		}),
		widget.NewToolbarAction(assets.Cross, func() {
			if selected == -1 {
				return
			}
			tg := s.Grades[selected]
			i := data.FindStudentGradeIndexByID(tg.ID)
			err := data.Delete(data.StudentGrades[i])
			if err != nil {
				wins.ErrWin(ui.App, err.Error())
			}
			s.GetGrades()
			list.Refresh()
		}),
		widget.NewToolbarAction(assets.Edit, func() {
			if selected == -1 {
				return
			}
			i := data.FindStudentGradeIndexByID(s.Grades[selected].ID)
			ui.EditStudentGradeWin(&data.StudentGrades[i])

		}),
		widget.NewToolbarAction(assets.Info, func() {
			if selected == -1 {
				return
			}
			ui.StudentGradeDetailsWin(&data.StudentGrades[selected])
		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(assets.Refresh, func() {
			s.GetGrades()
			list.Refresh()
		}),
	)

	content := container.NewBorder(bar, nil, nil, nil, list)
	w.SetContent(content)
	w.Show()
}

func (ui *ui) EditStudentGradeWin(g *data.StudentGrade) {
	i := data.FindGradeIndexByID(g.GradeID)
	grade := data.Grades[i]
	window := ui.App.NewWindow("Edit " + grade.Name)
	window.Resize(fyne.NewSize(500, 100))

	gradeNameLabel := widget.NewLabel(grade.Name)
	gradeSelectButton := widget.NewButton("Select a new grade", func() {
		w := ui.App.NewWindow("Select a grade")
		w.Resize(sizes.ListSize)
		var selected = -1
		list := ui.GetGradesList(&data.Grades)
		list.OnSelected = func(id widget.ListItemID) {
			selected = id
		}
		addGradeButton := widget.NewButton("Select grade", func() {
			if selected == -1 {
				return
			}
			g.GradeID = data.Grades[selected].ID
			err := g.Edit(g)
			if err != nil {
				wins.ErrWin(ui.App, err.Error())
			}
			gradeNameLabel.SetText(data.Grades[selected].Name)
			w.Close()
		})
		cancelButton := widget.NewButton("Cancel", func() {
			w.Close()
		})

		const gridNumber int = 2
		buttonsCont := container.NewAdaptiveGrid(gridNumber, cancelButton, addGradeButton)
		content := container.NewBorder(buttonsCont, nil, nil, nil, list)
		w.SetContent(content)
		w.Show()
	})
	const gridNumber = 2
	gradeSelectCont := container.NewAdaptiveGrid(gridNumber, gradeNameLabel, gradeSelectButton)

	startEntry := widget.NewEntry()
	startEntry.SetText(g.Start)
	endEntry := widget.NewEntry()
	endEntry.SetText(g.End)

	form := widget.NewForm(
		widget.NewFormItem("Current grade:", gradeSelectCont),
		widget.NewFormItem("Start:", startEntry),
		widget.NewFormItem("End:", endEntry),
	)
	form.OnSubmit = func() {
		g.Start = startEntry.Text
		g.End = endEntry.Text
		err := g.Edit(g)
		if err != nil {
			wins.ErrWin(ui.App, err.Error())
		}
		window.Close()
	}
	form.OnCancel = func() {
		window.Close()
	}

	window.SetContent(form)
	window.Show()
}
