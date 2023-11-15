package data

import (
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

func (d *DB_Str) EditStudent(id int, editedStudent Student) error {
	db, err := GetNewDb()
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()
	const EditStudentQuery string = `
		update students set Name = ?,Age = ?,DNI = ?,Phone_number = ?,ImagePath = ? 
		where student_id = ?
	`
	_, err = db.Exec(EditStudentQuery,
		editedStudent.Name,
		editedStudent.Age,
		editedStudent.DNI,
		editedStudent.Phone_number,
		editedStudent.ImageFilePath,
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

func (d *DB_Str) AddStudent(newStudent Student) (LastInsertId int, err error) {
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
		newStudent.Name,
		newStudent.Age,
		newStudent.DNI,
		newStudent.Phone_number,
		newStudent.ImageFilePath,
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
