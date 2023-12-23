package src

import "flag"

var (
	BuildForWindows    *bool = flag.Bool("compile-to-windows", false, "Compile for windows")
	BuildForLinux            = flag.Bool("compile-to-linux", false, "Compile to linux")
	ReleaseFlag              = flag.Bool("release", false, "Compile for all plattaforms and make windows zip")
	MakeWindowsZip           = flag.Bool("make-windows-zip", false, "Only makes the windows zip")
	TestRunning              = flag.Bool("test-run", false, "Test if the script can run")
	HelpFlag                 = flag.Bool("help", false, "Print this message")
	MakeWinInstaller         = flag.Bool("make-windows-installer", false, "Compile for windows to make a installer")
	MakeLinuxInstaller       = flag.Bool("make-linux-installer", false, "Compile for linux to make it installer")
)
