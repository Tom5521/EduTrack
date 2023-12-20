package data

import (
	"slices"
)

func FindStudentIndexByID(id uint) int {
	return slices.IndexFunc(Students, func(s Student) bool {
		return s.ID == id
	})
}

func FindCourseIndexByID(id uint) int {
	return slices.IndexFunc(Courses, func(g Course) bool {
		return g.ID == id
	})
}

func FindStudentCourseIndexByID(id uint) int {
	return slices.IndexFunc(StudentCourses, func(sg StudentCourse) bool {
		return sg.ID == id
	})
}

func GetCoursesNames() []string {
	var names []string
	for _, grade := range Courses {
		names = append(names, grade.Name)
	}
	return names
}

func (s Student) GetCourseNames() []string {
	s.GetCourses()
	var names []string
	for _, sgrade := range s.Courses {
		for _, grade := range Courses {
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

func FindCourseByName(name string) Course {
	for _, grade := range Courses {
		if name == grade.Name {
			return grade
		}
	}
	return Course{}
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

func FindCourseIndexbyName(name string) int {
	return slices.IndexFunc(Courses, func(g Course) bool {
		return g.Name == name
	})
}

func FindRecordIndexByName(name string) int {
	return slices.IndexFunc(Records, func(r Record) bool {
		return r.Name == name
	})
}

// Note: This is extremely slow.
func (s Student) FindCourseIndexByName(name string) int {
	return slices.IndexFunc(s.Courses, func(sc StudentCourse) bool {
		i := FindCourseIndexByID(sc.CourseID)
		if i == -1 {
			return false
		}
		c := Courses[i]
		return c.Name == name
	})
}

func FindCourseIndexByName(name string) int {
	return slices.IndexFunc(Courses, func(c Course) bool {
		return c.Name == name
	})
}

func GetStudentIDs() []uint {
	var ids []uint
	for _, student := range Students {
		ids = append(ids, student.ID)
	}
	return ids
}
