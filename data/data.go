/*
 * Copyright Tom5521(c) - All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package data

import (
	"fmt"
	"os"
	"os/user"
	"runtime"

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
	StudentsFile string
	Lang         string // TODO: Add multilanguage support

}

var (
	Config        = Config_str{}
	Students      []Student // A slice to hold student data.
	_, ConfigFile string    = getOSConfFile()
)

func LoadFiles() {
	GetConfData()
	GetYamlData()
}

func getOSConfFile() (dataYml string, ConfYml string) {
	if runtime.GOOS == "linux" || runtime.GOOS == "unix" {
		CurrentUser, err := user.Current()
		if err != nil {
			fmt.Println(err)
		}
		confDir := fmt.Sprintf("%v/.config/EduTrack", CurrentUser.HomeDir)
		if _, err := os.Stat(confDir); os.IsNotExist(err) {
			err := os.Mkdir(confDir, os.ModePerm)
			if err != nil {
				fmt.Println(err)
			}
		}
		return confDir + "/students.yml", confDir + "/config.yml"
	} else {
		return "students.yml", "config.yml"
	}
}

func LoadConf(d string) {
	Config.StudentsFile = d
	data, err := yaml.Marshal(Config)
	if err != nil {
		fmt.Println(err)
	}
	err = os.WriteFile(ConfigFile, data, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	GetYamlData()
}

func NewConfigurationFile() {
	var err error
	StudentsFile, confdir := getOSConfFile()
	ymlData, err := yaml.Marshal(Config_str{StudentsFile: StudentsFile})
	if err != nil {
		fmt.Println(err)
	}
	err = os.WriteFile(confdir, ymlData, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}

func GetConfData() {
	var err error
	_, confFile := getOSConfFile()
	if _, err := os.Stat(confFile); os.IsNotExist(err) {
		NewConfigurationFile()
	}
	data, err := os.ReadFile(confFile)
	if err != nil {
		fmt.Println(err)
	}
	yaml.Unmarshal(data, &Config)
}

func NewYamlStudentsFile() {
	var err error
	data, err := yaml.Marshal(Students)
	if err != nil {
		fmt.Println(err)
	}
	err = os.WriteFile(Config.StudentsFile, data, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}

// GetYamlData reads student data from the YAML configuration file.
func GetYamlData() {
	var err error

	if _, err := os.Stat(Config.StudentsFile); os.IsNotExist(err) {
		NewYamlStudentsFile()
	}

	data, err := os.ReadFile(Config.StudentsFile)
	if err != nil {
		fmt.Println(err)
	}
	yaml.Unmarshal(data, &Students)
}

// NewYmlFile creates a new YAML file and returns its data.

// SaveData saves student data to the YAML configuration file.
func SaveData() error {
	data, err := yaml.Marshal(Students)
	if err != nil {
		return err
	}
	err = os.WriteFile(Config.StudentsFile, data, os.ModePerm)

	return err
}

// Resave overwrites the YAML file with the provided student data and updates the in-memory data.
func Resave(writer []Student) {
	data, _ := yaml.Marshal(writer)
	os.WriteFile(Config.StudentsFile, data, os.ModePerm)
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
