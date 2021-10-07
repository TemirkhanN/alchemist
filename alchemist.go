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
	assets := GUI.Assets{}
	// todo shall filesystem be passed by reference or not?
	err := assets.RegisterAssets("assets/sprites", &spritesFs)

	if err != nil {
		log.Fatal(err)
	}

	window := GUI.CreateWindow(1024, 768)
	runGame(window, assets)
}

func runGame(window *GUI.Window, assets GUI.Assets) {
	ingredientsLayout := createIngredientsLayout(window, assets)
	mainLayout := createMortarLayoutLayout(window, assets, ingredientsLayout)

	window.AddLayer(mainLayout)
	window.AddLayer(ingredientsLayout)

	for !window.Closed() {
		window.Refresh()
	}
}

func createMortarLayoutLayout(window *GUI.Window, assets GUI.Assets, ingredientsLayout *GUI.Layer) *GUI.Layer {
	mortar := new(Mortar)

	alchemyLayoutSprite := assets.GetSprite("mortar-interface")
	addIngredientButtonSprite := assets.GetSprite("btn.add-ingredient")
	createPotionButtonSprite := assets.GetSprite("btn.create-potion")
	exitButtonSprite := assets.GetSprite("btn.exit")

	mortarLayout := new(GUI.Layer)

	mortarLayout.AddCanvas(window.CreateCanvas(alchemyLayoutSprite, GUI.Position{}))

	ingredientSelectors := []*GUI.Button{
		window.CreateButton(addIngredientButtonSprite, GUI.Position{X: 187, Y: 180}),
		window.CreateButton(addIngredientButtonSprite, GUI.Position{X: 187, Y: 250}),
		window.CreateButton(addIngredientButtonSprite, GUI.Position{X: 187, Y: 320}),
		window.CreateButton(addIngredientButtonSprite, GUI.Position{X: 187, Y: 390}),
	}

	for _, ingredientButton := range ingredientSelectors {
		ingredientButton.SetClickHandler(func() { ingredientsLayout.Show() })
		mortarLayout.AddCanvas(ingredientButton)
	}

	createPotionButton := window.CreateButton(createPotionButtonSprite, GUI.Position{X: 253, Y: 116})
	createPotionButton.SetClickHandler(func() { createPotion(mortar) })

	exitButton := window.CreateButton(exitButtonSprite, GUI.Position{X: 646, Y: 115})
	exitButton.SetClickHandler(func() { os.Exit(0) })

	mortarLayout.AddCanvas(createPotionButton)
	mortarLayout.AddCanvas(exitButton)
	mortarLayout.Show()

	return mortarLayout
}

func createIngredientsLayout(window *GUI.Window, assets GUI.Assets) *GUI.Layer {
	//ingredientRepository := initStorage()
	exitButtonSprite := assets.GetSprite("btn.exit")
	ingredientsLayoutSprite := assets.GetSprite("ingredients-interface")

	ingredientsLayout := new(GUI.Layer)
	ingredientsLayout.AddCanvas(window.CreateCanvas(ingredientsLayoutSprite, GUI.Position{}))

	closeIngredientsLayoutButton := window.CreateButton(exitButtonSprite, GUI.Position{X: 410, Y: 65})
	closeIngredientsLayoutButton.SetClickHandler(func() { ingredientsLayout.Hide() })

	ingredientsLayout.AddCanvas(closeIngredientsLayoutButton)
	ingredientsLayout.Hide()

	return ingredientsLayout
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
