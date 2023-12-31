package gui

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/EduTrack/assets"
	"github.com/Tom5521/EduTrack/pkg/data"
	"github.com/Tom5521/EduTrack/pkg/wins"
)

func (ui *ui) MainWin() {
	w := ui.App.NewWindow("EduTrack")
	w.SetMaster()
	wins.MaximizeWin(w)
	w.SetMainMenu(ui.MainMenu())

	var selected = -1
	ui.StudentList = ui.GetStudentsList(&data.Students)
	ui.StudentList.OnSelected = func(id widget.ListItemID) {
		selected = id
		ui.LoadStudentInfo(&data.Students[id])
	}
	ui.StudentTab = ui.TemplateUser()

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(assets.AddUser, func() {
			c := StudentForm{}
			c.Add = true
			ui.StudentForm(c)
		}),
		widget.NewToolbarAction(assets.DeleteStudent, func() {
			if selected == -1 {
				return
			}
			d := dialog.NewConfirm(
				po.Get("Please Confirm"),
				po.Get("Do you want to delete the student?"),
				func(b bool) {
					if !b {
						err := data.Delete(data.Students[selected])
						if err != nil {
							wins.ErrWin(ui.App, err.Error())
						}
						ui.StudentList.UnselectAll()
						ui.StudentTab.Objects = ui.TemplateUser().Objects
						ui.StudentTab.Refresh()
						selected = -1
					}
				},
				w,
			)
			d.SetConfirmText(po.Get("No"))
			d.SetDismissText(po.Get("Yes"))
			d.Show()
		}),
		widget.NewToolbarAction(assets.Edit, func() {
			if selected == -1 {
				return
			}
			s := &data.Students[selected]
			c := StudentForm{}
			c.Edit.Enable = true
			c.Edit.Student = s
			ui.StudentForm(c)
		}),
		widget.NewToolbarAction(assets.Info, func() {
			if selected == -1 {
				return
			}
			ui.StudentDetailsWin(&data.Students[selected])
		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(assets.ShowCourses, ui.CoursesMainWin),
		widget.NewToolbarAction(assets.Lens1, ui.SearchMainWin),
	)
	tabCont := container.NewBorder(nil, nil, nil, widget.NewSeparator(), ui.StudentTab)
	listCont := container.NewBorder(toolbar, nil, nil, nil, ui.StudentList)
	mainContainer := container.NewBorder(nil, nil, tabCont, nil, listCont)

	w.SetContent(mainContainer)
	w.ShowAndRun()
}
