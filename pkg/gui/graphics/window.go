package graphics

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"

	"github.com/TemirkhanN/alchemist/pkg/gui/geometry"
)

type WindowConfig struct {
	Title       string
	Width       float64
	Height      float64
	DebugMode   bool
	Position    geometry.Position
	ScrollSpeed uint8
}

type Window struct {
	graphics    *Layout
	window      *pixelgl.Window
	scrollSpeed float64
	debugMode   bool
}

func NewWindow(preset WindowConfig) *Window {
	cfg := pixelgl.WindowConfig{
		Title:                  preset.Title,
		Icon:                   nil,
		Bounds:                 pixel.R(0, 0, preset.Width, preset.Height),
		Position:               pixel.Vec{X: preset.Position.X(), Y: preset.Position.Y()},
		Monitor:                nil,
		Resizable:              false,
		Undecorated:            false,
		NoIconify:              false,
		AlwaysOnTop:            false,
		TransparentFramebuffer: false,
		VSync:                  true,
		Maximized:              false,
		Invisible:              false,
	}

	w, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	scrollSpeed := 1.0

	if preset.ScrollSpeed != 0 {
		if preset.ScrollSpeed > 10 {
			preset.ScrollSpeed = 10
		}

		scrollSpeed = float64(preset.ScrollSpeed)
	}

	window := &Window{
		window:      w,
		graphics:    NewLayer(preset.Width, preset.Height, true),
		scrollSpeed: scrollSpeed,
		debugMode:   preset.DebugMode,
	}

	return window
}

func (w Window) target() pixel.Target {
	return w.window
}

func (w Window) Draw() {
	w.graphics.Draw(w)
}

func (w Window) FillWithColor(color color.RGBA) {
	w.window.Clear(color)
}

func (w Window) Width() float64 {
	return w.graphics.Width()
}

func (w Window) Height() float64 {
	return w.graphics.Height()
}

func (w *Window) AddLayer(layer *Layout, position geometry.Position) {
	w.graphics.AddElement(layer, position)
}

func (w *Window) LeftButtonClicked() bool {
	return w.window.JustPressed(pixelgl.MouseButtonLeft) && w.window.MouseInsideWindow()
}

func (w Window) CursorPosition() geometry.Position {
	return geometry.NewPosition(w.window.MousePosition().X, w.window.MousePosition().Y)
}

func (w Window) Closed() bool {
	return w.window.Closed()
}

func (w *Window) Close() {
	if !w.Closed() {
		w.window.SetClosed(true)
	}
}

func (w Window) EndFrame() {
	w.window.Update()
}

func (w *Window) StartFrame() {
	if w.Closed() {
		return
	}

	if w.window.MouseInsideWindow() {
		cursorPosition := w.CursorPosition()
		w.handleVerticalScroll(w.graphics, cursorPosition, w.window.MouseScroll())
		w.handleMouseOut(w.graphics, cursorPosition)
		w.handleMouseOver(w.graphics, cursorPosition)

		if w.LeftButtonClicked() {
			w.handleLeftClick(w.graphics, cursorPosition)
		}
	}
}

func (w *Window) handleVerticalScroll(layer *Layout, cursorPosition geometry.Position, vector pixel.Vec) bool {
	if vector.Y == 0 {
		return true
	}

	if layer.IsUnderPosition(cursorPosition) {
		if layer.EmitVerticalScroll(vector.Y * w.scrollSpeed) {
			return true
		}
	}

	for i := len(layer.Elements()) - 1; i >= 0; i-- {
		element := layer.Elements()[i]
		if !element.IsVisible() {
			continue
		}

		childElement, isLayer := element.(*Layout)
		if isLayer && childElement.IsUnderPosition(cursorPosition) {
			if w.handleVerticalScroll(childElement, cursorPosition, vector) {
				return true
			}
		}
	}

	return false
}

func (w *Window) handleLeftClick(element Canvas, clickedPosition geometry.Position) bool {
	if !element.IsVisible() {
		return false
	}

	interactiveElement, isInteractiveElement := element.(InteractiveCanvas)
	if isInteractiveElement {
		if interactiveElement.IsUnderPosition(clickedPosition) {
			interactiveElement.EmitClick()

			return true
		}
	}

	// Interaction priority is LIFO. EmitClick over canvasB which is drawn over canvasA shall start from canvas B handle
	for i := len(element.Elements()) - 1; i >= 0; i-- {
		childElement := element.Elements()[i]
		if w.handleLeftClick(childElement, clickedPosition.Subtract(element.Position())) {
			// stop further propagation
			return true
		}
	}

	return false
}

func (w *Window) handleMouseOver(element Canvas, onPosition geometry.Position) bool {
	if !element.IsVisible() {
		return false
	}

	interactiveElement, isInteractiveElement := element.(InteractiveCanvas)
	if isInteractiveElement {
		if interactiveElement.IsUnderPosition(onPosition) {
			interactiveElement.EmitMouseOver()

			return true
		}
	}

	// Interaction priority is LIFO.
	for i := len(element.Elements()) - 1; i >= 0; i-- {
		childElement := element.Elements()[i]
		if w.handleMouseOver(childElement, onPosition.Subtract(element.Position())) {
			// stop further propagation
			return true
		}
	}

	return false
}

func (w *Window) handleMouseOut(element Canvas, lastCursorPosition geometry.Position) {
	if !element.IsVisible() {
		return
	}

	interactiveElement, isInteractiveElement := element.(InteractiveCanvas)
	if isInteractiveElement {
		if !interactiveElement.IsUnderPosition(lastCursorPosition) {
			interactiveElement.EmitMouseOut()
		}
	}

	// Interaction priority is LIFO. EmitClick over canvasB which is drawn over canvasA shall start from canvas B handle
	for i := len(element.Elements()) - 1; i >= 0; i-- {
		childElement := element.Elements()[i]
		w.handleMouseOut(childElement, lastCursorPosition.Subtract(element.Position()))
	}
}
