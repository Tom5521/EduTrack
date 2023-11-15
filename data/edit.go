package data

import (
	"fmt"
	"log"
)

func (d *DB_Str) DeleteFrom(table, column string, id int) error {
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
