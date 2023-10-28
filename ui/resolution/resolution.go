/*
 * Copyright Tom5521(c) - All Rights Reserved.
 *
 * This project is licenced under the MIT License.
 */

package resolution

import (
	"github.com/fstanis/screenresolution"
)

func GetResolution() (float32, float32) {
	resolution := screenresolution.GetPrimary()
	return float32(resolution.Width), float32(resolution.Height)
}
