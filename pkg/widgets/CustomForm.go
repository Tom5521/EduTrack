package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type CustomForm struct {
	widget.BaseWidget
	FormItems   []*widget.FormItem
	CustomItems *fyne.Container
	ButtonBox   *fyne.Container
	OnSubmit    func()
	OnCancel    func()
	SubmitText  string
	CancelText  string

	submitButton *widget.Button
	cancelButton *widget.Button
}

func (f *CustomForm) updateButtons() {
	if f.SubmitText == "" {
		f.SubmitText = "Submit"
	}
	if f.CancelText == "" {
		f.CancelText = "Cancel"
	}
	if f.OnCancel == nil {
		f.cancelButton.Hide()
	} else {
		f.cancelButton.SetText(f.CancelText)
		f.cancelButton.OnTapped = f.OnCancel
		f.cancelButton.Show()
	}
	if f.OnSubmit == nil {
		f.submitButton.Hide()
	} else {
		f.submitButton.SetText(f.SubmitText)
		f.submitButton.OnTapped = f.OnSubmit
		f.Show()
	}
	if f.OnCancel == nil && f.OnSubmit == nil {
		f.ButtonBox.Hide()
	} else {
		f.ButtonBox.Show()
	}
}

func NewForm(items ...*widget.FormItem) *CustomForm {
	f := &CustomForm{}
	f.FormItems = items

	f.ExtendBaseWidget(f)
	return f
}

func (f *CustomForm) CreateRenderer() fyne.WidgetRenderer {
	f.cancelButton = &widget.Button{OnTapped: f.OnCancel}
	f.submitButton = &widget.Button{OnTapped: f.OnSubmit, Importance: widget.HighImportance}
	if f.ButtonBox == nil {
		buttons := container.NewGridWithRows(1, f.cancelButton, f.submitButton)
		f.ButtonBox = container.NewBorder(nil, nil, nil, buttons)
	} else {
		f.ButtonBox.Show()
	}

	customItems := container.NewStack(f.CustomItems)
	if f.CustomItems == nil {
		customItems.Hide()
	} else {
		customItems.Show()
	}

	c := container.NewVBox(
		widget.NewForm(f.FormItems...),
		customItems,
		f.ButtonBox,
	)
	f.updateButtons()
	return widget.NewSimpleRenderer(c)
}
