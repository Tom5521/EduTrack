package installer

import (
	"runtime"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/EduTrack/internal/installer/install"
	"github.com/Tom5521/EduTrack/internal/installer/locales"
	"github.com/Tom5521/EduTrack/pkg/widgets"
	"github.com/ncruces/zenity"
)

func (ui *ui) MainWin() {
	ui.Window = ui.App.NewWindow("EduTrack Install Wizard")
	ui.Window.Resize(fyne.NewSize(500, 400))
	ui.Lang1()
	ui.Window.ShowAndRun()
}

var NewRichText = widget.NewRichTextFromMarkdown

func NewContent(items ...fyne.CanvasObject) *widgets.InstallerLayout {
	l := widgets.NewInstallerLayout(items...)
	l.CancelText = po.Get("Cancel")
	l.NextText = po.Get("Next")
	l.BackText = po.Get("Back")
	return l
}

func (ui *ui) Lang1() {
	var content *widgets.InstallerLayout
	langSelect := widget.NewSelect([]string{"Spanish", "English", "Portuguese"},
		func(s string) {
			if s == "Spanish" {
				po.Parse(locales.Read("es.po"))
			}
			if s == "English" {
				po.Parse(locales.Read("en.pot"))
			}
			if s == "Portuguese" {
				po.Parse(locales.Read("pt.po"))
			}
			if s != "" {
				ui.Window.SetTitle(po.Get("EduTrack Install Wizard"))
				content.DisableNext = false
				content.Update()
			}
		},
	)
	content = NewContent(
		NewRichText("# Welcome to the EduTrack installer!"),
		NewRichText("#### First of all, select the language to use during installation."),
		langSelect,
	)
	content.DisableNext = true
	content.TopItemsLayout = layout.NewVBoxLayout()
	content.HideBack = true
	content.OnNext = func() { ui.Licence2() }
	content.OnCancel = func() { ui.FinalMsgCont() }
	ui.Window.SetContent(content)
}

func (ui *ui) Licence2() {
	licenceCont := widgets.NewInstallerLayout(
		&widget.Entry{MultiLine: true, Text: MITLicence},
	)
	licenceCont.NextText = po.Get("Accept")
	licenceCont.BackText = po.Get("Deny")
	licenceCont.OnBack = func() { ui.SumContentWindows4() }
	licenceCont.OnNext = func() {
		if runtime.GOOS == "windows" {
			ui.Options3Windows()
		} else {
			ui.Options3Linux()
		}
	}
	licenceLabel := NewRichText(po.Get("# Licence"))
	content := NewContent(
		licenceLabel,
		licenceCont,
	)
	content.OnBack = func() { ui.Lang1() }
	content.OnCancel = func() { ui.FinalMsgCont() }
	content.TopItemsLayout = layout.NewBorderLayout(licenceLabel, nil, nil, nil)
	content.HideNext = true
	ui.Window.SetContent(content)
}

func (ui *ui) Options3Linux() {
	var userInstallCheck, rootInstallCheck *widget.Check
	userInstallCheck = widget.NewCheck("", func(b bool) {
		UserInstall = b
		rootInstallCheck.SetChecked(!b)
	})
	rootInstallCheck = widget.NewCheck("", func(b bool) {
		RootInstall = b
		userInstallCheck.SetChecked(!b)
	})
	userInstallCheck.SetChecked(UserInstall)
	rootInstallCheck.SetChecked(RootInstall)
	content := NewContent(
		NewRichText(po.Get("# Install Options")),
		NewRichText(po.Get("## Install Mode")),
		widget.NewForm(
			widget.NewFormItem(po.Get("User Install"), userInstallCheck),
			widget.NewFormItem(po.Get("Root Install"), rootInstallCheck),
		),
	)
	content.OnBack = func() { ui.Licence2() }
	content.OnNext = func() { ui.SumContentLinux4() }
	content.OnCancel = func() { ui.FinalMsgCont() }
	content.TopItemsLayout = layout.NewVBoxLayout()
	ui.Window.SetContent(content)
}

