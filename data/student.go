package data

import (
	"fmt"
	"log"
)

type Student struct {
	ID            int
	Name          string
	Age           int
	DNI           string
	Phone_number  string
	ImageFilePath string
}

func (s Student) Delete() error {
	DB.DeleteFrom("Students", "student_id", s.ID)
	err := DB.LoadStudents()
	if err != nil {
		log.Println(err)
	}
	return err
}

func (d *DB_Str) EditStudent(editedStudent Student) error {

	return nil
}

func (d *DB_Str) AddStudent(newStudent Student) (LastInsertId int, err error) {
	db, err := GetNewDb()
	if err != nil {
		log.Println(err)
		return -1, err
	}
	defer db.Close()
	const AddStudentQuery string = `
		insert into students (Name,Age,DNI,Phone_number,ImagePath)
		values (%v,%v,%v,%v,%v)
	`
	Query := fmt.Sprintf(AddStudentQuery,
		newStudent.Name,
		newStudent.Age,
		newStudent.DNI,
		newStudent.Phone_number,
		newStudent.ImageFilePath,
	)
	result, err := db.Exec(Query)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(id), nil
}
