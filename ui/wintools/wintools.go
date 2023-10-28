/*
 * Copyright Tom5521(c) - All Rights Reserved.
 *
 * This project is licenced under the MIT License.
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

func MaximizeWin(window fyne.Window) {
	resolution.GetResolution()
	window.Resize(sizes.FyneScreenSize)
}

func LoadProfileImg(file string) *canvas.Image {
	var image *canvas.Image
	if check, _ := files.CheckFile(file); !check {
		image = canvas.NewImageFromResource(iconloader.UserTemplateICON)
		image.SetMinSize(sizes.ProfileSize)
		return image
	}
	res, _ := fyne.LoadResourceFromPath(file)

	image = canvas.NewImageFromResource(res)
	image.ScaleMode = canvas.ImageScaleFastest
	image.SetMinSize(sizes.ProfileSize)

	return image

}
