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

func (d *DBStr) FindGradeIndexByID(id int) int {
	err := d.LoadGrade()
	if err != nil {
		log.Println(err)
	}
	for i, grade := range d.Grades {
		if grade.ID == id {
			return i
		}
	}
	return -1
}

func (s Student) Delete() error {
	err := DB.DeleteFrom("Students", "student_id", s.ID)
	if err != nil {
		log.Println(err)
		return err
	}
	err = DB.LoadStudents()
	if err != nil {
		log.Println(err)
	}
	return err
}

func (s Student) Edit(edStudent Student) error {
	db, err := GetNewDB()
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()
	const query string = `
		update students set Name = ?,Age = ?,DNI = ?,Phone_number = ?,ImagePath = ? 
		where student_id = ?
	`
	_, err = db.Exec(query,
		edStudent.Name,
		edStudent.Age,
		edStudent.DNI,
		edStudent.PhoneNumber,
		edStudent.ImageFilePath,
		s.ID,
	)
	if err != nil {
		log.Println(err)
		return err
	}
	err = DB.LoadStudents()
	if err != nil {
		log.Println(err)
	}
	return err
}

func (d *DBStr) EditStudent(id int, edStudent Student) error {
	i := d.FindStudentIndexByID(id)
	student := d.Students[i]
	err := student.Edit(edStudent)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (d *DBStr) AddStudent(nStudent Student) (int, error) {
	db, err := GetNewDB()
	if err != nil {
		log.Println(err)
		return -1, err
	}
	defer db.Close()
	const addStudentQuery string = `
		insert into students (Name,Age,DNI,Phone_number,ImagePath)
		values (?,?,?,?,?)
	`
	result, err := db.Exec(addStudentQuery,
		nStudent.Name,
		nStudent.Age,
		nStudent.DNI,
		nStudent.PhoneNumber,
		nStudent.ImageFilePath,
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
	db, err := GetNewDB()
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()
	const query string = `
		select * from student_grades where student_id = ?
	`
	rows, err := db.Query(query, s.ID)
	if err != nil {
		log.Println(err)
		return err
	}
	if rows.Err() != nil {
		log.Println(rows.Err())
		return rows.Err()
	}
	defer rows.Close()
	for rows.Next() {
		var grade StudentGrade
		if err = rows.Scan(
			&grade.StudentID,
			&grade.GradeID,
			&grade.Start,
			&grade.End,
			&grade.StudentGradeID,
		); err != nil {
			log.Println(err)
			return err
		}
		s.Grades = append(s.Grades, grade)
	}
	return nil
}

func (d DBStr) GetStudentDNIs() []string {
	var students []string
	for _, student := range d.Students {
		students = append(students, student.DNI)
	}
	return students
}

func (s Student) GetGradeNames() []string {
	var names []string
	for _, grade := range s.Grades {
		i := s.FindGradeIndexByID(grade.GradeID)
		if i == -1 {
			continue
		}
		names = append(names, DB.Grades[i].Name)
	}
	return names
}
func (d DBStr) FindStudentIndexByDNI(dni string) int {
	for i, student := range d.Students {
		if student.DNI == dni {
			return i
		}
	}
	return -1
}
