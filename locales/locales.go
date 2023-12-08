package locales

import "embed"

//go:embed files
var LocaleFiles embed.FS

type Locale struct {
	StudentInfo struct {
		Name        string `yaml:"Name"`
		Age         string `yaml:"Age"`
		DNI         string `yaml:"DNI"`
		PhoneNumber string `yaml:"Phone Number"`
	} `yaml:"StudentInfo"`
	MainMenu struct {
		File string `yaml:"File"`
		Edit string `yaml:"Edit"`
		Help string `yaml:"Help"`
	} `yaml:"MainMenu"`
	WindowTitles struct {
		StudentGradesWindows struct {
			SelectStartAndEnd string `yaml:"Select start and end"`
			DetailsOfX        string `yaml:"Details of X"`
			EditX             string `yaml:"Edit X"`
		} `yaml:"StudentGradesWindows"`
		GradesWindows struct {
			XGrades      string `yaml:"X Grades"`
			Grades       string `yaml:"Grades"`
			SelectAGrade string `yaml:"Select a grade"`
		} `yaml:"GradesWindows"`
		StudentWindows struct {
			Students string `yaml:"Students"`
		} `yaml:"StudentWindows"`
		RecordsWindows struct {
			Records string `yaml:"Records"`
		} `yaml:"RecordsWindows"`
		OtherWindows struct {
			Search string `yaml:"Search"`
		} `yaml:"OtherWindows"`
	} `yaml:"WindowTitles"`
	Dialogs struct {
		Student struct {
			DeleteStudentDialog string `yaml:"Delete Student Dialog"`
		} `yaml:"Student"`
	} `yaml:"Dialogs"`
}
