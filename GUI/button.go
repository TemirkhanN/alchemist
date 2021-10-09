package GUI

type Button struct {
	SpriteCanvas
	onclickfn func()
}

func (b *Button) SetClickHandler(handler func()) {
	b.onclickfn = handler
}

func (b *Button) EmitClick() {
	b.onclickfn()
}
