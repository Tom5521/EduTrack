//go/:build delete
// /+build delete

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
	var Db = &data.Db
	assert := assert.New(t)
	if len(Db.Grades) == 0 {
		_, err := Db.AddGrade(data.Grade{Name: "Angel", Info: "Test", Price: "100"})
		if err != nil {
			log.Println(err)
			assert.Fail("Error adding temporal grade")
		}
		log.Println(Db.Grades)
	}
	originalLen := len(Db.Grades)
	err := Db.Grades[0].Delete()
	if err != nil {
		assert.Fail("Error deleting grade", err)
	}
	fmt.Println(Db.Grades)
	assert.NotEqual(originalLen, len(Db.Grades), "Grades array not modified!")
}

func TestDeleteStudent(t *testing.T) {
	var Db = &data.Db
	assert := assert.New(t)
	if len(Db.Students) == 0 {
		_, err := Db.AddStudent(data.Student{Name: "Angel", DNI: "123", Age: 123, PhoneNumber: "123", ImageFilePath: "123"})
		if err != nil {
			log.Println(err)
			assert.Fail("Error adding temporal student")
		}
		log.Println(Db.Grades)
	}
	originalLen := len(Db.Students)
	err := Db.Students[0].Delete()
	if err != nil {
		assert.Fail("Error deleting student")
	}
	fmt.Println(Db.Students)
	assert.NotEqual(originalLen, len(Db.Grades))
}

func TestDeleteIn(t *testing.T) {
	var Db = &data.Db
	assert := assert.New(t)
	// Test grades
	if len(Db.Grades) == 0 {
		_, err := Db.AddGrade(data.Grade{Name: "Angel", Info: "Test", Price: "100"})
		if err != nil {
			log.Println(err)
			assert.Fail("Error adding temporal grade")
		}
		log.Println(Db.Grades)
	}
	// Test Students
	if len(Db.Students) == 0 {
		_, err := Db.AddStudent(data.Student{Name: "Angel", DNI: "123", Age: 123, PhoneNumber: "123", ImageFilePath: "123"})
		if err != nil {
			log.Println(err)
			assert.Fail("Error adding temporal student")
		}
		log.Println(Db.Grades)
	}

	originalStudentlen := len(Db.Students)
	originalGradelen := len(Db.Grades)

	// Exec Deleter
	data.Delete(Db.Students[0])
	data.Delete(Db.Grades[0])

	fmt.Println(Db.Students)
	fmt.Println(Db.Grades)

	assert.NotEqual(originalStudentlen, len(Db.Students), "Student table not modified!")
	assert.NotEqual(originalGradelen, len(Db.Grades), "Grades table not modified!")

}
