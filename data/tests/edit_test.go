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
	var db = &data.DB
	// assert := assert.New(t)
	fmt.Println(db.Grades)
	if len(db.Grades) == 0 {
		_, err := db.AddGrade(data.Grade{Name: "Angel", Info: "Test", Price: "100"})
		require.NoError(t, err)
		log.Println(db.Grades)
	}
	err := db.EditGrade(db.Grades[0].ID, data.Grade{Name: "Jonh Doe", Info: "Lorem Ipsum", Price: "100"})
	require.NoError(t, err)
	fmt.Println(db.Grades)
}

func TestEditStudent(t *testing.T) {
	var db = &data.DB
	// assert := assert.New(t)
	fmt.Println(db.Students)
	err := db.EditStudent(db.Students[0].ID, data.Student{Name: "Carlos pajas"})
	require.NoError(t, err)
	fmt.Println(db.Students)
}

func TestEditRecord(t *testing.T) {
	var db = &data.DB
	fmt.Println("Starter student len:", len(db.Students))
	fmt.Println("Starter records len:", len(db.Students[0].Records))
	// assert := assert.New(t)
	if len(db.Students) == 0 {
		_, err := db.AddStudent(data.Student{Name: "Angel", DNI: "123", Age: 123, PhoneNumber: "123", ImageFilePath: "123"})
		require.NoError(t, err)
	}
	student := &db.Students[0]
	err := student.LoadRecords()
	require.NoError(t, err)
	fmt.Println("records loaded Len:", len(db.Students[0].Records))
	tmpRecord := data.Record{Name: "Testt", Info: "Lorem ipsum", Date: "777", StudentID: student.ID}
	if len(student.Records) == 0 {
		_, err = student.AddRecord(tmpRecord)
		require.NoError(t, err)
	}
	tmpRecord.Info = "Edited for testing!"
	fmt.Println(student.Records[0])
	err = student.Records[0].Edit(tmpRecord)
	require.NoError(t, err)
	err = student.LoadRecords()
	require.NoError(t, err)
	fmt.Println(student.Records[0])
}

func TestEditStudentGrade(t *testing.T) {
	var db = &data.DB
	// assert := assert.New(t)
	require := require.New(t)
	if len(db.Students) == 0 {
		_, err := db.AddStudent(data.Student{Name: "Angel", DNI: "123", Age: 123, PhoneNumber: "123", ImageFilePath: "123"})
		require.NoError(err)
	}
	student := &db.Students[0]
	if len(db.Grades) == 0 {
		_, err := db.AddGrade(data.Grade{Name: "Angel", Info: "Test", Price: "100"})
		require.NoError(err)
	}
	err := student.LoadGrades()
	require.NoError(err)
	if len(db.Grades) == 0 {
		_, err = db.AddGrade(data.Grade{Name: "Curso2", Info: "Lorem Ipsum", Price: "0"})
		require.NoError(err)
	}
	dbgrade := db.Grades[0]
	if len(student.Grades) == 0 {
		_, err = student.AddGrade(data.StudentGrade{StudentID: student.ID, GradeID: dbgrade.ID, Start: "1234", End: "12344"})
		require.NoError(err)
	}
	grade := &student.Grades[0]
	fmt.Println(grade)
	err = grade.Edit(data.StudentGrade{
		StudentID: student.ID,
		GradeID:   dbgrade.ID,
		Start:     "Edited for testing!",
		End:       "12344"})
	fmt.Println(grade)
	require.NoError(err)
}
