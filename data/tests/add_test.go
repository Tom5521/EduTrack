/*
 * Copyright (c) 2023 Tom5521- All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package data_test

import (
	"EduTrack/data"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var DB = &data.DB

func TestAddGrade(t *testing.T) {
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

func TestAddStudent(t *testing.T) {
	assert := assert.New(t)
	fmt.Println(DB.Students)
	id, err := DB.AddStudent(data.Student{Name: "T", Age: 12, DNI: "123", PhoneNumber: "", ImageFilePath: ""})
	assert.Nil(err, "Error adding student", err, id)
	fmt.Println(DB.Students)

}
