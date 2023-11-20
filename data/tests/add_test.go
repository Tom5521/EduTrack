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

var lorepIpsum = `Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim labore culpa sint ad nisi Lorem pariatur mollit ex esse exercitation amet. Nisi anim cupidatat excepteur officia. Reprehenderit nostrud nostrud ipsum Lorem est aliquip amet voluptate voluptate dolor minim nulla est proident. Nostrud officia pariatur ut officia. Sit irure elit esse ea nulla sunt ex occaecat reprehenderit commodo officia dolor Lorem duis laboris cupidatat officia voluptate. Culpa proident adipisicing id nulla nisi laboris ex in Lorem sunt duis officia eiusmod. Aliqua reprehenderit commodo ex non excepteur duis sunt velit enim. Voluptate laboris sint cupidatat ullamco ut ea consectetur et est culpa et culpa duis.`

var Db = &data.Db

func TestAddGrade(t *testing.T) {
	assert := assert.New(t)
	originalLen := len(Db.Grades)
	id, err := Db.AddGrade(data.Grade{Name: "Curso2", Info: "Lorem Ipsum", Price: "0"})
	if err != nil {
		assert.Fail("Error adding grade:", err)
	}
	if id == -1 {
		assert.Fail("Error getting new grade id!")
	}
	fmt.Println("Grades:", Db.Grades)

	assert.NotEqual(originalLen, len(Db.Grades), "Grades array not modified!")
}

func TestAddStudent(t *testing.T) {
	assert := assert.New(t)
	fmt.Println(Db.Students)
	id, err := Db.AddStudent(data.Student{Name: "T", Age: 12, DNI: "123", PhoneNumber: "", ImageFilePath: ""})
	assert.Nil(err, "Error adding student", err, id)
	fmt.Println(Db.Students)
}

func TestAddRecord(t *testing.T) {
	assert := assert.New(t)

	if len(Db.Students) == 0 {
		_, err := Db.AddStudent(data.Student{Name: "T", Age: 12, DNI: "123", PhoneNumber: "", ImageFilePath: ""})
		if err != nil {
			fmt.Println(err)
			assert.Fail("Error adding tmp user")
		}
	}

	var student = &Db.Students[0]
	orgLen := len(student.Records)
	fmt.Println(student.Records)
	_, err := student.AddRecord(data.Record{Name: "TEST", Date: "777", Info: lorepIpsum})
	if err != nil {
		assert.Fail("Error adding a record to", student.Name, "|Error:", err)
	}
	assert.NotEqual(len(student.Records), orgLen, "Records table not modified!")
	fmt.Println(student.Records)

}

func TestAddStudentGrade(t *testing.T) {
	Db := &data.Db
	assert := assert.New(t)

	if len(Db.Students) == 0 {
		_, err := Db.AddStudent(data.Student{Name: "T", Age: 12, DNI: "123", PhoneNumber: "", ImageFilePath: ""})
		if err != nil {
			fmt.Println(err)
			assert.Fail("Error adding tmp user")
		}
	}
	if len(Db.Grades) == 0 {
		_, err := Db.AddGrade(data.Grade{Name: "Curso2", Info: "Lorem Ipsum", Price: "0"})
		if err != nil {
			fmt.Println(err)
			assert.Fail("Error adding temporal grade")
		}
	}
	student := &Db.Students[0]
	grade := &Db.Grades[0]
	fmt.Println(student.Grades)
	orgLen := len(student.Grades)
	_, err := student.AddGrade(data.StudentGrade{StudentID: student.ID, GradeID: grade.ID, Start: "1234", End: "12344"})
	fmt.Println(student.Grades)
	if err != nil {
		fmt.Println(err)
		assert.Fail("Error adding new student grade")
	}
	assert.NotEqual(orgLen, len(student.Grades))

}
