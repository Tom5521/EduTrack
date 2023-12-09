package data

func AddGrade(newGrade *Course) error {
	err := DB.Create(&newGrade).Error
	printErr(err)
	return LoadCourses()
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

func AddStudentCourse(newGrade *StudentCourse) error {
	err := DB.Create(&newGrade).Error
	printErr(err)
	return LoadStudentCourses()
}

// Student

func (s Student) AddCourse(newGrade *StudentCourse) error {
	newGrade.StudentID = s.ID
	err := AddStudentCourse(newGrade)
	s.GetCourses()
	return err
}

func (s Student) AddRecord(newRecord *Record) error {
	newRecord.StudentID = s.ID
	err := AddRecord(newRecord)
	s.GetRecords()
	return err
}
