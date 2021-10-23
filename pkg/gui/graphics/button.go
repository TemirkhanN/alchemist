package graphics

import (
	"github.com/TemirkhanN/alchemist/pkg/gui/geometry"
)

type InteractiveCanvas interface {
	EmitClick()
	SetClickHandler(func())
	EmitMouseOver()
	SetMouseOverHandler(func())
	EmitMouseOut()
	SetMouseOutHandler(func())
	Canvas
}

type Button struct {
	SpriteCanvas
	onclickFn     func()
	onmouseoverFn func()
	onmouseoutFn  func()
	hovered       bool
}

func NewButton(sprite *Sprite) *Button {
	return &Button{
		SpriteCanvas: SpriteCanvas{
			sprite: sprite,
			CommonCanvas: CommonCanvas{
				position:    geometry.ZeroPosition,
				visible:     true,
				needsRedraw: true,
			},
		},
		onclickFn:     nil,
		onmouseoverFn: nil,
		onmouseoutFn:  nil,
		hovered:       false,
	}
}

func (b *Button) SetClickHandler(handler func()) {
	b.onclickFn = handler
}

func (b *Button) EmitClick() {
	if b.onclickFn != nil {
		b.onclickFn()
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
