package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type InstallerLayout struct {
	widget.BaseWidget

	TopItems  []fyne.CanvasObject
	ButtonBox *fyne.Container

	TopItemsLayout fyne.Layout

	OnNext   func()
	OnBack   func()
	OnCancel func()

	BackText   string
	NextText   string
	CancelText string

	HideBack   bool
	HideNext   bool
	HideCancel bool

	DisableNext   bool
	DisableBack   bool
	DisableCancel bool

	nextButton   *widget.Button
	backButton   *widget.Button
	cancelButton *widget.Button
}

func (l *InstallerLayout) updateButtons() {
	// Set text.
	if l.NextText == "" {
		l.nextButton.SetText("Next")
	} else {
		l.nextButton.SetText(l.NextText)
	}
	if l.BackText == "" {
		l.backButton.SetText("Back")
	} else {
		l.backButton.SetText(l.BackText)
	}
	if l.CancelText == "" {
		l.cancelButton.SetText("Cancel")
	} else {
		l.cancelButton.SetText(l.CancelText)
	}

	// Disablers.
	if l.DisableNext {
		l.nextButton.Disable()
	} else {
		l.nextButton.Enable()
	}
	if l.DisableBack {
		l.backButton.Disable()
	} else {
		l.backButton.Enable()
	}
	if l.DisableCancel {
		l.cancelButton.Disable()
	} else {
		l.cancelButton.Enable()
	}
	// Hidders.
	if l.HideBack {
		l.backButton.Hide()
	} else {
		l.backButton.Show()
	}

	if l.HideNext {
		l.nextButton.Hide()
	} else {
		l.nextButton.Show()
	}
	if l.HideCancel {
		l.cancelButton.Hide()
	} else {
		l.cancelButton.Show()
	}
}

func NewInstallerLayout(items ...fyne.CanvasObject) *InstallerLayout {
	l := &InstallerLayout{
		TopItems: items,
	}

	l.ExtendBaseWidget(l)
	return l
}

func (l *InstallerLayout) Update() {
	l.updateButtons()
	l.Refresh()
}

func (l *InstallerLayout) CreateRenderer() fyne.WidgetRenderer {
	l.nextButton = &widget.Button{OnTapped: l.OnNext, Importance: widget.HighImportance}
	l.backButton = &widget.Button{OnTapped: l.OnBack}
	l.cancelButton = &widget.Button{OnTapped: l.OnCancel}

	l.ButtonBox = container.NewBorder(
		nil, nil,
		container.NewBorder(nil, nil, l.backButton, nil),
		container.NewBorder(nil, nil, nil, l.nextButton),
		container.NewBorder(nil, nil, nil, l.cancelButton),
	)
	if l.HideBack && l.HideNext && l.HideCancel {
		l.ButtonBox.Hide()
	} else {
		l.ButtonBox.Show()
	}

	if l.TopItemsLayout == nil {
		l.TopItemsLayout = layout.NewStackLayout()
	}

	c := container.NewBorder(
		nil, l.ButtonBox, nil, nil,
		&fyne.Container{Layout: l.TopItemsLayout, Objects: l.TopItems},
	)
	l.updateButtons()
	return widget.NewSimpleRenderer(c)
}
