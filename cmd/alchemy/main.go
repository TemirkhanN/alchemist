package main

import (
	_ "image/png"

	"github.com/faiface/pixel/pixelgl"

	"github.com/TemirkhanN/alchemist/internal/game"
)

func main() {
	pixelgl.Run(func() {
		game.Launch(1024, 768, 8, false)
	})
}
