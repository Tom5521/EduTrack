package data

import (
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
