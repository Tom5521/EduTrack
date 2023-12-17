package src

import "flag"

var (
	BuildForWindows *bool = flag.Bool("compile-to-windows", false, "Compile for windows")
	BuildForLinux         = flag.Bool("compile-to-linux", false, "Compile to linux")
	ReleaseFlag           = flag.Bool("release", false, "Compile for all plattaforms and make windows zip")
	MakeWindowsZip        = flag.Bool("make-windows-zip", false, "Only makes the windows zip")
)
