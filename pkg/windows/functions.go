package win

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func LoadProfileImg(file string) (*canvas.Image, error) {
	res, err := fyne.LoadResourceFromPath(file)

	if err != nil {
		return nil, err
	}

	image := canvas.NewImageFromResource(res)

	image.SetMinSize(ProfileSize)

	return image, nil

}
