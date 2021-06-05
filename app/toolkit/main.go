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
	data := dummyData()
	tasks := taskApp{data: data, visible: data.remaining()}
	w.SetContent(tasks.makeUI())
	if len(data.remaining()) > 0 {
		tasks.setTask(data.remaining()[0])
	}
	w.ShowAndRun()
}

func formatDate(date *time.Time) string  {
	if date == nil {
		return ""
	}
	return date.Format("02 Jan 06 15:04")
}

func dateValidator(s string) error {
	_, err := time.Parse("02 Jan 06 15:04", s)
	return err
}

type taskApp struct {
	data    *taskList
	visible []*task
	current *task

	title       *widget.Entry
	tasks       *widget.List
	description *widget.Entry
	due         *widget.Entry
	category    *widget.Select
	priority    *widget.RadioGroup
	completion  *widget.Slider
	// more will be added here
}

// refreshData 刷新数据
func (a *taskApp) refreshData()  {
	// 隐藏已完成的任务
	a.visible = a.data.remaining()
	a.tasks.Refresh()
}

func (a *taskApp) setTask(t *task) {
	a.current = t
	a.title.SetText(t.title)
	a.description.SetText(t.description)
	a.category.SetSelected(t.category)
	if t.priority == midPriority {
		a.priority.SetSelected("Mid")
	} else if t.priority == highPriority {
		a.priority.SetSelected("High")
	} else {
		a.priority.SetSelected("Low")
	}
	a.due.SetText(formatDate(t.due))
	a.completion.SetValue(t.completion)
}

func (a *taskApp) makeUI() fyne.CanvasObject {
	length := func() int {
		length := len(a.visible)
		return length
	}
	createItem := func() fyne.CanvasObject {
		b := widget.NewCheck("TODO Item x", func(b bool) {

		})
		return b
	}
	updateItem := func(i widget.ListItemID, c fyne.CanvasObject) {
		check := c.(*widget.Check)
		check.Text = a.visible[i].title
		check.OnChanged = func(done bool) {
			a.visible[i].done = done
			a.refreshData()
		}
		check.Refresh()
	}
	todos := widget.NewList(length, createItem, updateItem)

	a.tasks.OnSelected = func(id widget.ListItemID) {
		a.setTask(a.visible[id])
	}

	a.title = widget.NewEntry()
	a.title.OnChanged = func(s string) {
		if a.current == nil {
			return
		}

		a.current.title = s
		a.tasks.Refresh() // refresh list of titles
	}
	a.description = widget.NewMultiLineEntry()
	a.category = widget.NewSelect([]string{"Home"}, func(string) {})
	a.priority = widget.NewRadioGroup([]string{"Low", "Mid", "High"}, func(pri string) {
		if a.current == nil {
			return
		}

		if pri == "Mid" {
			a.current.priority = midPriority
		} else if pri == "High" {
			a.current.priority = highPriority
		} else {
			a.current.priority = lowPriority
		}
	})
	a.due = widget.NewEntry()
	a.due.Validator = dateValidator
	a.due.OnChanged = func(s string) {
		if a.current == nil {
			return
		}

		if s == "" {
			a.current.due = nil
		} else {
			date, err := time.Parse("02 Jan 06 15:04", s)
			if err != nil {
				a.current.due = &date
			}
		}
	}
	a.completion = widget.NewSlider(0, 100)

	details := widget.NewForm(
		widget.NewFormItem("Title", a.title),
		widget.NewFormItem("Description", a.description),
		widget.NewFormItem("Category", a.category),
		widget.NewFormItem("Priority", a.priority),
		widget.NewFormItem("Due", a.due),
		widget.NewFormItem("Completion", a.completion),
	)

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			task := &task{title: "New task"}
			a.data.add(task)
			a.data.refreshData()
		}),
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
