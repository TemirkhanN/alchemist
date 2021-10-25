package animation

import (
	"time"

	"github.com/TemirkhanN/alchemist/pkg/gui/geometry"
	"github.com/TemirkhanN/alchemist/pkg/gui/graphics"
)

func Move(canvas graphics.Canvas, to geometry.Position, speed int64) {
	from := canvas.Position()

	x := int(from.X())
	y := int(from.Y())

	completed := true

	if x != int(to.X()) {
		completed = false

		if to.X() > from.X() {
			x++
		} else {
			x--
		}
	}

	if y != int(to.Y()) {
		completed = false

		if to.Y() > from.Y() {
			y++
		} else {
			y--
		}
	}

	time.AfterFunc(time.Second/time.Duration(speed), func() {
		canvas.ChangePosition(geometry.NewPosition(float64(x), float64(y)))

		if completed {
			return
		}

		Move(canvas, to, speed)
	})
}
