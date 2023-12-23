//go:build windows
// +build windows

package install

import (
	_ "embed"
	"os"
	"os/user"

	"fyne.io/fyne/v2/widget"
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

//go:embed files/EduTrack.exe
var ExeFile []byte

//go:embed files/opengl32.dll
var DllFile []byte

func (i *InstallConf) CopyExeFile() {
	defer i.ProgressBar.SetValue(0.4)
	i.LogContainer.Add(widget.NewRichTextFromMarkdown(i.Po.Get("`Extracting exe file...`")))
	c := i.Windows
	if isNotExists(c.InstallPath) {
		mkdir(c.InstallPath)
	}
	file, err := os.Create(c.InstallPath + "/EduTrack.exe")
	if err != nil {
		errWin(err.Error())
	}
	defer file.Close()
	_, err = file.Write(ExeFile)
	if err != nil {
		errWin(err.Error())
	}
}

func (i *InstallConf) CopyDllFile() {
	i.LogContainer.Add(widget.NewRichTextFromMarkdown(i.Po.Get("`Extracting opengl32.dll...`")))
	c := i.Windows
	defer i.ProgressBar.SetValue(0.6)
	file, err := os.Create(c.InstallPath + "/opengl32.dll")
	if err != nil {
		errWin(err.Error())
	}
	defer file.Close()
	_, err = file.Write(DllFile)
	if err != nil {
		errWin(err.Error())
	}
}

func (i *InstallConf) CreateShortcut() {
	defer i.ProgressBar.SetValue(0.9)
	if !i.Windows.CreateDestktopShortcut {
		return
	}
	i.LogContainer.Add(widget.NewRichTextFromMarkdown(i.Po.Get("`Creating desktop shortcut...`")))
	usr, err := user.Current()
	if err != nil {
		errWin(err.Error())
	}
	c := i.Windows
	if c.CreateDestktopShortcut {
		target := c.InstallPath + "/EduTrack.exe"
		CreateShortcut(target, usr.HomeDir+"/Desktop/")
	}
}

func (i *InstallConf) Install() {
	defer i.ProgressBar.SetValue(1)
	i.CopyExeFile()
	i.CopyDllFile()
	i.CreateShortcut()
	i.LogContainer.Add(widget.NewRichTextFromMarkdown(i.Po.Get("`Installation finished!`")))
}

func CreateShortcut(src, dst string) error {
	ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED|ole.COINIT_SPEED_OVER_MEMORY)
	oleShellObject, err := oleutil.CreateObject("WScript.Shell")
	if err != nil {
		return err
	}
	defer oleShellObject.Release()
	wshell, err := oleShellObject.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return err
	}
	defer wshell.Release()
	cs, err := oleutil.CallMethod(wshell, "CreateShortcut", dst)
	if err != nil {
		return err
	}
	idispatch := cs.ToIDispatch()
	oleutil.PutProperty(idispatch, "TargetPath", src)
	oleutil.CallMethod(idispatch, "Save")
	return nil
}
