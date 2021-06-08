package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Button")
	button := widget.NewButtonWithIcon("Cancel", theme.CancelIcon(), func() {

	})
	button1 := widget.NewButtonWithIcon("Cancel", theme.CancelIcon(), func() {

	})
	button1.Importance = widget.HighImportance
	button2 := widget.NewButtonWithIcon("Cancel", theme.CancelIcon(), func() {

	})
	button2.Importance = widget.MediumImportance
	button3 := widget.NewButtonWithIcon("Cancel", theme.CancelIcon(), func() {

	})
	button3.Importance = widget.LowImportance
	w.SetContent(container.NewHBox(button, button1, button2, button3))
	w.ShowAndRun()
}
