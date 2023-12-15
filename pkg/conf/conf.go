package conf

import (
	"encoding/json"
	"log"
	"os"
	"os/user"
	"runtime"

	"github.com/ncruces/zenity"
)

var (
	ConfDir       = GetConfDir()
	ConfFile      = ConfDir + "/config.json"
	DefaultDBFile = ConfDir + "/database.db"
)

type Config struct {
	DatabaseFile string `json:"database"`
	Lang         string `json:"lang"`
	Theme        string `json:"theme"`
}

func GetConfData() Config {
	conf := Config{}
	data, err := os.ReadFile(ConfFile)
	if err != nil {
		log.Println("Error reading config file!", err)
	}
	err = json.Unmarshal(data, &conf)
	if err != nil {
		log.Println(err)
	}
	return conf
}

func GetConfDir() string {
	cu, err := user.Current()
	if err != nil {
		zenity.Error(err.Error())
	}
	unixOS := []string{"linux", "darwin", "freebsd", "openbsd", "netbsd"}
	isUnix := func() bool {
		for _, os := range unixOS {
			if runtime.GOOS == os {
				return true
			}
		}
		return false
	}
	if isUnix() {
		dir := cu.HomeDir + "/.config/EduTrack"
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			err = os.Mkdir(dir, os.ModePerm)
			if err != nil {
				zenity.Error(err.Error())
			}
		}
		return dir
	} else if runtime.GOOS == "windows" {
		dir := cu.HomeDir + "/Documents/EduTrack"
		if _, err = os.Stat(dir); os.IsNotExist(err) {
			err = os.Mkdir(dir, os.ModePerm)
			if err != nil {
				zenity.Error(err.Error())
			}
		}
		return dir
	}
	return "./"
}

func NewConfigurationFile() {
	var err error
	jsonData, err := json.Marshal(Config{DatabaseFile: DefaultDBFile, Theme: "Default", Lang: "en"})
	if err != nil {
		zenity.Error(err.Error())
	}
	err = os.WriteFile(ConfFile, jsonData, os.ModePerm)
	if err != nil {
		zenity.Error(err.Error())
	}
}
