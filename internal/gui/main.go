package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/EduTrack/assets"
	"github.com/Tom5521/EduTrack/locales"
	"github.com/Tom5521/EduTrack/pkg/data"
	"github.com/Tom5521/EduTrack/pkg/themes"
	"github.com/Tom5521/EduTrack/pkg/wins"
	"github.com/leonelquinteros/gotext"

	xtheme "fyne.io/x/fyne/theme"
)

var po *gotext.Po

type ui struct {
	App         fyne.App
	StudentTab  *fyne.Container
	StudentList *widget.List
	CoursesList *widget.List
}

type Gui struct {
	ui
}

func (ui *ui) setTheme() {
	var th fyne.Theme
	switch data.Config.Theme {
	case "Adwaita":
		th = xtheme.AdwaitaTheme()
	case "DarkRed":
		th = themes.DarkRed{}
	case "DarkBlue":
		th = themes.DarkBlue{}
	case "SimpleRed":
		th = themes.SimpleRed{}
	case "Default":
		return
	default:
		return
	}
	ui.App.Settings().SetTheme(th)
}

func Init() *Gui {
	ui := &Gui{}
	ui.App = app.New()
	ui.setTheme()
	assets.Load()
	po = locales.GetPo(data.Config.Lang)
	return ui
}

func (ui *ui) MainWin() {
	w := ui.App.NewWindow("EduTrack")
	w.SetMaster()
	wins.MaximizeWin(w)
	w.SetMainMenu(ui.MainMenu())

	var selected = -1
	ui.StudentList = ui.GetStudentsList(&data.Students, func(id int) {
		selected = id
		ui.LoadStudentInfo(&data.Students[id])
	})
	ui.StudentTab = ui.GetTemplateUser()

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(assets.AddUser, ui.AddStudentForm),
		widget.NewToolbarAction(assets.DeleteStudent, func() {
			if selected == -1 {
				return
			}
			dialog.ShowConfirm(
				po.Get("Please Confirm"),
				po.Get("Do you want to delete the student?"),
				func(b bool) {
					if b {
						err := data.Delete(data.Students[selected])
						if err != nil {
							wins.ErrWin(ui.App, err.Error())
						}
						ui.StudentList.UnselectAll()
						ui.StudentTab.Objects = ui.GetTemplateUser().Objects
						ui.StudentTab.Refresh()
						selected = -1
					}
				},
				w,
			)
		}),
		widget.NewToolbarAction(assets.Edit, func() {
			if selected == -1 {
				return
			}
			ui.EditStudentWindow(&data.Students[selected])
		}),
		widget.NewToolbarAction(assets.Info, func() {
			if selected == -1 {
				return
			}
			ui.StudentDetailsWin(&data.Students[selected])
		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(assets.Lens1, ui.SearchMainWin),
		widget.NewToolbarAction(assets.ShowCourses, ui.CoursesMainWin))
	tabCont := container.NewBorder(nil, nil, nil, widget.NewSeparator(), ui.StudentTab)
	listCont := container.NewBorder(toolbar, nil, nil, nil, ui.StudentList)
	mainContainer := container.NewBorder(nil, nil, tabCont, nil, listCont)

	w.SetContent(mainContainer)
	w.ShowAndRun()
}
