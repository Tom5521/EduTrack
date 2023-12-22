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
		RunCmd("sudo fyne-cross -os linux -release -dir ./cmd/EduTrack/")
	}
	err := MoveFile("EduTrack.tar.xz", "builds/EduTrack-linux64.tar.xz")
	if err != nil {
		fmt.Println(err)
	}
}
