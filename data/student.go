package data

import "slices"

// Student represents the structure of a student's data.
type Student struct {
	Name          string
	Age           string
	ID            string
	Phone_number  string
	ImageFilePath string
	Registers     []Register
	Grades        []StudentGrade
}

type StudentGrade struct {
	*Grade
	Start string
	End   string
}

func (s *StudentGrade) Convert(g *Grade, Star, End string) {
	s.Grade = g
}

// Student functions

func (s *Student) Overwrite(newStudent Student) {
	s = &newStudent
}

func (s *Student) EditName(newName string) {
	s.Name = newName
}

func (s *Student) EditAge(newAge string) {
	s.Age = newAge
}

func (s *Student) EditID(newID string) {
	s.ID = newID
}

func (s *Student) EditPhoneNumber(newPhone string) {
	s.Phone_number = newPhone
}

func (s *Student) EditImagePath(newPath string) {
	s.ImageFilePath = newPath
}

// Register functions

func (s *Student) AddRegister(r Register) {
	s.Registers = append(s.Registers, r)
}

func (s *Student) EditRegister(id int, r Register) {
	if id < 0 || id >= len(s.Registers) {
		return
	}
	s.Registers[id] = r
}

func (s Student) GetRegisterNames() []string {
	var registers []string
	for _, reg := range s.Registers {
		registers = append(registers, reg.Name)
	}
	return registers
}

// Grade Funtions

func (s *Student) AddGrade(g StudentGrade) {
	s.Grades = append(s.Grades, g)
}

func (s *Student) EditGrade(id int, newGrade StudentGrade) {
	if id < 0 || id >= len(s.Grades) {
		return
	}
	s.Grades[id] = newGrade
}

func (s *Student) DeleteGrade(id int) {
	if id < 0 || id >= len(s.Grades) {
		return
	}
	s.Grades = slices.Delete(s.Grades, id, id+1)
}

func (s Student) GetGradeNames() []string {
	var names []string
	for _, grade := range s.Grades {
		names = append(names, grade.Name)
	}
	return names
}

// Data structure funcs

// GetStudentsNames returns a slice of student names.
func (d Data_str) GetStudentNames() []string {
	var names []string
	for _, student := range d.Students {
		names = append(names, student.Name)
	}
	return names
}

// GetStudentIDs returns a slice of student IDs.
func (d Data_str) GetStudentIDs() []string {
	var IDs []string
	for _, student := range d.Students {
		IDs = append(IDs, student.ID)
	}
	return IDs
}

// FindStudentByID searches for a student by their ID and returns a pointer to the student if found.
func (d Data_str) FindStudentByID(studentID string) *Student {
	for _, student := range d.Students {
		if student.ID == studentID {
			return &student
		}
	}
	return nil
}

func (d Data_str) FindStudentIndexByID(studentID string) int {
	for i, student := range d.Students {
		if student.ID == studentID {
			return i
		}
	}
	return -1
}
