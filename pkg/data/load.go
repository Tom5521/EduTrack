package data

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

/*func catchErr(err error, run func(error)) {
	if err != nil {
		fmt.Println(err)
		run(err)
	}
}*/

func printErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func GetDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(Config.DatabaseFile), &gorm.Config{})
	printErr(err)
	return db
}

func LoadEverything() {
	printErr(LoadGrades())
	printErr(LoadStudents())
	printErr(LoadRecords())
	printErr(LoadStudentGrades())
}

func LoadGrades() error {
	return DB.Find(&Grades).Error
}
func LoadStudents() error {
	return DB.Find(&Students).Error
}
func LoadRecords() error {
	return DB.Find(&Records).Error
}
func LoadStudentGrades() error {
	return DB.Find(&StudentGrades).Error
}

// Student

func (s *Student) GetGrades() {
	var cleanGrades []StudentGrade
	s.Grades = cleanGrades
	for _, grade := range StudentGrades {
		if grade.StudentID == s.ID {
			s.Grades = append(s.Grades, grade)
		}
	}
}

func (s *Student) GetRecords() {
	var cleanRecords []Record
	s.Records = cleanRecords
	for _, record := range Records {
		if record.StudentID == s.ID {
			s.Records = append(s.Records, record)
		}
	}
}
