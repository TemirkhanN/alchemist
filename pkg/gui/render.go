package gui

import (
	"golang.org/x/image/colornames"
)

type Renderer interface {
	Render(window Window)
}

type CommonRenderer struct{}

func (cr CommonRenderer) Render(window *Window) {
	for !window.Closed() {
		window.HandleEvents()

		if window.graphics.NeedsRedraw() {
			window.window.Clear(colornames.White)
			cr.draw(window.graphics, window)
		}

		window.window.Update()
	}
}

func (cr CommonRenderer) draw(drawer Canvas, in *Window) {
	if in.debugMode {
		highlightElement(drawer, in)
	}

	// todo omg...
	sprite, isSprite := drawer.(*SpriteCanvas)
	if isSprite {
		cr.renderSpriteCanvas(*sprite, in)

		return
	}

	button, isButton := drawer.(*Button)
	if isButton {
		cr.renderSpriteCanvas(button.SpriteCanvas, in)

		return
	}

	layer, isLayer := drawer.(*Layer)
	if isLayer {
		cr.renderLayer(*layer, in)

		return
	}

	text, isText := drawer.(*TextCanvas)
	if isText {
		cr.renderTextCanvas(*text, in)

		return
	}
}

func (cr CommonRenderer) renderLayer(layer Layer, in *Window) {
	if !layer.visible {
		layer.needsRedraw = false

		return
	}

	for _, element := range layer.elements {
		if element.isVisible() && layer.canFullyFit(element) {
			cr.draw(element, in)
		}
	}

	layer.needsRedraw = false
}

func (cr CommonRenderer) renderSpriteCanvas(canvas SpriteCanvas, in *Window) {
	if !canvas.visible {
		return
	}

	in.draw(canvas.sprite, canvas.position)

	canvas.needsRedraw = false
}

func (cr CommonRenderer) renderTextCanvas(canvas TextCanvas, in *Window) {
	if !canvas.visible {
		return
	}

	textPosition := canvas.position.absolute(NewPosition(0, canvas.Height()-canvas.font.atlas.LineHeight()))
	in.drawText(canvas.text, textPosition, canvas.font)
	canvas.needsRedraw = false
}
