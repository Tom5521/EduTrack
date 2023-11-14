package data

import (
	"os"

	"gopkg.in/yaml.v3"
)

// SaveStudentsData saves student data to the YAML configuration file.
func SaveStudentsData() error {
	if _, err := os.Stat(Config.GradesFile); os.IsNotExist(err) {
		NewYamlStudentsFile()
	}
	data, err := yaml.Marshal(Data.Students)
	if err != nil {
		return err
	}
	err = os.WriteFile(Config.StudentsFile, data, os.ModePerm)

	return err
}

func SaveGradesData() error {
	if _, err := os.Stat(Config.GradesFile); os.IsNotExist(err) {
		NewGradesFile()
	}
	data, err := yaml.Marshal(Data.Grades)
	if err != nil {
		return err
	}
	err = os.WriteFile(Config.GradesFile, data, os.ModePerm)

	return err
}

// Resave overwrites the YAML file with the provided student data and updates the in-memory data.
func Resave(writer []Student) {
	data, _ := yaml.Marshal(writer)
	os.WriteFile(Config.StudentsFile, data, os.ModePerm)
	GetStundentData()
}
