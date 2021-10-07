package GUI

type Button struct {
	CommonCanvas
	onclickfn func()
}

func (b Button) IsUnderPosition(position Position) bool {
	if !b.visible {
		return false
	}

	buttonWidth := b.sprite.Picture().Bounds().W()
	buttonHeight := b.sprite.Picture().Bounds().H()

	bottomLeftX := b.position.X
	bottomLeftY := b.position.Y
	topRightX := b.position.X + buttonWidth
	topRightY := b.position.Y + buttonHeight

	if (position.X > bottomLeftX && position.X < topRightX) && (position.Y > bottomLeftY && position.Y < topRightY) {
		return true
	}

	return false
}

func (b *Button) SetClickHandler(handler func()) {
	b.onclickfn = handler
}

func (b *Button) Click() {
	b.onclickfn()
}
