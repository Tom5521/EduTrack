package data

import (
	"fmt"
	"os"
	"os/user"
	"runtime"

	"gopkg.in/yaml.v3"
)

func LoadFiles() {
	Config = GetConfData()
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

	yaml.Unmarshal(data, &Data.Grades)
}

func NewYamlStudentsFile() {
	var err error
	data, err := yaml.Marshal(Data.Students)
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
	data, err := yaml.Marshal(Data.Grades)
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
	yaml.Unmarshal(data, &Data.Students)
}
