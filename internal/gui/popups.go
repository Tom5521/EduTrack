package gui

import (
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/EduTrack/assets"
	"github.com/Tom5521/EduTrack/internal/pkg/sizes"
	"github.com/Tom5521/EduTrack/pkg/data"
	"github.com/Tom5521/EduTrack/pkg/wins"
)

func (ui *ui) AboutWin() {
	window := ui.App.NewWindow("About")
	label1 := widget.NewLabel("Created by:")
	link, _ := url.Parse("https://github.com/Tom5521")
	gitLabel := widget.NewHyperlink("Tom5521", link)
	licenceLabel := widget.NewLabel("Under the MIT license")
	authorCont := container.NewHBox(label1, gitLabel)
	logo := canvas.NewImageFromResource(assets.App)
	const w, h float32 = 300, 300
	logo.SetMinSize(fyne.NewSize(w, h))
	vbox1 := container.NewVBox(
		authorCont,
		licenceLabel,
		logo,
	)
	window.SetContent(vbox1)
	window.Show()
}

func (ui *ui) SearchWindow() {
	w := ui.App.NewWindow("Search Student")
	w.Resize(sizes.SearchSize)
	entry := widget.NewEntry()
	searchButton := widget.NewButton("Search", func() {
		studentDNI := entry.Text
		i := data.FindStudentIndexByDNI(studentDNI)
		if i == -1 {
			wins.ErrWin(ui.App, "Student not found!")
			return
		}
		student := data.Students[i]
		ui.LoadStudentInfo(&student)
		w.Close()
	})

	content := container.NewVBox(
		widget.NewLabel("Enter Student DNI:"),
		entry,
		searchButton,
	)

	w.SetContent(content)
	w.Show()
}