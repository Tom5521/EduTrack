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

	"github.com/stretchr/testify/require"
)

func TestEditGrade(t *testing.T) {
	data.LoadFiles()
	// assert := assert.New(t)
	fmt.Println(data.Grades)
	if len(data.Grades) == 0 {
		err := data.AddGrade(data.Grade{Name: "Angel", Info: "Test", Price: "100"})
		require.NoError(t, err)
		log.Println(data.Grades)
	}
	err := data.EditGrade(data.Grades[0].ID, data.Grade{Name: "Jonh Doe", Info: "Lorem Ipsum", Price: "100"})
	require.NoError(t, err)
	fmt.Println(data.Grades)
}

func TestEditStudent(t *testing.T) {
	// assert := assert.New(t)
	fmt.Println(data.Students)
	err := data.EditStudent(data.Students[0].ID, data.Student{Name: "Carlos pajas"})
	require.NoError(t, err)
	fmt.Println(data.Students)
}

func TestEditRecord(t *testing.T) {
	fmt.Println("Starter student len:", len(data.Students))
	fmt.Println("Starter records len:", len(data.Records))
	// assert := assert.New(t)
	if len(data.Students) == 0 {
		err := data.AddStudent(data.Student{Name: "Angel", DNI: "123", Age: 123, PhoneNumber: "123", ImageFilePath: "123"})
		require.NoError(t, err)
	}
	student := &data.Students[0]
	fmt.Println("records loaded Len:", len(data.Records))
	tmpRecord := data.Record{Name: "Testt", Info: "Lorem ipsum", Date: "777", StudentID: student.ID}
	if len(data.Records) == 0 {
		err := student.AddRecord(tmpRecord)
		require.NoError(t, err)
	}
	tmpRecord.Info = "Edited for testing!"
	fmt.Println(data.Records[0])
	err := data.Records[0].Edit(tmpRecord)
	require.NoError(t, err)
	data.LoadRecords()
	require.NoError(t, err)
	fmt.Println(data.Records[0])
}

func TestEditStudentGrade(t *testing.T) {
	// assert := assert.New(t)
	require := require.New(t)
	if len(data.Students) == 0 {
		err := data.AddStudent(data.Student{Name: "Angel", DNI: "123", Age: 123, PhoneNumber: "123", ImageFilePath: "123"})
		require.NoError(err)
	}
	student := &data.Students[0]
	if len(data.Grades) == 0 {
		err := data.AddGrade(data.Grade{Name: "Angel", Info: "Test", Price: "100"})
		require.NoError(err)
	}
	err := data.LoadGrades()
	require.NoError(err)
	if len(data.Grades) == 0 {
		err = data.AddGrade(data.Grade{Name: "Curso2", Info: "Lorem Ipsum", Price: "0"})
		require.NoError(err)
	}
	dbgrade := data.Grades[0]
	if len(data.Grades) == 0 {
		err = student.AddGrade(data.StudentGrade{StudentID: student.ID, GradeID: dbgrade.ID, Start: "1234", End: "12344"})
		require.NoError(err)
	}
	grade := &data.Grades[0]
	fmt.Println(grade)
	err = grade.Edit(data.StudentGrade{
		StudentID: student.ID,
		GradeID:   dbgrade.ID,
		Start:     "Edited for testing!",
		End:       "12344"})
	fmt.Println(grade)
	require.NoError(err)
}
