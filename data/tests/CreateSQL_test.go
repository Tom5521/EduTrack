package data_test

import (
	"EduTrack/data"
	"fmt"
	"testing"
)

func TestCreateSQL(t *testing.T) {
	err := data.CreateDatabase()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}
