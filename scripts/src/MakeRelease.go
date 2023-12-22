package src

import (
	"fmt"
	"runtime"
)

func MakeRelease() {
	fmt.Println("Compiling for linux...")
	CompileForLinux()
	fmt.Println("Compiling for windows...")
	CompileForWin()
	fmt.Println("Making windows zip")
	MakeWinZip()
	if runtime.GOOS == "linux" {
		RunCmd("notify-send \"Release compilation ended susefully\"")
	}
}
