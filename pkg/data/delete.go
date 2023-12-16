package data

func (c Course) Delete() error {
	err := DB.Delete(&c).Error
	printErr(err)
	for _, studentCourse := range StudentCourses {
		if studentCourse.CourseID == c.ID {
			err = studentCourse.Delete()
			printErr(err)
		}
	}
	return LoadCourses()
}

func (s Student) Delete() error {
	err := DB.Delete(&s).Error
	printErr(err)
	for _, record := range Records {
		if record.StudentID == s.ID {
			err = record.Delete()
			printErr(err)
		}
	}
	for _, studentCourse := range StudentCourses {
		if studentCourse.StudentID == s.ID {
			err = studentCourse.Delete()
			printErr(err)
		}
	}
	return LoadStudents()
}

func (s StudentCourse) Delete() error {
	err := DB.Delete(&s).Error
	printErr(err)
	return LoadStudentCourses()
}

func (r Record) Delete() error {
	err := DB.Delete(&r).Error
	printErr(err)
	return LoadRecords()
}

func (s *Student) DeleteRecord(id uint) error {
	i := FindRecordIndexByID(id)
	err := Delete(Records[i])
	printErr(err)
	s.GetRecords()
	return nil
}

type Deleter interface {
	Delete() error
}

func Delete(i Deleter) error {
	return i.Delete()
}
