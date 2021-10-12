package gui

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

// Debugger : crappy tools to debug layer shapes and etc.
type Debugger interface {
	debugDraw()
	setDebugTool(window *Window)
}

func (layer *Layer) isDebugModeOn() bool {
	return layer._debug != nil
}

func (layer *Layer) setDebugTool(window *Window) {
	layer._debug = window
}

func (layer *Layer) debugDraw() {
	if !layer.isDebugModeOn() {
		return
	}

	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(0, 0, 0)
	// ↑
	imd.Push(pixel.V(layer.Position().X(), layer.Position().Y()))
	imd.Push(pixel.V(layer.Position().X(), layer.Position().Y()+layer.Height()))
	// →
	imd.Push(pixel.V(layer.Position().X()+layer.Width(), layer.Position().Y()+layer.Height()))
	// ↓
	imd.Push(pixel.V(layer.Position().X()+layer.Width(), layer.Position().Y()))
	// ←
	imd.Push(pixel.V(layer.Position().X(), layer.Position().Y()))
	imd.Rectangle(1)
	imd.Draw(layer._debug.window)

	for _, childElement := range layer.elements {
		debugger, isDebugger := childElement.(Debugger)
		if isDebugger {
			debugger.setDebugTool(layer._debug)
		}
	}
}
