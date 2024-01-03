package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/EduTrack/assets"
	"github.com/Tom5521/EduTrack/internal/gui/credits"
)

func (ui *ui) AboutWin() {
	window := ui.App.NewWindow(po.Get("About"))
	label1 := widget.NewLabel(po.Get("Created by:"))
	gitLabel := widget.NewRichTextFromMarkdown("[Tom5521](https://github.com/Tom5521)")
	licenceLabel := widget.NewLabel(po.Get("Under %s licence", "MIT"))
	creditsButton := widget.NewButton(po.Get("CREDITS"), func() {
		const size1, size2 float32 = 800, 400
		credits.CreditsWindow(ui.App, fyne.NewSize(size1, size2)).Show()
	})
	versionLabel := widget.NewLabel(po.Get("Version: %s", fyne.CurrentApp().Metadata().Version))
	authorCont := container.NewHBox(label1, gitLabel)
	logo := canvas.NewImageFromResource(assets.App)
	const w, h float32 = 300, 300
	logo.SetMinSize(fyne.NewSize(w, h))
	vbox1 := container.NewVBox(
		authorCont,
		versionLabel,
		licenceLabel,
		creditsButton,
		logo,
	)
	window.SetContent(vbox1)
	window.Show()
}
