package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("ProgressBar")
	w.SetContent(makeUI())
	w.ShowAndRun()
}

func makeUI() fyne.CanvasObject {
	bar1 := widget.NewProgressBar()
	bar2 := widget.NewProgressBarInfinite()
	form := widget.NewForm(
		widget.NewFormItem("bar1", bar1),
		widget.NewFormItem("bar2", bar2),
	)
	return form
}
