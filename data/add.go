package data

func AddGrade(newGrade *Grade) error {
	err := DB.Create(&newGrade).Error
	printErr(err)
	return LoadGrades()
}

func AddStudent(newStudent *Student) error {
	err := DB.Create(&newStudent).Error
	printErr(err)
	return LoadStudents()
}

func AddRecord(newRecord *Record) error {
	err := DB.Create(&newRecord).Error
	printErr(err)
	return LoadRecords()
}

func AddStudentGrade(newGrade *StudentGrade) error {
	err := DB.Create(&newGrade).Error
	printErr(err)
	return LoadStudentGrades()
}

// Student

func (s Student) AddGrade(newGrade *StudentGrade) error {
	newGrade.StudentID = s.ID
	err := AddStudentGrade(newGrade)
	s.GetGrades()
	return err
}

func (s Student) AddRecord(newRecord *Record) error {
	newRecord.StudentID = s.ID
	err := AddRecord(newRecord)
	s.GetRecords()
	return err
}
