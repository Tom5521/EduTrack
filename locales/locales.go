package locales

type Locale struct {
	GeneralWords      map[string]string `yaml:"GeneralWords"`
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
		AboutWindows          map[string]string `yaml:"AboutWindow"`
	} `yaml:"WindowTitles"`
	Dialogs struct {
		General map[string]string `yaml:"General"`
		Student map[string]string `yaml:"Student"`
		Records map[string]string `yaml:"Records"`
	} `yaml:"Dialogs"`
	Buttons struct {
		MainWin               map[string]string `yaml:"MainWindow"`
		AddStudentWindow      map[string]string `yaml:"AddStudentWindow"`
		EditStudentWindow     map[string]string `yaml:"EditStudentWindow"`
		RecordWindow          map[string]string `yaml:"RecordWindow"`
		StudentRecordsMainWin map[string]string `yaml:"StudentRecordsMainWin"`
		SelectCourseWin       map[string]string `yaml:"SelectCourseWin"`
	} `yaml:"Buttons"`
	Labels struct {
		AddRecordWindow  map[string]string `yaml:"AddRecordWindow"`
		EditRecordWindow map[string]string `yaml:"EditRecordWindow"`
		RecordDetailsWin map[string]string `yaml:"RecordDetailsWin"`
	} `yaml:"Labels"`
	Errors map[string]string `yaml:"Errors"`
}
