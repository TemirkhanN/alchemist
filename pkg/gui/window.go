package gui

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Position struct {
	x float64
	y float64
}

func (p Position) X() float64 {
	return p.x
}

func (p Position) Y() float64 {
	return p.y
}

func (p Position) absolute(against Position) Position {
	return NewPosition(p.X()+against.X(), p.Y()+against.Y())
}

func (p Position) relative(against Position) Position {
	return NewPosition(p.X()-against.X(), p.Y()-against.Y())
}

func NewPosition(x float64, y float64) Position {
	return Position{x: x, y: y}
}

type WindowConfig struct {
	Title       string
	Width       float64
	Height      float64
	DebugMode   bool
	Position    Position
	ScrollSpeed uint8
}

type Window struct {
	graphics    *Layer
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
		graphics:    CreateLayer(preset.Width, preset.Height, true),
		scrollSpeed: scrollSpeed,
		debugMode:   preset.DebugMode,
	}

	return window
}

func (w Window) Width() float64 {
	return w.graphics.Width()
}

func (w Window) Height() float64 {
	return w.graphics.Height()
}

func (w *Window) AddLayer(layer *Layer, position Position) {
	w.graphics.AddElement(layer, position)
}

func (w *Window) LeftButtonClicked() bool {
	return w.window.JustPressed(pixelgl.MouseButtonLeft) && w.window.MouseInsideWindow()
}

func (w Window) CursorPosition() Position {
	return Position{
		x: w.window.MousePosition().X,
		y: w.window.MousePosition().Y,
	}
}

func (w Window) Closed() bool {
	return w.window.Closed()
}

func (w *Window) Close() {
	if !w.Closed() {
		w.window.SetClosed(true)
	}
}

func (w *Window) handleEvents() {
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

func (w *Window) handleVerticalScroll(layer *Layer, cursorPosition Position, vector pixel.Vec) bool {
	if vector.Y == 0 {
		return true
	}

	if layer.isUnderPosition(cursorPosition) {
		if layer.emitVerticalScroll(vector.Y * w.scrollSpeed) {
			return true
		}
	}

	for i := len(layer.Elements()) - 1; i >= 0; i-- {
		element := layer.Elements()[i]
		if !element.isVisible() {
			continue
		}

		childElement, isLayer := element.(*Layer)
		if isLayer && childElement.isUnderPosition(cursorPosition) {
			if w.handleVerticalScroll(childElement, cursorPosition, vector) {
				return true
			}
		}
	}

	return false
}

func (w *Window) handleLeftClick(graphics Canvas, clickedPosition Position) bool {
	if !graphics.isVisible() {
		return false
	}

	interactiveElement, isInteractiveElement := graphics.(InteractiveCanvas)
	if isInteractiveElement {
		if interactiveElement.IsUnderPosition(clickedPosition) {
			interactiveElement.EmitClick()

			return true
		}
	}

	// Interaction priority is LIFO. EmitClick over canvasB which is drawn over canvasA shall start from canvas B handle
	for i := len(graphics.Elements()) - 1; i >= 0; i-- {
		element := graphics.Elements()[i]
		if w.handleLeftClick(element, clickedPosition) {
			// stop further propagation
			return true
		}
	}

	return false
}

func (w *Window) handleMouseOver(graphics Canvas, onPosition Position) bool {
	if !graphics.isVisible() {
		return false
	}

	interactiveElement, isInteractiveElement := graphics.(InteractiveCanvas)
	if isInteractiveElement {
		if interactiveElement.IsUnderPosition(onPosition) {
			interactiveElement.EmitMouseOver()

			return true
		}
	}

	// Interaction priority is LIFO. EmitClick over canvasB which is drawn over canvasA shall start from canvas B handle
	for i := len(graphics.Elements()) - 1; i >= 0; i-- {
		element := graphics.Elements()[i]
		if w.handleMouseOver(element, onPosition) {
			// stop further propagation
			return true
		}
	}

	return false
}

func (w *Window) handleMouseOut(graphics Canvas, lastCursorPosition Position) {
	if !graphics.isVisible() {
		return
	}

	interactiveElement, isInteractiveElement := graphics.(InteractiveCanvas)
	if isInteractiveElement {
		if !interactiveElement.IsUnderPosition(lastCursorPosition) {
			interactiveElement.EmitMouseOut()
		}
	}

	// Interaction priority is LIFO. EmitClick over canvasB which is drawn over canvasA shall start from canvas B handle
	for i := len(graphics.Elements()) - 1; i >= 0; i-- {
		element := graphics.Elements()[i]
		w.handleMouseOut(element, lastCursorPosition)
	}
}

var ZeroPosition = Position{x: 0, y: 0}
