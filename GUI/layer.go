package GUI

type Layer struct {
	elements []Canvas
}

func (layer *Layer) Draw() {
	for _, element := range layer.elements {
		element.Draw()
	}
}

func (layer Layer) NeedsRedraw() bool {
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
