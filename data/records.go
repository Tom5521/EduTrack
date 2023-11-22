package data

import (
	"errors"
	"log"
)

type Record struct {
	ID        int
	StudentId int
	Name      string
	Date      string
	Info      string
}

func (s Student) FindRecordIndexByID(id int) (index int) {
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

func (s *Student) AddRecord(newR Record) (lastInsertID int, err error) {
	db, err := GetNewDb()
	if err != nil {
		log.Println(err)
		return -1, err
	}
	defer db.Close()
	const Query string = `
		insert into records (student_id,name,date,info)
		values (?,?,?,?)
	`
	result, err := db.Exec(Query,
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

func (r Record) Edit(NewRec Record) (err error) {
	db, err := GetNewDb()
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()
	const Query string = `
		update records set student_id = ?,name = ?,date = ?,info = ?
		where record_id = ?
`
	_, err = db.Exec(Query,
		NewRec.StudentId,
		NewRec.Name,
		NewRec.Date,
		NewRec.Info,
		r.ID,
	)
	if err != nil {
		log.Println(err)
		return err
	}

	i := Db.FindStudentIndexByID(NewRec.StudentId)
	if NewRec.StudentId != r.StudentId {
		i := Db.FindStudentIndexByID(r.StudentId)
		if i != -1 {
			err := Db.Students[i].LoadRecords()
			if err != nil {
				log.Println(err)
			}
		}
	}
	if i != -1 {
		err = Db.Students[i].LoadRecords()
		if err != nil {
			log.Println(err)
		}
	}
	return err
}

func (r Record) Delete() error {
	db, err := GetNewDb()
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()
	const Query string = `
		delete from records where record_id = ?
	`
	_, err = db.Exec(Query, r.ID)
	if err != nil {
		log.Println(err)
		return err
	}
	id := Db.FindStudentIndexByID(r.StudentId)
	if id == -1 {
		return errors.New("Can't find student ID")
	}
	err = Db.Students[id].LoadRecords()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
