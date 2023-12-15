package gui

import (
	"fyne.io/fyne/v2"
	"github.com/Tom5521/EduTrack/pkg/data"
)

func (ui *ui) MainMenu() *fyne.MainMenu {
	// Create the main menu
	menu := fyne.NewMainMenu(
		fyne.NewMenu(po.Get("File")),
		fyne.NewMenu(po.Get("Edit"),
			fyne.NewMenuItem(po.Get("Reload data"), func() {
				data.LoadEverything()
			}),
		),
		fyne.NewMenu(po.Get("Search"),
			fyne.NewMenuItem(po.Get("Search"), func() {
				ui.SearchMainWin()
			}),
			fyne.NewMenuItem(po.Get("Search Students"), func() {
				ui.SearchStudentsMainWin()
			}),
			fyne.NewMenuItem(po.Get("Search Records"), func() {
				ui.SearchRecordsMainWin()
			}),
			fyne.NewMenuItem(po.Get("Search Courses"), func() {
				ui.SearchCoursesMainWin()
			}),
		),
		fyne.NewMenu(po.Get("Help"),
			fyne.NewMenuItem(po.Get("About"), func() {
				ui.AboutWin()
			}),
		),
	)

	return menu
}
