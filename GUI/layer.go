package GUI

type Layer struct {
	elements []Drawer
	visible bool
	needsRedraw bool
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

func (layer *Layer) Draw() {
	if !layer.visible{
		return
	}

	for _, element := range layer.elements {
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
