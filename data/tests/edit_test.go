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

func TestEditGrade(t *testing.T) {
	var Db = &data.Db
	assert := assert.New(t)
	fmt.Println(Db.Grades)
	if len(Db.Grades) == 0 {
		_, err := Db.AddGrade(data.Grade{Name: "Angel", Info: "Test", Price: "100"})
		if err != nil {
			log.Println(err)
			assert.Fail("Error adding temporal grade")
		}
		log.Println(Db.Grades)
	}
	err := Db.EditGrade(Db.Grades[0].ID, data.Grade{Name: "Jonh Doe", Info: "Lorem Ipsum", Price: "100"})
	if err != nil {
		assert.Fail("Error editing grade", err)
	}
	fmt.Println(Db.Grades)
}

func TestEditStudent(t *testing.T) {
	var Db = &data.Db
	assert := assert.New(t)
	fmt.Println(Db.Students)
	err := Db.EditStudent(Db.Students[0].ID, data.Student{Name: "Carlos pajas"})
	assert.Nil(err, "Error modifying student", err)
	fmt.Println(Db.Students)
}

func TestEditRecord(t *testing.T) {
	var Db = &data.Db
	fmt.Println("Starter student len:", len(Db.Students))
	fmt.Println("Starter records len:", len(Db.Students[0].Records))
	assert := assert.New(t)
	if len(Db.Students) == 0 {
		_, err := Db.AddStudent(data.Student{Name: "Angel", DNI: "123", Age: 123, PhoneNumber: "123", ImageFilePath: "123"})
		if err != nil {
			log.Println(err)
			assert.Fail("Error adding temporal student")
		}
	}
	student := &Db.Students[0]
	err := student.LoadRecords()
	if err != nil {
		assert.Fail("Error loading records")
	}
	fmt.Println("records loaded Len:", len(Db.Students[0].Records))
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

func TestEditStudentGrade(t *testing.T) {
	var Db = &data.Db
	assert := assert.New(t)

	if len(Db.Students) == 0 {
		_, err := Db.AddStudent(data.Student{Name: "Angel", DNI: "123", Age: 123, PhoneNumber: "123", ImageFilePath: "123"})
		if err != nil {
			fmt.Println(err)
			assert.Fail("Error adding temporal student")
		}
	}
	student := &Db.Students[0]
	if len(Db.Grades) == 0 {
		_, err := Db.AddGrade(data.Grade{Name: "Angel", Info: "Test", Price: "100"})
		if err != nil {
			log.Println(err)
			assert.Fail("Error adding temporal grade")
		}
	}
	err := student.LoadGrades()
	if err != nil {
		assert.Fail("Error loading grades")
	}
	if len(Db.Grades) == 0 {
		_, err := Db.AddGrade(data.Grade{Name: "Curso2", Info: "Lorem Ipsum", Price: "0"})
		if err != nil {
			fmt.Println(err)
			assert.Fail("Error adding temporal grade")
		}
	}
	Dbgrade := Db.Grades[0]
	if len(student.Grades) == 0 {
		_, err := student.AddGrade(data.StudentGrade{StudentID: student.ID, GradeID: Dbgrade.ID, Start: "1234", End: "12344"})
		if err != nil {
			fmt.Println(err)
			assert.Fail("Error adding temporal grade")
		}
	}
	grade := &student.Grades[0]
	fmt.Println(grade)
	err = grade.Edit(data.StudentGrade{StudentID: student.ID, GradeID: Dbgrade.ID, Start: "Edited for testing!", End: "12344"})
	fmt.Println(grade)
	if err != nil {
		assert.Fail("Error editing StudentGrade!")
	}
}
