/*
 * Copyright (c) 2023 Tom5521- All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package wins

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/EduTrack/pkg/conf"
	"github.com/Tom5521/EduTrack/pkg/resolution"
	"github.com/ncruces/zenity"

	"github.com/Tom5521/EduTrack/assets"
	"github.com/Tom5521/EduTrack/internal/pkg/sizes"
)

// ImagePicker opens a file picker window to select an image file.
func ImagePicker(imageFilePath *string) {
	filter := []string{"*.png", "*.gif", "*.ico", "*.jpg", "*.webp"}
	const defaultPath string = ""
	ret, err := zenity.SelectFile(
		zenity.Filename(defaultPath),
		zenity.FileFilters{
			{
				Name:     "Image files",
				Patterns: filter,
				CaseFold: true,
			},
		})
	if err != nil {
		fmt.Println(err)
	} else {
		*imageFilePath = ret
	}
}

// FilePicker opens a file picker window to select a configuration file.
func FilePicker() string {
	filter := []string{"*.yml", "*.yaml"}
	const defaultPath string = ""
	ret, err := zenity.SelectFile(
		zenity.Filename(defaultPath),
		zenity.FileFilters{
			{
				Name:     "Yaml files",
				Patterns: filter,
				CaseFold: true,
			},
		})
	if err != nil {
		fmt.Println(err)
	}
	return ret
}

// ErrWin opens an error window with a message.
func ErrWin(app fyne.App, err string, clWindow ...fyne.Window) {
	var buttonText string
	if conf.Config.Lang == "es" {
		buttonText = "Aceptar"
	}
	if conf.Config.Lang == "en" || conf.Config.Lang == "" {
		buttonText = "Accept"
	}
	w := app.NewWindow("Error")
	w.RequestFocus()
	w.Resize(sizes.ErrSize)
	w.SetIcon(assets.Error)
	errlabel := widget.NewLabel(err)
	errlabel.TextStyle.Bold = true
	errlabel.Alignment = fyne.TextAlignCenter
	acceptButton := widget.NewButton(buttonText, func() {
		w.Close()
		if len(clWindow) > 0 {
			clWindow[0].Close()
		}
	})

	content := container.NewVBox(
		errlabel,
		acceptButton,
	)
	w.SetContent(content)
	w.SetMainMenu(w.MainMenu())
	w.Show()
}

// MaximizeWin resizes a given window to match the screen's resolution.
func MaximizeWin(window fyne.Window) {
	resolution.GetResolution()
	window.Resize(sizes.FyneScreenSize)
	window.CenterOnScreen()
}
