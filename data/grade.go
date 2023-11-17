/*
 * Copyright (c) 2023 Tom5521- All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package data

import (
	"log"

	_ "github.com/glebarez/go-sqlite"
)

type Grade struct {
	ID    int
	Name  string
	Info  string
	Price string
}

func (d *DbStr) AddGrade(NGrade Grade) (LastInsertId int, err error) {
	db, err := GetNewDb()
	if err != nil {
		log.Println(err)
		return -1, err
	}
	defer db.Close()

	const Query string = `
    insert into grades (name,info,price)
    values (?,?,?)`
	res, err := db.Exec(Query,
		NGrade.Name,
		NGrade.Info,
		NGrade.Price,
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
