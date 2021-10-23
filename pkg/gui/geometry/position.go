package geometry

type Position struct {
	x float64
	y float64
}

func NewPosition(x float64, y float64) Position {
	return Position{x: x, y: y}
}

func (p Position) X() float64 {
	return p.x
}

func (p Position) Y() float64 {
	return p.y
}

func (p Position) Add(position Position) Position {
	return NewPosition(p.X()+position.X(), p.Y()+position.Y())
}

func (p Position) Subtract(position Position) Position {
	return NewPosition(p.X()-position.X(), p.Y()-position.Y())
}

var ZeroPosition = Position{x: 0, y: 0}
