package data

func AddGrade(newGrade Grade) error {
	err := DB.Create(&newGrade).Error
	printErr(err)
	err = LoadGrades()
	return err
}

func AddStudent(newStudent Student) error {
	err := DB.Create(&newStudent).Error
	printErr(err)
	err = LoadStudents()
	return err
}

func AddRecord(newRecord Record) error {
	err := DB.Create(&newRecord).Error
	printErr(err)
	err = LoadRecords()
	return err
}

func AddStudentGrade(newGrade StudentGrade) error {
	err := DB.Create(&newGrade).Error
	printErr(err)
	err = LoadStudentGrades()
	return err
}

// Student

func (s Student) AddGrade(newGrade StudentGrade) error {
	newGrade.StudentID = s.ID
	return AddStudentGrade(newGrade)
}

func (s Student) AddRecord(newRecord Record) error {
	newRecord.StudentID = s.ID
	return AddRecord(newRecord)
}
