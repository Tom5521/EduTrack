package data

import "gorm.io/gorm"

type Record struct {
	gorm.Model
	StudentID uint
	Name      string
	Date      string
	Info      string
}

type Grade struct {
	gorm.Model
	Name  string
	Info  string
	Price string
}

type Student struct {
	gorm.Model
	Name          string
	Age           uint
	DNI           string
	PhoneNumber   string
	ImageFilePath string
	Grades        []StudentGrade `gorm:"-"`
	Records       []Record       `gorm:"-"`
}

type StudentGrade struct {
	gorm.Model
	GradeID   uint
	StudentID uint
	Start     string
	End       string
}

var Students []Student
var Grades []Grade
var StudentGrades []StudentGrade
var Records []Record

var DB *gorm.DB
