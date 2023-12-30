package data

import "gorm.io/gorm"

type Record struct {
	gorm.Model
	StudentID uint
	Name      string
	Date      string
	Info      string
}

type Course struct {
	gorm.Model
	Name  string
	Info  string
	Price string
}

type Student struct {
	gorm.Model
	Name           string
	Age            uint
	DNI            string
	PhoneNumber    string
	ImageFilePath  string
	StudentCourses []StudentCourse `gorm:"-"`
	StudentRecords []Record        `gorm:"-"`
}

type StudentCourse struct {
	gorm.Model
	CourseID  uint
	StudentID uint
	Start     string
	End       string
}

var Students []Student
var Courses []Course
var StudentCourses []StudentCourse
var Records []Record

var DB *gorm.DB
