package data

import (
	"fmt"

	"github.com/Tom5521/EduTrack/pkg/conf"
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

func OpenDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(conf.Config.DatabaseFile), &gorm.Config{})
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

func (s *Student) Courses() {
	var cleanGrades []StudentCourse
	s.StudentCourses = cleanGrades
	for _, grade := range StudentCourses {
		if grade.StudentID == s.ID {
			s.StudentCourses = append(s.StudentCourses, grade)
		}
	}
}

func (s *Student) Records() {
	var cleanRecords []Record
	s.StudentRecords = cleanRecords
	for _, record := range Records {
		if record.StudentID == s.ID {
			s.StudentRecords = append(s.StudentRecords, record)
		}
	}
}
