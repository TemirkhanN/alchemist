package GUI

type Layer struct {
	elements []Canvas
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

func (layer *Layer) Draw() {
	if !layer.visible{
		return
	}

	for _, element := range layer.elements {
		element.Draw()
	}
	layer.needsRedraw = false
}

func (layer Layer) NeedsRedraw() bool {
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

func (layer *Layer) AddCanvas(canvas Canvas) {
	layer.elements = append(layer.elements, canvas)
}
