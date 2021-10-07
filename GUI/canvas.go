package GUI

import "github.com/faiface/pixel"

type Canvas interface {
	Show()
	Hide()
	Draw()
	NeedsRedraw() bool
}

type InteractiveCanvas interface {
	Click()
	SetClickHandler(func())
	IsUnderPosition(position Position) bool
	Canvas
}

type CommonCanvas struct {
	position Position
	visible     bool
	sprite      *pixel.Sprite
	drawnOn     *Window
	needsRedraw bool
}

func (canvas CommonCanvas) NeedsRedraw() bool {
	return canvas.needsRedraw && canvas.visible
}

func (canvas CommonCanvas) Draw() {
	canvas.drawnOn.drawSprite(canvas.sprite, Position{X: canvas.position.X, Y: canvas.position.Y})
}

func (canvas *CommonCanvas) Show() {
	canvas.visible = true
}

func (canvas *CommonCanvas) Hide() {
	canvas.visible = false
}
