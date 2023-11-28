package data

import (
	"errors"
	"log"
)

type Record struct {
	ID        int
	StudentID int
	Name      string
	Date      string
	Info      string
}

func (s Student) FindRecordIndexByID(id int) int {
	for i, student := range s.Records {
		if student.ID == id {
			return i
		}
	}
	return -1
}

func (s *Student) DeleteRecord(id int) error {
	i := s.FindRecordIndexByID(id)
	err := s.Records[i].Delete()
	if err != nil {
		log.Println(err)
		return err
	}
	err = s.LoadRecords()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *Student) AddRecord(newR Record) (int, error) {
	db, err := GetNewDB()
	if err != nil {
		log.Println(err)
		return -1, err
	}
	defer db.Close()
	const query string = `
		insert into records (student_id,name,date,info)
		values (?,?,?,?)
	`
	result, err := db.Exec(query,
		s.ID,
		newR.Name,
		newR.Date,
		newR.Info,
	)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return -1, err
	}
	err = s.LoadRecords()
	if err != nil {
		log.Println(err)
		return int(id), err
	}
	return int(id), err
}

func (r Record) Edit(newRec Record) error {
	db, err := GetNewDB()
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()
	const query string = `
		update records set student_id = ?,name = ?,date = ?,info = ?
		where record_id = ?
`
	_, err = db.Exec(query,
		newRec.StudentID,
		newRec.Name,
		newRec.Date,
		newRec.Info,
		r.ID,
	)
	if err != nil {
		log.Println(err)
		return err
	}

	i := DB.FindStudentIndexByID(newRec.StudentID)
	if newRec.StudentID != r.StudentID {
		i = DB.FindStudentIndexByID(r.StudentID)
		if i != -1 {
			err = DB.Students[i].LoadRecords()
			if err != nil {
				log.Println(err)
			}
		}
	}
	if i != -1 {
		err = DB.Students[i].LoadRecords()
		if err != nil {
			log.Println(err)
		}
	}
	return err
}

func (r Record) Delete() error {
	db, err := GetNewDB()
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()
	const query string = `
		delete from records where record_id = ?
	`
	_, err = db.Exec(query, r.ID)
	if err != nil {
		log.Println(err)
		return err
	}
	id := DB.FindStudentIndexByID(r.StudentID)
	if id == -1 {
		return errors.New("can't find student ID")
	}
	err = DB.Students[id].LoadRecords()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