func (ui *ui) Options3Windows() {
	entry := widget.NewEntry()
	entry.SetText(SelectedPath)
	selectPathBt := widget.NewButton(po.Get("Select path"), func() {
		path, _ := zenity.SelectFile(zenity.Directory())
		SelectedPath = path
		entry.SetText(path)
	})
	content := NewContent(
		NewRichText(po.Get("# Install Options")),
		NewRichText(po.Get("## Select the installation path:")),
		container.NewBorder(nil, nil, nil, selectPathBt, entry),
		container.NewHBox(
			widget.NewCheck(
				po.Get("Create Desktop shortcut"),
				func(b bool) {
					CreateDesktopShortcut = b
				}),
		),
	)
	content.OnBack = func() { ui.Licence2() }
	content.OnNext = func() { ui.SumContentWindows4() }
	content.OnCancel = func() { ui.FinalMsgCont() }
	content.TopItemsLayout = layout.NewVBoxLayout()
	ui.Window.SetContent(content)
}

func (ui *ui) SumContentLinux4() {
	content := NewContent(
		NewRichText(po.Get("# Final summary")),
		widget.NewForm(
			widget.NewFormItem(po.Get("Install mode"), widget.NewLabel(func() string {
				if UserInstall {
					return po.Get("User")
				}
				if RootInstall {
					return po.Get("Root")
				}
				return po.Get("User (default)")
			}())),
		),
	)
	content.TopItemsLayout = layout.NewVBoxLayout()
	content.OnNext = func() { ui.Installing5() }
	content.OnBack = func() { ui.Options3Linux() }
	content.NextText = "Install"
	ui.Window.SetContent(content)
}

func (ui *ui) SumContentWindows4() {
	content := NewContent(
		NewRichText(po.Get("# Final summary")),
		widget.NewForm(
			widget.NewFormItem(po.Get("Install Path:"), widget.NewLabel(SelectedPath)),
			widget.NewFormItem(
				po.Get("Create desktop shortcut:"),
				widget.NewLabel(strconv.FormatBool(CreateDesktopShortcut)),
			),
		),
	)
	content.NextText = "Install"
	content.TopItemsLayout = layout.NewVBoxLayout()
	content.OnNext = func() { ui.Installing5() }
	content.OnBack = func() { ui.Options3Windows() }
	content.OnCancel = func() { ui.FinalMsgCont() }
	ui.Window.SetContent(content)
}

func (ui *ui) Installing5() {
	progressBar := widget.NewProgressBar()
	logContainer := container.NewVBox()
	titlesContent := container.NewVBox(
		NewRichText(po.Get("# Installing EduTrack...")),
		progressBar,
		NewRichText(po.Get("### This probably won't take more than a minute")),
		NewRichText(po.Get("**Log**")),
	)
	scrollContent := container.NewVScroll(
		logContainer,
	)
	content := NewContent(
		titlesContent,
		scrollContent,
	)

	go func() {
		i := &install.InstallConf{}
		i.Po = po
		i.LogContainer = logContainer
		i.ProgressBar = progressBar
		i.Windows.InstallPath = SelectedPath
		i.Windows.CreateDestktopShortcut = CreateDesktopShortcut
		i.Linux.RootInstall = RootInstall
		i.Linux.UserInstall = UserInstall
		i.Install()
		content.DisableNext = false
		content.Update()
	}()

	content.OnNext = func() {
		ui.FinalMsgCont()
	}
	content.DisableCancel = true
	content.DisableBack = true
	content.DisableNext = true
	content.TopItemsLayout = layout.NewBorderLayout(titlesContent, nil, nil, nil)
	ui.Window.SetContent(content)
}

func (ui *ui) FinalMsgCont() {
	const link = "https://github.com/Tom5521/EduTrack"
	content := NewContent(
		NewRichText(po.Get("## Thank you for installing EduTrack!")),
		NewRichText(po.Get("Please check the github [page](%s)", link)),
	)
	content.NextText = po.Get("Finish")
	content.HideBack = true
	content.HideCancel = true
	content.OnNext = func() {
		ui.App.Quit()
	}
	content.TopItemsLayout = layout.NewVBoxLayout()
	ui.Window.SetContent(content)
}
