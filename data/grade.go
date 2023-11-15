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

func (g Grade) Delete() error {
	err := DB.DeleteFrom("grades", "grade_id", g.ID)
	err = DB.LoadGrade()
	if err != nil {
		log.Println(err)
	}
	return nil
}
