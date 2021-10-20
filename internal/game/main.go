package game

import (
	"strings"

	"github.com/TemirkhanN/alchemist/assets"
	"github.com/TemirkhanN/alchemist/pkg/alchemy/alchemist"
	"github.com/TemirkhanN/alchemist/pkg/alchemy/ingredient"
	"github.com/TemirkhanN/alchemist/pkg/gui"
)

func Launch(windowWidth float64, windowHeight float64) {
	window := gui.NewWindow(gui.WindowConfig{
		Title:       "Alchemist",
		Width:       windowWidth,
		Height:      windowHeight,
		DebugMode:   false,
		Position:    gui.ZeroPosition,
		ScrollSpeed: 8,
	})

	alchemyLevelHardcoded := 26
	luckLevelHardcoded := 5
	mortar := alchemist.NewMortar(alchemist.EquipmentNovice)
	player := alchemist.NewAlchemist(alchemyLevelHardcoded, luckLevelHardcoded, mortar)

	mainLayout := newPrimaryLayout(window, player)
	backpackLayout := newBackpackLayout(window, player)

	window.AddLayer(mainLayout.graphics, gui.ZeroPosition)
	window.AddLayer(backpackLayout.graphics, gui.ZeroPosition)

	for !window.Closed() {
		window.Refresh()
	}
}

// GetIngredientSprite todo move to more appropriate place.
func getIngredientSprite(ingr ingredient.Ingredient) *gui.Sprite {
	spriteName := "ingr." + strings.ReplaceAll(strings.ToLower(ingr.Name()), "'", "")

	return assets.TESAssets.GetSprite(spriteName)
}

var potionEffectFrameSize = gui.FrameSize{
	LeftBottom: gui.NewPosition(10, 15),
	RightTop:   gui.NewPosition(35, 40),
}

var tesOblivion24Font = gui.LoadFont("TESOblivionFont", "assets/font/Kingthings Petrock.ttf", 24)
