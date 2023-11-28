/*
 * Copyright (c) 2023 Tom5521- All Rights Reserved.
 *
 * This project is licensed under the MIT License.
 */

package data

import (
	"log"
)

func (d *DBStr) DeleteFrom(table, column string, id int) error {
	db, err := GetNewDB()
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()
	const deleteQuery string = `
		delete from ? where ? = ?
	`
	_, err = db.Exec(deleteQuery, table, column, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
