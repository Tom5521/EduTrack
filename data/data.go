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

type Config_str struct {
	DataFile string
}

var (
	Config     = GetConfigurationData()
	Students   []Student              // A slice to hold student data.
	DataName   string    = "data.yml" // Default configuration file name.
	ConfigName string    = "config.yml"
)

// LoadConf sets the configuration file name.
func LoadConf(conf string) {
	DataName = conf
	GetYamlData()
}

func GetConfigurationData() Config_str {
	data := Config_str{}
	bytes := fileChecks(ConfigName)
	yaml.Unmarshal(bytes, &data)
	return data
}

func fileChecks(DataName string) []byte {
	var (
		err       error
		data_file []byte
	)
	if check, _ := file.CheckFile(DataName); !check {
		data_file = NewYmlFile(DataName)
	} else {
		data_file, err = os.ReadFile(DataName)
		if err != nil {
			return nil
		}
	}
	return data_file
}

// GetYamlData reads student data from the YAML configuration file.
func GetYamlData() {
	data_file := fileChecks(DataName)
	yaml.Unmarshal(data_file, &Students)
}

// NewYmlFile creates a new YAML file and returns its data.
func NewYmlFile(file_field string) []byte {
	_, err := os.Create(file_field)
	if err != nil {
		return nil
	}
	data, err := os.ReadFile(file_field)
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
	err = os.WriteFile(DataName, data, os.ModePerm)

	return err
}

// Resave overwrites the YAML file with the provided student data and updates the in-memory data.
func Resave(writer []Student) {
	data, _ := yaml.Marshal(writer)
	os.WriteFile(DataName, data, os.ModePerm)
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
