package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/EduTrack/assets"
	"github.com/Tom5521/EduTrack/pkg/data"
	"github.com/Tom5521/EduTrack/pkg/themes"
	"github.com/Tom5521/EduTrack/pkg/wins"

	xtheme "fyne.io/x/fyne/theme"
)

type ui struct {
	App         fyne.App
	StudentTab  *fyne.Container
	StudentList *widget.List
	GradesList  *widget.List
}

type Gui struct {
	ui
}

func Init() *Gui {
	ui := &Gui{}
	ui.App = app.New()
	var th fyne.Theme
	switch data.Config.Theme {
	case "Adwaita":
		th = xtheme.AdwaitaTheme()
	case "DarkRed":
		th = themes.DarkRed{}
	case "DarkBlue":
		th = themes.DarkBlue{}
	}
	ui.App.Settings().SetTheme(th)
	assets.Load()
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
			ui.DeleteStudentWin(&data.Students[selected])
			selected = -1
		}),
		widget.NewToolbarAction(assets.Edit, func() {
			if selected == -1 {
				return
			}
			ui.EditFormWindow(&data.Students[selected])
		}),

		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(assets.Lens1, ui.SearchWindow),
		widget.NewToolbarAction(assets.ShowGrades, ui.GradesMainWin))

	tabCont := container.NewBorder(nil, nil, nil, widget.NewSeparator(), ui.StudentTab)
	listCont := container.NewBorder(toolbar, nil, nil, nil, ui.StudentList)
	mainContainer := container.NewBorder(nil, nil, tabCont, nil, listCont)

	w.SetContent(mainContainer)

	w.ShowAndRun()
}
