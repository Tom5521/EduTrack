package data

import (
	"errors"
	"fmt"
	"log"
)

type DbStr struct {
	Students []Student
	Grades   []Grade
}

func (d *DbStr) Update() error {
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

func InitDB() DbStr {
	Config = GetConfData()
	db := DbStr{}
	db.LoadGrade()
	db.LoadStudents()
	return db
}

func (d DbStr) FindStudentByID(id int) (Student, error) {
	for _, student := range d.Students {
		if student.ID == id {
			return student, nil
		}
	}
	return Student{}, errors.New(fmt.Sprintf("Can't find student by id <%v>", id))
}
