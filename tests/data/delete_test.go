//go:build delete
// +build delete

/*
 * Copyright (c) 2023 Tom5521- All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package data_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/Tom5521/EduTrack/pkg/data"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeleteGrade(t *testing.T) {
	data.LoadFiles()
	assert := assert.New(t)
	require := require.New(t)
	if len(data.Courses) == 0 {
		err := data.AddCourse(&data.Course{Name: "Angel", Info: "Test", Price: "100"})
		require.NoError(err)
		log.Println(data.Courses)
	}
	originalLen := len(data.Courses)
	err := data.Courses[0].Delete()
	if err != nil {
		assert.Fail("Error deleting grade", err)
	}
	fmt.Println(data.Courses)
	assert.NotEqual(originalLen, len(data.Courses), "Grades array not modified!")
}

func TestDeleteStudent(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)
	if len(data.Students) == 0 {
		err := data.AddStudent(&data.Student{Name: "Angel", DNI: "123", Age: 123, PhoneNumber: "123", ImageFilePath: "123"})
		require.NoError(err)
	}
	originalLen := len(data.Students)
	err := data.Students[0].Delete()
	require.NoError(err)
	fmt.Println(data.Students)
	assert.NotEqual(originalLen, len(data.Courses))
}

func TestDeleteIn(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)
	// Test grades
	if len(data.Courses) == 0 {
		err := data.AddCourse(&data.Course{Name: "Angel", Info: "Test", Price: "100"})
		require.NoError(err)
		log.Println(data.Courses)
	}
	// Test Students
	if len(data.Students) == 0 {
		err := data.AddStudent(&data.Student{Name: "Angel", DNI: "123", Age: 123, PhoneNumber: "123", ImageFilePath: "123"})
		require.NoError(err)
		log.Println(data.Courses)
	}

	originalStudentlen := len(data.Students)
	originalGradelen := len(data.Courses)

	// Exec Deleter
	err := data.Delete(data.Students[0])
	require.NoError(err)
	err = data.Delete(data.Courses[0])
	require.NoError(err)

	fmt.Println(data.Students)
	fmt.Println(data.Courses)

	assert.NotEqual(originalStudentlen, len(data.Students), "Student table not modified!")
	assert.NotEqual(originalGradelen, len(data.Courses), "Grades table not modified!")
}
