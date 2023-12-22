package src

import (
	"flag"
	"fmt"
	"os"
)

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
