/*
 * Copyright Tom5521(c) - All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package data_test

import (
	"EduTrack/data"
	"errors"
	"os"
	"testing"
)

var Data = &data.Data

func Test_LoadData(t *testing.T) {
	data.LoadFiles()
	if data.Config.StudentsFile == "" {
		t.Fatal()
	}
}

func Test_GetData(t *testing.T) {
	t.Log(data.Config)
	t.Log("GetGradesNamesP:", Data.GetGradesNames())
	t.Log("GetStudentIDs:", Data.GetStudentIDs())
	t.Log("GetStudentNames:", Data.GetStudentNames())
}

func Test_CreateConfFile(t *testing.T) {
	data.NewConfigurationFile()
	if _, err := os.Stat(data.Config.StudentsFile); os.IsNotExist(err) {
		t.Fail()
	}
}

func Test_Notify(t *testing.T) {
	data.NotifyError("ERROR TEST", errors.New("TEST ERR VAR"))
}
