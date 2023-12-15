package conf

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"runtime"

	"github.com/ncruces/zenity"
)

var (
	Config        Conf
	ConfDir       = GetConfDir()
	ConfFile      = ConfDir + "/config.json"
	DefaultDBFile = ConfDir + "/database.db"
)

type Conf struct {
	DatabaseFile string `json:"database"`
	Lang         string `json:"lang"`
	Theme        string `json:"theme"`
}

func errWin(err error, optText ...string) {
	if err != nil {
		var otxt string
		if len(optText) >= 1 {
			otxt = optText[0]
		}
		text := fmt.Sprintf(otxt, err.Error())
		if zenity.Error(text) != nil {
			fmt.Println(err)
		}
	}
}

func GetConfData() Conf {
	conf := Conf{}
	data, err := os.ReadFile(ConfFile)
	errWin(err, "Error reading config file!")
	err = json.Unmarshal(data, &conf)
	errWin(err)
	return conf
}

func GetConfDir() string {
	cu, err := user.Current()
	errWin(err)
	mkdir := func(dir string) {
		err = os.Mkdir(dir, os.ModePerm)
		errWin(err)
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
		if _, err = os.Stat(dir); os.IsNotExist(err) {
			mkdir(dir)
		}
		return dir
	} else if runtime.GOOS == "windows" {
		dir := cu.HomeDir + "/Documents/EduTrack"
		if _, err = os.Stat(dir); os.IsNotExist(err) {
			mkdir(dir)
		}
		return dir
	}
	return "./"
}
func (c *Conf) Update() {
	jsonData, err := json.Marshal(c)
	errWin(err)
	err = os.WriteFile(ConfFile, jsonData, os.ModePerm)
	errWin(err)
	newc := GetConfData()
	c = &newc
	fmt.Println(c)
}

func NewConfigurationFile() {
	var err error
	jsonData, err := json.Marshal(Conf{DatabaseFile: DefaultDBFile, Theme: "Adwaita", Lang: "en"})
	errWin(err)
	err = os.WriteFile(ConfFile, jsonData, os.ModePerm)
	errWin(err)
}
