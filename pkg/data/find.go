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

func FindStudentGradeIndexByID(id uint) int {
	for i, grade := range StudentGrades {
		if grade.ID == id {
			return i
		}
	}
	return -1
}

func GetGradesNames() []string {
	var names []string
	for _, grade := range Grades {
		names = append(names, grade.Name)
	}
	return names
}

func (s Student) GetGradesNames() []string {
	s.GetGrades()
	var names []string
	for _, sgrade := range s.Grades {
		for _, grade := range Grades {
			if sgrade.ID == grade.ID {
				names = append(names, grade.Name)
			}
		}
	}
	return names
}

func GetStudentDNIs() []string {
	var dnis []string
	for _, student := range Students {
		dnis = append(dnis, student.DNI)
	}
	return dnis
}

func FindGradeByName(name string) Grade {
	for _, grade := range Grades {
		if name == grade.Name {
			return grade
		}
	}
	return Grade{}
}

func FindStudentIndexByDNI(dni string) int {
	for i, student := range Students {
		if student.DNI == dni {
			return i
		}
	}
	return -1
}

func FindRecordIndexByID(id uint) int {
	for i, record := range Records {
		if record.ID == id {
			return i
		}
	}
	return -1
}

func FindStudentIndexByName(name string) int {
	for i, student := range Students {
		if student.Name == name {
			return i
		}
	}
	return -1
}

func FindGradeIndexbyName(name string) int {
	for i, grade := range Grades {
		if grade.Name == name {
			return i
		}
	}
	return -1
}

func FindRecordIndexByName(name string) int {
	for i, record := range Records {
		if record.Name == name {
			return i
		}
	}
	return -1
}
