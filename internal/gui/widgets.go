package gui

import (
	"fyne.io/fyne/v2"
	"github.com/Tom5521/EduTrack/pkg/data"
)

func (ui *ui) MainMenu() *fyne.MainMenu {
	// Create the main menu
	menu := fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("Add Student", func() {
				ui.AddStudentForm()
			}),
		),
		fyne.NewMenu("Edit",
			fyne.NewMenuItem("Reload data", func() {
				data.LoadEverything()
			}),
		),
		fyne.NewMenu("Help",
			fyne.NewMenuItem("About", func() {
				ui.AboutWin()
			}),
		),
	)
	return menu
}
