/*
 * Copyright (c) 2023 Tom5521- All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package data

import (
	"fmt"
	"log"
)

type DBStr struct {
	Students []Student
	Grades   []Grade
}

func (d *DBStr) Update() error {
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
	for _, student := range d.Students {
		err = student.LoadGrades()
		if err != nil {
			log.Println(err)
		}
		err = student.LoadRecords()
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func InitDB() DBStr {
	CheckFiles()
	Config = GetConfData()
	db := DBStr{}
	err := db.LoadGrade()
	if err != nil {
		log.Println(err)
	}
	err = db.LoadStudents()
	if err != nil {
		log.Println(err)
	}
	for _, student := range db.Students {
		err = student.LoadRecords()
		if err != nil {
			log.Println(err)
		}
		err = student.LoadGrades()
		if err != nil {
			log.Println(err)
		}
	}
	return db
}

func (d *DBStr) FindStudentIndexByID(id int) int {
	for i, s := range d.Students {
		if s.ID == id {
			return i
		}
	}
	return -1
}

func (d DBStr) FindStudentByID(id int) (Student, error) {
	for _, student := range d.Students {
		if student.ID == id {
			return student, nil
		}
	}
	return Student{}, fmt.Errorf("can't find student by id <%v>", id)
}

func (d DBStr) GetGradesNames() []string {
	var grades []string
	for _, grade := range d.Grades {
		grades = append(grades, grade.Name)
	}
	return grades
}
