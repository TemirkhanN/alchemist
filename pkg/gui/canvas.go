package gui

import (
	"strings"
)

type Drawer interface {
	Show()
	Hide()
	Draw()
	Width() float64
	Height() float64
	isVisible() bool
	NeedsRedraw() bool
	Elements() []Drawer
	Position() Position
	setPosition(position Position)
}

type Canvas interface {
	IsUnderPosition(position Position) bool
	Position() Position
	Drawer
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
	text     string
	font     Font
	maxWidth float64
	CommonCanvas
}

type SpriteCanvas struct {
	sprite *Sprite
	CommonCanvas
}

func (canvas *CommonCanvas) NeedsRedraw() bool {
	return canvas.needsRedraw && canvas.visible
}

func (canvas *CommonCanvas) Show() {
	canvas.visible = true
}

func (canvas *CommonCanvas) Hide() {
	canvas.visible = false
}

func (canvas *CommonCanvas) Draw() {}

func (canvas *CommonCanvas) isVisible() bool {
	return canvas.visible
}

func (canvas CommonCanvas) Width() float64 {
	return 0
}

func (canvas CommonCanvas) Height() float64 {
	return 0
}

func (canvas *CommonCanvas) Position() Position {
	return canvas.position
}

func (canvas *CommonCanvas) setPosition(position Position) {
	canvas.position = position
}

func (canvas CommonCanvas) IsUnderPosition(Position) bool {
	return false
}

func (canvas CommonCanvas) Elements() []Drawer {
	return nil
}

func (canvas *SpriteCanvas) Draw() {
	if !canvas.visible {
		return
	}

	canvas.sprite.draw(canvas.drawnOn, canvas.position)

	if canvas.drawnOn.debugMode {
		highlightElement(canvas, canvas.drawnOn)
	}

	canvas.needsRedraw = false
}

func (canvas *SpriteCanvas) IsUnderPosition(position Position) bool {
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

func (canvas *TextCanvas) Draw() {
	if !canvas.visible {
		return
	}

	if canvas.drawnOn.debugMode {
		// debug draws canvas frame. Not text frame itself.
		highlightElement(canvas, canvas.drawnOn)
	}

	textPosition := canvas.position.absolute(NewPosition(0, canvas.Height()-canvas.font.atlas.LineHeight()))
	canvas.drawnOn.drawText(canvas.text, textPosition, canvas.font)
	canvas.needsRedraw = false
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
