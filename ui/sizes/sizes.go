/*
 * Copyright (c) 2023 Tom5521- All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package sizes

import (
	"EduTrack/ui/resolution"

	"fyne.io/fyne/v2"
)

// These constants define various sizes used in the application's user interface.
var (
	RecSize                   = fyne.NewSize(620, 520)                  // Size for registration window.
	FormSize                  = fyne.NewSize(500, 240)                  // Size for the form window.
	PickSize                  = fyne.NewSize(600, 400)                  // Size for file picker window.
	ErrSize                   = fyne.NewSize(400, 80)                   // Size for error window.
	ProfileSize               = fyne.NewSize(300, 300)                  // Size for a student profile window.
	SearchSize                = fyne.NewSize(300, 100)                  // Size for student search window.
	ScreenWidth, ScreenHeight = resolution.GetResolution()              // Screen resolution in pixels.
	FyneScreenSize            = fyne.NewSize(ScreenWidth, ScreenHeight) // Size of the Fyne application window.
	RecInfoSize               = fyne.NewSize(200, 100)
	ListSize                  = fyne.NewSize(400, 200)
)
