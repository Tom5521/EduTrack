/*
 * Copyright Tom5521(c) - All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package data

import (
	"os"

	"github.com/Tom5521/MyGolangTools/file"
	"gopkg.in/yaml.v3"
)

// Student represents the structure of a student's data.
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

var Students []Student // A slice to hold student data.

var ConfName string = "Data.yml" // Default configuration file name.

// LoadConf sets the configuration file name.
func LoadConf(conf string) {
	ConfName = conf
	GetYamlData()
}

// GetYamlData reads student data from the YAML configuration file.
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

// NewYmlFile creates a new YAML file and returns its data.
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

// SaveData saves student data to the YAML configuration file.
func SaveData() error {
	data, err := yaml.Marshal(Students)
	if err != nil {
		return err
	}
	err = os.WriteFile(ConfName, data, os.ModePerm)

	return err
}

// Resave overwrites the YAML file with the provided student data and updates the in-memory data.
func Resave(writer []Student) {
	data, _ := yaml.Marshal(writer)
	os.WriteFile(ConfName, data, os.ModePerm)
	GetYamlData()
}

// GetNames returns a slice of student names.
func GetNames() []string {
	var names []string
	for _, student := range Students {
		names = append(names, student.Name)
	}
	return names
}

// GetIDs returns a slice of student IDs.
func GetIDs() []string {
	var IDs []string
	for _, student := range Students {
		IDs = append(IDs, student.ID)
	}
	return IDs
}

// FindStudentByID searches for a student by their ID and returns a pointer to the student if found.
func FindStudentByID(studentID string) *Student {
	for _, student := range Students {
		if student.ID == studentID {
			return &student
		}
	}
	return nil
}

