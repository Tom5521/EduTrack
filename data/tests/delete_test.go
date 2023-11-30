//asdgo:build delete
// dd+dbuild delete

/*
 * Copyright (c) 2023 Tom5521- All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package data_test

import (
	"EduTrack/data"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeleteGrade(t *testing.T) {
	data.LoadFiles()
	assert := assert.New(t)
	require := require.New(t)
	if len(data.Grades) == 0 {
		err := data.AddGrade(data.Grade{Name: "Angel", Info: "Test", Price: "100"})
		require.NoError(err)
		log.Println(data.Grades)
	}
	originalLen := len(data.Grades)
	err := data.Grades[0].Delete()
	if err != nil {
		assert.Fail("Error deleting grade", err)
	}
	fmt.Println(data.Grades)
	assert.NotEqual(originalLen, len(data.Grades), "Grades array not modified!")
}

func TestDeleteStudent(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)
	if len(data.Students) == 0 {
		err := data.AddStudent(data.Student{Name: "Angel", DNI: "123", Age: 123, PhoneNumber: "123", ImageFilePath: "123"})
		require.NoError(err)
	}
	originalLen := len(data.Students)
	err := data.Students[0].Delete()
	require.NoError(err)
	fmt.Println(data.Students)
	assert.NotEqual(originalLen, len(data.Grades))
}

func TestDeleteIn(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)
	// Test grades
	if len(data.Grades) == 0 {
		err := data.AddGrade(data.Grade{Name: "Angel", Info: "Test", Price: "100"})
		require.NoError(err)
		log.Println(data.Grades)
	}
	// Test Students
	if len(data.Students) == 0 {
		err := data.AddStudent(data.Student{Name: "Angel", DNI: "123", Age: 123, PhoneNumber: "123", ImageFilePath: "123"})
		require.NoError(err)
		log.Println(data.Grades)
	}

	originalStudentlen := len(data.Students)
	originalGradelen := len(data.Grades)

	// Exec Deleter
	err := data.Delete(data.Students[0])
	require.NoError(err)
	err = data.Delete(data.Grades[0])
	require.NoError(err)

	fmt.Println(data.Students)
	fmt.Println(data.Grades)

	assert.NotEqual(originalStudentlen, len(data.Students), "Student table not modified!")
	assert.NotEqual(originalGradelen, len(data.Grades), "Grades table not modified!")
}
