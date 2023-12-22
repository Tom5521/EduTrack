package src

import (
	"fmt"
	"os"
)

func SetEnvForWin() {
	setEnv := func(key, v string) {
		err := os.Setenv(key, v)
		if err != nil {
			fmt.Println(err)
		}
	}
	setEnv("GGO_ENABLED", "1")
	setEnv("CC", "/usr/bin/x86_64-w64-mingw32-gcc")
	setEnv("CXX", "/usr/bin/x86_64-w64-mingw32-c++")
	setEnv("GOOS", "windows")
}
