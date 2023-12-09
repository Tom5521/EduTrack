package locales

type Locale struct {
	CourseInfo        map[string]string `yaml:"CourseInfo"`
	StudentCourseInfo map[string]string `yaml:"StudentCourseInfo"`
	StudentInfo       map[string]string `yaml:"StudentInfo"`
	MainMenu          map[string]string `yaml:"MainMenu"`
	WindowTitles      struct {
		StudentCoursesWindows map[string]string `yaml:"StudentCoursesWindows"`
		CoursesWindows        map[string]string `yaml:"CoursesWindows"`
		StudentWindows        map[string]string `yaml:"StudentWindows"`
		RecordsWindows        map[string]string `yaml:"RecordsWindows"`
		SearchWindows         map[string]string `yaml:"SearchWindows"`
	} `yaml:"WindowTitles"`
	Dialogs struct {
		Student map[string]string `yaml:"Student"`
	} `yaml:"Dialogs"`
}