package gui

type Layer struct {
	elements    []Drawer
	visible     bool
	needsRedraw bool
	width       float64
	height      float64
}

func NewLayer(width float64, height float64, visible bool) *Layer {
	return &Layer{
		visible:     visible,
		needsRedraw: true,
		width:       width,
		height:      height,
		elements:    nil,
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

	renderedHeight := 0.0
	for _, element := range layer.elements {
		renderedHeight += element.Height()
		if renderedHeight > layer.Height() {
			break
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

func (layer *Layer) AddElement(canvas Drawer) {
	layer.elements = append(layer.elements, canvas)
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

// Height todo fix invalid calculation because of different positioning. Elements on layer are not positioned in rows.
func (layer *Layer) Height() float64 {
	if layer.height != 0.0 {
		return layer.height
	}

	calculatedHeight := 0.0
	for _, element := range layer.elements {
		calculatedHeight += element.Height()
	}

	return calculatedHeight
}
