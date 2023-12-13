/*
 * Copyright (c) 2023 Tom5521- All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package assets

import (
	"embed"
	"log"

	"fyne.io/fyne/v2"
)

//go:embed IconsLight
var LightFS embed.FS

//go:embed IconsDark
var DarkFS embed.FS

//go:embed Shared
var commonFS embed.FS

func cmReader(file string) []byte {
	r, err := commonFS.ReadFile("Shared/" + file)
	if err != nil {
		log.Println(err)
	}
	return r
}

func Read(file string) []byte {
	var fs embed.FS
	var themedir string
	if fyne.CurrentApp().Settings().ThemeVariant() == 1 {
		themedir = "IconsLight/"
		fs = LightFS
	} else {
		themedir = "IconsDark/"
		fs = DarkFS
	}
	r, err := fs.ReadFile(themedir + file)
	if err != nil {
		log.Println(err)
		return cmReader("Placeholder.png")
	}
	return r
}

func getResource(resourceName string) *fyne.StaticResource {
	return &fyne.StaticResource{
		StaticName:    resourceName,
		StaticContent: Read(resourceName),
	}
}

func getCmResource(name string) *fyne.StaticResource {
	return &fyne.StaticResource{
		StaticName:    name,
		StaticContent: cmReader(name),
	}
}
