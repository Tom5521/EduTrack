package data

import (
	"os"

	"github.com/Tom5521/MyGolangTools/file"
	"gopkg.in/yaml.v3"
)

var Students []Student

type Student struct {
	Name          string
	Age           int
	ID            string
	Phone_number  string
	ImageFilePath string
}

var (
	confName string = "students.yml"
)

func GetYamlData() {
	var (
		err       error
		data_file []byte
	)
	if check, _ := file.CheckFile(confName); !check {
		data_file = NewYmlFile()
	} else {
		data_file, err = os.ReadFile(confName)
		if err != nil {
			return
		}
	}
	yaml.Unmarshal(data_file, &Students)
}

func NewYmlFile() []byte {
	_, err := os.Create(confName)
	if err != nil {
		return nil
	}
	data, err := os.ReadFile(confName)
	if err != nil {
		return nil
	}
	return data
}

func SaveData() error {
	data, err := yaml.Marshal(Students)
	if err != nil {
		return err
	}
	err = os.WriteFile(confName, data, os.ModePerm)

	return err
}

func GetNames() []string {
	var names []string
	for _, student := range Students {
		names = append(names, student.Name)
	}
	return names
}

func GetIDs() []string {
	var IDs []string
	for _, student := range Students {
		IDs = append(IDs, student.ID)
	}
	return IDs
}

func FindStudentByID(studentID string) *Student {
	for _, student := range Students {
		if student.ID == studentID {
			return &student
		}
	}
	return nil
}
