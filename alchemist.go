package main

import (
	"embed"
	"fmt"
	"github.com/TemirkhanN/alchemist/GUI"
	"github.com/faiface/pixel/pixelgl"
	_ "image/png"
	"log"
	"os"
)

//go:embed assets/sprites/*.png
var spritesFs embed.FS

func main() {
	pixelgl.Run(launch)
}

func launch() {
	ingredientRepository := initStorage()
	assets := GUI.Assets{}
	// todo shall filesystem be passed by reference or not?
	err := assets.RegisterAssets("assets/sprites", &spritesFs)

	if err != nil {
		log.Fatal(err)
	}

	window := GUI.CreateWindow(1024, 768)
	runGame(window, assets, ingredientRepository)
}

func runGame(window *GUI.Window, assets GUI.Assets, storage IngredientFinder) {
	mortar := new(Mortar)

	alchemyWindowSprite := assets.GetSprite("mortar-interface")
	addIngredientButtonSprite := assets.GetSprite("btn.add-ingredient")
	createPotionButtonSprite := assets.GetSprite("btn.create-potion")
	exitButtonSprite := assets.GetSprite("btn.exit")

	alchemyWindowLayer := GUI.Layer{}

	mortarBackground := window.CreateCanvas(alchemyWindowSprite, GUI.Position{})
	alchemyWindowLayer.AddCanvas(mortarBackground)

	ingredientSelectors := []*GUI.Button{
		window.CreateButton(addIngredientButtonSprite, GUI.Position{X: 187, Y: 180}),
		window.CreateButton(addIngredientButtonSprite, GUI.Position{X: 187, Y: 250}),
		window.CreateButton(addIngredientButtonSprite, GUI.Position{X: 187, Y: 320}),
		window.CreateButton(addIngredientButtonSprite, GUI.Position{X: 187, Y: 390}),
	}

	for _, ingredientButton := range ingredientSelectors {
		ingredientButton.SetClickHandler(func() {fmt.Println("Show ingredient list window")})
		alchemyWindowLayer.AddCanvas(ingredientButton)
	}

	createPotionButton := window.CreateButton(createPotionButtonSprite, GUI.Position{X: 253, Y: 116})
	createPotionButton.SetClickHandler(func() { createPotion(mortar) })

	exitButton := window.CreateButton(exitButtonSprite, GUI.Position{X: 646, Y: 115})
	exitButton.SetClickHandler(func() { os.Exit(0) })

	alchemyWindowLayer.AddCanvas(createPotionButton)
	alchemyWindowLayer.AddCanvas(exitButton)

	window.AddLayer(alchemyWindowLayer)

	for !window.Closed() {
		window.Refresh()
	}
}

func createPotion(mortar *Mortar) {
	potion, err := mortar.Pestle()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Congratulations! You've created a potion!")
	fmt.Println(potion.Description())
}
