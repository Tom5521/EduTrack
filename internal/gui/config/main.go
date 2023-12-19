package config

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/EduTrack/locales"
	"github.com/Tom5521/EduTrack/pkg/conf"
	"github.com/Tom5521/EduTrack/pkg/passwd"
	"github.com/Tom5521/EduTrack/pkg/themes"
	"github.com/Tom5521/EduTrack/pkg/widgets"
	"github.com/Tom5521/EduTrack/pkg/wins"
	"github.com/leonelquinteros/gotext"
	"github.com/ncruces/zenity"
	"github.com/thanhpk/randstr"

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
	passwdEntry := &widget.Entry{Password: true}
	passwdEntry.SetPlaceHolder(randstr.Hex(16))
	passwdCheck := widget.NewCheck(po.Get("Enabled:"), func(b bool) {
		if b {
			passwdEntry.Enable()
		} else {
			passwdEntry.Disable()
		}
		tmpConf.Password.Enabled = b
	})
	passwdCheck.SetChecked(tmpConf.Password.Enabled)

	passwdButton := widget.NewButton(po.Get("Set password"), func() {
		if passwdEntry.Text == "" {
			wins.ErrWin(app, po.Get("Password field is empty!"))
			return
		}
		p := passwd.Password(passwdEntry.Text)
		newHash, err := p.ToHash()
		if err != nil {
			wins.ErrWin(app, err.Error())
			return
		}
		tmpConf.Password.Hash = string(newHash)
	})
	databaseLabel := widget.NewLabel(tmpConf.DatabaseFile)
	m := widgets.NewForm()
	m.CustomItems = container.NewVBox(
		getCenteredLabel("General Options"),
		widget.NewForm(
			widget.NewFormItem(po.Get("Language:"), langSelect),
			widget.NewFormItem(po.Get("Theme:"), themeSelect),
		),
		getCenteredLabel("Database Options"),
		widget.NewForm(
			widget.NewFormItem(po.Get("Current database route:"), databaseLabel),
		),
		widget.NewButton(po.Get("Set database route"), func() {
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
		}),
		getCenteredLabel(po.Get("Password Options")),
		container.NewAdaptiveGrid(2, passwdCheck, passwdEntry),
		passwdButton,
	)

	m.SubmitText = po.Get("Apply changes")
	m.CancelText = po.Get("Close")

	m.OnSubmit = func() {
		tmpConf.Update()
		conf.Config = conf.GetConfData()
		SetTheme(app)
		applyedChanges = true
	}
	m.OnCancel = func() {
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

	w.SetContent(m)
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
