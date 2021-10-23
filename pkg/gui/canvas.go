package gui

import (
	"strings"
)

type Canvas interface {
	Show()
	Hide()
	Width() float64
	Height() float64
	isVisible() bool
	NeedsRedraw() bool
	Elements() []Canvas
	Position() Position
	IsUnderPosition(position Position) bool
	setPosition(position Position)
}

type InteractiveCanvas interface {
	EmitClick()
	SetClickHandler(func())
	EmitMouseOver()
	SetMouseOverHandler(func())
	EmitMouseOut()
	SetMouseOutHandler(func())
	Canvas
}

type CommonCanvas struct {
	position    Position
	visible     bool
	needsRedraw bool
	drawnOn     *Window
}

type TextCanvas struct {
	CommonCanvas
	text     string
	font     Font
	maxWidth float64
}

type SpriteCanvas struct {
	CommonCanvas
	sprite *Sprite
}

func (canvas CommonCanvas) NeedsRedraw() bool {
	return canvas.needsRedraw && canvas.visible
}

func (canvas *CommonCanvas) Show() {
	canvas.visible = true
}

func (canvas *CommonCanvas) Hide() {
	canvas.visible = false
}

func (canvas CommonCanvas) isVisible() bool {
	return canvas.visible
}

func (canvas CommonCanvas) Width() float64 {
	return 0
}

func (canvas CommonCanvas) Height() float64 {
	return 0
}

func (canvas CommonCanvas) Position() Position {
	return canvas.position
}

func (canvas *CommonCanvas) setPosition(position Position) {
	canvas.position = position
}

func (canvas CommonCanvas) IsUnderPosition(Position) bool {
	return false
}

func (canvas CommonCanvas) Elements() []Canvas {
	return nil
}

func (canvas SpriteCanvas) IsUnderPosition(position Position) bool {
	buttonWidth := canvas.sprite.Width()
	buttonHeight := canvas.sprite.Height()

	bottomLeftX := canvas.position.X()
	bottomLeftY := canvas.position.Y()
	topRightX := canvas.position.X() + buttonWidth
	topRightY := canvas.position.Y() + buttonHeight

	posX := position.X()
	posY := position.Y()

	if (posX > bottomLeftX && posX < topRightX) && (posY > bottomLeftY && posY < topRightY) {
		return true
	}

	return false
}

func (canvas SpriteCanvas) Width() float64 {
	return canvas.sprite.Width()
}

func (canvas SpriteCanvas) Height() float64 {
	return canvas.sprite.Height()
}

func (canvas *SpriteCanvas) ChangeSprite(withSprite *Sprite) {
	canvas.sprite = withSprite
	canvas.needsRedraw = true
}

func (canvas TextCanvas) Width() float64 {
	return canvas.maxWidth
}

func (canvas TextCanvas) Height() float64 {
	lineBreaks := 1 + strings.Count(canvas.text, "\n")

	return canvas.font.atlas.LineHeight() * float64(lineBreaks)
}

func (canvas *TextCanvas) ChangeText(text string) {
	canvas.text = text
	canvas.AddLineBreaks()
	canvas.needsRedraw = true
}

func (canvas *TextCanvas) AddLineBreaks() {
	parts := strings.Split(canvas.text, " ")

	lineLength := 0.0

	for i, part := range parts {
		wordLength := canvas.font.calculateWidthInPixels(part)
		if lineLength+wordLength <= canvas.maxWidth || i == 0 {
			lineLength += wordLength

			continue
		}

		// we add line breaks only after at least first word
		parts[i-1] += "\n"
		lineLength = wordLength
		canvas.position.y += canvas.font.atlas.LineHeight()
	}

	canvas.text = strings.Join(parts, " ")
}
