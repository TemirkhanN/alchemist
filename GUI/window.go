package GUI

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Position struct {
	X float64
	Y float64
}

type Window struct {
	layers []*Layer
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

func (w *Window) AddLayer(layer *Layer) {
	w.layers = append(w.layers, layer)
}

func (w *Window) CreateCanvas(sprite *pixel.Sprite, position Position) *CommonCanvas {
	return &CommonCanvas{
		position:    position,
		visible:     true,
		sprite:      sprite,
		drawnOn:     w,
		needsRedraw: true,
	}
}

func (w *Window) CreateButton(sprite *pixel.Sprite, position Position) *Button {
	return &Button{
		CommonCanvas: CommonCanvas{
			position:    position,
			sprite:      sprite,
			drawnOn:     w,
			visible:     true,
			needsRedraw: true,
		},
		onclickfn: nil,
	}
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
	if w.Closed() {
		return
	}

	w.handleClicks()
	w.draw()
	w.window.Update()
}

func (w *Window) draw() {
	for _, layer := range w.layers {
		if layer.NeedsRedraw() {
			w.window.Clear(colornames.White)
			for _, layer := range w.layers {
				layer.Draw()
			}
			break
		}
	}
}

func (w *Window) handleClicks() {
	if !w.LeftButtonClicked() {
		return
	}
	// Interaction priority is LIFO. Click over canvasB which is drawn over canvasA shall start from canvas B handle
	for i := len(w.layers) - 1; i >= 0; i-- {
		layer := w.layers[i]
		if !layer.visible {
			continue
		}
		for j := len(layer.elements) - 1; j >= 0; j-- {
			element := layer.elements[j]
			if !element.IsUnderPosition(w.ClickedPosition()) {
				continue
			}

			interactiveElement, isInteractiveElement := element.(InteractiveCanvas)
			if isInteractiveElement {
				interactiveElement.Click()
			}
			// stop further propagation
			return
		}
	}
}

func (w *Window) drawSprite(sprite *pixel.Sprite, position Position) {
	fromLeftBottomCorner := pixel.Vec{
		X: sprite.Picture().Bounds().Center().X + position.X,
		Y: sprite.Picture().Bounds().Center().Y + position.Y,
	}

	sprite.Draw(w.window, pixel.IM.Moved(fromLeftBottomCorner))
}
