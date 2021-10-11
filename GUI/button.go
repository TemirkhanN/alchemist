package GUI

type Button struct {
	SpriteCanvas
	onclickFn     func()
	onmouseoverFn func()
	onmouseoutFn  func()
	hovered       bool
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
