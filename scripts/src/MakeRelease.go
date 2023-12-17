package src

import (
	"flag"
	"fmt"
	"os"
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
func MakeWinZip() {
	if IsNotExist("tmp/") {
		mkdir("tmp")
	}
	fmt.Println("Downloading opengl32.dll...")
	if IsNotExist("tmp/opengl32.7z") {
		url := "https://downloads.fdossena.com/geth.php?r=mesa64-latest"
		DonwloadFile("tmp/opengl32.7z", url)
	}

	fmt.Println("Unzipping 7z file...")
	if IsNotExist("tmp/opengl") {
		mkdir("tmp/opengl")
	}
	if IsNotExist("tmp/opengl/opengl32.dll") {
		RunCmd("7z e -o\"tmp/opengl\" tmp/opengl32.7z")
	}

	fmt.Println("Compressing for windows")
	if remDir := "builds/EduTrack-win64.zip"; IsExists(remDir) {
		err := os.Remove(remDir)
		if err != nil {
			fmt.Println(err)
		}
	}
	if IsNotExist("EduTrack-win64/") {
		mkdir("EduTrack-win64/")
	}
	CopyFile("builds/EduTrack.exe", "EduTrack-win64/EduTrack.exe")
	CopyFile("tmp/opengl/opengl32.dll", "EduTrack-win64/opengl32.dll")
	RunCmd("zip -r builds/EduTrack-win64.zip EduTrack-win64/")
	os.RemoveAll("EduTrack-win64/")
}

func Main() {
	flag.Parse()
	if IsNotExist("builds/") {
		mkdir("builds")
	}
	parserFlags := []*bool{BuildForLinux, BuildForWindows, ReleaseFlag, MakeWindowsZip, TestRunning, HelpFlag}
	catchBadFlags := func(flags []*bool) bool {
		var trueValues int
		for _, parser := range flags {
			if *parser {
				trueValues++
			}
		}
		return trueValues > 1
	}
	if catchBadFlags(parserFlags) {
		fmt.Println("There are arguments that cannot be used with others!")
		flag.PrintDefaults()
		return
	}
	if *BuildForLinux {
		fmt.Println("Compiling for linux...")
		CompileForLinux()
	}
	if *BuildForWindows {
		fmt.Println("Compiling for windows...")
		CompileForWin()
		MakeWinZip()
	}
	if *ReleaseFlag {
		fmt.Println("Making release...")
		MakeRelease()
	}
	if *MakeWindowsZip {
		fmt.Println("Making windows zip...")
		MakeWinZip()
	}
	if *TestRunning {
		fmt.Println("Can run!")
		fmt.Println(os.Getwd())
	}
	if *HelpFlag {
		flag.PrintDefaults()
	}
}
