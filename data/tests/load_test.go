package data_test

import (
	"EduTrack/data"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_LoadDatabase(t *testing.T) {
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

func Test_CreateDatabase(t *testing.T) {
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
