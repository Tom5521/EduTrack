//go:build dev
// +build dev

package gui

import (
	"fyne.io/fyne/v2"
	"github.com/Tom5521/EduTrack/tests/data/fillers"
)

func GetDevMenu(m *fyne.MainMenu) {
	menu := fyne.NewMenu("Dev",
		fyne.NewMenuItem("Fill students", func() {
			fillers.Student()
		}),
		fyne.NewMenuItem("Fill courses", func() {
			fillers.Course()
		}),
		fyne.NewMenuItem("Fill records", func() {
			fillers.Record()
		}),
		fyne.NewMenuItem("Fill Student Courses", func() {
			fillers.StudentCourse()
		}),
		fyne.NewMenuItem("Custom function", nil),
	)
	m.Items = append(m.Items, menu)
	m.Refresh()
}
