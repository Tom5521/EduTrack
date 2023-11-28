/*
 * Copyright Tom5521(c) 2023 - All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package data_test

import (
	"EduTrack/data"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestLoadDatabase(t *testing.T) {
	data.LoadFiles()
	var db = &data.DB
	require := require.New(t)
	// assert := assert.New(t)
	data.LoadFiles()
	var seconds = 3
	/*if len(os.Args) != 1 {
		newTime, err := strconv.Atoi(os.Args[1])
		require.NoErrorf(err, "Error converting <%v> to int", os.Args[1])
		seconds = newTime
	}*/
	for ts := seconds; ts != 0; ts-- {
		time.Sleep(1 * time.Second)
		err := db.LoadStudents()
		require.NoError(err)
		err = db.LoadGrade()
		require.NoError(err)
		fmt.Println("Grades:", db.Grades)
		fmt.Println("Students:", db.Students)
		fmt.Println("Seconds left:", ts)
	}
}

var Config = &data.Config

func TestCreateDatabase(t *testing.T) {
	// assert := assert.New(t)
	require := require.New(t)
	data.LoadFiles()
	if _, err := os.Stat(Config.DatabaseFile); os.IsNotExist(err) {
		err = data.CreateDatabase()

		require.NoError(err, "Error creating database!:")
	}
	if _, err := os.Stat(Config.DatabaseFile); os.IsNotExist(err) {
		require.NoError(err, "Error creating/checking database file!")
	}
}

func TestLoadStudentRecords(t *testing.T) {
	var db = &data.DB
	// assert := assert.New(t)
	require := require.New(t)
	if len(db.Students) == 0 {
		_, err := db.AddStudent(data.Student{Name: "Angel", DNI: "123", Age: 123, PhoneNumber: "123", ImageFilePath: "123"})
		require.NoError(err)
	}

	fmt.Println(db.Students[0])
	fmt.Println(db.Students[0].Records)
	err := db.LoadStudents()
	require.NoError(err)
	err = db.Students[0].LoadRecords()
	require.NoError(err)
	fmt.Println(db.Students[0])
	fmt.Println(db.Students[0].Records)
}

func TestLoadAllStudentRecords(t *testing.T) {
	var db = &data.DB
	// assert := assert.New(t)
	for _, student := range db.Students {
		err := student.LoadRecords()
		require.NoError(t, err)
		fmt.Printf("Student Name:%v\nRecords:%v\n", student.Name, student.Records)
	}
	fmt.Println(db.Students[0].Records)
}
