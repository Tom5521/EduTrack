package data

func (s *Student) Edit(newStudent Student) error {
	s.Name = newStudent.Name
	s.DNI = newStudent.DNI
	s.Age = newStudent.Age
	s.PhoneNumber = newStudent.PhoneNumber
	s.ImageFilePath = newStudent.ImageFilePath
	return DB.Save(s).Error
}

func (r *Record) Edit(newRecord Record) error {
	r.Name = newRecord.Name
	r.StudentID = newRecord.StudentID
	r.Info = newRecord.Info
	r.Date = newRecord.Date
	return DB.Save(r).Error
}

func (g *Grade) Edit(newGrade Grade) error {
	g.Name = newGrade.Name
	g.Info = newGrade.Info
	g.Price = newGrade.Price
	return DB.Save(g).Error
}

func (s *StudentGrade) Edit(nSG StudentGrade) error {
	s.StudentID = nSG.StudentID
	s.GradeID = nSG.GradeID
	s.Start = nSG.Start
	s.End = nSG.End
	return DB.Save(s).Error
}

// Global

func EditStudentGrade(id uint, newStudentGrade StudentGrade) error {

}

func EditStudent(id uint, newStudent Student) error {
	i := FindStudentIndexByID(id)
	return Students[i].Edit(newStudent)
}
func EditGrade(id uint, newGrade Grade) error {
	i := FindGradeIndexByID(id)
	return Grades[i].Edit(newGrade)
}

func EditStudentByStruct(s *Student, newStudent Student) error {
	return s.Edit(newStudent)
}

func EditGradeByStruct(g *Grade, newGrade Grade) error {
	return g.Edit(newGrade)
}
