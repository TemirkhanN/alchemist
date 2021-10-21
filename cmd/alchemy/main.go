package main

import (
	_ "image/png"

	"github.com/faiface/pixel/pixelgl"

	"github.com/TemirkhanN/alchemist/internal/game"
	"github.com/TemirkhanN/alchemist/pkg/alchemy/alchemist"
)

func main() {
	hardcodedLevel := 26
	hardcodedLuck := 5
	hardcodedMortarLevel := alchemist.EquipmentNovice
	launcher := game.NewGame(hardcodedLevel, hardcodedLuck, hardcodedMortarLevel)

	pixelgl.Run(func() {
		launcher.Launch(1024, 768, 8, false)
	})
}
