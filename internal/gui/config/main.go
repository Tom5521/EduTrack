package config

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/EduTrack/locales"
	"github.com/Tom5521/EduTrack/pkg/conf"
	"github.com/Tom5521/EduTrack/pkg/themes"
	"github.com/leonelquinteros/gotext"
	"github.com/ncruces/zenity"

	xtheme "fyne.io/x/fyne/theme"
)

func MainWin(app fyne.App, po *gotext.Po) {
	var applyedChanges bool
	tmpConf := conf.Config
	w := app.NewWindow(po.Get("Configurations"))
	const size1, size2 float32 = 600, 100
	w.Resize(fyne.NewSize(size1, size2))
	langSelect := widget.NewSelect(locales.Languages, func(s string) {
		if s == "Español" {
			tmpConf.Lang = "es"
		}
		if s == "English" {
			tmpConf.Lang = "en"
		}
		if s == "Português" {
			tmpConf.Lang = "pt"
		}
	})
	langSelect.PlaceHolder = po.Get("Select a language")
	langSelect.Selected = func() string {
		current := po.Get("Select a language")
		if tmpConf.Lang == "es" {
			current = "Español"
		}
		if tmpConf.Lang == "en" {
			current = "English"
		}
		if tmpConf.Lang == "pt" {
			current = "Português"
		}
		return current
	}()
	themesList := []string{"Adwaita", "Default", "SimpleRed", "DarkRed", "DarkBlue"}
	themeSelect := widget.NewSelect(themesList, func(s string) {
		if s != "" {
			tmpConf.Theme = s
		}
	})
	themeSelect.PlaceHolder = po.Get("Select a theme")
	themeSelect.Selected = tmpConf.Theme

	getCenteredLabel := func(text string) *widget.Label {
		l := widget.NewLabel(po.Get(text))
		l.Alignment = fyne.TextAlignCenter
		l.TextStyle.Bold = true
		return l
	}
	databaseLabel := widget.NewLabel(tmpConf.DatabaseFile)
	mainForm := widget.NewForm(
		widget.NewFormItem("", getCenteredLabel("General Options")),
		widget.NewFormItem(po.Get("Language:"), langSelect),
		widget.NewFormItem(po.Get("Theme:"), themeSelect),
		widget.NewFormItem("", getCenteredLabel("Database Options")),
		widget.NewFormItem(po.Get("Current database route:"), databaseLabel),
		widget.NewFormItem("", widget.NewButton(po.Get("Set database route"), func() {
			db, err := zenity.SelectFile(
				zenity.Filename("database.db"),
				zenity.FileFilters{
					{po.Get("Database file"), []string{"*.db"}, true},
				},
			)
			if err != nil {
				return
			}
			tmpConf.DatabaseFile = db
			databaseLabel.SetText(tmpConf.DatabaseFile)
		})),
	)

	mainForm.SubmitText = po.Get("Apply changes")
	mainForm.CancelText = po.Get("Close")

	mainForm.OnSubmit = func() {
		tmpConf.Update()
		conf.Config = conf.GetConfData()
		SetTheme(app)
		applyedChanges = true
	}
	mainForm.OnCancel = func() {
		onOk := func() {
			po.Parse(locales.GetParser(tmpConf.Lang))
			w.Close()
		}
		if applyedChanges {
			dialog.ShowCustomWithoutButtons(
				po.Get("The changes will be noticeable after restarting the program"),
				container.NewCenter(widget.NewButton(po.Get("Ok"), onOk)),
				w,
			)
		} else {
			w.Close()
		}
	}

	w.SetContent(mainForm)
	w.Show()
}

func SetTheme(app fyne.App) {
	var th fyne.Theme
	switch conf.Config.Theme {
	case "Adwaita":
		th = xtheme.AdwaitaTheme()
	case "DarkRed":
		th = themes.DarkRed{}
	case "DarkBlue":
		th = themes.DarkBlue{}
	case "SimpleRed":
		th = themes.SimpleRed{}
	case "Default":
		th = theme.DefaultTheme()
	default:
		th = theme.DefaultTheme()
	}
	app.Settings().SetTheme(th)
}
