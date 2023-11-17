/*
 * Copyright Tom5521(c) 2023 - All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package data_test

import (
	"EduTrack/data"
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLoadDatabase(t *testing.T) {
	var DB = &data.DB
	assert := assert.New(t)
	data.LoadFiles()
	var seconds int = 3
	if len(os.Args) != 1 {
		newTime, err := strconv.Atoi(os.Args[1])
		assert.NotNilf(err, "Error converting <%v> to int", os.Args[1])
		seconds = newTime
	}
	for t := seconds; t != 0; t-- {
		time.Sleep(1 * time.Second)
		DB.LoadStudents()
		DB.LoadGrade()
		fmt.Println("Grades:", DB.Grades)
		fmt.Println("Students:", DB.Students)
		fmt.Println("Seconds left:", t)
	}
}

var Config = &data.Config

func TestCreateDatabase(t *testing.T) {
	assert := assert.New(t)
	data.LoadFiles()
	if _, err := os.Stat(Config.DatabaseFile); os.IsNotExist(err) {
		err := data.CreateDatabase()
		if err != nil {
			assert.NotNil(err, "Error creating database!:", err)
		}
	}
	if _, err := os.Stat(Config.DatabaseFile); os.IsNotExist(err) {
		assert.NotNil(err, "Error creating/checking database file!")
	}
}

func TestLoadStudentRecords(t *testing.T) {
	var DB = &data.DB
	assert := assert.New(t)
	if len(DB.Students) == 0 {
		_, err := DB.AddStudent(data.Student{Name: "Angel", DNI: "123", Age: 123, PhoneNumber: "123", ImageFilePath: "123"})
		if err != nil {
			log.Println(err)
			assert.Fail("Error adding temporal student")
		}
	}

	fmt.Println(DB.Students[0])
	fmt.Println(DB.Students[0].Records)
	DB.LoadStudents()
	err := DB.Students[0].LoadRecords()
	if err != nil {
		assert.Fail("Error loading records", err.Error())
	}
	fmt.Println(DB.Students[0])
	fmt.Println(DB.Students[0].Records)
}

func TestLoadAllStudentRecords(t *testing.T) {
	var DB = &data.DB
	//assert := assert.New(t)
	for _, student := range DB.Students {
		student.LoadRecords()
		fmt.Printf("Student Name:%v\nRecords:%v\n", student.Name, student.Records)
	}
	fmt.Println(DB.Students[0].Records)
}
