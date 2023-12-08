/*
 * Copyright (c) 2023 Tom5521- All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package sizes

import (
	"github.com/Tom5521/EduTrack/pkg/resolution"

	"fyne.io/fyne/v2"
)

// These constants define various sizes used in the application's user interface.

const (
	startEndW, startEndH float32 = formW, 0
	recW, recH           float32 = 620, 520
	formW, formH         float32 = 500, 240
	pickW, pickH         float32 = 600, 400
	errW, errH           float32 = 400, 80
	profileW, profileH   float32 = 300, 300
	searchW, searchH     float32 = 300, 100
	recInfoW, recInfoH   float32 = 200, 100
	listW, listH         float32 = 400, 200
)

var screenWidth, screenHeight = resolution.GetResolution() // Screen resolution in pixels.
var FyneScreenSize = fyne.NewSize(screenWidth, screenHeight)
var (
	RecSize      = fyne.NewSize(recW, recH)         // Size for record window.
	FormSize     = fyne.NewSize(formW, formH)       // Size for the form window.
	PickSize     = fyne.NewSize(pickW, pickH)       // Size for file picker window.
	ErrSize      = fyne.NewSize(errW, errH)         // Size for error window.
	ProfileSize  = fyne.NewSize(profileW, profileH) // Size for a student profile window.
	SearchSize   = fyne.NewSize(searchW, searchH)   // Size for student search window.
	StartEndSize = fyne.NewSize(startEndW, startEndH)

	// Size of the Fyne application window.
	RecInfoSize = fyne.NewSize(recInfoW, recInfoH)
	ListSize    = fyne.NewSize(listW, listH)
)
