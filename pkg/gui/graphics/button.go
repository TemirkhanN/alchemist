package graphics

import (
	"github.com/TemirkhanN/alchemist/pkg/gui/geometry"
)

type InteractiveCanvas interface {
	EmitClick(position geometry.Position)
	SetClickHandler(func(position geometry.Position))
	EmitMouseOver()
	SetMouseOverHandler(func())
	EmitMouseOut()
	SetMouseOutHandler(func())
	Canvas
}

type Button struct {
	SpriteCanvas
	onclickFn     func(position geometry.Position)
	onmouseoverFn func()
	onmouseoutFn  func()
	hovered       bool
}

func NewButton(sprite *Sprite) *Button {
	return &Button{
		SpriteCanvas: SpriteCanvas{
			sprite: sprite,
			CommonCanvas: CommonCanvas{
				position: geometry.ZeroPosition,
				visible:  true,
			},
		},
		onclickFn:     nil,
		onmouseoverFn: nil,
		onmouseoutFn:  nil,
		hovered:       false,
	}
}

func (b *Button) SetClickHandler(handler func(position geometry.Position)) {
	b.onclickFn = handler
}

func (b *Button) EmitClick(position geometry.Position) {
	if b.onclickFn != nil {
		b.onclickFn(position)
	}
}

func (b *Button) SetMouseOverHandler(handler func()) {
	b.onmouseoverFn = handler
}

func (b *Button) EmitMouseOver() {
	if !b.hovered && b.onmouseoverFn != nil {
		b.hovered = true
		b.onmouseoverFn()
	}
}

func (b *Button) SetMouseOutHandler(handler func()) {
	b.onmouseoutFn = handler
}

func (b *Button) EmitMouseOut() {
	if b.hovered && b.onmouseoutFn != nil {
		b.hovered = false
		b.onmouseoutFn()
	}
}
