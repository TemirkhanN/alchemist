package gui

type Scroll struct {
	currentOffsetFromTop float64
	maximumOffsetFromTop float64
	isAvailable          bool
}

type Layer struct {
	elements    []Canvas
	visible     bool
	needsRedraw bool
	width       float64
	height      float64
	position    Position
	scroll      Scroll
}

func (layer *Layer) IsUnderPosition(position Position) bool {
	buttonWidth := layer.Width()
	buttonHeight := layer.Height()

	bottomLeftX := layer.position.X()
	bottomLeftY := layer.position.Y()
	topRightX := layer.position.X() + buttonWidth
	topRightY := layer.position.Y() + buttonHeight

	posX := position.X()
	posY := position.Y()

	if (posX > bottomLeftX && posX < topRightX) && (posY > bottomLeftY && posY < topRightY) {
		return true
	}

	return false
}

func (layer *Layer) Show() {
	layer.visible = true
	layer.needsRedraw = true
}

func (layer *Layer) Hide() {
	layer.visible = false
	layer.needsRedraw = true
}

func (layer *Layer) isVisible() bool {
	return layer.visible
}

func (layer *Layer) NeedsRedraw() bool {
	if !layer.visible {
		return false
	}

	if layer.needsRedraw {
		return true
	}

	for _, element := range layer.elements {
		if !element.isVisible() || !layer.canFullyFit(element) {
			continue
		}

		if element.NeedsRedraw() {
			return true
		}
	}

	return false
}

func (layer *Layer) Elements() []Canvas {
	return layer.elements
}

func (layer *Layer) AddElement(drawer Canvas, relativePosition Position) {
	if layer.scroll.isAvailable && relativePosition.Y() < 0 {
		offset := -1 * relativePosition.Y()
		if offset > layer.scroll.maximumOffsetFromTop {
			layer.scroll.maximumOffsetFromTop = offset
		}
	}

	drawer.setPosition(layer.position.absolute(relativePosition))
	layer.elements = append(layer.elements, drawer)
}

func (layer *Layer) Clear() {
	layer.elements = nil
}

func (layer Layer) Width() float64 {
	return layer.width
}

func (layer Layer) Height() float64 {
	return layer.height
}

func (layer *Layer) setPosition(position Position) {
	previousPosition := layer.position
	layer.position = position

	for _, element := range layer.elements {
		element.setPosition(element.Position().relative(previousPosition).absolute(layer.position))
	}

	layer.needsRedraw = true
}

func (layer Layer) Position() Position {
	return layer.position
}

func (layer Layer) canFullyFit(element Canvas) bool {
	// element is placed left from the layer
	if layer.Position().X() > element.Position().X() {
		return false
	}

	// element is not fitting on layer width
	if layer.Position().X()+layer.Width() < element.Position().X()+element.Width() {
		return false
	}

	// element is placed below the layer
	if layer.Position().Y() > element.Position().Y() {
		return false
	}

	if layer.Position().Y()+layer.Height() < element.Position().Y()+element.Height() {
		return false
	}

	return true
}

func (layer *Layer) isUnderPosition(position Position) bool {
	buttonWidth := layer.Width()
	buttonHeight := layer.Height()

	bottomLeftX := layer.position.X()
	bottomLeftY := layer.position.Y()
	topRightX := layer.position.X() + buttonWidth
	topRightY := layer.position.Y() + buttonHeight

	posX := position.X()
	posY := position.Y()

	if (posX > bottomLeftX && posX < topRightX) && (posY > bottomLeftY && posY < topRightY) {
		return true
	}

	return false
}

func (layer Layer) isScrollable() bool {
	if !layer.visible || !layer.scroll.isAvailable {
		return false
	}

	return true
}

func (layer *Layer) emitVerticalScroll(vector float64) bool {
	if !layer.isScrollable() {
		return false
	}

	// We can scroll but there is no space for that
	if layer.scroll.maximumOffsetFromTop < layer.Height() ||
		(layer.scroll.currentOffsetFromTop <= 0 && vector > 0) ||
		(layer.scroll.currentOffsetFromTop >= layer.scroll.maximumOffsetFromTop && vector < 0) {
		return true
	}

	layer.scroll.currentOffsetFromTop -= vector

	if layer.scroll.currentOffsetFromTop < 0 {
		vector = layer.scroll.currentOffsetFromTop
		layer.scroll.currentOffsetFromTop = 0
	}

	if layer.scroll.maximumOffsetFromTop < layer.scroll.currentOffsetFromTop {
		vector = layer.scroll.maximumOffsetFromTop - layer.scroll.currentOffsetFromTop
		layer.scroll.currentOffsetFromTop = layer.scroll.maximumOffsetFromTop
	}

	offset := NewPosition(0, vector)
	for _, element := range layer.elements {
		element.setPosition(element.Position().relative(offset))
	}

	layer.needsRedraw = true

	return true
}
