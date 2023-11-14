package data

import "time"

type Register struct {
	Date string
	Name string
	Data string
}

// Register functions
func NewRegister(Name, Data string, Date ...string) Register {
	getNow := func() string {
		return time.Now().Format("2006-01-02 15:04:05")
	}
	reg := Register{}
	reg.Date = getNow()
	reg.Data = Data
	reg.Name = Name
	if Name == "" {
		reg.Name = getNow()
	}
	if len(Date) != 0 {
		if Date[0] != "" {
			reg.Date = Date[0]
		}
	}
	return reg
}
func (r *Register) Overwrite(newRegister Register) {
	r = &newRegister
}

func (r *Register) EditDate(newDate string) {
	r.Date = newDate
}
func (r *Register) EditName(newName string) {
	r.Name = newName
}
func (r *Register) EditData(newData string) {
	r.Data = newData
}
