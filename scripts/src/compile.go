package src

import (
	"fmt"
	"runtime"
)

const mainDir = "./cmd/EduTrack/"

func CompileForWin() {
	if runtime.GOOS == "linux" {
		SetEnvForWin()
	}
	cmd := "fyne package --os windows --exe builds/EduTrack.exe --release --src " + mainDir
	RunCmd(cmd)
}

func CompileForLinux() {
	if runtime.GOOS != "windows" {
		RunCmd("fyne package --os linux --release --src ./cmd/EduTrack/")
	} else {
		RunCmd("fyne-cross -os linux -release -dir ./cmd/EduTrack/")
	}
	err := MoveFile("EduTrack.tar.xz", "builds/EduTrack-linux64.tar.xz")
	if err != nil {
		fmt.Println(err)
	}
}

func CompileLinuxIstaller() {
	if IsNotExist("internal/installer/install/files") {
		mkdir("internal/installer/install/files")
	}
	CopyFile("builds/EduTrack-linux64.tar.xz", "internal/installer/install/files/EduTrack-linux64.tar.xz")
	RunCmd("go build -o builds/EduTrackInstaller cmd/Installer/main_linux.go")
}

func CompileWindowsInstaller() {
	if IsNotExist("internal/installer/install/files") {
		mkdir("internal/installer/install/files")
	}
	CopyFile("tmp/opengl/opengl32.dll", "internal/installer/install/files/opengl32.dll")
	CopyFile("builds/EduTrack.exe", "internal/installer/install/files/EduTrack.exe")
	if runtime.GOOS != "windows" {
		SetEnvForWin()
	}
	RunCmd("go build -o builds/EduTrack-Windows-Installer.exe -ldflags -H=windowsgui cmd/Installer/main_windows.go")
}
