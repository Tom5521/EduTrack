package sizes

import (
	"EduTrack/ui/resolution"

	"fyne.io/fyne/v2"
)

var (
	FormSize                  = fyne.NewSize(500, 240)
	PickSize                  = fyne.NewSize(600, 400)
	ErrSize                   = fyne.NewSize(400, 80)
	ProfileSize               = fyne.NewSize(300, 300)
	SearchSize                = fyne.NewSize(300, 100)
	ScreenWidth, ScreenHeight = resolution.GetResolution()
	FyneScreenSize            = fyne.NewSize(ScreenWidth, ScreenHeight)
)
