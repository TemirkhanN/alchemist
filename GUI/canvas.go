package GUI

import "github.com/faiface/pixel"

type Drawer interface {
	Show()
	Hide()
	Draw()
	Width() float64
	Height() float64
	IsVisible() bool
	NeedsRedraw() bool
	Elements() []Drawer
}

type Canvas interface {
	IsUnderPosition(position Position) bool
	Position() Position
	Drawer
}

type InteractiveCanvas interface {
	EmitClick()
	SetClickHandler(func())
	Canvas
}

type CommonCanvas struct {
	position    Position
	visible     bool
	needsRedraw bool
	drawnOn     *Window
}

type TextCanvas struct {
	text string
	CommonCanvas
}

type SpriteCanvas struct {
	sprite *pixel.Sprite
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

func (canvas *CommonCanvas) IsVisible() bool {
	return canvas.visible
}

func (canvas *CommonCanvas) Draw() {

}

func (canvas *CommonCanvas) Position() Position {
	return canvas.position
}

func (canvas *CommonCanvas) IsUnderPosition(position Position) bool {
	return false
}

func (canvas *CommonCanvas) Elements() []Drawer {
	return nil
}

func (canvas *SpriteCanvas) Draw() {
	if !canvas.visible {
		return
	}
	canvas.drawnOn.drawSprite(canvas.sprite, Position{X: canvas.position.X, Y: canvas.position.Y})
	canvas.needsRedraw = false
}

func (canvas *SpriteCanvas) IsUnderPosition(position Position) bool {
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

func (canvas *SpriteCanvas) Width() float64 {
	return canvas.sprite.Picture().Bounds().W()
}

func (canvas *SpriteCanvas) Height() float64 {
	return canvas.sprite.Picture().Bounds().H()
}

func (canvas *SpriteCanvas) ChangeSprite(withSprite *pixel.Sprite) {
	canvas.sprite = withSprite
	canvas.needsRedraw = true
}

func (canvas *TextCanvas) Draw() {
	if !canvas.visible {
		return
	}

	canvas.drawnOn.DrawText(canvas.text, canvas.position)
	canvas.needsRedraw = false
}

func (canvas *TextCanvas) Width() float64 {
	// todo
	return 0
}

func (canvas *TextCanvas) Height() float64 {
	// todo
	return 0
}

func (canvas *TextCanvas) ChangeText(text string) {
	canvas.text = text
	canvas.needsRedraw = true
}
