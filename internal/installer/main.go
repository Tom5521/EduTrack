package installer

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/EduTrack/internal/installer/locales"
	"github.com/Tom5521/EduTrack/pkg/widgets"
	"github.com/ncruces/zenity"
)

var (
	SelectedPath       string
	CreateDesktopEntry bool
	InstallDir         string
)

func (ui *ui) MainWin() {
	ui.Window = ui.App.NewWindow("EduTrack Installer")
	ui.Window.Resize(fyne.NewSize(500, 400))
	ui.Content1()
	ui.Window.ShowAndRun()
}

func NewRichText(txt string) *widget.RichText {
	return widget.NewRichTextFromMarkdown(txt)
}

func (ui *ui) Content1() {
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
				ui.Window.SetTitle(po.Get("EduTrack Installer"))
				content.DisableNext = false
				content.Update()
			}
		},
	)
	content = widgets.NewInstallerLayout(
		NewRichText("# Welcome to the EduTrack installer!"),
		NewRichText("#### First of all, select the language to use during installation."),
		langSelect,
	)
	content.DisableNext = true
	content.TopItemsLayout = layout.NewVBoxLayout()
	content.HideBack = true
	content.OnNext = func() {
		ui.Content2()
	}
	ui.Window.SetContent(content)
}

func (ui *ui) Content2() {
	licenceCont := widgets.NewInstallerLayout(
		&widget.Entry{MultiLine: true, Text: MITLicence},
	)
	licenceCont.NextText = po.Get("Accept")
	licenceCont.BackText = po.Get("Deny")
	licenceCont.OnBack = func() {
		ui.FinalContent()
	}
	licenceCont.OnNext = func() {
		ui.Content3()
	}
	content := widgets.NewInstallerLayout(
		NewRichText(po.Get("# Licence")),
		licenceCont,
	)
	content.OnBack = func() {
		ui.Content1()
	}
	content.HideNext = true
	ui.Window.SetContent(content)
}

func (ui *ui) Content3() {
	entry := widget.NewEntry()
	entry.SetText("")
	selectPathBt := widget.NewButton(po.Get("Select path"), func() {
		path, _ := zenity.SelectFile(zenity.Directory())
		SelectedPath = path
	})
	content := widgets.NewInstallerLayout(
		NewRichText(po.Get("## Select the installation path:")),
		container.NewBorder(nil, nil, nil, selectPathBt, entry),
		container.NewHBox(
			widget.NewCheck(
				po.Get("Create Desktop entry"),
				func(b bool) {
					CreateDesktopEntry = b
				}),
		),
	)
	content.TopItemsLayout = layout.NewVBoxLayout()
	ui.Window.SetContent(content)
}

func (ui *ui) FinalContent() {}

const MITLicence string = `
MIT License

Copyright (c) 2023 Tom5521

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
`
