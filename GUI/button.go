package GUI

type Button struct {
	SpriteCanvas
	onclickFn func()
	onmouseoverFn func()
}

func (b *Button) SetClickHandler(handler func()) {
	b.onclickFn = handler
}

func (b *Button) EmitClick() {
	b.onclickFn()
}

func (b *Button) SetMouseOverHandler(handler func()) {
	b.onmouseoverFn = handler
}

func (b *Button) EmitMouseOver() {
	b.onmouseoverFn()
}
