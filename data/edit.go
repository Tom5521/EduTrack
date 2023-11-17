/*
 * Copyright (c) 2023 Tom5521- All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package data

import (
	"fmt"
	"log"
)

func (d *DbStr) DeleteFrom(table, column string, id int) error {
	db, err := GetNewDb()
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()
	const DeleteQuery string = `
		delete from %v where %v = %v
	`
	_, err = db.Exec(fmt.Sprintf(DeleteQuery, table, column, id))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (d *DbStr) EditGrade(id int, EdGrade Grade) error {
	db, err := GetNewDb()
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()
	const Query string = `
		update grades set Name = ?,info = ?,price = ? 
		where grade_id = ?
	`
	_, err = db.Exec(Query,
		EdGrade.Name,
		EdGrade.Info,
		EdGrade.Price,
		id,
	)
	if err != nil {
		log.Println(err)
		return err
	}
	err = d.LoadGrade()
	if err != nil {
		log.Println(err)
	}
	return nil
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

	i := DB.FindStudentIndexByID(NewRec.StudentId)
	if NewRec.StudentId != r.StudentId {
		i := DB.FindStudentIndexByID(r.StudentId)
		if i != -1 {
			err := DB.Students[i].LoadRecords()
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
