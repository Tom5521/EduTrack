//go:build mage && (windows || linux)

package main

import (
	"os"
	"runtime"

	"github.com/Tom5521/GoNotes/pkg/messages"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var (
	TemporalDir = "./tmp"
	Mesa64Url   = "https://downloads.fdossena.com/geth.php?r=mesa64-latest"
	MainDir     = "./cmd/EduTrack/"
	WindowsEnv  = windowsEnv()
	// Mage config.
	Aliases = map[string]interface{}{
		"user-install": UserInstall,
	}
)

type Build mg.Namespace

func Test() error {
	return movefile("cmd/EduTrack/builds/EduTrack.exe", "builds/EduTrack.exe")
	return nil
}

func copyfile(src, dest string) error {
	source, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	err = os.WriteFile(dest, source, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func movefile(src, dest string) error {
	err := copyfile(src, dest)
	if err != nil {
		return err
	}
	err = os.Remove(src)
	if err != nil {
		return err
	}
	return nil
}

func downloadWinFiles() error {
	if _, err := os.Stat(TemporalDir); os.IsNotExist(err) {
		err = os.Mkdir(TemporalDir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	if _, err := os.Stat(TemporalDir + "/opengl32.7z"); os.IsNotExist(err) {
		err = sh.RunV("wget", "-O", "tmp/opengl32.7z", Mesa64Url)
		if err != nil {
			return err
		}
	}
	if _, err := os.Stat(TemporalDir + "/opengl32.dll"); os.IsNotExist(err) {
		err = os.Chdir(TemporalDir)
		if err != nil {
			return err
		}
		if err = sh.RunV("7z", "e", "opengl32.7z"); err != nil {
			return err
		}
		err = os.Chdir("..")
		if err != nil {
			return err
		}
	}
	return nil
}

func windowsEnv() map[string]string {
	var env map[string]string
	if runtime.GOOS == "linux" {
		env = map[string]string{
			"CC":          "/usr/bin/x86_64-w64-mingw32-gcc",
			"CXX":         "/usr/bin/x86_64-w64-mingw32-c++",
			"CGO_ENABLED": "1",
		}
	}
	return env
}

func checkdir() error {
	if _, err := os.Stat("builds"); os.IsNotExist(err) {
		err = os.Mkdir("builds", os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func (Build) LinuxInstaller() error {
	if err := checkdir(); err != nil {
		return err
	}
	return nil
}

func (Build) WindowsInstaller() error {
	if err := checkdir(); err != nil {
		return err
	}
	if _, err := os.Stat("tmp/opengl32.7z"); os.IsNotExist(err) {
		if err = downloadWinFiles(); err != nil {
			return err
		}
	}
	return nil
}

func (Build) Installer() error {
	if err := checkdir(); err != nil {
		return err
	}
	return nil
}

func (Build) Windows() error {
	if err := checkdir(); err != nil {
		return err
	}
	err := sh.RunWithV(WindowsEnv, "fyne", "package", "--os", "windows", "--release",
		"--tags", "windows", "--src", MainDir, "--exe", "builds/EduTrack.exe")
	if err != nil {
		return err
	}
	err = movefile(MainDir+"/builds/EduTrack.exe", "./builds/EduTrack.exe")
	if err != nil {
		return err
	}
	return nil
}

func (Build) Linux() error {
	if err := checkdir(); err != nil {
		return err
	}
	return nil
}

func Install() error {
	return nil
}

func UserInstall() error {
	return nil
}

func Clean() {
	var errorList []error
	appendErr := func(err error) {
		if err != nil {
			errorList = append(errorList, err)
		}
	}
	err := sh.Rm("tmp")
	appendErr(err)
	err = sh.Rm("builds")
	appendErr(err)
	err = sh.Rm("./cmd/EduTrack/EduTrack")
	appendErr(err)
	err = sh.Rm("./cmd/EduTrack/EduTrack.exe")
	appendErr(err)
	err = sh.Rm("./cmd/Installer/EduTrack Installer.exe")
	appendErr(err)
	err = sh.Rm("./cmd/Installer/EduTrack Installer")
	appendErr(err)
	for _, e := range errorList {
		messages.Warning(e.Error())
	}
}
