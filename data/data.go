/*
 * Copyright Tom5521(c) - All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package data

import (
	"github.com/ncruces/zenity"
)

type Data_str struct {
	Grades   []*Grade
	Students []Student
}

var Data = Data_str{}

var Config = Config_str{}

var _, ConfigFile, _ string = getOSConfFile()

func NotifyError(text string, err error) {
	zenity.Notify(text + "::" + err.Error())
}
