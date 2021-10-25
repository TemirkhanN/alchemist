package render

import (
	"github.com/TemirkhanN/alchemist/pkg/gui/graphics"
)

type Renderer interface {
	Render(window graphics.Window)
}

type CommonRenderer struct{}

func (cr CommonRenderer) Render(window *graphics.Window) {
	for !window.Closed() {
		window.StartFrame()
		window.Draw()
		window.EndFrame()
	}
}
