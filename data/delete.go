package data

func (g Grade) Delete() error {
	err := DB.Delete(&g).Error
	printErr(err)
	return LoadGrades()
}

func (s Student) Delete() error {
	err := DB.Delete(&s).Error
	printErr(err)
	return LoadStudents()
}

func (s StudentGrade) Delete() error {
	err := DB.Delete(&s).Error
	printErr(err)
	return LoadStudentGrades()
}

func (r Record) Delete() error {
	err := DB.Delete(&r).Error
	printErr(err)
	return LoadRecords()
}
