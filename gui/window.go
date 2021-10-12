package gui

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

// Position todo make it mutable from package only.
type Position struct {
	X float64
	Y float64
}

type Window struct {
	graphics *Layer
	window   *pixelgl.Window
}

func CreateWindow(width float64, height float64, debugMode ...bool) *Window {
	cfg := pixelgl.WindowConfig{
		Title:                  "Alchemist",
		Icon:                   nil,
		Bounds:                 pixel.R(0, 0, width, height),
		Position:               pixel.Vec{X: 0, Y: 0},
		Monitor:                nil,
		Resizable:              false,
		Undecorated:            false,
		NoIconify:              false,
		AlwaysOnTop:            false,
		TransparentFramebuffer: false,
		VSync:                  true,
		Maximized:              false,
		Invisible:              false,
	}

	w, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	window := &Window{window: w, graphics: NewLayer(width, height, true)}

	if len(debugMode) != 0 && debugMode[0] {
		window.graphics.setDebugTool(window)
	}

	return window
}

func (w Window) Width() float64 {
	return w.graphics.Width()
}

func (w Window) Height() float64 {
	return w.graphics.Height()
}

func (w *Window) AddLayer(layer *Layer, position Position) {
	w.graphics.AddElement(layer, position)
}

func (w *Window) CreateSpriteCanvas(sprite *Sprite) *SpriteCanvas {
	return &SpriteCanvas{
		sprite: sprite,
		CommonCanvas: CommonCanvas{
			position:    ZeroPosition,
			visible:     true,
			drawnOn:     w,
			needsRedraw: true,
		},
	}
}

func (w *Window) CreateTextCanvas(text string) *TextCanvas {
	return &TextCanvas{
		text: text,
		CommonCanvas: CommonCanvas{
			position:    ZeroPosition,
			visible:     true,
			needsRedraw: true,
			drawnOn:     w,
		},
	}
}

func (w *Window) CreateButton(sprite *Sprite) *Button {
	return &Button{
		SpriteCanvas: SpriteCanvas{
			sprite: sprite,
			CommonCanvas: CommonCanvas{
				position:    ZeroPosition,
				drawnOn:     w,
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

func (w *Window) LeftButtonClicked() bool {
	return w.window.JustPressed(pixelgl.MouseButtonLeft) && w.window.MouseInsideWindow()
}

func (w Window) CursorPosition() Position {
	return Position{
		X: w.window.MousePosition().X,
		Y: w.window.MousePosition().Y,
	}
}

func (w Window) Closed() bool {
	return w.window.Closed()
}

func (w *Window) Refresh() {
	if w.Closed() {
		return
	}

	// only if w.window.MouseInsideWindow()
	cursorPosition := w.CursorPosition()
	w.handleMouseOut(w.graphics, cursorPosition)
	w.handleMouseOver(w.graphics, cursorPosition)

	if w.LeftButtonClicked() {
		w.handleLeftClick(w.graphics, cursorPosition)
	}

	w.draw()
	w.window.Update()
}

func (w *Window) draw() {
	if !w.graphics.NeedsRedraw() {
		return
	}

	w.window.Clear(colornames.White)
	w.graphics.Draw()
}

func (w *Window) handleLeftClick(graphics Drawer, clickedPosition Position) bool {
	if !graphics.isVisible() {
		return false
	}

	interactiveElement, isInteractiveElement := graphics.(InteractiveCanvas)
	if isInteractiveElement {
		if interactiveElement.IsUnderPosition(clickedPosition) {
			interactiveElement.EmitClick()

			return true
		}
	}

	// Interaction priority is LIFO. EmitClick over canvasB which is drawn over canvasA shall start from canvas B handle
	for i := len(graphics.Elements()) - 1; i >= 0; i-- {
		element := graphics.Elements()[i]
		if w.handleLeftClick(element, clickedPosition) {
			// stop further propagation
			return true
		}
	}

	return false
}

func (w *Window) handleMouseOver(graphics Drawer, onPosition Position) bool {
	if !graphics.isVisible() {
		return false
	}

	interactiveElement, isInteractiveElement := graphics.(InteractiveCanvas)
	if isInteractiveElement {
		if interactiveElement.IsUnderPosition(onPosition) {
			interactiveElement.EmitMouseOver()

			return true
		}
	}

	// Interaction priority is LIFO. EmitClick over canvasB which is drawn over canvasA shall start from canvas B handle
	for i := len(graphics.Elements()) - 1; i >= 0; i-- {
		element := graphics.Elements()[i]
		if w.handleMouseOver(element, onPosition) {
			// stop further propagation
			return true
		}
	}

	return false
}

func (w *Window) handleMouseOut(graphics Drawer, lastCursorPosition Position) {
	if !graphics.isVisible() {
		return
	}

	interactiveElement, isInteractiveElement := graphics.(InteractiveCanvas)
	if isInteractiveElement {
		if !interactiveElement.IsUnderPosition(lastCursorPosition) {
			interactiveElement.EmitMouseOut()
		}
	}

	// Interaction priority is LIFO. EmitClick over canvasB which is drawn over canvasA shall start from canvas B handle
	for i := len(graphics.Elements()) - 1; i >= 0; i-- {
		element := graphics.Elements()[i]
		w.handleMouseOut(element, lastCursorPosition)
	}
}

func (w *Window) drawSprite(sprite *Sprite, position Position) {
	sprite.draw(w, position)
}

func (w *Window) drawText(textValue string, position Position) {
	basicTxt := text.New(pixel.V(position.X, position.Y), basicAtlas)

	fmt.Fprintln(basicTxt, textValue)

	basicTxt.DrawColorMask(w.window, pixel.IM, colornames.Black)
}

var (
	basicAtlas   = text.NewAtlas(basicfont.Face7x13, text.ASCII)
	ZeroPosition = Position{X: 0, Y: 0}
)
