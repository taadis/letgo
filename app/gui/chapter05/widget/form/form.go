package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Form")
	w.SetContent(makeUI())
	w.ShowAndRun()
}

func makeUI() fyne.CanvasObject {
	form := widget.NewForm(
		widget.NewFormItem("Username", widget.NewEntry()),
		widget.NewFormItem("Password", widget.NewPasswordEntry()),
	)
	form.OnCancel = func() {
		fmt.Println("Cancelled")
	}
	form.OnSubmit = func() {
		fmt.Println("Form Submitted")
	}
	return form
}
