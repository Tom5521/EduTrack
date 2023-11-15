package data_test

import (
	"EduTrack/data"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_EditGrade(t *testing.T) {
	assert := assert.New(t)
	fmt.Println(DB.Grades)
	err := DB.EditGrade(11, data.Grade{Name: "Jonh Doe", Info: "Lorem Ipsum", Price: "100"})
	if err != nil {
		assert.Fail("Error editing grade", err)
	}
	fmt.Println(DB.Grades)
}
