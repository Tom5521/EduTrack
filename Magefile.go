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
	TmpDir     = "./tmp"
	Mesa64Url  = "https://downloads.fdossena.com/geth.php?r=mesa64-latest"
	MainDir    = "./cmd/EduTrack/"
	WindowsEnv = windowsEnv()
	build      = Build{}
)

type Build mg.Namespace
type Install mg.Namespace
type Uninstall mg.Namespace

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
	if _, err := os.Stat(TmpDir); os.IsNotExist(err) {
		err = os.Mkdir(TmpDir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	if _, err := os.Stat(TmpDir + "/opengl32.7z"); os.IsNotExist(err) {
		err = sh.RunV("wget", "-O", "tmp/opengl32.7z", Mesa64Url)
		if err != nil {
			return err
		}
	}
	if _, err := os.Stat(TmpDir + "/opengl32.dll"); os.IsNotExist(err) {
		err = os.Chdir(TmpDir)
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
	if _, err := os.Stat("builds/EduTrack-linux64.tar.xz"); os.IsNotExist(err) {
		err = build.Linux()
		if err != nil {
			return err
		}
	}
	err := copyfile("./builds/EduTrack-linux64.tar.xz", "./internal/installer/install/files/EduTrack-linux64.tar.xz")
	if err != nil {
		return err
	}
	err = sh.RunV("go", "build", "-v", "-o", "builds/EduTrack-Installer-linux64", "./cmd/Installer/main_linux.go")
	if err != nil {
		return err
	}
	return nil
}

func (Build) All() error {
	err := build.Linux()
	if err != nil {
		return err
	}
	err = build.Windows()
	if err != nil {
		return err
	}
	err = build.LinuxInstaller()
	if err != nil {
		return err
	}
	err = build.WindowsInstaller()
	if err != nil {
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
	err := copyfile("tmp/opengl32.dll", "./internal/installer/install/files/opengl32.dll")
	if err != nil {
		return err
	}
	if _, err = os.Stat("builds/EduTrack.exe"); os.IsNotExist(err) {
		err = build.Windows()
		if err != nil {
			return err
		}
	}
	err = copyfile("builds/EduTrack.exe", "./internal/installer/install/files/EduTrack.exe")
	if err != nil {
		return err
	}
	err = sh.RunWith(WindowsEnv, "fyne", "package", "--os", "windows", "--release", "--src", "./cmd/Installer/",
		"--exe", "builds/EduTrack-Installer-win64.exe", "--tags", "windows")
	if err != nil {
		return err
	}
	err = movefile("./cmd/Installer/builds/EduTrack-Installer-win64.exe", "builds/EduTrack-Installer-win64.exe")
	if err != nil {
		return err
	}
	return nil
}

// Compile the program to be distributed on windows, NOTE: This will only return an .exe of the program, the installation in windows can only be done through the installer.
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

// Compile the program to be distributed on linux.
func (Build) Linux() error {
	if err := checkdir(); err != nil {
		return err
	}
	err := sh.RunV("fyne", "package", "--os", "linux", "--release", "--tags", "linux", "--src", MainDir)
	if err != nil {
		return err
	}
	err = movefile("EduTrack.tar.xz", "builds/EduTrack-linux64.tar.xz")
	if err != nil {
		return err
	}
	return nil
}

func setupLinuxMake() error {
	if _, err := os.Stat("builds/EduTrack-linux64.tar.xz"); os.IsNotExist(err) {
		err = build.LinuxInstaller()
		if err != nil {
			return err
		}
	}
	err := os.Chdir("builds")
	if err != nil {
		return err
	}
	if _, err = os.Stat("Makefile"); os.IsNotExist(err) {
		err = sh.RunV("tar", "-xvf", "EduTrack-linux64.tar.xz")
		if err != nil {
			return err
		}
	}

	return nil
}

// NOTE: Only works in linux, in windows you will have to use the installer.
func (Install) Root() error {
	err := setupLinuxMake()
	if err != nil {
		return err
	}
	err = sh.RunV("sudo", "make", "install")
	if err != nil {
		return err
	}
	err = os.Chdir("..")
	if err != nil {
		return err
	}
	return nil
}

// NOTE: Only works in linux, in windows you will have to use the installer.
func (Install) User() error {
	err := setupLinuxMake()
	if err != nil {
		return err
	}
	err = sh.RunV("make", "user-install")
	if err != nil {
		return err
	}
	err = os.Chdir("..")
	if err != nil {
		return err
	}
	return nil
}

// Delete temporary directories, compilation files, etc, It leaves it as if it had just been cloned.
func Clean() {
	var errorList []error
	rm := func(src string) {
		err := sh.Rm(src)
		if err != nil {
			errorList = append(errorList, err)
		}
	}
	rm("tmp")
	rm("builds")
	rm("./cmd/EduTrack/EduTrack")
	rm("./cmd/EduTrack/EduTrack.exe")
	rm("./cmd/Installer/EduTrack Installer.exe")
	rm("./cmd/Installer/EduTrack Installer")
	rm("./cmd/EduTrack/builds/")
	rm("./cmd/Installer/builds/")
	for _, e := range errorList {
		messages.Warning(e.Error())
	}
}

func MakeWindowsZip() error {
	var zipDir = "windows-tmp"
	if _, err := os.Stat(zipDir); os.IsNotExist(err) {
		err = os.Mkdir(zipDir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	if _, err := os.Stat("tmp/opengl32.dll"); os.IsNotExist(err) {
		err = downloadWinFiles()
		if err != nil {
			return err
		}
	}
	err := copyfile("tmp/opengl32.dll", zipDir+"/opengl32.dll")
	if err != nil {
		return err
	}
	if _, err = os.Stat("builds/EduTrack.exe"); os.IsNotExist(err) {
		err = build.Windows()
		if err != nil {
			return err
		}
	}
	err = copyfile("builds/EduTrack.exe", zipDir+"/EduTrack.exe")
	if err != nil {
		return err
	}
	err = copyfile("README.md", zipDir+"/README.md")
	if err != nil {
		return err
	}
	if _, err = os.Stat("builds/EduTrack-win64.zip"); os.IsExist(err) {
		err = os.Remove("builds/EduTrack-win64.zip")
		if err != nil {
			return err
		}
	}
	err = os.Chdir(zipDir)
	if err != nil {
		return err
	}

	err = sh.RunV("zip", "-r", "../builds/EduTrack-win64.zip", ".")
	if err != nil {
		return err
	}
	err = os.Chdir("..")
	if err != nil {
		return err
	}
	err = os.RemoveAll(zipDir)
	if err != nil {
		return err
	}
	return nil
}

func (Uninstall) User() error {
	err := setupLinuxMake()
	if err != nil {
		return err
	}
	err = sh.RunV("make", "user-uninstall")
	if err != nil {
		return err
	}
	err = os.Chdir("..")
	if err != nil {
		return err
	}
	return nil
}

func (Uninstall) Root() error {
	err := setupLinuxMake()
	if err != nil {
		return err
	}
	err = sh.RunV("sudo", "make", "uninstall")
	if err != nil {
		return err
	}
	err = os.Chdir("..")
	if err != nil {
		return err
	}
	return nil
}
