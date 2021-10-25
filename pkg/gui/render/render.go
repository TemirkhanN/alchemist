package render

import (
	"golang.org/x/image/colornames"

	"github.com/TemirkhanN/alchemist/pkg/gui/geometry"
	"github.com/TemirkhanN/alchemist/pkg/gui/graphics"
)

type Renderer interface {
	Render(window graphics.Window)
}

type CommonRenderer struct{}

func (cr CommonRenderer) Render(window *graphics.Window) {
	for !window.Closed() {
		window.StartFrame()

		window.FillWithColor(colornames.White)
		cr.draw(window.Graphics(), window)

		window.EndFrame()
	}
}

func (cr CommonRenderer) draw(drawer graphics.Canvas, in *graphics.Window) {
	// todo omg...
	sprite, isSprite := drawer.(*graphics.SpriteCanvas)
	if isSprite {
		cr.renderSpriteCanvas(*sprite, in)

		return
	}

	button, isButton := drawer.(*graphics.Button)
	if isButton {
		cr.renderSpriteCanvas(button.SpriteCanvas, in)

		return
	}

	layer, isLayer := drawer.(*graphics.Layer)
	if isLayer {
		cr.renderLayer(*layer, in)

		return
	}

	textCanvas, isText := drawer.(*graphics.TextCanvas)
	if isText {
		cr.renderTextCanvas(*textCanvas, in)

		return
	}
}

func (cr CommonRenderer) renderLayer(layer graphics.Layer, in *graphics.Window) {
	if !layer.IsVisible() {
		return
	}

	for _, element := range layer.Elements() {
		if !element.IsVisible() {
			continue
		}

		if layer.IsScrollable() && !layer.CanFullyFit(element) {
			continue
		}

		cr.draw(element, in)
	}
}

func (cr CommonRenderer) renderSpriteCanvas(canvas graphics.SpriteCanvas, in *graphics.Window) {
	if !canvas.IsVisible() {
		return
	}

	fromLeftBottomCorner := geometry.NewPosition(
		canvas.Width()/2+canvas.Position().X(),
		canvas.Height()/2+canvas.Position().Y(),
	)

	canvas.Sprite().Draw(in, fromLeftBottomCorner)
}

func (cr CommonRenderer) renderTextCanvas(canvas graphics.TextCanvas, in *graphics.Window) {
	if !canvas.IsVisible() {
		return
	}

	textPosition := canvas.Position().Add(geometry.NewPosition(0, canvas.Height()-canvas.Height()))

	canvas.Draw(in, textPosition)
}
