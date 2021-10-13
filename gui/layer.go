package gui

type Scroll struct {
	currentOffset float64
	maximumOffset float64
	isAvailable   bool
}

type Layer struct {
	elements    []Drawer
	visible     bool
	needsRedraw bool
	width       float64
	height      float64
	position    Position
	scroll      Scroll
	_debug      *Window
}

func NewLayer(width float64, height float64, visible bool, scrollable ...bool) *Layer {
	scroll := Scroll{currentOffset: 0, maximumOffset: 0, isAvailable: false}
	if len(scrollable) == 1 && scrollable[0] {
		scroll.isAvailable = true
	}

	return &Layer{
		elements:    nil,
		visible:     visible,
		needsRedraw: true,
		width:       width,
		height:      height,
		position:    ZeroPosition,
		_debug:      nil,
		scroll:      scroll,
	}
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

func (layer *Layer) Draw() {
	if !layer.visible {
		return
	}

	if layer.isDebugModeOn() {
		layer.debugDraw()
	}

	for _, element := range layer.elements {
		if !element.isVisible() {
			continue
		}

		if layer.canFullyFit(element) {
			element.Draw()
		}
	}

	layer.needsRedraw = false
}

func (layer *Layer) NeedsRedraw() bool {
	if layer.needsRedraw {
		return true
	}

	for _, element := range layer.elements {
		if element.NeedsRedraw() {
			return true
		}
	}

	return false
}

func (layer *Layer) Elements() []Drawer {
	return layer.elements
}

func (layer *Layer) AddElement(drawer Drawer, relativePosition Position) {
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

func (layer Layer) actualHeight() float64 {
	actualHeight := 0.0
	for _, element := range layer.elements {
		actualHeight += element.Height()
	}

	return actualHeight
}

func (layer *Layer) setPosition(position Position) {
	layer.position = position
}

func (layer Layer) Position() Position {
	return layer.position
}

func (layer Layer) canFullyFit(element Drawer) bool {
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
	return layer.scroll.isAvailable && layer.visible
}

func (layer *Layer) emitVerticalScroll(vector float64) bool {
	if !layer.isScrollable() {
		return false
	}

	layerMaxHeight := layer.actualHeight()
	if layerMaxHeight <= layer.Height() {
		return false
	}

	// We can scroll but there is no space for that
	if (layer.scroll.currentOffset == 0 && vector > 0) ||
		(layer.scroll.currentOffset == layerMaxHeight && vector < 0) {
		return true
	}

	layer.scroll.currentOffset -= vector

	if layer.scroll.currentOffset < 0 {
		layer.scroll.currentOffset = 0
	}

	if layerMaxHeight < layer.scroll.currentOffset {
		layer.scroll.currentOffset = layerMaxHeight
	}

	offset := NewPosition(0, vector)
	for _, element := range layer.elements {
		element.setPosition(element.Position().relative(offset))
	}

	layer.needsRedraw = true

	return true
}
