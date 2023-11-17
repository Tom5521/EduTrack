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

var DB = &data.DB

func TestEditGrade(t *testing.T) {
	assert := assert.New(t)
	fmt.Println(DB.Grades)
	if len(DB.Grades) == 0 {
		_, err := DB.AddGrade(data.Grade{Name: "Angel", Info: "Test", Price: "100"})
		if err != nil {
			log.Println(err)
			assert.Fail("Error adding temporal grade")
		}
		log.Println(DB.Grades)
	}
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

func TestEditRecord(t *testing.T) {
	fmt.Println("Starter student len:", len(DB.Students))
	fmt.Println("Starter records len:", len(DB.Students[0].Records))
	assert := assert.New(t)
	if len(DB.Students) == 0 {
		_, err := DB.AddStudent(data.Student{Name: "Angel", DNI: "123", Age: 123, PhoneNumber: "123", ImageFilePath: "123"})
		if err != nil {
			log.Println(err)
			assert.Fail("Error adding temporal student")
		}
	}
	student := &DB.Students[0]
	err := student.LoadRecords()
	if err != nil {
		assert.Fail("Error loading records")
	}
	fmt.Println("records loaded Len:", len(DB.Students[0].Records))
	tmpRecord := data.Record{Name: "Testt", Info: "Lorem ipsum", Date: "777", StudentId: student.ID}
	if len(student.Records) == 0 {
		_, err := student.AddRecord(tmpRecord)
		if err != nil {
			assert.Fail("Error adding new record for test", err)
		}
	}
	tmpRecord.Info = "Edited for testing!"
	fmt.Println(student.Records[0])
	err = student.Records[0].Edit(tmpRecord)
	student.LoadRecords()
	fmt.Println(student.Records[0])
}
