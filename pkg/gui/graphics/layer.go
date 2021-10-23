package graphics

import "github.com/TemirkhanN/alchemist/pkg/gui/geometry"

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
	position    geometry.Position
	scroll      Scroll
}

func NewLayer(width float64, height float64, visible bool, scrollable ...bool) *Layer {
	scroll := Scroll{currentOffsetFromTop: 0, maximumOffsetFromTop: 0, isAvailable: false}
	if len(scrollable) == 1 && scrollable[0] {
		scroll.isAvailable = true
	}

	return &Layer{
		elements:    nil,
		visible:     visible,
		needsRedraw: true,
		width:       width,
		height:      height,
		position:    geometry.ZeroPosition,
		scroll:      scroll,
	}
}

func (layer *Layer) IsUnderPosition(position geometry.Position) bool {
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

func (layer *Layer) IsVisible() bool {
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
		if !element.IsVisible() || !layer.CanFullyFit(element) {
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

func (layer *Layer) AddElement(drawer Canvas, relativePosition geometry.Position) {
	if layer.scroll.isAvailable && relativePosition.Y() < 0 {
		offset := -1 * relativePosition.Y()
		if offset > layer.scroll.maximumOffsetFromTop {
			layer.scroll.maximumOffsetFromTop = offset
		}
	}

	drawer.setPosition(layer.position.Add(relativePosition))
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

func (layer *Layer) setPosition(position geometry.Position) {
	previousPosition := layer.position
	layer.position = position

	for _, element := range layer.elements {
		element.setPosition(element.Position().Subtract(previousPosition).Add(layer.position))
	}

	layer.needsRedraw = true
}

func (layer Layer) Position() geometry.Position {
	return layer.position
}

func (layer Layer) CanFullyFit(element Canvas) bool {
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

func (layer Layer) isScrollable() bool {
	if !layer.visible || !layer.scroll.isAvailable {
		return false
	}

	return true
}

// EmitVerticalScroll todo delegate to another system.
func (layer *Layer) EmitVerticalScroll(vector float64) bool {
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

	offset := geometry.NewPosition(0, vector)
	for _, element := range layer.elements {
		element.setPosition(element.Position().Subtract(offset))
	}

	layer.needsRedraw = true

	return true
}
