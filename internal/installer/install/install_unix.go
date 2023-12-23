//go:build unix
// +build unix

package install

import (
	_ "embed"
	"io"
	"os"

	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/CmdRunTools/command"
	"github.com/ncruces/zenity"
	"github.com/yi-ge/unxz"
)

//go:embed files/EduTrack-linux64.tar.xz
var TarFile []byte

func (i *InstallConf) Untar() {
	defer i.ProgressBar.SetValue(0.3)
	i.LogContainer.Add(widget.NewRichTextFromMarkdown(i.Po.Get("`Unzipping package to /tmp...`")))
	chdir("/tmp")
	err := ExtractTarXz(TarFile)
	if err != nil {
		errWin(err.Error())
	}
}

func RootMake() {
	_, passwd, err := zenity.Password()
	if err != nil {
		errWin(err.Error())
		panic(err)
	}
	cmd := command.InitCmd("sudo -S make install")
	cmd.UseBashShell(true)
	cmd.CustomStd(false, true, true)
	ncmd := cmd.GetExec()
	stdin, err := ncmd.StdinPipe()
	if err != nil {
		errWin(err.Error())
	}
	go func() {
		defer stdin.Close()
		_, err = io.WriteString(stdin, passwd)
		if err != nil {
			errWin(err.Error())
		}
	}()
	err = ncmd.Run()
	if err != nil {
		errWin(err.Error())
	}
}

func UserMake() {
	cmd := command.InitCmd("make user-install")
	cmd.UseBashShell(true)
	cmd.CustomStd(true, true, true)
	err := cmd.Run()
	if err != nil {
		errWin(err.Error())
	}
}

func (i *InstallConf) Make() {
	defer i.ProgressBar.SetValue(0.8)
	i.LogContainer.Add(widget.NewRichTextFromMarkdown(i.Po.Get("`Running make...`")))
	chdir("EduTrack/")
	if i.Linux.RootInstall {
		i.LogContainer.Add(widget.NewRichTextFromMarkdown(i.Po.Get("`[root] make install`")))
		RootMake()
	} else {
		i.LogContainer.Add(widget.NewRichTextFromMarkdown(i.Po.Get("`[user] make user-install`")))
		UserMake()
	}
}

func (i *InstallConf) Install() {
	defer i.ProgressBar.SetValue(1)
	i.Untar()
	i.Make()
	i.LogContainer.Add(widget.NewRichTextFromMarkdown(i.Po.Get("`Installation finished!`")))
}

func ExtractTarXz(embeddedData []byte) error {
	filename := "tmp_file.tar.xz"
	folder := "./EduTrack"
	err := os.WriteFile(filename, embeddedData, os.ModePerm)
	if err != nil {
		errWin(err.Error())
	}
	if isNotExists(folder) {
		mkdir(folder)
	}
	u := unxz.New(filename, folder)
	return u.Extract()
}
