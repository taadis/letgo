package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Toolkit")

	helloLabel := widget.NewLabel("Hello World!")
	opButton := widget.NewButton("Hi!", func() {
		helloLabel.SetText("Hello Toolkit!")
	})
	c := container.NewVBox(helloLabel, opButton)
	w.SetContent(c)
	w.ShowAndRun()
}
