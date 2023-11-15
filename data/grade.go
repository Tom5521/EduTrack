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

func (d *DB_Str) AddGrade(newGrade Grade) (LastInsertId int, err error) {
	db, err := GetNewDb()
	if err != nil {
		log.Println(err)
		return -1, err
	}
	defer db.Close()

	const AddGradeQuery string = `
    insert into grades (name,info,price)
    values (?,?,?)`
	result, err := db.Exec(AddGradeQuery,
		newGrade.Name,
		newGrade.Info,
		newGrade.Price,
	)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	lastInsert, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return -1, err
	}
	err = d.LoadGrade()
	if err != nil {
		log.Println(err)
	}
	return int(lastInsert), err
}

func (d *DB_Str) EditGrade(id int, editedGrade Grade) error {
	db, err := GetNewDb()
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()
	const EditGradeQuery string = `
		update grades set Name = ?,info = ?,price = ? where grade_id = ?
	`
	_, err = db.Exec(EditGradeQuery,
		editedGrade.Name,
		editedGrade.Info,
		editedGrade.Price,
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
	err = DB.LoadGrade()
	if err != nil {
		log.Println(err)
	}
	return nil
}
