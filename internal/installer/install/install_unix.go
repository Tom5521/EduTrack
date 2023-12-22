//go:build unix
// +build unix

package install

import (
	"io"
	"os"

	"github.com/Tom5521/CmdRunTools/command"
	"github.com/ncruces/zenity"
	"github.com/yi-ge/unxz"
)

func (i *InstallConf) Untar() {
	defer i.ProgressBar.SetValue(0.4)
	f, err := Files.ReadFile("files/EduTrack-linux64.tar.xz")
	if err != nil {
		errWin(err.Error())
	}
	chdir("/tmp")
	err = ExtractTarXz(f)
	if err != nil {
		errWin(err.Error())
	}
}

func (i *InstallConf) Make() {
	defer i.ProgressBar.SetValue(0.8)
	chdir("EduTrack/")
	if i.Linux.RootInstall {
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
	} else {
		cmd := command.InitCmd("make user-install")
		cmd.UseBashShell(true)
		cmd.CustomStd(true, true, true)
		err := cmd.Run()
		if err != nil {
			errWin(err.Error())
		}
	}
}

func (i *InstallConf) Install() {
	i.Untar()
	i.Make()
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