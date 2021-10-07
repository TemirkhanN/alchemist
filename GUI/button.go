package GUI

type Button struct {
	CommonCanvas
	onclickfn func()
}

func (b *Button) SetClickHandler(handler func()) {
	b.onclickfn = handler
}

func (b *Button) Click() {
	b.onclickfn()
}
