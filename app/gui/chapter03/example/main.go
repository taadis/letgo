package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.New()
	w := a.NewWindow("Snake")
	game = setupGame()
	w.SetContent(game)
	w.Resize(fyne.NewSize(200, 200))
	w.SetFixedSize(true)
	go runGame()
	w.ShowAndRun()
}

type snakePart struct {
	x, y float32
}

var (
	snakeParts []snakePart
	game       *fyne.Container
)

func setupGame() *fyne.Container {
	var segments []fyne.CanvasObject
	for i := 0; i < 10; i++ {
		seg := snakePart{9, float32(5 + i)}
		snakeParts = append(snakeParts, seg)

		r := canvas.NewRectangle(color.RGBA{G: 0x66, A: 0xff})
		r.Resize(fyne.NewSize(10, 10))
		r.Move(fyne.NewPos(90, float32(50+i*10)))
		segments = append(segments, r)
	}
	//head := canvas.NewRectangle(&color.RGBA{G: 0x66, A: 0xff})
	//head.Resize(fyne.NewSize(10, 10))
	//head.Move(fyne.NewPos(snakeParts[0].x*10, snakeParts[0].y*10))
	//segments = append(segments, head)
	return container.NewWithoutLayout(segments...)
}

func refreshGame() {
	for i, seg := range snakeParts {
		rect := game.Objects[i]
		rect.Move(fyne.NewPos(seg.x*10, seg.y*10))
	}
	game.Refresh()
}

func runGame() {
	for {
		time.Sleep(250 * time.Millisecond)
		for i := len(snakeParts) - 1; i >= 1; i-- {
			snakeParts[i] = snakeParts[i-1]
		}
		snakeParts[0].y--
		refreshGame()
	}
}
