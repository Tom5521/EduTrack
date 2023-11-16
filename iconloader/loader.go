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
	DevICON_Dark          fyne.Resource = Dev_Dark
	DownloadICON_Dark     fyne.Resource = Download_Dark
	ErrorICON_Dark        fyne.Resource = Error_Dark
	InstallICON_Dark      fyne.Resource = Install_Dark
	SaveICON_Dark         fyne.Resource = Save_Dark
	RestartICON_Dark      fyne.Resource = Restart_Dark
	InfoICON_Dark         fyne.Resource = Info_Dark
	UninstallICON_Dark    fyne.Resource = Uninstall_Dark
	UserTemplateICON_Dark fyne.Resource = TemplateUser_Dark
	AppICON_Dark          fyne.Resource = App_dark
)

// Light Icons
var (
	DevICON_Light          fyne.Resource = Dev_Light
	DownloadICON_Light     fyne.Resource = Download_Light
	ErrorICON_Light        fyne.Resource = Error_Light
	InstallICON_Light      fyne.Resource = Install_Light
	SaveICON_Light         fyne.Resource = Save_Light
	RestartICON_Light      fyne.Resource = Restart_Light
	InfoICON_Light         fyne.Resource = Info_Light
	UninstallICON_Light    fyne.Resource = Uninstall_Light
	UserTemplateICON_Light fyne.Resource = TemplateUser_Light
	AppICON_Light          fyne.Resource = App_light
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
	AppICON          fyne.Resource
)

// No-Theme Icons
var (
	PlaceholderICON fyne.Resource = Placeholder
)

// SetThemeIcons sets themed icons based on the Fyne theme (light or dark).
func SetThemeIcons(thememode fyne.ThemeVariant) {
	if thememode == 1 { // Set light mode
		DevICON = DevICON_Light
		InstallICON = InstallICON_Light
		DownloadICON = DownloadICON_Light
		ErrorICON = ErrorICON_Light
		SaveICON = SaveICON_Light
		RestartICON = RestartICON_Light
		InfoICON = InfoICON_Light
		UninstallICON = UninstallICON_Light
		UserTemplateICON = UserTemplateICON_Light
		AppICON = AppICON_Light
	} else if thememode == 0 { // Set dark mode
		DevICON = DevICON_Dark
		InstallICON = InstallICON_Dark
		DownloadICON = DownloadICON_Dark
		ErrorICON = ErrorICON_Dark
		SaveICON = SaveICON_Dark
		RestartICON = RestartICON_Dark
		InfoICON = InfoICON_Dark
		UninstallICON = UninstallICON_Dark
		UserTemplateICON = UserTemplateICON_Dark
		AppICON = AppICON_Dark
	}
}
