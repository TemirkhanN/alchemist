package gui

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
)

type Renderer interface {
	Render(window Window)
}

type CommonRenderer struct{}

func (cr CommonRenderer) Render(window *Window) {
	for !window.Closed() {
		window.handleEvents()

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

	textCanvas, isText := drawer.(*TextCanvas)
	if isText {
		cr.renderTextCanvas(*textCanvas, in)

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

	fromLeftBottomCorner := pixel.Vec{
		X: canvas.sprite.src.Frame().W()/2 + canvas.position.X(),
		Y: canvas.sprite.src.Frame().H()/2 + canvas.position.Y(),
	}

	canvas.sprite.src.Draw(in.window, pixel.IM.Moved(fromLeftBottomCorner))

	canvas.needsRedraw = false
}

func (cr CommonRenderer) renderTextCanvas(canvas TextCanvas, in *Window) {
	if !canvas.visible {
		return
	}

	textPosition := canvas.position.absolute(NewPosition(0, canvas.Height()-canvas.font.atlas.LineHeight()))

	basicTxt := text.New(pixel.V(textPosition.X(), textPosition.Y()), canvas.font.atlas)

	fmt.Fprintln(basicTxt, canvas.text)

	basicTxt.DrawColorMask(in.window, pixel.IM, colornames.Sienna)

	canvas.needsRedraw = false
}
