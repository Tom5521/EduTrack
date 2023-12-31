package gui

import (
	"strconv"

	"github.com/Tom5521/EduTrack/pkg/data"
)

func itoa[T ~int | ~uint](i T) string {
	return strconv.Itoa(int(i))
}

func atoi[T ~string](s T) int {
	ret, err := strconv.Atoi(string(s))
	if err != nil {
		return -1
	}
	return ret
}

func validateStudentDNI(check string, list []string) bool {
	var contains bool
	for _, v := range list {
		if v == check {
			contains = true
			break
		}
	}
	return contains
}

// validateStudentForm checks if all required form fields are not empty.
func validateStudentForm(s data.Student) bool {
	if s.Age == 0 {
		return false
	}
	if s.DNI == "" {
		return false
	}
	if s.PhoneNumber == "" {
		return false
	}
	if s.Name == "" {
		return false
	}
	return true
}
