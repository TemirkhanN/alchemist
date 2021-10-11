package GUI

import (
	"github.com/faiface/pixel"
)

type Sprite struct {
	src *pixel.Sprite
}

func (s *Sprite) Width() float64 {
	return s.src.Picture().Bounds().W()
}

func (s *Sprite) Height() float64 {
	return s.src.Picture().Bounds().H()
}

func (s *Sprite) draw(window *Window, position Position) {
	fromLeftBottomCorner := pixel.Vec{
		X: s.src.Picture().Bounds().Center().X + position.X,
		Y: s.src.Picture().Bounds().Center().Y + position.Y,
	}

	s.src.Draw(window.window, pixel.IM.Moved(fromLeftBottomCorner))
}
