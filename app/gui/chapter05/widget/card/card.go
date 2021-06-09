package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Card")
	content := makeUI()
	w.SetContent(content)
	w.ShowAndRun()
}

func makeUI() fyne.CanvasObject {
	card1 := widget.NewCard("Card Title", "Subtitle", widget.NewLabel("Content"))
	card1.Image = canvas.NewImageFromResource(theme.FyneLogo())
	return card1
}
