package graphics

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"

	"github.com/TemirkhanN/alchemist/pkg/gui/geometry"
)

type Canvas interface {
	Show()
	Hide()
	Width() float64
	Height() float64
	IsVisible() bool
	Elements() []Canvas
	Position() geometry.Position
	IsUnderPosition(position geometry.Position) bool
	ChangePosition(position geometry.Position)
	Draw(on Layer)
}

type CommonCanvas struct {
	position geometry.Position
	visible  bool
}

func (canvas *CommonCanvas) Show() {
	canvas.visible = true
}

func (canvas *CommonCanvas) Hide() {
	canvas.visible = false
}

func (canvas CommonCanvas) IsVisible() bool {
	return canvas.visible
}

func (canvas CommonCanvas) Width() float64 {
	return 0
}

func (canvas CommonCanvas) Height() float64 {
	return 0
}

func (canvas CommonCanvas) Position() geometry.Position {
	return canvas.position
}

func (canvas *CommonCanvas) ChangePosition(position geometry.Position) {
	canvas.position = position
}

func (canvas CommonCanvas) IsUnderPosition(geometry.Position) bool {
	return false
}

func (canvas CommonCanvas) Elements() []Canvas {
	return nil
}

type SpriteCanvas struct {
	CommonCanvas
	sprite *Sprite
}

func NewSpriteCanvas(sprite *Sprite) *SpriteCanvas {
	return &SpriteCanvas{
		sprite: sprite,
		CommonCanvas: CommonCanvas{
			position: geometry.ZeroPosition,
			visible:  true,
		},
	}
}

func (canvas SpriteCanvas) Sprite() *Sprite {
	return canvas.sprite
}

func (canvas SpriteCanvas) IsUnderPosition(position geometry.Position) bool {
	width := canvas.sprite.Width()
	height := canvas.sprite.Height()

	bottomLeftX := canvas.position.X()
	bottomLeftY := canvas.position.Y()
	topRightX := canvas.position.X() + width
	topRightY := canvas.position.Y() + height

	posX := position.X()
	posY := position.Y()

	if (posX >= bottomLeftX && posX <= topRightX) && (posY >= bottomLeftY && posY <= topRightY) {
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
}

type TextCanvas struct {
	CommonCanvas
	text     string
	font     Font
	color    color.Color
	maxWidth float64
}

func NewTextCanvas(text string, font Font, maxWidth float64, fontColor ...color.Color) *TextCanvas {
	if len(fontColor) == 0 {
		fontColor = append(fontColor, color.Black)
	}

	canvas := &TextCanvas{
		color:    fontColor[0],
		text:     text,
		font:     font,
		maxWidth: maxWidth,
		CommonCanvas: CommonCanvas{
			position: geometry.ZeroPosition,
			visible:  true,
		},
	}
	canvas.AddLineBreaks()

	return canvas
}

func (canvas TextCanvas) Text() string {
	return canvas.text
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
}

func (canvas *TextCanvas) AddLineBreaks() {
	parts := strings.Split(canvas.text, " ")

	lineLength := 0.0
	offset := canvas.position.Y()

	for i, part := range parts {
		wordLength := canvas.font.calculateWidthInPixels(part)
		if lineLength+wordLength <= canvas.maxWidth || i == 0 {
			lineLength += wordLength

			continue
		}

		// we add line breaks only after at least first word
		parts[i-1] += "\n"
		lineLength = wordLength
		offset += canvas.font.atlas.LineHeight()
	}

	canvas.position = geometry.NewPosition(canvas.position.X(), offset)

	canvas.text = strings.Join(parts, " ")
}

func (canvas TextCanvas) Draw(on Layer) {
	if !canvas.IsVisible() {
		return
	}

	textPosition := canvas.Position().Add(geometry.NewPosition(0, canvas.Height()-canvas.Height()))

	basicTxt := text.New(pixel.V(textPosition.X(), textPosition.Y()), canvas.font.atlas)

	fmt.Fprintln(basicTxt, canvas.Text())

	basicTxt.DrawColorMask(on.target(), pixel.IM, canvas.color)
}

func (canvas SpriteCanvas) Draw(on Layer) {
	if !canvas.IsVisible() {
		return
	}

	fromLeftBottomCorner := geometry.NewPosition(
		canvas.Width()/2+canvas.Position().X(),
		canvas.Height()/2+canvas.Position().Y(),
	)

	canvas.sprite.src.Draw(on.target(), pixel.IM.Moved(pixel.Vec{
		X: fromLeftBottomCorner.X(),
		Y: fromLeftBottomCorner.Y(),
	}))
}
