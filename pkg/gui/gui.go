package gui

func CreateLayer(width float64, height float64, visible bool, scrollable ...bool) *Layer {
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
		position:    ZeroPosition,
		scroll:      scroll,
	}
}

func CreateSpriteCanvas(sprite *Sprite) *SpriteCanvas {
	return &SpriteCanvas{
		sprite: sprite,
		CommonCanvas: CommonCanvas{
			position:    ZeroPosition,
			visible:     true,
			needsRedraw: true,
		},
	}
}

func CreateTextCanvas(text string, font Font, maxWidth float64) *TextCanvas {
	canvas := &TextCanvas{
		text:     text,
		font:     font,
		maxWidth: maxWidth,
		CommonCanvas: CommonCanvas{
			position:    ZeroPosition,
			visible:     true,
			needsRedraw: true,
		},
	}
	canvas.AddLineBreaks()

	return canvas
}

func CreateButton(sprite *Sprite) *Button {
	return &Button{
		SpriteCanvas: SpriteCanvas{
			sprite: sprite,
			CommonCanvas: CommonCanvas{
				position:    ZeroPosition,
				visible:     true,
				needsRedraw: true,
			},
		},
		onclickFn:     nil,
		onmouseoverFn: nil,
		onmouseoutFn:  nil,
		hovered:       false,
	}
}
