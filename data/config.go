package data

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config_str struct {
	GradesFile   string
	StudentsFile string
	Lang         string // TODO: Add multilanguage support
}

func GetConfData() Config_str {
	var err error
	conf := Config_str{}
	_, confFile, _ := getOSConfFile()
	if _, err := os.Stat(confFile); os.IsNotExist(err) {
		NewConfigurationFile()
	}
	data, err := os.ReadFile(confFile)
	if err != nil {
		NotifyError("Error reading config file!", err)
	}
	yaml.Unmarshal(data, &conf)
	return conf
}
