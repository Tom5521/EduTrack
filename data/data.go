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

	"github.com/ncruces/zenity"
	"gopkg.in/yaml.v3"
)

type Grade struct {
	Name  string
	Info  string
	Price string
	Start string
	End   string
}

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
	Grades []Grade
}

type Config_str struct {
	GradesFile   string
	StudentsFile string
	Lang         string // TODO: Add multilanguage support
}

var (
	Grades           []Grade
	Config           = Config_str{}
	Students         []Student // A slice to hold student data.
	_, ConfigFile, _ string    = getOSConfFile()
)

func NotifyError(text string, err error) {
	zenity.Notify(text + "::" + err.Error())
}

func LoadFiles() {
	GetConfData()
	GetStundentData()
	GetGradesData()
}

func getOSConfFile() (dataYml string, ConfYml string, gradesYml string) {
	if runtime.GOOS == "linux" || runtime.GOOS == "unix" {
		CurrentUser, err := user.Current()
		if err != nil {
			fmt.Println(err)
		}
		confDir := fmt.Sprintf("%v/.config/EduTrack", CurrentUser.HomeDir)
		if _, err := os.Stat(confDir); os.IsNotExist(err) {
			err := os.Mkdir(confDir, os.ModePerm)
			if err != nil {
				NotifyError("Error creating ~/.config/EduTrack/", err)
			}
		}
		return confDir + "/students.yml", confDir + "/config.yml", confDir + "/grades.yml"
	} else {
		return "students.yml", "config.yml", "grades.yml"
	}
}

func LoadConf(d string) {
	Config.StudentsFile = d
	data, err := yaml.Marshal(Config)
	if err != nil {
		NotifyError("Error marshalling the configuration", err)
	}
	err = os.WriteFile(ConfigFile, data, os.ModePerm)
	if err != nil {
		NotifyError("Error loading configuration", err)
	}
	GetStundentData()
}

func NewConfigurationFile() {
	var err error
	StudentsFile, confdir, gradesyml := getOSConfFile()
	ymlData, err := yaml.Marshal(Config_str{StudentsFile: StudentsFile, GradesFile: gradesyml})
	if err != nil {
		NotifyError("Error marshalling new configuration file", err)
	}
	err = os.WriteFile(confdir, ymlData, os.ModePerm)
	if err != nil {
		NotifyError("Error writing config file", err)
	}
}

func GetGradesData() {
	var err error
	if _, err := os.Stat(Config.GradesFile); os.IsNotExist(err) {
		NewGradesFile()
	}

	data, err := os.ReadFile(Config.GradesFile)
	if err != nil {
		NotifyError("Error reading grades file", err)
	}

	yaml.Unmarshal(data, &Grades)
}

func GetConfData() {
	var err error
	_, confFile, _ := getOSConfFile()
	if _, err := os.Stat(confFile); os.IsNotExist(err) {
		NewConfigurationFile()
	}
	data, err := os.ReadFile(confFile)
	if err != nil {
		NotifyError("Error reading config file!", err)
	}
	yaml.Unmarshal(data, &Config)
}

func NewYamlStudentsFile() {
	var err error
	data, err := yaml.Marshal(Students)
	if err != nil {
		NotifyError("Error marshalling students file", err)
	}
	err = os.WriteFile(Config.StudentsFile, data, os.ModePerm)
	if err != nil {
		NotifyError("Error writing new students file", err)
	}
}

func NewGradesFile() {
	var err error
	data, err := yaml.Marshal(Grades)
	if err != nil {
		NotifyError("Error marshalling grades var", err)
	}
	err = os.WriteFile(Config.GradesFile, data, os.ModePerm)
	if err != nil {
		NotifyError("Error writing new grades file", err)
	}
}

// GetStundentData reads student data from the YAML configuration file.
func GetStundentData() {
	var err error

	if _, err := os.Stat(Config.StudentsFile); os.IsNotExist(err) {
		NewYamlStudentsFile()
	}

	data, err := os.ReadFile(Config.StudentsFile)
	if err != nil {
		NotifyError("Error reading students file", err)
	}
	yaml.Unmarshal(data, &Students)
}

// SaveStudentsData saves student data to the YAML configuration file.
func SaveStudentsData() error {
	if _, err := os.Stat(Config.GradesFile); os.IsNotExist(err) {
		NewYamlStudentsFile()
	}
	data, err := yaml.Marshal(Students)
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
	data, err := yaml.Marshal(Grades)
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

// GetStudentsNames returns a slice of student names.
func GetStudentNames() []string {
	var names []string
	for _, student := range Students {
		names = append(names, student.Name)
	}
	return names
}

// GetStudentIDs returns a slice of student IDs.
func GetStudentIDs() []string {
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

func GetGradesNames() []string {
	var grades []string
	for _, grade := range Grades {
		grades = append(grades, grade.Name)
	}
	return grades
}
