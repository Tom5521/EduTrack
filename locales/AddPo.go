//go:build AddPo
// +build AddPo

// Run this with go run -tags AddPo locales/AddPo.go

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type YamlData struct {
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

func ReadFile(f string) YamlData {
	ret := YamlData{}
	file, _ := os.ReadFile(f)
	yaml.Unmarshal(file, &ret)
	return ret
}

const Template string = `

#: %s
msgid "%s"
msgstr ""
`

var dirFlag = flag.Bool("dir", false, "Add po msgid to another directory")

func main() {
	flag.Parse()
	if *dirFlag {
		if len(os.Args) < 3 {
			fmt.Println("Not enough arguments to dir flag!")
			fmt.Println(os.Args)
			return
		}
		os.Chdir(os.Args[2])
		yml := ReadFile("./last_add.yml")
		fmt.Println(os.Args[2])
		dirs, _ := GetFilesInDirectory("./po")
		txt := fmt.Sprintf(Template, yml.Route, yml.Msgid)
		for _, name := range dirs {
			AddLineToFile(name, txt)
		}
		fmt.Println(fmt.Println(txt))
		return
	}
	yml := ReadFile("locales/last_add.yml")
	dirs, _ := GetFilesInDirectory("locales/po/")
	for _, filename := range dirs {
		AddLineToFile(filename, fmt.Sprintf(Template, yml.Route, yml.Msgid))
	}
}
