package gui

import (
	"fmt"
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/EduTrack/assets"
	"github.com/Tom5521/EduTrack/internal/pkg/sizes"
	"github.com/Tom5521/EduTrack/internal/pkg/tools"
	"github.com/Tom5521/EduTrack/pkg/conf"
	"github.com/Tom5521/EduTrack/pkg/data"
	"github.com/Tom5521/EduTrack/pkg/widgets"
	"github.com/Tom5521/EduTrack/pkg/wins"
)

type StudentForm struct {
	Edit struct {
		Enable  bool
		Student *data.Student
	}
	Add bool
}

func (ui ui) GetStudentsList(students *[]data.Student) *widget.List {
	list := widget.NewList(
		func() int {
			return len(*students)
		},
		func() fyne.CanvasObject {
			return &widget.Label{}
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			s := *students
			o.(*widget.Label).SetText(s[i].Name)
		},
	)
	return list
}

func (ui *ui) LoadStudentInfo(s *data.Student) {
	var currentColor color.Color
	if fyne.CurrentApp().Settings().ThemeVariant() == 1 {
		currentColor = color.Black
	} else {
		currentColor = color.White
	}

	itemLabel := func(inputText any, limit int) *canvas.Text {
		text := fmt.Sprint(inputText)
		if len(text) > limit {
			text = text[:limit] + "..."
		}
		label := canvas.NewText(text, currentColor)
		if conf.Config.Lang == "es" {
			label.TextSize = 18
		}
		if conf.Config.Lang == "en" {
			label.TextSize = 20
		}
		return label
	}
	tagLabel := func(inputText any) *canvas.Text {
		text := fmt.Sprint((inputText))
		label := itemLabel(po.Get(text), 90)
		label.TextSize = 20
		label.TextStyle.Bold = true
		return label
	}

	var nhbx = container.NewHBox

	image := tools.LoadProfileImg(s.ImageFilePath)

	// Name Label
	nameCont := nhbx(tagLabel("Name:"), itemLabel(s.Name, 21))
	// Age Label
	ageCont := nhbx(tagLabel("Age:"), itemLabel(s.Age, 20))
	// DNI Label
	dniCont := nhbx(tagLabel("DNI:"), itemLabel(s.DNI, 20))
	// Phone Label
	phoneCont := nhbx(tagLabel("Phone Number:"), itemLabel(s.PhoneNumber, 20))

	dataContainer := container.NewVBox(nameCont, ageCont, dniCont, phoneCont)
	showRecordsBt := widget.NewButton(po.Get("Show Student Records"), func() {
		ui.StudentRecordsMainWin(s)
	})
	showCoursesBt := widget.NewButton(po.Get("Show Student Courses"), func() {
		ui.StudentCoursesMainWin(s)
	})

	const gridNumber int = 1
	content := container.NewVBox(
		image,
		dataContainer,
		container.NewAdaptiveGrid(
			gridNumber,
			showRecordsBt,
			showCoursesBt,
		),
	)
	ui.StudentTab.Objects = []fyne.CanvasObject{content}
}

func getAgeEntry(app fyne.App, ageEntry *widget.Entry) uint {
	text := ageEntry.Text
	ret := uint(atoi(text))
	if (text != "0") && (ret == math.MaxUint) || (atoi(text) == -1) {
		wins.ErrWin(app, po.Get("Overflow in age entry"))
		ret = math.MaxUint
	}
	return ret
}

