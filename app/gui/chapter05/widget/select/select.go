package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Select")
	w.SetContent(makeUI())
	w.ShowAndRun()
}

func makeUI() fyne.CanvasObject {
	sel := widget.NewSelect([]string{"Item 1", "Item 2"}, func(s string) {
		fmt.Println("Selected", s)
	})
	return sel
}
