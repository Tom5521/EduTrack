/*
 * Copyright (c) 2023 Tom5521- All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

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
	for _, student := range db.Students {
		err := student.LoadRecords()
		if err != nil {
			log.Println(err)
		}
	}
	return db
}

func (d *DbStr) FindStudentIndexByID(id int) (index int) {
	for i, s := range d.Students {
		if s.ID == id {
			return i
		}
	}
	return -1
}

func (d DbStr) FindStudentByID(id int) (Student, error) {
	for _, student := range d.Students {
		if student.ID == id {
			return student, nil
		}
	}
	return Student{}, errors.New(fmt.Sprintf("Can't find student by id <%v>", id))
}

func (d DbStr) GetGradesNames() []string {
	var grades []string
	for _, grade := range d.Grades {
		grades = append(grades, grade.Name)
	}
	return grades
}
