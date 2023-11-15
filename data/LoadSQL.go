package data

import (
	"database/sql"
	"log"

	_ "github.com/glebarez/go-sqlite"
)

func GetNewDb() (*sql.DB, error) {
	db, err := sql.Open("sqlite", Config.DatabaseFile)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return db, nil
}

func (d *DbStr) LoadStudents() error {
	db, err := GetNewDb()
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()
	const GetStudentsQuery string = `
		select * from students
	`
	rows, err := db.Query(GetStudentsQuery)
	if err != nil {
		log.Println(err)
		return err
	}
	defer rows.Close()
	var students []Student
	for rows.Next() {
		var student Student
		if err := rows.Scan(
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

func (d *DbStr) LoadGrade() error {
	db, err := GetNewDb()
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()
	const GetGradesQuery string = `
		select * from grades
	`
	rows, err := db.Query(GetGradesQuery)
	if err != nil {
		log.Println(err)
		return err
	}
	defer rows.Close()
	var grades []Grade
	for rows.Next() {
		var grade Grade
		if err := rows.Scan(
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
