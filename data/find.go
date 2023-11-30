package data

func FindStudentIndexByID(id uint) int {
	for i, student := range Students {
		if student.ID == id {
			return i
		}
	}
	return -1
}

func FindGradeIndexByID(id uint) int {
	for i, grades := range Grades {
		if grades.ID == id {
			return i
		}
	}
	return -1
}
