/*
 * Copyright (c) 2023 Tom5521- All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package data

import (
	"log"
)

type Student struct {
	ID            int
	Name          string
	Age           int
	DNI           string
	PhoneNumber   string
	ImageFilePath string
	Records       []Record
	Grades        []StudentGrade
}

func (d *DbStr) FindGradeIndexByID(id int) (index int) {
	d.LoadGrade()
	for i, grade := range d.Grades {
		if grade.ID == id {
			return i
		}
	}
	return -1
}

func (s Student) Delete() error {
	DB.DeleteFrom("Students", "student_id", s.ID)
	err := DB.LoadStudents()
	if err != nil {
		log.Println(err)
	}
	return err
}

func (d *DbStr) EditStudent(id int, EdStudent Student) error {
	db, err := GetNewDb()
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()
	const Query string = `
		update students set Name = ?,Age = ?,DNI = ?,Phone_number = ?,ImagePath = ? 
		where student_id = ?
	`
	_, err = db.Exec(Query,
		EdStudent.Name,
		EdStudent.Age,
		EdStudent.DNI,
		EdStudent.PhoneNumber,
		EdStudent.ImageFilePath,
		id,
	)
	if err != nil {
		log.Println(err)
		return err
	}
	err = d.LoadStudents()
	if err != nil {
		log.Println(err)
	}
	return err
}

func (d *DbStr) AddStudent(NStudent Student) (LastInsertId int, err error) {
	db, err := GetNewDb()
	if err != nil {
		log.Println(err)
		return -1, err
	}
	defer db.Close()
	const AddStudentQuery string = `
		insert into students (Name,Age,DNI,Phone_number,ImagePath)
		values (?,?,?,?,?)
	`
	result, err := db.Exec(AddStudentQuery,
		NStudent.Name,
		NStudent.Age,
		NStudent.DNI,
		NStudent.PhoneNumber,
		NStudent.ImageFilePath,
	)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	err = d.LoadStudents()
	if err != nil {
		log.Println(err)
	}
	return int(id), err
}

func (s *Student) LoadGrades() error {
	db, err := GetNewDb()
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()
	const Query string = `
		select * from student_grades where student_id = ?
	`
	rows, err := db.Query(Query, s.ID)
	if err != nil {
		log.Println(err)
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var grade StudentGrade
		if err := rows.Scan(
			&grade.Student_id,
			&grade.ID,
			&grade.Start,
			&grade.End,
		); err != nil {
			log.Println(err)
			return err
		}
		s.Grades = append(s.Grades, grade)
	}
	return nil
}
