package main

import (
	"embed"
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

// todo mistake number 1 - GUI shall be predefined structure. Not constructed dynamically.
func runGame(window *GUI.Window, assets GUI.Assets) {
	mortar := new(Mortar)
	mortar.alchemyLevel = MortarLevel(APPRENTICE)

	ingredientsLayout := createIngredientsLayout(window, assets, mortar)
	mainLayout := createMortarLayoutLayout(window, assets, ingredientsLayout, mortar)

	window.AddLayer(mainLayout)
	window.AddLayer(ingredientsLayout)

	for !window.Closed() {
		window.Refresh()
	}
}

func createMortarLayoutLayout(window *GUI.Window, assets GUI.Assets, ingredientsLayout *GUI.Layer, mortar *Mortar) *GUI.Layer {
	alchemyLayoutSprite := assets.GetSprite("mortar-interface")
	addIngredientButtonSprite := assets.GetSprite("btn.add-ingredient")
	createPotionButtonSprite := assets.GetSprite("btn.create-potion")
	exitButtonSprite := assets.GetSprite("btn.exit")

	mortarLayout := new(GUI.Layer)

	// todo
	descriptionText := window.CreateTextCanvas("Description here", GUI.Position{X: 555, Y: 430})

	mortarLayout.AddCanvas(window.CreateSpriteCanvas(alchemyLayoutSprite, GUI.Position{}))

	ingredientSelectors := []*GUI.Button{
		window.CreateButton(addIngredientButtonSprite, GUI.Position{X: 187, Y: 180}),
		window.CreateButton(addIngredientButtonSprite, GUI.Position{X: 187, Y: 250}),
		window.CreateButton(addIngredientButtonSprite, GUI.Position{X: 187, Y: 320}),
		window.CreateButton(addIngredientButtonSprite, GUI.Position{X: 187, Y: 390}),
	}

	for _, ingredientButton := range ingredientSelectors {
		ingredientButton.SetClickHandler(func() {
			ingredientsLayout.Show()
		})
		mortarLayout.AddCanvas(ingredientButton)
	}

	createPotionButton := window.CreateButton(createPotionButtonSprite, GUI.Position{X: 253, Y: 116})
	createPotionButton.SetClickHandler(func() {
		potion, err := mortar.Pestle()
		if err != nil {
			log.Fatal(err)
		}
		descriptionText.ChangeText(potion.Description())
	})

	exitButton := window.CreateButton(exitButtonSprite, GUI.Position{X: 646, Y: 115})
	exitButton.SetClickHandler(func() { os.Exit(0) })

	mortarLayout.AddCanvas(createPotionButton)
	mortarLayout.AddCanvas(exitButton)
	mortarLayout.AddCanvas(descriptionText)
	mortarLayout.Show()

	return mortarLayout
}

func createIngredientsLayout(window *GUI.Window, assets GUI.Assets, mortar *Mortar) *GUI.Layer {
	backpack := initStorage()
	exitButtonSprite := assets.GetSprite("btn.exit")
	ingredientsLayoutSprite := assets.GetSprite("ingredients-interface")

	ingredientsLayout := new(GUI.Layer)
	ingredientsLayout.AddCanvas(window.CreateSpriteCanvas(ingredientsLayoutSprite, GUI.Position{}))

	closeIngredientsLayoutButton := window.CreateButton(exitButtonSprite, GUI.Position{X: 410, Y: 65})
	closeIngredientsLayoutButton.SetClickHandler(func() { ingredientsLayout.Hide() })

	lastIngredientPosition := GUI.Position{X: 50, Y: 500}
	for _, ingredient := range backpack.All() {
		button := window.CreateButton(assets.GetSprite(ingredient.sprite), lastIngredientPosition)
		button.SetClickHandler(addIngredient(ingredient, mortar, ingredientsLayout))

		ingredientsLayout.AddCanvas(button)
		lastIngredientPosition.Y -= 64
	}

	ingredientsLayout.AddCanvas(closeIngredientsLayoutButton)
	ingredientsLayout.Hide()

	return ingredientsLayout
}

// todo
func addIngredient(ingredient Ingredient, mortar *Mortar, layer *GUI.Layer) func () {
	return func () {
		err := mortar.AddIngredient(ingredient)
		// TODO remove after creating structured GUI interface
		if err != nil {
			log.Fatal(err)
		}
		layer.Hide()
	}
}
