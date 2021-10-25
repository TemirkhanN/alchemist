package graphics

import (
	"image"

	"github.com/faiface/pixel"

	"github.com/TemirkhanN/alchemist/pkg/gui/geometry"
)

type FrameSize struct {
	LeftBottom geometry.Position
	RightTop   geometry.Position
}

func (f FrameSize) toRectangle() pixel.Rect {
	rectangle := pixel.R(f.LeftBottom.X(), f.LeftBottom.Y(), f.RightTop.X(), f.RightTop.Y())

	return rectangle.Norm()
}

func (f FrameSize) Width() float64 {
	return f.RightTop.X() - f.LeftBottom.X()
}

func (f FrameSize) Height() float64 {
	return f.RightTop.Y() - f.LeftBottom.Y()
}

type Sprite struct {
	src *pixel.Sprite
}

func NewSprite(image image.Image) *Sprite {
	pic := pixel.PictureDataFromImage(image)

	return &Sprite{src: pixel.NewSprite(pic, pic.Bounds())}
}

func (s *Sprite) Width() float64 {
	return s.src.Frame().W()
}

func (s *Sprite) Height() float64 {
	return s.src.Frame().H()
}

func (s Sprite) Frame(frame FrameSize) *Sprite {
	return &Sprite{src: pixel.NewSprite(s.src.Picture(), frame.toRectangle())}
}
