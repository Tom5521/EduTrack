package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/EduTrack/assets"
	"github.com/Tom5521/EduTrack/internal/pkg/sizes"
	"github.com/Tom5521/EduTrack/pkg/data"
	"github.com/Tom5521/EduTrack/pkg/wins"
)

func (ui *ui) SearchStudentsMainWin() {
	w := ui.App.NewWindow(po.Get("Search Students"))
	w.Resize(sizes.SearchSize)
	dniEntry := widget.NewEntry()
	label := widget.NewLabel(po.Get("Enter Student name:"))
	label.Alignment = fyne.TextAlignCenter
	dniSearchButton := widget.NewButton(po.Get("Search"), func() {
		studentDNI := dniEntry.Text
		i := data.StudentIndexByDNI(studentDNI)
		if i == -1 {
			wins.ErrWin(ui.App, po.Get("Student not found!"))
			return
		}
		student := data.Students[i]
		ui.LoadStudentInfo(&student)
	})

	dniLabel := widget.NewLabel(po.Get("Enter Student DNI:"))
	dniLabel.Alignment = fyne.TextAlignCenter

	byDNIContainer := container.NewVBox(
		dniLabel,
		dniEntry,
		dniSearchButton,
	)

	nameEntry := widget.NewEntry()
	nameSearchButton := widget.NewButton(po.Get("Search"), func() {
		studentName := nameEntry.Text
		i := data.StudentIndexByName(studentName)
		if i == -1 {
			wins.ErrWin(ui.App, po.Get("Student not found!"))
			return
		}
		student := data.Students[i]
		ui.LoadStudentInfo(&student)
	})

	byNameContainer := container.NewVBox(
		label,
		nameEntry,
		nameSearchButton,
	)

	content := container.NewAppTabs(
		container.NewTabItem(po.Get("Search by name"), byNameContainer),
		container.NewTabItem(po.Get("Search by DNI"), byDNIContainer),
	)

	w.SetContent(content)
	w.Show()
}

func (ui *ui) SearchCoursesMainWin() {
	w := ui.App.NewWindow(po.Get("Search Courses"))
	w.Resize(sizes.SearchSize)

	label := widget.NewLabel(po.Get("Enter a course name:"))
	label.Alignment = fyne.TextAlignCenter
	searchCourseEntry := widget.NewEntry()
	searchCourseButton := widget.NewButton(po.Get("Search"), func() {
		text := searchCourseEntry.Text
		i := data.CourseIndexbyName(text)
		if i == -1 {
			wins.ErrWin(ui.App, po.Get("Course not found!"))
			return
		}
		ui.CourseDetailsWin(&data.Courses[i])
	})

	content := container.NewVBox(
		label,
		searchCourseEntry,
		searchCourseButton,
	)
	w.SetContent(content)

	w.Show()
}

func (ui *ui) SearchRecordsMainWin() {
	w := ui.App.NewWindow(po.Get("Search Records"))
	w.Resize(sizes.SearchSize)
	label := widget.NewLabel(po.Get("Enter a record name:"))
	label.Alignment = fyne.TextAlignCenter
	entry := widget.NewEntry()
	button := widget.NewButton(po.Get("Search"), func() {
		text := entry.Text
		i := data.RecordIndexByName(text)
		if i == -1 {
			wins.ErrWin(ui.App, po.Get("Record not found!"))
			return
		}
		record := data.Records[i]
		ui.EditRecordData(record.ID)
	})
	content := container.NewVBox(
		label,
		entry,
		button,
	)

	w.SetContent(content)
	w.Show()
}

func (ui *ui) SelectStudentWin(onConfirm func(selected int)) {
	var selectetList = -1
	w := ui.App.NewWindow(po.Get("Select student"))
	w.Resize(sizes.FormSize)
	list := ui.GetStudentsList(&data.Students)
	list.OnSelected = func(id widget.ListItemID) {
		selectetList = id
	}

	selectBt := widget.NewButton(po.Get("Select"), func() {
		if selectetList == -1 {
			return
		}
		onConfirm(selectetList)
		w.Close()
	})
	mainBox := container.NewBorder(nil, container.NewCenter(selectBt), nil, nil, list)
	w.SetContent(mainBox)
	w.Show()
}

func (ui *ui) SearchStudentCoursesMainWin() {
	w := ui.App.NewWindow(po.Get("Search student courses"))
	w.Resize(sizes.SearchSize)

	var student data.Student

	selectStudentBt := widget.NewButton(po.Get("Select student"), func() {
		ui.SelectStudentWin(func(id int) {
			student = data.Students[id]
		})
	})
	label := widget.NewLabel(po.Get("Enter course name to search"))
	label.Alignment = fyne.TextAlignCenter
	entry := widget.NewEntry()
	searchButton := widget.NewButton(po.Get("Search"), func() {
		i := student.CourseIndexByName(entry.Text)
		if i == -1 {
			wins.ErrWin(ui.App, po.Get("Student Course not found!"))
			return
		}
		ui.StudentCourseDetailsWin(&data.StudentCourses[i])
	})
	mainBox := container.NewVBox(
		selectStudentBt,
		label,
		entry,
		searchButton,
	)
	w.SetContent(mainBox)
	w.Show()
}

func (ui *ui) SearchMainWin() {
	const size1, size2 float32 = 300, 400
	w := ui.App.NewWindow(po.Get("Search"))
	w.Resize(fyne.NewSize(size1, size2))
	w.SetIcon(assets.Lens1)
	const gridNumber int = 1

	searchStudentsBt := widget.NewButtonWithIcon(po.Get("Search in students"), assets.Lens1, func() {
		ui.SearchStudentsMainWin()
	})
	searchCoursesBt := widget.NewButtonWithIcon(po.Get("Search in courses"), assets.Lens1, func() {
		ui.SearchCoursesMainWin()
	})
	searchRecordsBt := widget.NewButtonWithIcon(po.Get("Search in records"), assets.Lens1, func() {
		ui.SearchRecordsMainWin()
	})
	searchStudentCoursesBt := widget.NewButtonWithIcon(po.Get("Search in student courses"), assets.Lens1, func() {
		ui.SearchStudentCoursesMainWin()
	})
	content := container.NewAdaptiveGrid(gridNumber,
		searchStudentsBt,
		searchCoursesBt,
		searchRecordsBt,
		searchStudentCoursesBt,
	)

	w.SetContent(content)

	w.Show()
}
