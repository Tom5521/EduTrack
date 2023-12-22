package install

import (
	"embed"
	"fmt"
	"os"

	"fyne.io/fyne/v2/widget"
	"github.com/ncruces/zenity"
)

//go:embed files
var Files embed.FS

type InstallConf struct {
	ProgressBar *widget.ProgressBar
	Windows     struct {
		CreateDestktopShortcut bool
		InstallPath            string
	}
	Linux struct {
		UserInstall bool
		RootInstall bool
	}
}

func errWin(txt string) {
	err := zenity.Error(txt)
	if err != nil {
		fmt.Println(err)
	}
}

func isExists(dir string) bool {
	_, err := os.Stat(dir)
	return os.IsExist(err)
}
func isNotExists(dir string) bool {
	_, err := os.Stat(dir)
	return os.IsNotExist(err)
}

func mkdir(dir string) {
	err := os.Mkdir(dir, os.ModePerm)
	if err != nil {
		errWin(err.Error())
	}
}

func chdir(dir string) {
	err := os.Chdir(dir)
	if err != nil {
		errWin(err.Error())
	}
}