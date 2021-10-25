package graphics

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"

	"github.com/TemirkhanN/alchemist/pkg/gui/geometry"
)

type Layer interface {
	target() pixel.Target
}

type Scroll struct {
	currentOffsetFromTop float64
	maximumOffsetFromTop float64
	isAvailable          bool
}

type Layout struct {
	elements []Canvas
	visible  bool
	width    float64
	height   float64
	position geometry.Position
	scroll   Scroll
	drawnOn  *pixelgl.Canvas
}

func NewLayer(width float64, height float64, visible bool, scrollable ...bool) *Layout {
	scroll := Scroll{currentOffsetFromTop: 0, maximumOffsetFromTop: 0, isAvailable: false}
	if len(scrollable) == 1 && scrollable[0] {
		scroll.isAvailable = true
	}

	return &Layout{
		elements: nil,
		visible:  visible,
		width:    width,
		height:   height,
		position: geometry.ZeroPosition,
		scroll:   scroll,
		drawnOn:  pixelgl.NewCanvas(pixel.R(0, 0, width, height)),
	}
}

func (l Layout) target() pixel.Target {
	return l.drawnOn
}

func (l *Layout) IsUnderPosition(position geometry.Position) bool {
	buttonWidth := l.Width()
	buttonHeight := l.Height()

	bottomLeftX := l.position.X()
	bottomLeftY := l.position.Y()
	topRightX := l.position.X() + buttonWidth
	topRightY := l.position.Y() + buttonHeight

	posX := position.X()
	posY := position.Y()

	if (posX > bottomLeftX && posX < topRightX) && (posY > bottomLeftY && posY < topRightY) {
		return true
	}

	return false
}

func (l *Layout) Show() {
	l.visible = true
}

func (l *Layout) Hide() {
	l.visible = false
}

func (l *Layout) IsVisible() bool {
	return l.visible
}

func (l *Layout) Elements() []Canvas {
	return l.elements
}

func (l *Layout) AddElement(drawer Canvas, relativePosition geometry.Position) {
	if l.scroll.isAvailable && relativePosition.Y() < 0 {
		offset := -1 * relativePosition.Y()
		if offset > l.scroll.maximumOffsetFromTop {
			l.scroll.maximumOffsetFromTop = offset
		}
	}

	drawer.ChangePosition(relativePosition)
	l.elements = append(l.elements, drawer)
}

func (l *Layout) Clear() {
	l.elements = nil
	l.drawnOn.Clear(color.Transparent)
}

func (l Layout) Width() float64 {
	return l.width
}

func (l Layout) Height() float64 {
	return l.height
}

func (l *Layout) ChangePosition(position geometry.Position) {
	l.position = position
}

func (l Layout) Position() geometry.Position {
	return l.position
}

func (l Layout) intersects(element Canvas) bool {
	// element is placed from left or right side from layer
	if element.Position().X()+element.Width() < 0 || element.Position().X() > l.Width() {
		return false
	}

	// element is placed below the l
	if element.Position().Y()+element.Height() < 0 || element.Position().Y() > l.Height() {
		return false
	}

	return true
}

func (l Layout) IsScrollable() bool {
	if !l.visible || !l.scroll.isAvailable {
		return false
	}

	return true
}

// EmitVerticalScroll todo delegate to another system.
func (l *Layout) EmitVerticalScroll(vector float64) bool {
	if !l.IsScrollable() {
		return false
	}

	// We can scroll but there is no space for that
	if l.scroll.maximumOffsetFromTop < l.Height() ||
		(l.scroll.currentOffsetFromTop <= 0 && vector > 0) ||
		(l.scroll.currentOffsetFromTop >= l.scroll.maximumOffsetFromTop && vector < 0) {
		return true
	}

	l.scroll.currentOffsetFromTop -= vector

	if l.scroll.currentOffsetFromTop < 0 {
		vector = l.scroll.currentOffsetFromTop
		l.scroll.currentOffsetFromTop = 0
	}

	if l.scroll.maximumOffsetFromTop < l.scroll.currentOffsetFromTop {
		vector = l.scroll.maximumOffsetFromTop - l.scroll.currentOffsetFromTop
		l.scroll.currentOffsetFromTop = l.scroll.maximumOffsetFromTop
	}

	offset := geometry.NewPosition(0, vector)
	for _, element := range l.elements {
		element.ChangePosition(element.Position().Subtract(offset))
	}

	return true
}

func (l *Layout) Draw(on Layer) {
	if !l.IsVisible() {
		return
	}

	l.drawnOn.Clear(color.Transparent)

	for _, element := range l.Elements() {
		if !element.IsVisible() || !l.intersects(element) {
			continue
		}

		element.Draw(l)
	}

	fromLeftBottomCorner := geometry.NewPosition(
		l.Width()/2+l.Position().X(),
		l.Height()/2+l.Position().Y(),
	)

	l.drawnOn.Draw(on.target(), pixel.IM.Moved(pixel.Vec{
		X: fromLeftBottomCorner.X(),
		Y: fromLeftBottomCorner.Y(),
	}))
}
