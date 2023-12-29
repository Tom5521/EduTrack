//go:build linux
// +build linux

//go:generate cp ../../../builds/EduTrack-linux64.tar.xz files/

package install

import (
	_ "embed"
	"fmt"
	"os"

	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/CmdRunTools/command"
	"github.com/ncruces/zenity"
	"github.com/yi-ge/unxz"
)

//go:embed files/EduTrack-linux64.tar.xz
var TarFile []byte

func (i *InstallConf) newRichLine(txt string, args ...interface{}) *widget.RichText {
	return widget.NewRichTextFromMarkdown(i.Po.Get(txt, args...))
}

func (i *InstallConf) Untar() {
	defer i.ProgressBar.SetValue(0.3)
	i.LogContainer.Add(i.newRichLine("`Unzipping package to /tmp...`"))
	chdir("/tmp")
	err := ExtractTarXz(TarFile)
	if err != nil {
		i.ErrorOut = err
		txt := fmt.Sprintf("```\n%s\n```", err.Error())
		i.LogContainer.Add(widget.NewRichTextFromMarkdown(txt))
	}
}

func askPwd() string {
	_, passwd, err := zenity.Password()
	if err != nil {
		errWin(err.Error())
		askPwd()
	}
	return passwd
}

func RootMake() (string, error) {
	passwd := askPwd()

	cmd := command.NewSudoCmd("make install", passwd)
	out, err := cmd.Out()
	if err != nil {
		return fmt.Sprintf("Command:%s\n%s\n%s", cmd.Input, out, err.Error()), err
	}
	return out, err
}

func UserMake() (string, error) {
	cmd := command.NewCmd("make user-install")
	cmd.Std(true, true, true)
	out, err := cmd.Out()
	if err != nil {
		return fmt.Sprintf("Command:%s\n%s\n%s", cmd.Input, out, err.Error()), err
	}
	return out, err
}

func (i *InstallConf) Make() {
	defer i.ProgressBar.SetValue(0.8)
	i.LogContainer.Add(i.newRichLine("`Running make...`"))
	chdir("EduTrack/")
	if i.Linux.RootInstall {
		i.LogContainer.Add(i.newRichLine("`[root] make install`"))
		out, err := RootMake()
		if err != nil {
			i.ErrorOut = err
		}
		txt := fmt.Sprintf("```\n%s\n```", out)
		i.LogContainer.Add(widget.NewRichTextFromMarkdown(txt))
	} else {
		i.LogContainer.Add(i.newRichLine("`[user] make user-install`"))
		out, err := UserMake()
		if err != nil {
			i.ErrorOut = err
		}
		txt := fmt.Sprintf("```\n%s\n```", out)
		i.LogContainer.Add(widget.NewRichTextFromMarkdown(txt))
	}
}

func (i *InstallConf) Install() {
	defer i.ProgressBar.SetValue(1)
	i.Untar()
	i.Make()
	if i.ErrorOut == nil {
		i.LogContainer.Add(i.newRichLine("`Installation finished!`"))
	} else {
		i.LogContainer.Add(i.newRichLine("`Installation completed with errors.`"))
	}
}

func ExtractTarXz(embeddedData []byte) error {
	filename := "tmp_file.tar.xz"
	folder := "./EduTrack"
	err := os.WriteFile(filename, embeddedData, os.ModePerm)
	if err != nil {
		return err
	}
	if isNotExists(folder) {
		mkdir(folder)
	}
	u := unxz.New(filename, folder)
	return u.Extract()
}
