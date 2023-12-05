package gui

import (
	"fmt"
	"slices"

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

func (ui *ui) StartEndWin(start, end *string) {
	w := ui.App.NewWindow("Select start and end")
	startEntry := widget.NewEntry()
	endEntry := widget.NewEntry()
	form := widget.NewForm(
		widget.NewFormItem("Start:", startEntry),
		widget.NewFormItem("End:", endEntry),
	)

	form.OnSubmit = func() {
		*start = startEntry.Text
		*end = endEntry.Text
		w.Close()
	}
	w.SetContent(form)

	w.Show()
}

func (ui *ui) SelectGradeWin(s *data.Student) {
	w := ui.App.NewWindow("Select a grade!")

	// Selection vars
	var addedSelected, toAddSelected = -1, -1
	fmt.Println(addedSelected)

	// Temporal lists
	tmpToAddList := data.Grades

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
		var start, end string
		ui.StartEndWin(&start, &end)
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
		slices.Delete(tmpToAddList, toAddSelected, toAddSelected+1)
		s.GetGrades()
		addedList.UnselectAll()
	}
	quitGrade := func() {

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

func (ui *ui) AddStudentGradeWin(student *data.Student) {
	w := ui.App.NewWindow("Add Grade")

	currrentGradesLabel := widget.NewLabel("Current labels")
	currrentGradesLabel.Alignment = fyne.TextAlignCenter
	currentGradesList := ui.GetStudentGradesList(&student.Grades)

	var selectedID int

	bar := widget.NewToolbar(
		widget.NewToolbarAction(assets.Plus, func() {
			ui.SelectGradeWin(student)
			fmt.Println(selectedID)
		}),
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
			ui.SelectGradeWin(s)
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
