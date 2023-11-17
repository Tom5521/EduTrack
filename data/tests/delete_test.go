//go:build delete
// +build delete

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
)

func TestDeleteGrade(t *testing.T) {
	var DB = &data.DB
	assert := assert.New(t)
	if len(DB.Grades) == 0 {
		_, err := DB.AddGrade(data.Grade{Name: "Angel", Info: "Test", Price: "100"})
		if err != nil {
			log.Println(err)
			assert.Fail("Error adding temporal grade")
		}
		log.Println(DB.Grades)
	}
	originalLen := len(DB.Grades)
	err := DB.Grades[0].Delete()
	if err != nil {
		assert.Fail("Error deleting grade", err)
	}
	fmt.Println(DB.Grades)
	assert.NotEqual(originalLen, len(DB.Grades), "Grades array not modified!")
}

func TestDeleteStudent(t *testing.T) {
	var DB = &data.DB
	assert := assert.New(t)
	if len(DB.Students) == 0 {
		_, err := DB.AddStudent(data.Student{Name: "Angel", DNI: "123", Age: 123, PhoneNumber: "123", ImageFilePath: "123"})
		if err != nil {
			log.Println(err)
			assert.Fail("Error adding temporal student")
		}
		log.Println(DB.Grades)
	}
	originalLen := len(DB.Students)
	err := DB.Students[0].Delete()
	if err != nil {
		assert.Fail("Error deleting student")
	}
	fmt.Println(DB.Students)
	assert.NotEqual(originalLen, len(DB.Grades))
}

func TestDeleteIn(t *testing.T) {
	var DB = &data.DB
	assert := assert.New(t)
	// Test grades
	if len(DB.Grades) == 0 {
		_, err := DB.AddGrade(data.Grade{Name: "Angel", Info: "Test", Price: "100"})
		if err != nil {
			log.Println(err)
			assert.Fail("Error adding temporal grade")
		}
		log.Println(DB.Grades)
	}
	// Test Students
	if len(DB.Students) == 0 {
		_, err := DB.AddStudent(data.Student{Name: "Angel", DNI: "123", Age: 123, PhoneNumber: "123", ImageFilePath: "123"})
		if err != nil {
			log.Println(err)
			assert.Fail("Error adding temporal student")
		}
		log.Println(DB.Grades)
	}

	originalStudentlen := len(DB.Students)
	originalGradelen := len(DB.Grades)

	// Exec Deleter
	data.Delete(DB.Students[0])
	data.Delete(DB.Grades[0])

	fmt.Println(DB.Students)
	fmt.Println(DB.Grades)

	assert.NotEqual(originalStudentlen, len(DB.Students), "Student table not modified!")
	assert.NotEqual(originalGradelen, len(DB.Grades), "Grades table not modified!")

}
