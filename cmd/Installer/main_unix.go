//go:build unix
// +build unix

package main

import "github.com/Tom5521/EduTrack/internal/installer"

func main() {
	gui := installer.InitGUI()
	gui.MainWin()
}
