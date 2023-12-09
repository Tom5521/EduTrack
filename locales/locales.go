package locales

type Locale struct {
	GradeInfo        map[string]string `yaml:"GradeInfo"`
	StudentGradeInfo map[string]string `yaml:"StudentGradeInfo"`
	StudentInfo      map[string]string `yaml:"StudentInfo"`
	MainMenu         map[string]string `yaml:"MainMenu"`
	WindowTitles     struct {
		StudentGradesWindows map[string]string `yaml:"StudentGradesWindows"`
		GradesWindows        map[string]string `yaml:"GradesWindows"`
		StudentWindows       map[string]string `yaml:"StudentWindows"`
		RecordsWindows       map[string]string `yaml:"RecordsWindows"`
		SearchWindows        map[string]string `yaml:"SearchWindows"`
	} `yaml:"WindowTitles"`
	Dialogs struct {
		Student map[string]string `yaml:"Student"`
	} `yaml:"Dialogs"`
}
