package gui

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/EduTrack/assets"
	"github.com/Tom5521/EduTrack/internal/pkg/sizes"
	"github.com/Tom5521/EduTrack/internal/pkg/tools"
	"github.com/Tom5521/EduTrack/pkg/data"
	"github.com/Tom5521/EduTrack/pkg/wins"
)

func (ui ui) GetStudentsList(students *[]data.Student, onSelected func(id widget.ListItemID)) *widget.List {
	list := widget.NewList(
		func() int {
			return len(*students)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("--")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			s := *students
			o.(*widget.Label).SetText(s[i].Name)
		},
	)
	list.OnSelected = onSelected
	return list
}

func (ui *ui) LoadStudentInfo(s *data.Student) {
	var currentColor color.Color
	if fyne.CurrentApp().Settings().ThemeVariant() == 1 {
		currentColor = color.Black
	} else {
		currentColor = color.White
	}

	itemLabel := func(inputText any) *canvas.Text {
		text := fmt.Sprint(inputText)
		maxTextSize := 19
		if len(text) > maxTextSize {
			text = text[:maxTextSize] + "..."
		}
		label := canvas.NewText(text, currentColor)
		label.TextSize = 20
		return label
	}
	tagLabel := func(inputText any) *canvas.Text {
		text := fmt.Sprint(inputText)
		label := itemLabel(text)
		label.TextSize = 20
		label.TextStyle.Bold = true
		return label
	}

	var nhbx = container.NewHBox

	image := tools.LoadProfileImg(s.ImageFilePath)

	// Name Label
	nameLabel := tagLabel("Name: ")
	nameLabel.TextStyle.Bold = true
	nameCont := nhbx(nameLabel, itemLabel(s.Name))
	// Age Label
	ageCont := nhbx(tagLabel("Age:"), itemLabel(s.Age))
	// DNI Label
	dniCont := nhbx(tagLabel("DNI:"), itemLabel(s.DNI))
	// Phone Label
	phoneCont := nhbx(tagLabel("Phone number:"), itemLabel(s.PhoneNumber))

	dataContainer := container.NewVBox(nameCont, ageCont, dniCont, phoneCont)
	showRecordsBt := widget.NewButton("Show student records", func() {
		ui.StudentRecordsMainWin(s)
	})
	showGradesBt := widget.NewButton("Show student grades", func() {
		ui.StudentGradesMainWin(s)
	})

	const gridNumber int = 1
	content := container.NewVBox(image,
		dataContainer,
		container.NewAdaptiveGrid(gridNumber,
			showRecordsBt,
			showGradesBt,
		),
	)
	ui.StudentTab.Objects = []fyne.CanvasObject{content}
}

func (ui *ui) AddStudentForm() {
	var imagePath string
	var gradesStr []data.Grade
	window := ui.App.NewWindow("Add a student")
	window.Resize(sizes.FormSize)

	// Initialize form fields
	nameEntry := widget.NewEntry()
	ageEntry := widget.NewEntry()
	dniEntry := widget.NewEntry()
	phoneEntry := widget.NewEntry()

	imageButton := widget.NewButton("Select Image", func() {
		wins.ImagePicker(ui.App, &imagePath)
	})
	deleteImgBtn := widget.NewButton("Delete Current Image", func() {
		imagePath = ""
	})

	getStGrade := func() string {
		var grades string
		for _, g := range gradesStr {
			grades += g.Name + ","
		}
		return grades
	}

	studentGradesLabel := widget.NewLabel("")
	studentGradesLabel.SetText(getStGrade())
	gradeSelect := widget.NewSelect(data.GetGradesNames(), func(s string) {
		gradesStr = append(gradesStr, data.FindGradeByName(s))

		studentGradesLabel.SetText(getStGrade())
	})

	nameForm := widget.NewFormItem("Name:", nameEntry)
	idForm := widget.NewFormItem("DNI:", dniEntry)
	ageForm := widget.NewFormItem("Age:", ageEntry)
	phoneForm := widget.NewFormItem("Phone:", phoneEntry)
	gradeForm := widget.NewFormItem("Select Grades:", gradeSelect)
	gradesShowForm := widget.NewFormItem("Grades:", studentGradesLabel)
	const gridNumber int = 2
	imageForm := widget.NewFormItem("Image:", container.NewAdaptiveGrid(gridNumber, imageButton, deleteImgBtn))

	form := widget.NewForm(
		nameForm,
		idForm,
		ageForm,
		phoneForm,
		gradeForm,
		gradesShowForm,
		imageForm,
	)
	form.OnSubmit = func() {
		StGrades := func() []data.StudentGrade {
			var stgrades []data.StudentGrade
			for _, grade := range gradesStr {
				tmpGrade := data.StudentGrade{GradeID: grade.ID}
				stgrades = append(stgrades, tmpGrade)
			}
			return stgrades
		}()

		newStudent := data.Student{
			Name:          nameEntry.Text,
			Age:           uint(atoi(ageEntry.Text)),
			DNI:           dniEntry.Text,
			PhoneNumber:   phoneEntry.Text,
			ImageFilePath: imagePath,
			Grades:        StGrades,
		}
		// Validate form fields
		if !checkValues(newStudent) {
			wins.ErrWin(ui.App, "Some value in the form is empty")
			return
		}
		if existsDNI(dniEntry.Text, data.GetStudentDNIs()) {
			wins.ErrWin(ui.App, "The DNI already exists")
			return
		}

		// Add a new student

		err := data.AddStudent(&newStudent)
		if err != nil {
			wins.ErrWin(ui.App, err.Error())
		}
		ui.StudentList.Refresh()
		s := data.Students[data.FindStudentIndexByID(newStudent.ID)]
		ui.LoadStudentInfo(&s)
		window.Close()
	}
	form.OnCancel = func() {
		window.Close()
	}
	window.SetContent(form)
	window.Show()
}