func (ui *ui) StudentForm(c StudentForm) {
	studentToEdit := c.Edit.Student
	var imagePath = func() string {
		if c.Edit.Enable {
			return studentToEdit.ImageFilePath
		}
		return ""
	}()
	getWinTitle := func() string {
		var winTitle string
		if c.Add {
			winTitle = po.Get("Add a student")
		}
		if c.Edit.Enable {
			winTitle = po.Get("Edit %s", studentToEdit.Name)
		}
		return winTitle
	}
	w := ui.App.NewWindow(getWinTitle())
	w.Resize(sizes.FormSize)

	// Initialize form fields
	nameEntry := widget.NewEntry()
	ageEntry := widget.NewEntry()
	dniEntry := widget.NewEntry()
	phoneEntry := widget.NewEntry()

	imagePathLabel := widget.NewLabel(imagePath)
	if c.Edit.Enable {
		nameEntry.SetText(studentToEdit.Name)
		ageEntry.SetText(itoa(studentToEdit.Age))
		dniEntry.SetText(studentToEdit.DNI)
		phoneEntry.SetText(studentToEdit.PhoneNumber)
	}

	imageButton := widget.NewButton(po.Get("Select Image"), func() {
		wins.ImagePicker(&imagePath)
		imagePathLabel.SetText(imagePath)
	})
	deleteImgBtn := widget.NewButton(po.Get("Delete Current Image"), func() {
		imagePath = ""
		imagePathLabel.SetText(imagePath)
	})

	const gridNumber int = 2
	form := widgets.NewForm(
		&widget.FormItem{Text: po.Get("Name:"), Widget: nameEntry, HintText: po.Get("First and last name")},
		widget.NewFormItem(po.Get("DNI:"), dniEntry),
		widget.NewFormItem(po.Get("Age:"), ageEntry),
		widget.NewFormItem(po.Get("Phone Number:"), phoneEntry),
		widget.NewFormItem(po.Get("Image Path"), imagePathLabel),
	)
	form.CustomItems = container.NewAdaptiveGrid(gridNumber, imageButton, deleteImgBtn)
	form.SubmitText = po.Get("Submit")
	form.CancelText = po.Get("Cancel")

	form.OnCancel = func() {
		w.Close()
	}

	form.OnSubmit = func() {
		n := data.Student{
			Name:          nameEntry.Text,
			Age:           getAgeEntry(ui.App, ageEntry),
			DNI:           dniEntry.Text,
			PhoneNumber:   phoneEntry.Text,
			ImageFilePath: imagePath,
		}
		if !checkValues(n) {
			wins.ErrWin(ui.App, po.Get("Some value in this form is empty"))
			return
		}
		if c.Edit.Enable {
			if dniEntry.Text != studentToEdit.DNI {
				if existsDNI(dniEntry.Text, data.GetStudentDNIs()) {
					wins.ErrWin(ui.App, po.Get("The DNI already exists"))
					return
				}
			}
			err := studentToEdit.Edit(&n)
			if err != nil {
				wins.ErrWin(ui.App, err.Error())
			}
		}
		if c.Add {
			if existsDNI(dniEntry.Text, data.GetStudentDNIs()) {
				wins.ErrWin(ui.App, po.Get("The DNI already exists"))
				return
			}
			// Add a new student
			err := data.AddStudent(&n)
			if err != nil {
				wins.ErrWin(ui.App, err.Error())
			}
		}
		ui.StudentList.Refresh()
		var idToLoad uint
		if c.Add {
			idToLoad = n.ID
		}
		if c.Edit.Enable {
			idToLoad = studentToEdit.ID
		}
		i := data.FindStudentIndexByID(idToLoad)
		if i == -1 {
			wins.ErrWin(ui.App, po.Get("Student not found!"))
			return
		}
		s := data.Students[i]
		ui.LoadStudentInfo(&s)
		w.Close()
	}
	w.SetContent(form)
	w.Show()
}

func (ui *ui) StudentDetailsWin(s *data.Student) {
	w := ui.App.NewWindow(po.Get("Details of %s", s.Name))
	w.Resize(sizes.FormSize)

	form := widgets.NewForm(
		widget.NewFormItem(po.Get("Name:"), widget.NewLabel(s.Name)),
		widget.NewFormItem(po.Get("Age:"), widget.NewLabel(itoa(s.Age))),
		widget.NewFormItem(po.Get("DNI:"), widget.NewLabel(s.DNI)),
		widget.NewFormItem(po.Get("Phone Number:"), widget.NewLabel(s.PhoneNumber)),
	)
	form.CustomItems = container.NewVBox(
		widget.NewButton(po.Get("Show student courses"), func() { ui.StudentCoursesMainWin(s) }),
		widget.NewButton(po.Get("Show student records"), func() { ui.StudentRecordsMainWin(s) }),
	)

	form.SubmitText = po.Get("Close")
	form.OnSubmit = func() {
		w.Close()
	}
	w.SetContent(form)
	w.Show()
}

func (ui ui) GetTemplateUser() *fyne.Container {
	image := canvas.NewImageFromResource(assets.UserTemplate)
	image.SetMinSize(sizes.ProfileSize)

	content := container.NewVBox(image)
	return content
}
