package installer

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/Tom5521/EduTrack/internal/installer/locales"
	"github.com/leonelquinteros/gotext"
)

type ui struct {
	App    fyne.App
	Window fyne.Window
}

// export ui.
type Gui struct {
	ui
}

var po *gotext.Po

func InitGUI() *Gui {
	app := app.New()
	g := &Gui{}
	g.App = app
	g.Window = app.NewWindow("EduTrack installer")
	po = locales.GetPo("en")
	return g
}
