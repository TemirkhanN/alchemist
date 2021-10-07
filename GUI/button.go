package GUI

import (
	"github.com/faiface/pixel"
)

type Button struct {
	positionX   float64
	positionY   float64
	visible     bool
	sprite      *pixel.Sprite
	drawnOn     *Window
	onclickfn   func()
	needsRedraw bool
}

func (b Button) NeedsRedraw() bool {
	return b.needsRedraw && b.visible
}

func (b Button) IsUnderPosition(position Position) bool {
	if !b.visible {
		return false
	}

	buttonWidth := b.sprite.Picture().Bounds().W()
	buttonHeight := b.sprite.Picture().Bounds().H()

	bottomLeftX := b.positionX
	bottomLeftY := b.positionY
	topRightX := b.positionX + buttonWidth
	topRightY := b.positionY + buttonHeight

	if (position.X > bottomLeftX && position.X < topRightX) && (position.Y > bottomLeftY && position.Y < topRightY) {
		return true
	}

	return false
}

func (b *Button) Draw() {
	if !b.NeedsRedraw() {
		return
	}

	b.drawnOn.DrawSprite(b.sprite, Position{X: b.positionX, Y: b.positionY})
	b.needsRedraw = false
}

func (b *Button) Enable() {
	b.visible = true
}

func (b *Button) Disable() {
	b.visible = false
}

func (b *Button) SetClickHandler(handler func()) {
	b.onclickfn = handler
}

func (b *Button) Click() {
	b.onclickfn()
}
