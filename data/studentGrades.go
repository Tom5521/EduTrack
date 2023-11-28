package data

import (
	"errors"
	"log"
)

type StudentGrade struct {
	StudentGradeID int
	GradeID        int
	StudentID      int
	Start          string
	End            string
}

func (s StudentGrade) Edit(newData StudentGrade) error {
	db, err := GetNewDB()
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()
	const query string = `
    update student_grades set student_id = ?,start = ?,end = ?,grade_id = ?
    where student_grade_id = ?
    `
	_, err = db.Exec(query,
		newData.StudentID,
		newData.Start,
		newData.End,
		newData.GradeID,
		s.StudentGradeID,
	)
	if err != nil {
		log.Println(err)
		return err
	}
	id := DB.FindStudentIndexByID(s.StudentID)
	if id == -1 {
		return errors.New("can't find student index")
	}
	err = DB.Students[id].LoadGrades()
	if err != nil {
		log.Println(err)
	}

	return err
}

func (s *Student) AddGrade(newGrade StudentGrade) (int, error) {
	db, err := GetNewDB()
	if err != nil {
		log.Println(err)
		return -1, err
	}
	defer db.Close()
	const query string = `
		insert into student_grades (student_id,grade_id,start,end)
		values (?,?,?,?)
	`
	result, err := db.Exec(query, newGrade.StudentID, newGrade.GradeID, newGrade.Start, newGrade.End)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return -1, err
	}
	err = s.LoadGrades()
	if err != nil {
		log.Println(err)
	}
	return int(id), err
}

func (s Student) FindGradeIndexByID(id int) int {
	for i, grade := range s.Grades {
		if grade.StudentGradeID == id {
			return i
		}
	}

	return -1
}
