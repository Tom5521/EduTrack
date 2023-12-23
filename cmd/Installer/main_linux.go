//go:build linux
// +build linux

package main

import (
	"github.com/Tom5521/EduTrack/internal/installer"
)

func main() {
	gui := installer.InitGUI()
	gui.MainWin()
}
