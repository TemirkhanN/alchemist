package game

import (
	"log"
	"strings"

	"github.com/TemirkhanN/alchemist/assets"
	"github.com/TemirkhanN/alchemist/pkg/alchemy/alchemist"
	"github.com/TemirkhanN/alchemist/pkg/alchemy/ingredient"
	"github.com/TemirkhanN/alchemist/pkg/gui"
)

var (
	gameAssets = func() *gui.Assets {
		loadedAssets := new(gui.Assets)
		err := loadedAssets.RegisterAssets("sprites", assets.SpritesFs)
		if err != nil {
			log.Fatal(err)
		}

		return loadedAssets
	}()

	potionEffectFrameSize = gui.FrameSize{
		LeftBottom: gui.NewPosition(10, 15),
		RightTop:   gui.NewPosition(35, 40),
	}

	tesOblivion24Font = gui.LoadFont("TESOblivionFont", "assets/font/Kingthings Petrock.ttf", 24)
)

// GetIngredientSprite todo move to more appropriate place.
func getIngredientSprite(ingr ingredient.Ingredient) *gui.Sprite {
	spriteName := "ingr." + strings.ReplaceAll(strings.ToLower(ingr.Name()), "'", "")

	return gameAssets.GetSprite(spriteName)
}

type Game struct {
	alchemistLevel int
	alchemistLuck  int
	mortarLevel    alchemist.EquipmentLevel
}

func NewGame(alchemistLevel int, alchemistLuck int, mortarLevel alchemist.EquipmentLevel) *Game {
	if alchemistLevel < 1 || alchemistLevel > 100 || alchemistLuck < 1 || alchemistLuck > 100 {
		log.Fatal("alchemist level/luck is invalid")
	}

	return &Game{
		alchemistLevel: alchemistLevel,
		alchemistLuck:  alchemistLuck,
		mortarLevel:    mortarLevel,
	}
}

func (g *Game) Launch(windowWidth float64, windowHeight float64, scrollSpeed uint8, debugMode bool) {
	window := gui.NewWindow(gui.WindowConfig{
		Title:       "Alchemist",
		Width:       windowWidth,
		Height:      windowHeight,
		DebugMode:   debugMode,
		Position:    gui.ZeroPosition,
		ScrollSpeed: scrollSpeed,
	})

	mortar := alchemist.NewMortar(g.mortarLevel)
	player := alchemist.NewAlchemist(g.alchemistLevel, g.alchemistLuck, mortar)

	primaryScreen := newPrimaryLayout(window, player)
	backpackScreen := newBackpackLayout(window, player)

	window.AddLayer(primaryScreen.graphics, gui.ZeroPosition)
	window.AddLayer(backpackScreen.graphics, gui.ZeroPosition)

	for !window.Closed() {
		window.Refresh()
	}
}
