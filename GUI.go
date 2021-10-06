package main

import (
	"embed"
	_ "embed"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"image"
)

type Button struct {
	positionX   float64
	positionY   float64
	visible     bool
	sprite      *pixel.Sprite
	drawnOn     *pixelgl.Window
	onclickfn   func()
	needsRedraw bool
}

type Interactive interface {
	Enable()
	Disable()
	Click()
	IsUnderPosition(vec pixel.Vec) bool
	Drawer
}

type Drawer interface {
	Draw()
	NeedsRedraw() bool
}

func (b Button) NeedsRedraw() bool {
	return b.needsRedraw && b.visible
}

func (b Button) IsUnderPosition(point pixel.Vec) bool {
	if !b.visible {
		return false
	}

	buttonWidth := b.sprite.Picture().Bounds().W()
	buttonHeight := b.sprite.Picture().Bounds().H()

	bottomLeftX := b.positionX
	bottomLeftY := b.positionY
	topRightX := b.positionX + buttonWidth
	topRightY := b.positionY + buttonHeight

	if (point.X > bottomLeftX && point.X < topRightX) && (point.Y > bottomLeftY && point.Y < topRightY) {
		return true
	}

	return false
}

func (b *Button) Draw() {
	if !b.NeedsRedraw() {
		return
	}

	placeSprite(b.sprite, b.drawnOn, b.positionX, b.positionY)
	b.needsRedraw = false
}

func (b *Button) Enable() {
	b.visible = true
}

func (b *Button) Disable() {
	b.visible = false
}

func (b *Button) Click() {
	b.onclickfn()
}

func createWindow(width float64, height float64) *pixelgl.Window {
	cfg := pixelgl.WindowConfig{
		Title:  "Alchemist",
		Bounds: pixel.R(0, 0, width, height),
		VSync:  true,
	}

	window, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	return window
}

func placeSprite(sprite *pixel.Sprite, target *pixelgl.Window, x float64, y float64) {
	position := pixel.Vec{
		X: sprite.Picture().Bounds().Center().X + x,
		Y: sprite.Picture().Bounds().Center().Y + y,
	}

	sprite.Draw(target, pixel.IM.Moved(position))
}

func placeButton(sprite *pixel.Sprite, target *pixelgl.Window, x float64, y float64) *Button {
	button := &Button{
		positionX:   x,
		positionY:   y,
		sprite:      sprite,
		drawnOn:     target,
		visible:     true,
		onclickfn:   func() {},
		needsRedraw: true,
	}

	button.Draw()
	button.needsRedraw = true

	return button
}

//go:embed assets/sprites/*.png
var sprites embed.FS
func loadSprite(spriteName string) *pixel.Sprite {
	file, err := sprites.Open("assets/sprites/" + spriteName + ".png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	pic := pixel.PictureDataFromImage(img)

	return pixel.NewSprite(pic, pic.Bounds())
}

/*
func drawText() {
	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt := text.New(pixel.V(100, 500), basicAtlas)

	fmt.Fprintln(basicTxt, "Some text")

	basicTxt.Draw(window, pixel.IM)
}
*/
