package data_test

import (
	"EduTrack/data"
	"os"
	"testing"
)

func Test_LoadData(t *testing.T) {
	data.LoadFiles()
	if data.Config.StudentsFile == "" {
		t.Fatal()
	}
}

func Test_GetData(t *testing.T) {
	t.Log(data.Config)
	t.Log(data.GetGradesNames())
	t.Log(data.GetStudentIDs())
	t.Log(data.GetStudentNames())
}

func Test_CreateConfFile(t *testing.T) {
	data.NewConfigurationFile()
	if _, err := os.Stat(data.Config.StudentsFile); os.IsNotExist(err) {
		t.Fail()
	}
}
