package resolution

import (
	"github.com/fstanis/screenresolution"
)

func GetResolution() (float32, float32) {
	resolution := screenresolution.GetPrimary()
	return float32(resolution.Width), float32(resolution.Height)
}
