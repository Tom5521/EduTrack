package src

import (
	"fmt"
	"os"
)

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
	binDir := "cmd/EduTrack/builds/EduTrack.exe"
	CopyFile(binDir, "builds/EduTrack.exe")
	CopyFile(binDir, "EduTrack-win64/EduTrack.exe")
	CopyFile("tmp/opengl/opengl32.dll", "EduTrack-win64/opengl32.dll")
	RunCmd("zip -r builds/EduTrack-win64.zip EduTrack-win64/")
	os.RemoveAll("EduTrack-win64/")
}
