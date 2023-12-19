package gui

import (
	"fmt"
	"os/exec"
	"runtime"

	"fyne.io/fyne/v2"
	"github.com/Tom5521/EduTrack/internal/gui/config"
	"github.com/Tom5521/EduTrack/pkg/data"
	"github.com/ncruces/zenity"
)

func (ui *ui) MainMenu() *fyne.MainMenu {
	// Create the main menu
	menu := fyne.NewMainMenu(
		fyne.NewMenu(po.Get("File")),
		fyne.NewMenu(po.Get("Edit"),
			fyne.NewMenuItem(po.Get("Reload data"), func() {
				data.LoadEverything()
			}),
			fyne.NewMenuItem(po.Get("Configuration"), func() {
				config.Init(ui.App, po)
			}),
		),
		fyne.NewMenu(po.Get("Search"),
			fyne.NewMenuItem(po.Get("Search"), func() {
				ui.SearchMainWin()
			}),
			fyne.NewMenuItem(po.Get("Search Students"), func() {
				ui.SearchStudentsMainWin()
			}),
			fyne.NewMenuItem(po.Get("Search Records"), func() {
				ui.SearchRecordsMainWin()
			}),
			fyne.NewMenuItem(po.Get("Search Courses"), func() {
				ui.SearchCoursesMainWin()
			}),
			fyne.NewMenuItem(po.Get("Search student courses"), func() {
				ui.SearchStudentCoursesMainWin()
			}),
		),
		fyne.NewMenu(po.Get("Help"),
			fyne.NewMenuItem(po.Get("About"), func() {
				ui.AboutWin()
			}),
			fyne.NewMenuItem(po.Get("Documentation"), func() {
				var cmd string
				var args []string
				switch runtime.GOOS {
				case "windows":
					cmd = "cmd"
					args = []string{"/c", "start"}
				case "darwin":
					cmd = "open"
				default:
					cmd = "xdg-open"
				}
				args = append(args, "https://github.com/Tom5521/EduTrack/")
				err := exec.Command(cmd, args...).Start()
				if err != nil {
					if zenity.Error(err.Error()) != nil {
						fmt.Println(err)
					}
				}
			}),
		),
	)

	return menu
}
