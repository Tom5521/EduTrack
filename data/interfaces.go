package data

type Deleter interface {
	Delete() error
}

func Delete(i Deleter) error {
	return i.Delete()
}
