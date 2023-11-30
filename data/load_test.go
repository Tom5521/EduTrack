package data_test

import (
	"EduTrack/data"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	data.LoadFiles()
	data.LoadEverything()
	err := data.AddGrade(data.Grade{Name: "org", Info: "org2", Price: "132"})
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	err = data.EditStruct(data.Grades[0], data.Grade{Name: "Edited for testing!", Info: "TESTT", Price: "999+"})
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}
