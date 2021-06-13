package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Label")
	w.SetContent(makeUI())
	w.ShowAndRun()
}

func makeUI() fyne.CanvasObject {
	label := widget.NewLabel("Text label")
	return label
}
