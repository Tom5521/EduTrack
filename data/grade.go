/*
 * Copyright (c) 2023 Tom5521- All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package data

import (
	"log"
)

type Grade struct {
	ID    int
	Name  string
	Info  string
	Price string
}

func (d *DBStr) AddGrade(nGrade Grade) (int, error) {
	db, err := GetNewDB()
	if err != nil {
		log.Println(err)
		return -1, err
	}
	defer db.Close()

	const query string = `
    insert into grades (name,info,price)
    values (?,?,?)`
	res, err := db.Exec(query,
		nGrade.Name,
		nGrade.Info,
		nGrade.Price,
	)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	lastIns, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
		return -1, err
	}
	err = d.LoadGrade()
	if err != nil {
		log.Println(err)
	}
	return int(lastIns), err
}

func (d *DBStr) EditGrade(id int, edGrade Grade) error {
	db, err := GetNewDB()
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()
	const query string = `
		update grades set Name = ?,info = ?,price = ? 
		where grade_id = ?
	`
	_, err = db.Exec(query,
		edGrade.Name,
		edGrade.Info,
		edGrade.Price,
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

func (g Grade) Delete() error {
	err := DB.DeleteFrom("grades", "grade_id", g.ID)
	if err != nil {
		log.Println(err)
		return err
	}
	err = DB.LoadGrade()
	if err != nil {
		log.Println(err)
	}
	return nil
}

func (d DBStr) FindGradeByName(name string) Grade {
	for _, grade := range d.Grades {
		if grade.Name == name {
			return grade
		}
	}
	return Grade{}
}
