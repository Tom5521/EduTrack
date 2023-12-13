package gui

import (
	"fyne.io/fyne/v2"
	"github.com/Tom5521/EduTrack/pkg/data"
)

func (ui *ui) MainMenu() *fyne.MainMenu {
	// Create the main menu
	menu := fyne.NewMainMenu(
		fyne.NewMenu(po.Get("File"),
			fyne.NewMenuItem(po.Get("Add a Student"), func() {
				ui.AddStudentForm()
			}),
		),
		fyne.NewMenu(po.Get("Edit"),
			fyne.NewMenuItem(po.Get("Reload data"), func() {
				data.LoadEverything()
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
