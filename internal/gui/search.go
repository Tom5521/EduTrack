package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/EduTrack/assets"
	"github.com/Tom5521/EduTrack/internal/pkg/sizes"
	"github.com/Tom5521/EduTrack/pkg/data"
	"github.com/Tom5521/EduTrack/pkg/wins"
)

func (ui *ui) SearchStudentsMainWin() {
	w := ui.App.NewWindow("Search students")
	w.Resize(sizes.SearchSize)
	dniEntry := widget.NewEntry()
	dniSearchButton := widget.NewButton("Search", func() {
		studentDNI := dniEntry.Text
		i := data.FindStudentIndexByDNI(studentDNI)
		if i == -1 {
			wins.ErrWin(ui.App, "Student not found!")
			return
		}
		student := data.Students[i]
		ui.LoadStudentInfo(&student)
	})

	byDNIContainer := container.NewVBox(
		widget.NewLabel("Enter Student DNI:"),
		dniEntry,
		dniSearchButton,
	)

	nameEntry := widget.NewEntry()
	nameSearchButton := widget.NewButton("Search", func() {
		studentName := nameEntry.Text
		i := data.FindStudentIndexByName(studentName)
		if i == -1 {
			wins.ErrWin(ui.App, "Student not found!")
			return
		}
		student := data.Students[i]
		ui.LoadStudentInfo(&student)
	})

	byNameContainer := container.NewVBox(
		widget.NewLabel("Enter Student name:"),
		nameEntry,
		nameSearchButton,
	)

	content := container.NewAppTabs(
		container.NewTabItem("Search by name", byNameContainer),
		container.NewTabItem("Search by DNI", byDNIContainer),
	)

	w.SetContent(content)
	w.Show()
}

func (ui *ui) SearchGradesMainWin() {

}

func (ui *ui) SearchRecordsMainWin() {

}

func (ui *ui) SearchStudentGradesMainWin() {

}

func (ui *ui) SearchMainWin() {
	const size1, size2 float32 = 300, 400
	w := ui.App.NewWindow("Search")
	w.Resize(fyne.NewSize(size1, size2))
	w.SetIcon(assets.Lens1)
	const gridNumber int = 1

	searchStudentsButton := widget.NewButtonWithIcon("Search in students", assets.Lens1, func() {
		ui.SearchStudentsMainWin()
	})
	searchGradesButton := widget.NewButtonWithIcon("Search in grades", assets.Lens1, func() {
		ui.SearchGradesMainWin()
	})
	searchRecordsButton := widget.NewButtonWithIcon("Search in records", assets.Lens1, func() {
		ui.SearchRecordsMainWin()
	})
	searchStudentGradesButton := widget.NewButtonWithIcon("Search in student grades", assets.Lens1, func() {
		ui.SearchStudentGradesMainWin()
	})
	content := container.NewAdaptiveGrid(gridNumber,
		searchStudentsButton,
		searchGradesButton,
		searchRecordsButton,
		searchStudentGradesButton,
	)

	w.SetContent(content)

	w.Show()
}
