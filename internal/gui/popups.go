package gui

import (
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/EduTrack/assets"
	"github.com/Tom5521/EduTrack/internal/pkg/credits"
)

func (ui *ui) AboutWin() {
	window := ui.App.NewWindow("About")
	label1 := widget.NewLabel("Created by:")
	link, _ := url.Parse("https://github.com/Tom5521")
	gitLabel := widget.NewHyperlink("Tom5521", link)
	licenceLabel := widget.NewLabel("Under the MIT license")
	creditsButton := widget.NewButton("CREDITS", func() {
		const size1, size2 float32 = 800, 400
		credits.CreditsWindow(ui.App, fyne.NewSize(size1, size2)).Show()
	})
	authorCont := container.NewHBox(label1, gitLabel)
	logo := canvas.NewImageFromResource(assets.App)
	const w, h float32 = 300, 300
	logo.SetMinSize(fyne.NewSize(w, h))
	vbox1 := container.NewVBox(
		authorCont,
		licenceLabel,
		creditsButton,
		logo,
	)
	window.SetContent(vbox1)
	window.Show()
}
