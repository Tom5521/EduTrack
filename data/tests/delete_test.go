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
	var db = &data.DB
	assert := assert.New(t)
	require := require.New(t)
	if len(db.Grades) == 0 {
		_, err := db.AddGrade(data.Grade{Name: "Angel", Info: "Test", Price: "100"})
		require.NoError(err)
		log.Println(db.Grades)
	}
	originalLen := len(db.Grades)
	err := db.Grades[0].Delete()
	if err != nil {
		assert.Fail("Error deleting grade", err)
	}
	fmt.Println(db.Grades)
	assert.NotEqual(originalLen, len(db.Grades), "Grades array not modified!")
}

func TestDeleteStudent(t *testing.T) {
	var db = &data.DB
	assert := assert.New(t)
	require := require.New(t)
	if len(db.Students) == 0 {
		_, err := db.AddStudent(data.Student{Name: "Angel", DNI: "123", Age: 123, PhoneNumber: "123", ImageFilePath: "123"})
		require.NoError(err)
	}
	originalLen := len(db.Students)
	err := db.Students[0].Delete()
	require.NoError(err)
	fmt.Println(db.Students)
	assert.NotEqual(originalLen, len(db.Grades))
}

func TestDeleteIn(t *testing.T) {
	var db = &data.DB
	assert := assert.New(t)
	require := require.New(t)
	// Test grades
	if len(db.Grades) == 0 {
		_, err := db.AddGrade(data.Grade{Name: "Angel", Info: "Test", Price: "100"})
		require.NoError(err)
		log.Println(db.Grades)
	}
	// Test Students
	if len(db.Students) == 0 {
		_, err := db.AddStudent(data.Student{Name: "Angel", DNI: "123", Age: 123, PhoneNumber: "123", ImageFilePath: "123"})
		require.NoError(err)
		log.Println(db.Grades)
	}

	originalStudentlen := len(db.Students)
	originalGradelen := len(db.Grades)

	// Exec Deleter
	err := data.Delete(db.Students[0])
	require.NoError(err)
	err = data.Delete(db.Grades[0])
	require.NoError(err)

	fmt.Println(db.Students)
	fmt.Println(db.Grades)

	assert.NotEqual(originalStudentlen, len(db.Students), "Student table not modified!")
	assert.NotEqual(originalGradelen, len(db.Grades), "Grades table not modified!")
}
