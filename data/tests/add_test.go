package data_test

import (
	"EduTrack/data"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var DB = &data.DB

func Test_AddGrade(t *testing.T) {
	assert := assert.New(t)
	originalLen := len(DB.Grades)
	id, err := DB.AddGrade(data.Grade{Name: "Curso2", Info: "Lorem Ipsum", Price: "0"})
	if err != nil {
		assert.Fail("Error adding grade:", err)
	}
	if id == -1 {
		assert.Fail("Error getting new grade id!")
	}
	fmt.Println("Grades:", DB.Grades)

	assert.NotEqual(originalLen, len(DB.Grades), "Grades array not modified!")
}

func Test_AddStudent(t *testing.T) {

}
