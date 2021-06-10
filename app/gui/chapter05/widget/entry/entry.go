package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Entry")
	w.SetContent(makeUI())
	w.ShowAndRun()
}

func makeUI() fyne.CanvasObject {
	entry1 := widget.NewEntry()
	entry1.PlaceHolder = "Entry"
	entry2 := widget.NewPasswordEntry()
	entry2.PlaceHolder = "PasswordEntry"
	entry3 := widget.NewMultiLineEntry()
	entry3.PlaceHolder = "MultiLineEntry"
	return container.NewVBox(entry1, entry2, entry3)
}
