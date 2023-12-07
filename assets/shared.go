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
var lightFS embed.FS

//go:embed IconsDark
var darkFS embed.FS

//go:embed Shared
var commonFS embed.FS

func cmReader(file string) []byte {
	r, err := commonFS.ReadFile("Shared/" + file)
	if err != nil {
		log.Println(err)
	}
	return r
}

func read(file string) []byte {
	var fs embed.FS
	var themedir string
	if fyne.CurrentApp().Settings().ThemeVariant() == 1 {
		themedir = "IconsLight/"
		fs = lightFS
	} else {
		themedir = "IconsDark/"
		fs = darkFS
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
		StaticContent: read(resourceName),
	}
}

func getCmResource(name string) *fyne.StaticResource {
	return &fyne.StaticResource{
		StaticName:    name,
		StaticContent: cmReader(name),
	}
}
