package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Toolkit")
	//w.SetContent(canvas.NewText("Text", color.Black))
	w.SetContent(widget.NewIcon(theme.ContentCopyIcon()))
	w.ShowAndRun()
}

func makeSign() fyne.CanvasObject {
	bg := canvas.NewCircle(color.NRGBA{R: 255, A: 255})
	bg.StrokeColor = color.White
	bg.StrokeWidth = 5
	bg.Resize(fyne.NewSize(100, 100))
	bg.Move(fyne.NewPos(10, 10))

	bar := canvas.NewRectangle(color.White)
	bar.Resize(fyne.NewSize(80, 20))
	bar.Move(fyne.NewPos(20, 50))

	c := container.NewWithoutLayout(bg, bar)
	return c
}

func main2() {
	a := app.New()
	w := a.NewWindow("Toolkit")
	w.SetContent(makeSign())
	w.SetPadded(false)
	w.Resize(fyne.NewSize(120, 120))
	w.ShowAndRun()
}

func main1() {
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
