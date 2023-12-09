package data

func (s *Student) Edit(newStudent *Student) error {
	s.Name = newStudent.Name
	s.DNI = newStudent.DNI
	s.Age = newStudent.Age
	s.PhoneNumber = newStudent.PhoneNumber
	s.ImageFilePath = newStudent.ImageFilePath
	return DB.Save(s).Error
}

func (r *Record) Edit(newRecord *Record) error {
	r.Name = newRecord.Name
	r.StudentID = newRecord.StudentID
	r.Info = newRecord.Info
	r.Date = newRecord.Date
	return DB.Save(r).Error
}

func (g *Course) Edit(newGrade *Course) error {
	g.Name = newGrade.Name
	g.Info = newGrade.Info
	g.Price = newGrade.Price
	return DB.Save(g).Error
}

func (s *StudentCourse) Edit(nSG *StudentCourse) error {
	s.StudentID = nSG.StudentID
	s.CourseID = nSG.CourseID
	s.Start = nSG.Start
	s.End = nSG.End
	return DB.Save(s).Error
}

// Global

func EditStudentCourse(id uint, newStudentGrade StudentCourse) error {
	i := FindStudentCourseIndexByID(id)
	return StudentCourses[i].Edit(&newStudentGrade)
}

func EditStudent(id uint, newStudent Student) error {
	i := FindStudentIndexByID(id)
	return Students[i].Edit(&newStudent)
}
func EditGrade(id uint, newGrade Course) error {
	i := FindCourseIndexByID(id)
	return Courses[i].Edit(&newGrade)
}

func EditStudentByStruct(s *Student, newStudent Student) error {
	return s.Edit(&newStudent)
}

func EditGradeByStruct(g *Course, newGrade Course) error {
	return g.Edit(&newGrade)
}
