package gui

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

func highlightElement(drawer Canvas, window *Window) {
	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(0, 0, 0)
	// ↑
	imd.Push(pixel.V(drawer.Position().X(), drawer.Position().Y()))
	imd.Push(pixel.V(drawer.Position().X(), drawer.Position().Y()+drawer.Height()))
	// →
	imd.Push(pixel.V(drawer.Position().X()+drawer.Width(), drawer.Position().Y()+drawer.Height()))
	// ↓
	imd.Push(pixel.V(drawer.Position().X()+drawer.Width(), drawer.Position().Y()))
	// ←
	imd.Push(pixel.V(drawer.Position().X(), drawer.Position().Y()))
	imd.Rectangle(1)
	imd.Draw(window.window)
}
