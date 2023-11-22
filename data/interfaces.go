package data

import "log"

type Deleter interface {
	Delete() error
}

func Delete(i Deleter) error {
	err := i.Delete()
	if err != nil {
		log.Println(err)
	}
	return err
}
