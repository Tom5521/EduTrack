//go:build AddPo
// +build AddPo

// Run this with go run -tags AddPo locales/AddPo.go

package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type yamlData struct {
	Route string `yaml:"route"`
	Msgid string `yaml:"msgid"`
}

func GetFilesInDirectory(directory string) ([]string, error) {
	var files []string

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}
func AddLineToFile(filename, line string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = fmt.Fprintln(writer, line)
	if err != nil {
		return err
	}

	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}

func ReadFile() yamlData {
	ret := yamlData{}
	file, _ := os.ReadFile("locales/last_add.yml")
	yaml.Unmarshal(file, &ret)
	return ret
}

const Template string = `
#: %s
msgid "%s"
msgstr ""
`

func main() {
	yml := ReadFile()
	dirs, _ := GetFilesInDirectory("locales/po/")
	for _, filename := range dirs {
		AddLineToFile(filename, fmt.Sprintf(Template, yml.Route, yml.Msgid))
	}
}
