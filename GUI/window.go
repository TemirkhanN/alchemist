package GUI

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Position struct {
	X float64
	Y float64
}

type Window struct {
	window *pixelgl.Window
}

func CreateWindow(width float64, height float64) *Window {
	cfg := pixelgl.WindowConfig{
		Title:  "Alchemist",
		Bounds: pixel.R(0, 0, width, height),
		VSync:  true,
	}

	window, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	return &Window{window: window}
}

func (w *Window) DrawSprite(sprite *pixel.Sprite, position Position) {
	fromLeftBottomCorner := pixel.Vec{
		X: sprite.Picture().Bounds().Center().X + position.X,
		Y: sprite.Picture().Bounds().Center().Y + position.Y,
	}

	sprite.Draw(w.window, pixel.IM.Moved(fromLeftBottomCorner))
}

func (w *Window) CreateButton(sprite *pixel.Sprite, position Position) *Button {
	button := &Button{
		positionX:   position.X,
		positionY:   position.Y,
		sprite:      sprite,
		drawnOn:     w,
		visible:     true,
		onclickfn:   func() {},
		needsRedraw: true,
	}

	button.Draw()
	button.needsRedraw = true

	return button
}

func (w Window) LeftButtonClicked() bool {
	return w.window.JustPressed(pixelgl.MouseButtonLeft) && w.window.MouseInsideWindow()
}

func (w Window) ClickedPosition() Position {
	return Position{
		X: w.window.MousePosition().X,
		Y: w.window.MousePosition().Y,
	}
}

func (w Window) Closed() bool {
	return w.window.Closed()
}

func (w *Window) Refresh() {
	w.window.Update()
}
