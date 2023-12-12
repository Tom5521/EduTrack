package gui

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/EduTrack/assets"
	"github.com/Tom5521/EduTrack/internal/pkg/sizes"
	"github.com/Tom5521/EduTrack/pkg/data"
	"github.com/Tom5521/EduTrack/pkg/wins"
)

var (
	CourseWins   map[string]string
	CourseLocals map[string]string
)

func (ui ui) GetCoursesList(courses *[]data.Course) *widget.List {
	list := widget.NewList(
		func() int {
			return len(*courses)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			mod := *courses
			o.(*widget.Label).SetText(mod[i].Name)
		},
	)

	return list
}

func (ui *ui) EditCourse(c *data.Course) {
	window := ui.App.NewWindow(fmt.Sprintf(CourseWins["Edit X"], c.Name))
	window.Resize(sizes.FormSize)

	nameEntry := widget.NewEntry()
	nameEntry.SetText(c.Name)

	priceEntry := widget.NewEntry()
	priceEntry.SetText(c.Price)

	infoEntry := widget.NewMultiLineEntry()
	infoEntry.SetText(c.Info)

	form := widget.NewForm(
		widget.NewFormItem(CourseLocals["Name"], nameEntry),
		widget.NewFormItem(CourseLocals["Price"], priceEntry),
		widget.NewFormItem(CourseLocals["Info"], infoEntry),
	)

	form.OnSubmit = func() {
		newGrade := data.Course{
			Name:  nameEntry.Text,
			Price: priceEntry.Text,
			Info:  infoEntry.Text,
		}
		err := c.Edit(&newGrade)
		if err != nil {
			log.Println(err)
			wins.ErrWin(ui.App, err.Error())
		}
		window.Close()
	}

	window.SetContent(form)
	window.Show()
}

func (ui *ui) CourseDetailsWin(c *data.Course) {
	window := ui.App.NewWindow(c.Name)
	window.Resize(sizes.FormSize)

	editButton := widget.NewButton(locale.GeneralWords["Edit"], func() {
		ui.EditCourse(c)
		window.Close()
	})
	deleteButton := widget.NewButton(locale.GeneralWords["Delete"], func() {
		err := data.Delete(c)
		if err != nil {
			wins.ErrWin(ui.App, err.Error())
		}
		ui.CoursesList = ui.GetCoursesList(&data.Courses)
		ui.CoursesList.Refresh()
		window.Close()
	})
	const gridNumber int = 2

	form := widget.NewForm(
		widget.NewFormItem(CourseLocals["Name"], widget.NewLabel(c.Name)),
		widget.NewFormItem(CourseLocals["Price"], widget.NewLabel(c.Price)),
		widget.NewFormItem(CourseLocals["Info"], widget.NewLabel(c.Info)),
		widget.NewFormItem("", container.NewAdaptiveGrid(gridNumber, deleteButton, editButton)),
	)

	window.SetContent(form)
	window.Show()
}

func (ui *ui) AddCourse() {
	window := ui.App.NewWindow(CourseWins["Add a Course"])
	window.Resize(sizes.FormSize)
	courseEntry := widget.NewEntry()
	priceEntry := widget.NewEntry()
	infoEntry := widget.NewMultiLineEntry()

	coruseFormInput := widget.NewFormItem(CourseLocals["Name"], courseEntry)
	priceFormInput := widget.NewFormItem(CourseLocals["Price"], priceEntry)
	infoFormInput := widget.NewFormItem(CourseLocals["Info"], infoEntry)

	form := widget.NewForm(
		coruseFormInput,
		priceFormInput,
		infoFormInput,
	)
	form.OnSubmit = func() {
		if courseEntry.Text == "" {
			wins.ErrWin(ui.App, locale.Errors["Couse name entry is empty"])
			return
		}
		if priceEntry.Text == "" {
			wins.ErrWin(ui.App, locale.Errors["Info entry is empty"])
			return
		}
		if func() bool {
			for _, grade := range data.Courses {
				if grade.Name == courseEntry.Text {
					return true
				}
			}
			return false
		}() {
			wins.ErrWin(ui.App, locale.Errors["This course already exists"])
			return
		}
		newGrade := data.Course{
			Name:  courseEntry.Text,
			Info:  infoEntry.Text,
			Price: priceEntry.Text,
		}
		err := data.AddCourse(&newGrade)
		if err != nil {
			wins.ErrWin(ui.App, err.Error())
		}
		ui.CoursesList = ui.GetCoursesList(&data.Courses)
		window.Close()
	}
	content := container.NewVBox(form)
	window.SetContent(content)
	window.Show()
}

func (ui *ui) CoursesMainWin() {
	w := ui.App.NewWindow(CourseWins["Courses"])
	w.Resize(sizes.ListSize)

	selected := -1

	list := ui.GetCoursesList(&data.Courses)

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(assets.Plus, func() {
			ui.AddCourse()
			list.Refresh()
		}),
		widget.NewToolbarAction(assets.DeleteCourse, func() {
			if selected == -1 {
				return
			}
			err := data.Delete(data.Courses[selected])
			if err != nil {
				wins.ErrWin(ui.App, err.Error())
				return
			}
			list.UnselectAll()
			selected = -1
		}),
		widget.NewToolbarAction(assets.ShowCourses, func() {
			if selected == -1 {
				return
			}
			ui.CourseDetailsWin(&data.Courses[selected])
		}),
		widget.NewToolbarAction(assets.Edit, func() {
			if selected == -1 {
				return
			}
			ui.EditCourse(&data.Courses[selected])
			list.Refresh()
		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(assets.Refresh, func() {
			list.UnselectAll()
			list.Refresh()
		}),
	)

	list.OnSelected = func(id widget.ListItemID) {
		selected = id
	}

	content := container.NewBorder(toolbar, nil, nil, nil, list)
	w.SetIcon(assets.ShowCourses)
	w.SetContent(content)

	w.Show()
}
