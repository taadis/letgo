package main

import (
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hyperlink")
	w.SetContent(makeUI())
	w.ShowAndRun()
}

func makeUI() fyne.CanvasObject {
	link, _ := url.Parse("https://fyne.io")
	return widget.NewHyperlink("fyne.io", link)
}
