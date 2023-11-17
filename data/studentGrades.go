package data

import "log"

type StudentGrade struct {
	Grade
	Student_id int
	Start      string
	End        string
}

func (s StudentGrade) Edit(struct{ Student_id, Start, End string }) error {
	db, err := GetNewDb()
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()
	const Query string = `
    update  
    `

	return nil
}
