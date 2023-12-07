package data

import (
	"slices"
)

func FindStudentIndexByID(id uint) int {
	return slices.IndexFunc(Students, func(s Student) bool {
		return s.ID == id
	})
}

func FindGradeIndexByID(id uint) int {
	return slices.IndexFunc(Grades, func(g Grade) bool {
		return g.ID == id
	})
}

func FindStudentGradeIndexByID(id uint) int {
	return slices.IndexFunc(StudentGrades, func(sg StudentGrade) bool {
		return sg.ID == id
	})
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
	return slices.IndexFunc(Students, func(s Student) bool {
		return s.DNI == dni
	})
}

func FindRecordIndexByID(id uint) int {
	return slices.IndexFunc(Records, func(r Record) bool {
		return r.ID == id
	})
}

func FindStudentIndexByName(name string) int {
	return slices.IndexFunc(Students, func(s Student) bool {
		return s.Name == name
	})
}

func FindGradeIndexbyName(name string) int {
	return slices.IndexFunc(Grades, func(g Grade) bool {
		return g.Name == name
	})
}

func FindRecordIndexByName(name string) int {
	return slices.IndexFunc(Records, func(r Record) bool {
		return r.Name == name
	})
}
