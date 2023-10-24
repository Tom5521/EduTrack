package windowtools

import (
	"EduTrack/pkg/resolution"

	"fyne.io/fyne/v2"
)

var MaxSize fyne.Size = fyne.NewSize(resolution.GetResolution())

func MaximizeWin(window fyne.Window) {
	resolution.GetResolution()
	//window.SetFixedSize(true)
	window.Resize(MaxSize)
}
