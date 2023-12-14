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

func (ui ui) GetStudentCoursesList(g *[]data.StudentCourse) *widget.List {
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
			if len(data.Courses) != 0 {
				index := data.FindCourseIndexByID(data.StudentCourses[i].CourseID)
				if index == -1 {
					return
				}
				name = data.Courses[index].Name
			}
			o.(*widget.Label).SetText(name)
		},
	)
	return list
}

func (ui *ui) StudentCourseDetailsWin(sc *data.StudentCourse) {
	getCourse := func() *data.Course {
		i := data.FindCourseIndexByID(sc.CourseID)
		return &data.Courses[i]
	}
	c := getCourse()

	window := ui.App.NewWindow(po.Get("Details of %s", c.Name))
	window.Resize(sizes.ListSize)

	courseNameLabel := widget.NewLabel(c.Name)
	coursePricePMLabel := widget.NewLabel(c.Price)
	courseStartLabel := widget.NewLabel(sc.Start)
	courseEndLabel := widget.NewLabel(sc.End)
	courseInfoLabel := widget.NewMultiLineEntry()
	courseInfoLabel.SetText(c.Info)

	nameForm := widget.NewFormItem(po.Get("Name:"), courseNameLabel)
	priceForm := widget.NewFormItem(po.Get("Price:"), coursePricePMLabel)
	startForm := widget.NewFormItem(po.Get("Start:"), courseStartLabel)
	endForm := widget.NewFormItem(po.Get("End:"), courseEndLabel)
	infoForm := widget.NewFormItem(po.Get("Info:"), courseInfoLabel)

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
	form.SubmitText = po.Get("Close")

	window.SetContent(form)
	window.Show()
}

func (ui *ui) StartEndWin(submitFunc func(start, end string)) {
	w := ui.App.NewWindow(po.Get("Select start and end"))
	w.Resize(sizes.StartEndSize)
	startEntry := widget.NewEntry()
	endEntry := widget.NewEntry()
	form := widget.NewForm(
		widget.NewFormItem(po.Get("Start:"), startEntry),
		widget.NewFormItem(po.Get("End:"), endEntry),
	)

	form.OnSubmit = func() {
		submitFunc(startEntry.Text, endEntry.Text)
		w.Close()
	}
	w.SetContent(form)

	w.Show()
}

