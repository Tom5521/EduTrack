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
	Db.DeleteFrom("Students", "student_id", s.ID)
	err := Db.LoadStudents()
	if err != nil {
		log.Println(err)
	}
	return err
}

func (s Student) Edit(EdStudent Student) error {
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
		s.ID,
	)
	if err != nil {
		log.Println(err)
		return err
	}
	err = Db.LoadStudents()
	if err != nil {
		log.Println(err)
	}
	return err
}

func (d *DbStr) EditStudent(id int, EdStudent Student) error {
	i := d.FindStudentIndexByID(id)
	student := d.Students[i]
	err := student.Edit(EdStudent)
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

func (d DbStr) GetStudentDNIs() []string {
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
		names = append(names, Db.Grades[i].Name)
	}
	return names
}
func (d DbStr) FindStudentIndexByDNI(dni string) (index int) {
	for i, student := range d.Students {
		if student.DNI == dni {
			return i
		}
	}
	return -1
}
