package data

import (
	"slices"
)

func StudentIndexByID(id uint) int {
	return slices.IndexFunc(Students, func(s Student) bool {
		return s.ID == id
	})
}

func CourseIndexByID(id uint) int {
	return slices.IndexFunc(Courses, func(g Course) bool {
		return g.ID == id
	})
}

func StudentCourseIndexByID(id uint) int {
	return slices.IndexFunc(StudentCourses, func(sg StudentCourse) bool {
		return sg.ID == id
	})
}

func CoursesNames() []string {
	var names []string
	for _, grade := range Courses {
		names = append(names, grade.Name)
	}
	return names
}

func (s Student) CourseNames() []string {
	s.Courses()
	var names []string
	for _, sgrade := range s.StudentCourses {
		for _, grade := range Courses {
			if sgrade.ID == grade.ID {
				names = append(names, grade.Name)
			}
		}
	}
	return names
}

func StudentDNIs() []string {
	var dnis []string
	for _, student := range Students {
		dnis = append(dnis, student.DNI)
	}
	return dnis
}

func CourseByName(name string) Course {
	for _, grade := range Courses {
		if name == grade.Name {
			return grade
		}
	}
	return Course{}
}

func StudentIndexByDNI(dni string) int {
	return slices.IndexFunc(Students, func(s Student) bool {
		return s.DNI == dni
	})
}

func RecordIndexByID(id uint) int {
	return slices.IndexFunc(Records, func(r Record) bool {
		return r.ID == id
	})
}

func StudentIndexByName(name string) int {
	return slices.IndexFunc(Students, func(s Student) bool {
		return s.Name == name
	})
}

func CourseIndexbyName(name string) int {
	return slices.IndexFunc(Courses, func(g Course) bool {
		return g.Name == name
	})
}

func RecordIndexByName(name string) int {
	return slices.IndexFunc(Records, func(r Record) bool {
		return r.Name == name
	})
}

// Note: This is extremely slow.
func (s Student) CourseIndexByName(name string) int {
	return slices.IndexFunc(s.StudentCourses, func(sc StudentCourse) bool {
		i := CourseIndexByID(sc.CourseID)
		if i == -1 {
			return false
		}
		c := Courses[i]
		return c.Name == name
	})
}

func CourseIndexByName(name string) int {
	return slices.IndexFunc(Courses, func(c Course) bool {
		return c.Name == name
	})
}

func StudentIDs() []uint {
	var ids []uint
	for _, student := range Students {
		ids = append(ids, student.ID)
	}
	return ids
}
