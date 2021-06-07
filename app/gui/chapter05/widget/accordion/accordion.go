package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Accordion")
	acc := widget.NewAccordion(
		widget.NewAccordionItem("A", widget.NewLabel("Hidden")),
		widget.NewAccordionItem("B", widget.NewLabel("Shown item")),
		widget.NewAccordionItem("C", widget.NewLabel("End")),
	)
	acc.Items[1].Open = true
	w.SetContent(acc)
	w.ShowAndRun()
}