func (ui *ui) EditFormWindow(s *data.Student) {
	window := ui.App.NewWindow("Edit " + s.Name)
	window.Resize(sizes.FormSize)

	// Initialize form fields
	var imagePath = s.ImageFilePath
	nameEntry := widget.NewEntry()
	nameEntry.SetText(s.Name)
	ageEntry := widget.NewEntry()
	ageEntry.SetText(itoa(s.Age))
	dniEntry := widget.NewEntry()
	dniEntry.SetText(s.DNI)
	phoneEntry := widget.NewEntry()
	phoneEntry.SetText(s.PhoneNumber)

	imageLabel := widget.NewLabel(imagePath)

	nameForm := widget.NewFormItem("Name:", nameEntry)
	ageForm := widget.NewFormItem("Age:", ageEntry)
	dniForm := widget.NewFormItem("DNI:", dniEntry)

	phoneForm := widget.NewFormItem("Phone:", phoneEntry)

	deleteImgButton := widget.NewButton("Delete selected image", func() {
		imagePath = ""
		imageLabel.SetText(imagePath)
	})
	selectImgButton := widget.NewButton("Select student image", func() {
		wins.ImagePicker(ui.App, &imagePath)
		imageLabel.SetText(imagePath)
	})

	const gridNumber int = 2
	imgCont := container.NewAdaptiveGrid(gridNumber, deleteImgButton, selectImgButton)
	imgForm := widget.NewFormItem("", imgCont)
	imgInfoForm := widget.NewFormItem("Image path:", imageLabel)

	form := widget.NewForm(
		nameForm,
		ageForm,
		dniForm,
		phoneForm,
		imgForm,
		imgInfoForm,
	)
	form.OnSubmit = func() {
		edited := data.Student{
			Name:          nameEntry.Text,
			Age:           uint(atoi(ageEntry.Text)),
			DNI:           dniEntry.Text,
			PhoneNumber:   phoneEntry.Text,
			ImageFilePath: imagePath,
		}
		// Validate form fields
		if !checkValues(edited) {
			wins.ErrWin(ui.App, "Some value in the form is empty")
			return
		}
		if dniEntry.Text != s.DNI {
			if existsDNI(dniEntry.Text, data.GetStudentDNIs()) {
				wins.ErrWin(ui.App, "The DNI already exists")
				return
			}
		}
		err := s.Edit(&edited)
		if err != nil {
			wins.ErrWin(ui.App, err.Error())
		}
		ui.StudentList.UnselectAll()
		ui.LoadStudentInfo(&data.Students[data.FindStudentIndexByID(s.ID)])
		window.Close()
	}
	form.OnCancel = func() {
		window.Close()
	}

	window.SetContent(form)
	window.Show()
}

func (ui *ui) DeleteStudentWin(s *data.Student) {
	window := ui.App.NewWindow("Delete Student")
	const gridNumber int = 2
	content := container.NewVBox(
		widget.NewLabel("Are you sure you want to delete the student?"),
		container.NewAdaptiveGrid(gridNumber,
			widget.NewButton("Yes", func() {
				err := data.Delete(s)
				if err != nil {
					wins.ErrWin(ui.App, err.Error())
				}
				ui.StudentList.UnselectAll()
				ui.StudentTab = ui.GetTemplateUser()
				ui.StudentTab.Refresh()
				window.Close()
			}),
			widget.NewButton("No", func() {
				window.Close()
			}),
		))
	window.SetContent(content)
	window.Show()
}

func (ui ui) GetTemplateUser() *fyne.Container {
	image := canvas.NewImageFromResource(assets.UserTemplate)
	image.SetMinSize(sizes.ProfileSize)
	dataForm := widget.NewForm(
		widget.NewFormItem("Name:", widget.NewLabel("--")),
		widget.NewFormItem("Age:", widget.NewLabel("--")),
		widget.NewFormItem("DNI:", widget.NewLabel("--")),
		widget.NewFormItem("Phone number:", widget.NewLabel("--")),
	)
	content := container.NewVBox(image, dataForm)
	return content
}
