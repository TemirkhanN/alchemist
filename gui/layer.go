package gui

type Layer struct {
	elements    []Drawer
	visible     bool
	needsRedraw bool
	width       float64
	height      float64
	position    Position
	_debug      *Window
}

func NewLayer(width float64, height float64, visible bool) *Layer {
	return &Layer{
		elements:    nil,
		visible:     visible,
		needsRedraw: true,
		width:       width,
		height:      height,
		position:    ZeroPosition,
		_debug:      nil,
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

		if !layer.canFullyFit(element) {
			element.Hide()
		}

		element.Draw()
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
	absPosition := NewPosition(layer.position.X()+relativePosition.X(), layer.position.Y()+relativePosition.Y())

	drawer.setPosition(absPosition)
	layer.elements = append(layer.elements, drawer)
}

func (layer *Layer) Clear() {
	layer.elements = nil
}

func (layer *Layer) SetSize(width float64, height float64) {
	layer.width = width
	layer.height = height
}

func (layer *Layer) Width() float64 {
	return layer.width
}

func (layer *Layer) Height() float64 {
	return layer.height
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