func getNoAddedCourses(s *data.Student) []data.Course {
	var grades []data.Course
	for _, grade := range data.Courses {
		var found bool
		for _, sg := range s.Courses {
			if grade.ID == sg.CourseID {
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

func (ui *ui) SelectCourseWin(s *data.Student) {
	w := ui.App.NewWindow(po.Get("Select a course"))
	const size1, size2 float32 = 600, 500
	w.Resize(fyne.NewSize(size1, size2))

	// Selection vars
	var addedSelected, toAddSelected = -1, -1
	// Temporal lists
	tmpToAddList := getNoAddedCourses(s)

	// Lists (widgets)
	toAddList := ui.GetCoursesList(&tmpToAddList)
	toAddList.OnSelected = func(id widget.ListItemID) {
		toAddSelected = id
	}
	addedList := ui.GetStudentCoursesList(&s.Courses)
	addedList.OnSelected = func(id widget.ListItemID) {
		addedSelected = id
	}

	addGrade := func() {
		if toAddSelected == -1 {
			return
		}
		ui.StartEndWin(func(start, end string) {
			studentGrade := data.StudentCourse{
				Start:     start,
				End:       end,
				CourseID:  tmpToAddList[toAddSelected].ID,
				StudentID: s.ID,
			}
			err := data.AddStudentCourse(&studentGrade)
			if err != nil {
				wins.ErrWin(ui.App, err.Error())
			}
			tmpToAddList = slices.Delete(tmpToAddList, toAddSelected, toAddSelected+1)
			s.GetCourses()
			addedList.Refresh()
			toAddList.Refresh()
		})
	}
	quitGrade := func() {
		if addedSelected == -1 {
			return
		}
		i := data.FindStudentCourseIndexByID(s.Courses[addedSelected].ID)
		err := data.Delete(data.StudentCourses[i])
		if err != nil {
			wins.ErrWin(ui.App, err.Error())
		}
		s.GetCourses()
		addedList.Refresh()
		tmpToAddList = getNoAddedCourses(s)
		toAddList.Refresh()
	}

	addToButton := widget.NewButton(
		po.Get("Add course to student"),
		addGrade,
	)
	quitToButton := widget.NewButton(
		po.Get("Delete from student"), quitGrade,
	)

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

func (ui *ui) StudentCoursesMainWin(s *data.Student) {
	w := ui.App.NewWindow(po.Get("%s Courses", s.Name))
	const size float32 = 300
	w.Resize(fyne.NewSize(size, size))
	s.GetCourses()
	var selected = -1
	list := ui.GetStudentCoursesList(&s.Courses)
	list.OnSelected = func(id widget.ListItemID) {
		selected = id
	}

	bar := widget.NewToolbar(
		widget.NewToolbarAction(assets.Plus, func() {
			ui.SelectCourseWin(s)
			list.Refresh()
		}),
		widget.NewToolbarAction(assets.Cross, func() {
			if selected == -1 {
				return
			}
			tg := s.Courses[selected]
			i := data.FindStudentCourseIndexByID(tg.ID)
			err := data.Delete(data.StudentCourses[i])
			if err != nil {
				wins.ErrWin(ui.App, err.Error())
			}
			s.GetCourses()
			list.Refresh()
		}),
		widget.NewToolbarAction(assets.Edit, func() {
			if selected == -1 {
				return
			}
			i := data.FindStudentCourseIndexByID(s.Courses[selected].ID)
			ui.EditStudentCourseWin(&data.StudentCourses[i])
		}),
		widget.NewToolbarAction(assets.Info, func() {
			if selected == -1 {
				return
			}
			ui.StudentCourseDetailsWin(&data.StudentCourses[selected])
		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(assets.Refresh, func() {
			s.GetCourses()
			list.Refresh()
		}),
	)

	content := container.NewBorder(bar, nil, nil, nil, list)
	w.SetContent(content)
	w.Show()
}

func (ui *ui) EditStudentCourseWin(sc *data.StudentCourse) {
	i := data.FindCourseIndexByID(sc.CourseID)
	course := data.Courses[i]
	window := ui.App.NewWindow(po.Get("Edit %s", course.Name))
	const size1, size2 float32 = 500, 100
	window.Resize(fyne.NewSize(size1, size2))

	courseNameLabel := widget.NewLabel(course.Name)
	courseSelectButton := widget.NewButton(
		po.Get("Select a new course"), func() {
			w := ui.App.NewWindow(po.Get("Select a course"))
			w.Resize(sizes.ListSize)
			var selected = -1
			list := ui.GetCoursesList(&data.Courses)
			list.OnSelected = func(id widget.ListItemID) {
				selected = id
			}
			addGradeButton := widget.NewButton(po.Get("Select course"), func() {
				if selected == -1 {
					return
				}
				sc.CourseID = data.Courses[selected].ID
				err := sc.Edit(sc)
				if err != nil {
					wins.ErrWin(ui.App, err.Error())
				}
				courseNameLabel.SetText(data.Courses[selected].Name)
				w.Close()
			})
			cancelButton := widget.NewButton(po.Get("Cancel"), func() {
				w.Close()
			})

			const gridNumber int = 2
			buttonsCont := container.NewAdaptiveGrid(gridNumber, cancelButton, addGradeButton)
			content := container.NewBorder(buttonsCont, nil, nil, nil, list)
			w.SetContent(content)
			w.Show()
		})
	const gridNumber = 2
	courseSelectCont := container.NewAdaptiveGrid(gridNumber, courseNameLabel, courseSelectButton)

	startEntry := widget.NewEntry()
	startEntry.SetText(sc.Start)
	endEntry := widget.NewEntry()
	endEntry.SetText(sc.End)

	form := widget.NewForm(
		widget.NewFormItem(po.Get("Name:"), courseSelectCont),
		widget.NewFormItem(po.Get("Start:"), startEntry),
		widget.NewFormItem(po.Get("End:"), endEntry),
	)
	form.OnSubmit = func() {
		sc.Start = startEntry.Text
		sc.End = endEntry.Text
		err := sc.Edit(sc)
		if err != nil {
			wins.ErrWin(ui.App, err.Error())
		}
		window.Close()
	}
	form.OnCancel = func() {
		window.Close()
	}
	form.SubmitText = po.Get("Submit")
	form.CancelText = po.Get("Cancel")

	window.SetContent(form)
	window.Show()
}
