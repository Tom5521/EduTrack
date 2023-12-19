package src

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func RunCmd(command string) {
	c := strings.Fields(command)
	var cmd *exec.Cmd = exec.Command(c[0], c...)
	if runtime.GOOS == "linux" {
		cmd = exec.Command("bash", "-c", command)
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func SetEnv(key, v string) {
	err := os.Setenv(key, v)
	if err != nil {
		fmt.Println(err)
	}
}

func SetEnvForWin() {
	SetEnv("GGO_ENABLED", "1")
	SetEnv("CC", "/usr/bin/x86_64-w64-mingw32-gcc")
	SetEnv("CXX", "/usr/bin/x86_64-w64-mingw32-c++")
	SetEnv("GOOS", "windows")
}

var mainDir = "./cmd/EduTrack/"

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
	MoveFile("EduTrack.tar.xz", "builds/EduTrack-linux64.tar.xz")

}
