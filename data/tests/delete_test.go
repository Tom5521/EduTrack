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

func TestDeleteGrade(t *testing.T) {
	assert := assert.New(t)
	originalLen := len(DB.Grades)
	err := DB.Grades[0].Delete()
	if err != nil {
		assert.Fail("Error deleting grade", err)
	}
	fmt.Println(DB.Grades)
	assert.NotEqual(originalLen, len(DB.Grades), "Grades array not modified!")
}

func TestDeleteStudent(t *testing.T) {
	assert := assert.New(t)
	originalLen := len(DB.Students)
	err := DB.Students[0].Delete()
	if err != nil {
		assert.Fail("Error deleting student")
	}
	fmt.Println(DB.Students)
	assert.NotEqual(originalLen, len(DB.Grades))
}

func TestDeleteInterface(t *testing.T) {
	assert := assert.New(t)

	originalStudentlen := len(DB.Students)
	originalGradelen := len(DB.Grades)

	data.Deletes(DB.Students[0])
	data.Deletes(DB.Grades[0])

	fmt.Println(DB.Students)
	fmt.Println(DB.Grades)

	assert.NotEqual(originalStudentlen, len(DB.Students), "Student table not modified!")
	assert.NotEqual(originalGradelen, len(DB.Grades), "Grades table not modified!")

}
