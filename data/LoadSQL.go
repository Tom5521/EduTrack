/*
 * Copyright (c) 2023 Tom5521- All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package data

import (
	"database/sql"
	"log"

	_ "github.com/glebarez/go-sqlite"
)

func GetNewDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", Config.DatabaseFile)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return db, nil
}

func (d *DBStr) LoadStudents() error {
	db, err := GetNewDB()
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()
	const getStudentsQuery string = `
		select * from students
	`
	rows, err := db.Query(getStudentsQuery)
	if err != nil {
		log.Println(err)
		return err
	}
	if rows.Err() != nil {
		log.Println(rows.Err())
		return rows.Err()
	}
	defer rows.Close()
	var students []Student
	for rows.Next() {
		var student Student
		if err = rows.Scan(
			&student.ID,
			&student.Name,
			&student.DNI,
			&student.Age,
			&student.PhoneNumber,
			&student.ImageFilePath,
		); err != nil {
			log.Println(err)
			return err
		}
		students = append(students, student)
	}
	d.Students = students
	return nil
}

func (d *DBStr) LoadGrade() error {
	db, err := GetNewDB()
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()
	const getGradesQuery string = `
		select * from grades
	`
	rows, err := db.Query(getGradesQuery)
	if err != nil {
		log.Println(err)
		return err
	}
	if rows.Err() != nil {
		log.Println(rows.Err())
		return rows.Err()
	}
	defer rows.Close()
	var grades []Grade
	for rows.Next() {
		var grade Grade
		if err = rows.Scan(
			&grade.ID,
			&grade.Name,
			&grade.Info,
			&grade.Price,
		); err != nil {
			log.Println(err)
			return err
		}
		grades = append(grades, grade)
	}
	d.Grades = grades
	return nil
}

func (s *Student) LoadRecords() error {
	db, err := GetNewDB()
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()
	const query = `
	  select * from records where student_id = ?
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
	var records []Record
	for rows.Next() {
		var record Record
		if err = rows.Scan(
			&record.ID,
			&record.StudentID,
			&record.Name,
			&record.Date,
			&record.Info,
		); err != nil {
			log.Println(err)
			return err
		}
		records = append(records, record)
	}
	s.Records = records

	return nil
}
