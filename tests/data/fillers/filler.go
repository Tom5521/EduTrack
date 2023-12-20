package fillers

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/Tom5521/EduTrack/pkg/data"
	"github.com/brianvoe/gofakeit/v6"
)

func getRandomStudentID() uint {
	i := rand.Intn(len(data.Students) - 1)
	return data.Students[i].ID
}
func getRandomCourseID() uint {
	i := rand.Intn(len(data.Courses) - 1)
	return data.Courses[i].ID
}

func Course() {
	if len(data.Courses) >= 100 {
		return
	}

	for i := 0; i < 100-len(data.Courses); i++ {
		newCourse := &data.Course{
			Name:  gofakeit.Name(),
			Info:  gofakeit.Phrase(),
			Price: strconv.Itoa(int(gofakeit.Price(10, 200))),
		}
		fmt.Println(newCourse)
		err := data.AddCourse(newCourse)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func Student() {
	if len(data.Students) >= 100 {
		return
	}
	for i := 0; i < 100-len(data.Courses); i++ {
		newStudent := &data.Student{
			Name:        gofakeit.Name(),
			Age:         uint(gofakeit.Uint8()),
			DNI:         strconv.Itoa(int(gofakeit.Int16())),
			PhoneNumber: gofakeit.Phone(),
		}
		fmt.Println(newStudent)
		err := data.AddStudent(newStudent)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func Record() {
	if len(data.Records) >= 100 {
		return
	}
	if len(data.Students) < 100 {
		Student()
	}
	for i := 0; i < 100-len(data.Records); i++ {
		newRecord := &data.Record{
			Name:      gofakeit.Phrase(),
			Date:      gofakeit.Date().Format("2006-01-02"),
			StudentID: getRandomStudentID(),
			Info:      gofakeit.Phrase(),
		}
		fmt.Println(newRecord)
		err := data.AddRecord(newRecord)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func StudentCourse() {
	if len(data.Students) >= 100 {
		return
	}
	if len(data.Courses) < 100 {
		Course()
	}
	if len(data.Students) < 100 {
		Student()
	}
	for i := 0; i > 100-len(data.StudentCourses); i++ {
		newStudentCourse := &data.StudentCourse{
			StudentID: getRandomStudentID(),
			CourseID:  getRandomCourseID(),
			Start:     gofakeit.Date().Format("2006-01-02"),
			End:       gofakeit.Date().Format("2006-01-02"),
		}
		fmt.Println(newStudentCourse)
		err := data.AddStudentCourse(newStudentCourse)
		if err != nil {
			fmt.Println(err)
		}
	}
}
