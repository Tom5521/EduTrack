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

func (d *DBStr) DeleteFrom(table, column string, key int) error {
	db, err := GetNewDB()
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()
	const deleteQuery string = `
		delete from "%v" where "%v" = %v
	`
	// For some reason that I can't understand the db.Exec doesn't accept the arguments correctly, so I'll just use fmt.Sprintf for this.
	_, err = db.Exec(fmt.Sprintf(deleteQuery, table, column, key))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
