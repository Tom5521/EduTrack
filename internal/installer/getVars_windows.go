//go:build windows
// +build windows

package installer

import "github.com/luisiturrios/gowin"

var DefaultWinInstallPath = func() string {
	f := gowin.ShellFolders{Context: gowin.ALL}
	return f.ProgramFiles()
}()
