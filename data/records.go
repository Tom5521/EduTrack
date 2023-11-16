package data

import "log"

type Record struct {
	ID        int
	StudentId int
	Name      string
	Date      string
	Info      string
}

func (s *Student) LoadRecords() error {
	db, err := GetNewDb()
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()
	const Query = `
	  select * from records where student_id = ?
	`

	rows, err := db.Query(Query, s.ID)
	if err != nil {
		log.Println(err)
		return err
	}
	defer rows.Close()
	var records []Record
	for rows.Next() {
		var record Record
		if err := rows.Scan(
			&record.ID,
			&record.StudentId,
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
