/*
 * Copyright (c) 2023 Tom5521- All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package data_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/Tom5521/EduTrack/pkg/data"
	"github.com/Tom5521/EduTrack/tests/data/fillers"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var lorepIpsum = `Lorem ipsum dolor sit amet,
	officia excepteur ex fugiat reprehenderit enim labore culpa sint ad nisi Lorem pariatur mollit ex 
	esse exercitation amet. Nisi anim cupidatat excepteur officia. 
	Reprehenderit nostrud nostrud ipsum Lorem est aliquip amet voluptate 
	voluptate dolor minim nulla est proident. Nostrud officia pariatur ut officia.
	Sit irure elit esse ea nulla sunt ex occaecat reprehenderit commodo officia dolor Lorem duis laboris cupidatat 
	officia voluptate. Culpa proident adipisicing id nulla nisi laboris ex in Lorem sunt duis officia eiusmod.
	Aliqua reprehenderit commodo ex non excepteur duis sunt velit enim. Voluptate laboris sint cupidatat ullamco 
	ut ea consectetur et est culpa et culpa duis.`

func TestAddCourse(t *testing.T) {
	data.LoadFiles()
	require := require.New(t)
	fillers.Course()
	assert := assert.New(t)
	originalLen := len(data.Courses)
	newCourse := &data.Course{
		Name:  gofakeit.Name(),
		Info:  gofakeit.Phrase(),
		Price: strconv.Itoa(int(gofakeit.Price(10, 200))),
	}
	err := data.AddCourse(newCourse)
	require.NoError(err)
	fmt.Println("Courses:", data.Courses)

	assert.NotEqual(originalLen, len(data.Courses), "Course array not modified!")
}

func TestAddStudent(t *testing.T) {
	// assert := assert.New(t)
	require := require.New(t)
	fillers.Student()
	fmt.Println(data.Students)
	err := data.AddStudent(&data.Student{Name: "T", Age: 12, DNI: "123", PhoneNumber: "", ImageFilePath: ""})
	require.NoErrorf(err, "Error adding student || %v ||studentID:%v", err)
	fmt.Println(data.Students)
}

func TestAddRecord(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)
	fillers.Record()
	if len(data.Students) == 0 {
		err := data.AddStudent(&data.Student{Name: "T", Age: 12, DNI: "123", PhoneNumber: "", ImageFilePath: ""})
		require.NoError(err)
	}

	var student = &data.Students[0]
	orgLen := len(data.Records)
	fmt.Println(data.Records)
	err := student.AddRecord(&data.Record{Name: "TEST", Date: "777", Info: lorepIpsum})
	require.NoErrorf(err, "Error adding a record to %v | Error: %v", student.Name, err)

	assert.NotEqual(len(data.Records), orgLen, "Records table not modified!")
	fmt.Println(data.Records)
}

func TestAddStudentCourse(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)
	fillers.StudentCourse()
	if len(data.Students) == 0 {
		err := data.AddStudent(&data.Student{Name: "T", Age: 12, DNI: "123", PhoneNumber: "", ImageFilePath: ""})
		require.NoError(err)
	}
	if len(data.Courses) == 0 {
		err := data.AddCourse(&data.Course{Name: "Curso2", Info: "Lorem Ipsum", Price: "0"})
		require.NoError(err)
	}
	student := &data.Students[0]
	grade := &data.Courses[0]
	fmt.Println(data.Courses)
	orgLen := len(data.Courses)
	err := student.AddCourse(&data.StudentCourse{StudentID: student.ID, CourseID: grade.ID, Start: "1234", End: "12344"})
	fmt.Println(data.Courses)
	require.NoError(err)
	assert.Equal(orgLen, len(data.Courses))
}
