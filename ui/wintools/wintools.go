/*
 * Copyright Tom5521(c) - All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package wintools

import (
	"EduTrack/iconloader"
	"EduTrack/ui/resolution"
	"EduTrack/ui/sizes"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	files "github.com/Tom5521/MyGolangTools/file"
)

// MaximizeWin resizes a given window to match the screen's resolution.
func MaximizeWin(window fyne.Window) {
	resolution.GetResolution()
	window.Resize(sizes.FyneScreenSize)
}

// LoadProfileImg loads and returns a profile image from the specified file path, or a default image if the file doesn't exist.
func LoadProfileImg(file string) *canvas.Image {
	var image *canvas.Image

	// Check if the file exists; if not, use a default user template icon.
	if check, _ := files.CheckFile(file); !check {
		image = canvas.NewImageFromResource(iconloader.UserTemplateICON)
		image.SetMinSize(sizes.ProfileSize)
		return image
	}

	// Load the image from the specified file path.
	res, _ := fyne.LoadResourceFromPath(file)

	image = canvas.NewImageFromResource(res)
	image.ScaleMode = canvas.ImageScaleFastest
	image.SetMinSize(sizes.ProfileSize)

	return image
}

