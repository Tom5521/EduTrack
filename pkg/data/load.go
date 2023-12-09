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
	printErr(LoadCourses())
	printErr(LoadStudents())
	printErr(LoadRecords())
	printErr(LoadStudentCourses())
}

func LoadCourses() error {
	return DB.Find(&Courses).Error
}
func LoadStudents() error {
	return DB.Find(&Students).Error
}
func LoadRecords() error {
	return DB.Find(&Records).Error
}
func LoadStudentCourses() error {
	return DB.Find(&StudentCourses).Error
}

// Student

func (s *Student) GetCourses() {
	var cleanGrades []StudentCourse
	s.Courses = cleanGrades
	for _, grade := range StudentCourses {
		if grade.StudentID == s.ID {
			s.Courses = append(s.Courses, grade)
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
