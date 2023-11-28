/*
 * Copyright (c) 2023 Tom5521- All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package iconloader

import (
	"fyne.io/fyne/v2"
)

// Dark Icons

var (
	DevICONDark          fyne.Resource = Dev_Dark
	DownloadICONDark     fyne.Resource = Download_Dark
	ErrorICONDark        fyne.Resource = Error_Dark
	InstallICONDark      fyne.Resource = Install_Dark
	SaveICONDark         fyne.Resource = Save_Dark
	RestartICONDark      fyne.Resource = Restart_Dark
	InfoICONDark         fyne.Resource = Info_Dark
	UninstallICONDark    fyne.Resource = Uninstall_Dark
	UserTemplateICONDark fyne.Resource = TemplateUser_Dark
	AddUserICONDark      fyne.Resource = AddUser_Dark
	AppICONDark          fyne.Resource = App_dark
)

// Light Icons

var (
	DevICONLight          fyne.Resource = Dev_Light
	DownloadICONLight     fyne.Resource = Download_Light
	ErrorICONLight        fyne.Resource = Error_Light
	InstallICONLight      fyne.Resource = Install_Light
	SaveICONLight         fyne.Resource = Save_Light
	RestartICONLight      fyne.Resource = Restart_Light
	InfoICONLight         fyne.Resource = Info_Light
	UninstallICONLight    fyne.Resource = Uninstall_Light
	UserTemplateICONLight fyne.Resource = TemplateUser_Light
	AddUserICONLight      fyne.Resource = AddUser_Light
	AppICONLight          fyne.Resource = App_light
)

// Themed Icons

var (
	DevICON          fyne.Resource
	DownloadICON     fyne.Resource
	ErrorICON        fyne.Resource
	InstallICON      fyne.Resource
	SaveICON         fyne.Resource
	RestartICON      fyne.Resource
	InfoICON         fyne.Resource
	UninstallICON    fyne.Resource
	UserTemplateICON fyne.Resource
	AddUserICON      fyne.Resource
	AppICON          fyne.Resource
)

// No-Theme Icons

var (
	PlaceholderICON fyne.Resource = Placeholder
)

// SetThemeIcons sets themed icons based on the Fyne theme (light or dark).
func SetThemeIcons(thememode fyne.ThemeVariant) {
	if thememode == 1 { // Set light mode
		DevICON = DevICONLight
		InstallICON = InstallICONLight
		DownloadICON = DownloadICONLight
		ErrorICON = ErrorICONLight
		SaveICON = SaveICONLight
		RestartICON = RestartICONLight
		InfoICON = InfoICONLight
		UninstallICON = UninstallICONLight
		UserTemplateICON = UserTemplateICONLight
		AddUserICON = AddUserICONLight
		AppICON = AppICONLight
	} else if thememode == 0 { // Set dark mode
		DevICON = DevICONDark
		InstallICON = InstallICONDark
		DownloadICON = DownloadICONDark
		ErrorICON = ErrorICONDark
		SaveICON = SaveICONDark
		RestartICON = RestartICONDark
		InfoICON = InfoICONDark
		UninstallICON = UninstallICONDark
		UserTemplateICON = UserTemplateICONDark
		AddUserICON = AddUserICONDark
		AppICON = AppICONDark
	}
}
