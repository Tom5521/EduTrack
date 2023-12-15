/*
 * Copyright Tom5521(c) 2023 - All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package data_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/Tom5521/EduTrack/pkg/conf"
	"github.com/Tom5521/EduTrack/pkg/data"

	"github.com/stretchr/testify/require"
)

func TestLoadDatabase(t *testing.T) {
	data.LoadFiles()
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
		err := data.LoadStudents()
		require.NoError(err)
		err = data.LoadCourses()
		require.NoError(err)
		fmt.Println("Grades:", data.Courses)
		fmt.Println("Students:", data.Students)
		fmt.Println("Seconds left:", ts)
	}
}

var Config = &conf.Config

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
	// assert := assert.New(t)
	require := require.New(t)
	if len(data.Students) == 0 {
		err := data.AddStudent(&data.Student{Name: "Angel", DNI: "123", Age: 123, PhoneNumber: "123", ImageFilePath: "123"})
		require.NoError(err)
	}

	fmt.Println(data.Students[0])
	fmt.Println(data.Students[0].Records)
	err := data.LoadStudents()
	require.NoError(err)
	fmt.Println(data.Students[0])
	fmt.Println(data.Students[0].Records)
}
