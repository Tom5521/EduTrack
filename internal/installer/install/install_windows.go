//go:build windows
// +build windows

package install

import (
	"os"

	"github.com/jxeng/shortcut"
)

func (i *InstallConf) CopyFile() {
	defer i.ProgressBar.SetValue(0.4)
	c := i.Windows
	if isNotExists(c.InstallPath) {
		mkdir(c.InstallPath)
	}
	bytedata, err := files.ReadFile("files/EduTrack.exe")
	if err != nil {
		errWin(err.Error())
	}
	file, err := os.Create(c.InstallPath + "/EduTrack.exe")
	if err != nil {
		errWin(err.Error())
	}
	defer file.Close()
	_, err = file.Write(bytedata)
}

func (i *InstallConf) CreateShortcut() {
	defer i.ProgressBar.SetValue(0.8)
	c := i.Windows
	if c.CreateDestktopShortcut {
		target := c.InstallPath + "/EduTrack.exe"
		shortcut.CreateDesktopShortcut("EduTrack", target, target)
	}
}

func (i *InstallConf) Install() {
	defer i.ProgressBar.SetValue(1)
	CopyFile()
	CreateShortcut()
}
