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

func TestEditGrade(t *testing.T) {
	assert := assert.New(t)
	fmt.Println(DB.Grades)
	err := DB.EditGrade(DB.Grades[0].ID, data.Grade{Name: "Jonh Doe", Info: "Lorem Ipsum", Price: "100"})
	if err != nil {
		assert.Fail("Error editing grade", err)
	}
	fmt.Println(DB.Grades)
}

func TestEditStudent(t *testing.T) {
	assert := assert.New(t)
	fmt.Println(DB.Students)
	err := DB.EditStudent(DB.Students[0].ID, data.Student{Name: "Carlos pajas"})
	assert.Nil(err, "Error modifying student", err)
	fmt.Println(DB.Students)
}
