package GUI

import "github.com/faiface/pixel"

type Canvas interface {
	Show()
	Hide()
	Draw()
	NeedsRedraw() bool
	IsUnderPosition(position Position) bool
}

type InteractiveCanvas interface {
	Click()
	SetClickHandler(func())
	Canvas
}

type CommonCanvas struct {
	position    Position
	visible     bool
	sprite      *pixel.Sprite
	drawnOn     *Window
	needsRedraw bool
}

func (canvas CommonCanvas) NeedsRedraw() bool {
	return canvas.needsRedraw && canvas.visible
}

func (canvas *CommonCanvas) Draw() {
	if !canvas.visible {
		return
	}
	canvas.drawnOn.drawSprite(canvas.sprite, Position{X: canvas.position.X, Y: canvas.position.Y})
	canvas.needsRedraw = false
}

func (canvas *CommonCanvas) Show() {
	canvas.visible = true
}

func (canvas *CommonCanvas) Hide() {
	canvas.visible = false
}

func (canvas CommonCanvas) IsUnderPosition(position Position) bool {
	buttonWidth := canvas.sprite.Picture().Bounds().W()
	buttonHeight := canvas.sprite.Picture().Bounds().H()

	bottomLeftX := canvas.position.X
	bottomLeftY := canvas.position.Y
	topRightX := canvas.position.X + buttonWidth
	topRightY := canvas.position.Y + buttonHeight

	if (position.X > bottomLeftX && position.X < topRightX) && (position.Y > bottomLeftY && position.Y < topRightY) {
		return true
	}

	return false
}
