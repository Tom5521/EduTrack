package data

import (
	"errors"
	"fmt"
	"log"
)

type DB_Str struct {
	Students []Student
	Grades   []Grade
}

func (d *DB_Str) Update() error {
	err := d.LoadGrade()
	if err != nil {
		log.Println(err)
		return err
	}
	err = d.LoadStudents()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func InitDB() DB_Str {
	Config = GetConfData()
	db := DB_Str{}
	db.LoadGrade()
	db.LoadStudents()
	return db
}

func (d DB_Str) FindStudentByID(id int) (Student, error) {
	for _, student := range d.Students {
		if student.ID == id {
			return student, nil
		}
	}
	return Student{}, errors.New(fmt.Sprintf("Can't find student by id <%v>", id))
}
