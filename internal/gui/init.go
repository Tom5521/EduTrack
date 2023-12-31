package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/EduTrack/assets"
	"github.com/Tom5521/EduTrack/internal/gui/config"
	"github.com/Tom5521/EduTrack/locales"
	"github.com/Tom5521/EduTrack/pkg/conf"
	"github.com/leonelquinteros/gotext"
)

var po *gotext.Po

type ui struct {
	App         fyne.App
	StudentTab  *fyne.Container
	StudentList *widget.List
	CoursesList *widget.List
}

type Gui struct {
	ui
}

func Init() *Gui {
	ui := &Gui{}
	ui.App = app.New()
	config.SetTheme(ui.App)
	assets.Load()
	po = locales.GetPo(conf.Config.Lang)
	return ui
}
