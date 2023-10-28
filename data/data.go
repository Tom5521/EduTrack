/*
 * Copyright Tom5521(c) - All Rights Reserved.
 *
 * This project is licenced under the MIT License.
 */

package data

import (
	"os"

	"github.com/Tom5521/MyGolangTools/file"
	"gopkg.in/yaml.v3"
)

var Students []Student

type Student struct {
	Name          string
	Age           int
	ID            string
	Phone_number  string
	ImageFilePath string
	Register      []struct {
		Date string
		Name string
		Data string
	}
}

var (
	ConfName string = "Data.yml"
)

func LoadConf(conf string) {
	ConfName = conf
	GetYamlData()
}

func GetYamlData() {
	var (
		err       error
		data_file []byte
	)
	if check, _ := file.CheckFile(ConfName); !check {
		data_file = NewYmlFile()
	} else {
		data_file, err = os.ReadFile(ConfName)
		if err != nil {
			return
		}
	}
	yaml.Unmarshal(data_file, &Students)
}

func NewYmlFile() []byte {
	_, err := os.Create(ConfName)
	if err != nil {
		return nil
	}
	data, err := os.ReadFile(ConfName)
	if err != nil {
		return nil
	}
	return data
}

func SaveData() error {
	data, err := yaml.Marshal(Students)
	if err != nil {
		return err
	}
	err = os.WriteFile(ConfName, data, os.ModePerm)

	return err
}
func Resave(writer []Student) {
	data, _ := yaml.Marshal(writer)
	os.WriteFile(ConfName, data, os.ModePerm)
	GetYamlData()
}

func GetNames() []string {
	var names []string
	for _, student := range Students {
		names = append(names, student.Name)
	}
	return names
}

func GetIDs() []string {
	var IDs []string
	for _, student := range Students {
		IDs = append(IDs, student.ID)
	}
	return IDs
}

func FindStudentByID(studentID string) *Student {
	for _, student := range Students {
		if student.ID == studentID {
			return &student
		}
	}
	return nil
}
