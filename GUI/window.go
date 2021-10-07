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
	layers []Layer
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

func (w *Window) AddLayer(layer Layer) {
	w.layers = append(w.layers, layer)
}

func (w *Window) CreateCanvas(sprite *pixel.Sprite, position Position) *CommonCanvas {
	return &CommonCanvas{
		position:    position,
		visible:     true,
		sprite:      sprite,
		drawnOn:     w,
		needsRedraw: false,
	}
}

func (w *Window) CreateButton(sprite *pixel.Sprite, position Position) *Button {
	button := &Button{
		CommonCanvas: CommonCanvas{
			position:    position,
			sprite:      sprite,
			drawnOn:     w,
			visible:     true,
			needsRedraw: true,
		},
		onclickfn: nil,
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
	if w.Closed() {
		return
	}

	needRedraw := true
	leftClickHandled := false
	for _, layer := range w.layers {
		// Interaction priority is LIFO. Click over canvasB which is drawn over canvasA shall start from canvas B handle
		for i := len(layer.elements) - 1; i >= 0; i-- {
			element := layer.elements[i]
			interactiveElement, isInteractiveElement := element.(InteractiveCanvas)
			if !leftClickHandled && isInteractiveElement && w.LeftButtonClicked() {
				if interactiveElement.IsUnderPosition(w.ClickedPosition()) {
					interactiveElement.Click()
					// First item that handles click stops further propagation
					leftClickHandled = true
				}
			}

			if !needRedraw && element.NeedsRedraw() {
				needRedraw = true
			}
		}
	}

	if !needRedraw {
		return
	}

	w.window.Clear(colornames.White)
	for _, layer := range w.layers {
		layer.Draw()
	}

	w.window.Update()
}

func (w *Window) drawSprite(sprite *pixel.Sprite, position Position) {
	fromLeftBottomCorner := pixel.Vec{
		X: sprite.Picture().Bounds().Center().X + position.X,
		Y: sprite.Picture().Bounds().Center().Y + position.Y,
	}

	sprite.Draw(w.window, pixel.IM.Moved(fromLeftBottomCorner))
}
