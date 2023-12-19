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

var (
	po      *gotext.Po
	app     fyne.App
	tmpConf conf.Conf
)

func getCenteredLabel(text string) *widget.Label {
	l := widget.NewLabel(po.Get(text))
	l.Alignment = fyne.TextAlignCenter
	l.TextStyle.Bold = true
	return l
}

func Init(newApp fyne.App, newPo *gotext.Po) {
	tmpConf = conf.Config
	po = newPo
	app = newApp
	mainWin()
}

func passwdUI() *fyne.Container {
	passwdEntry := &widget.Entry{Password: true}
	passwdEntry.SetPlaceHolder(randstr.Hex(16))
	if !tmpConf.Password.Enabled {
		passwdEntry.Disable()
	}
	passwdCheck := widget.NewCheck(po.Get("Enabled"), func(b bool) {
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

	return container.NewVBox(
		getCenteredLabel("Password Options"),
		container.NewBorder(nil, nil, passwdCheck, nil, passwdEntry),
		passwdButton,
	)
}

func databaseUI() *fyne.Container {
	databaseLabel := widget.NewLabel(tmpConf.DatabaseFile)
	return container.NewVBox(
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
	)
}

func generalUI() *fyne.Container {
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

	return container.NewVBox(
		getCenteredLabel("General Options"),
		widget.NewForm(
			widget.NewFormItem(po.Get("Language:"), langSelect),
			widget.NewFormItem(po.Get("Theme:"), themeSelect),
		),
	)
}

func mainWin() {
	var applyedChanges bool
	w := app.NewWindow(po.Get("Configuration's"))
	const size1, size2 float32 = 400, 100
	w.Resize(fyne.NewSize(size1, size2))

	m := widgets.NewForm()
	m.CustomItems = container.NewVBox(
		generalUI(),
		databaseUI(),
		passwdUI(),
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
