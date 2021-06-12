package main

import (
	"fmt"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

const timeLayout = "2006-01-02 15:04:05"

func main() {
	a := app.New()
	w := a.NewWindow("Toolkit")
	w.SetContent(makeTimeUI())
	w.ShowAndRun()
}

func makeTimeUI() fyne.CanvasObject {
	tsEntry := widget.NewEntry()
	tsEntry.Validator = func(text string) error {
		i, err := strconv.ParseInt(text, 10, 0)
		if err != nil {
			return err
		}

		if i < 0 {
			return fmt.Errorf("invalid input")
		}

		return nil
	}
	dateTimeEntry := widget.NewEntry()

	form := widget.NewForm(
		widget.NewFormItem("Timestamp:", tsEntry),
		widget.NewFormItem("DateTime:", dateTimeEntry),
	)
	form.OnCancel = func() {
		tsEntry.SetText("")
		dateTimeEntry.SetText("")
	}
	form.OnSubmit = func() {
		i, _ := strconv.ParseInt(tsEntry.Text, 10, 0)
		s := time.Unix(i, 0).Format(timeLayout)
		dateTimeEntry.SetText(s)
	}

	return form
}
