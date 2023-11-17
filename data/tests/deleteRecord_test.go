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
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteRecords(t *testing.T) {
	var DB = &data.DB
	assert := assert.New(t)
	if len(DB.Students) == 0 {
		_, err := DB.AddStudent(data.Student{Name: "Angel", DNI: "123", Age: 123, PhoneNumber: "123", ImageFilePath: "123"})
		if err != nil {
			assert.Fail("Error adding temporal student")
		}
	}

	student := &DB.Students[0]
	err := student.LoadRecords()
	if err != nil {
		assert.Fail("Error loading student records")
	}
	if len(student.Records) == 0 {
		_, err := student.AddRecord(data.Record{Name: "Testt", Info: "Lorem ipsum", Date: "777"})
		if err != nil {
			assert.Fail("Error adding new record for test", err)
		}
	}
	log.Println(student.Records)
	student.Records[0].Delete()
	log.Println(student.Records)

}
