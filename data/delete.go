package data

import (
	"errors"
	"log"
)

func Delete(i Deleter) {
	i.Delete()
}

type Deleter interface {
	Delete() error
}

func (g Grade) Delete() error {
	err := DB.DeleteFrom("grades", "grade_id", g.ID)
	err = DB.LoadGrade()
	if err != nil {
		log.Println(err)
	}
	return nil
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
	id := DB.FindStudentIndexByID(r.StudentId)
	if id == -1 {
		return errors.New("Can't find student ID")
	}
	err = DB.Students[id].LoadRecords()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
