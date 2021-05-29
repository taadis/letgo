package main

import (
	"image/color"
	"net/url"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("TODO List")
	w.SetContent(makeUI())
	w.ShowAndRun()
}

func makeUI() fyne.CanvasObject {
	length := func() int {
		return 5
	}
	createItem := func() fyne.CanvasObject {
		b := widget.NewCheck("TODO Item x", func(b bool) {

		})
		return b
	}
	updateItem := func(widget.ListItemID, fyne.CanvasObject) {}
	todos := widget.NewList(length, createItem, updateItem)

	details := widget.NewForm(
		widget.NewFormItem("Title", widget.NewEntry()),
		widget.NewFormItem("Description", widget.NewMultiLineEntry()),
		widget.NewFormItem("Category", widget.NewSelect([]string{"Home"}, func(s string) {})),
		widget.NewFormItem("Priority", widget.NewRadioGroup([]string{"Low", "Mid", "High"}, func(string) {})),
		widget.NewFormItem("Due", widget.NewEntry()),
		widget.NewFormItem("Completion", widget.NewSlider(0, 100)),
	)

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {}),
	)

	return container.NewBorder(
		toolbar,
		nil,
		todos,
		nil,
		details,
	)
}

func main8() {
	a := app.New()
	w := a.NewWindow("Manual layout - 手动布局")
	square := canvas.NewRectangle(color.Black)
	square.Move(fyne.NewPos(10, 10))
	square.Resize(fyne.NewSize(90, 90))
	circle := canvas.NewCircle(color.Transparent)
	circle.StrokeColor = color.Gray{Y: 128}
	circle.StrokeWidth = 5
	circle.Move(fyne.NewPos(70, 70))
	circle.Resize(fyne.NewSize(120, 120))
	box := container.NewWithoutLayout(square, circle)
	box.Resize(fyne.NewSize(200, 200))
	w.SetContent(box)
	w.ShowAndRun()
}

func main7() {
	a := app.New()
	w := a.NewWindow("Main")
	w.SetContent(widget.NewLabel("Main"))
	go showAnother(a)
	w.ShowAndRun()
}

func showAnother(a fyne.App) {
	time.Sleep(5 * time.Second)

	w := a.NewWindow("Shown later")
	w.SetContent(widget.NewLabel("5 seconds later"))
	w.Resize(fyne.NewSize(200, 200))
	w.Show()
	time.Sleep(2 * time.Second)
	w.Close()
}

func main6() {
	a := app.New()
	w := a.NewWindow("Hyperlink - 超链接")
	bugURL, _ := url.Parse("https://github.com/fyne-io/fyne/issue/new")
	w.SetContent(widget.NewHyperlink("Report a bug", bugURL))
	w.ShowAndRun()
}

func main5() {
	a := app.New()
	w := a.NewWindow("Gradient - 渐变")
	//c := canvas.NewHorizontalGradient(color.White, color.Black)
	c := canvas.NewVerticalGradient(color.White, color.Black)
	w.SetContent(c)
	w.ShowAndRun()
}

func generate(x, y, w, h int) color.Color {
	if (x/20)%2 == (y/20)%2 {
		return color.White
	}
	return color.Black
}

func main4() {
	a := app.New()
	w := a.NewWindow("Toolkit")
	w.SetContent(canvas.NewRasterWithPixels(generate))
	w.ShowAndRun()
}

func main3() {
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
