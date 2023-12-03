package tools

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"github.com/Tom5521/EduTrack/assets"
	"github.com/Tom5521/EduTrack/internal/pkg/sizes"

	files "github.com/Tom5521/MyGolangTools/file"
)

// LoadProfileImg loads and returns a profile image from the specified file path,
// or a default image if the file doesn't exist.
func LoadProfileImg(file string) *canvas.Image {
	var image *canvas.Image

	// Check if the file exists; if not, use a default user template icon.
	if check, _ := files.CheckFile(file); !check {
		image = canvas.NewImageFromResource(assets.UserTemplate)
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
