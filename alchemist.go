package main

import (
	"fmt"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	_ "image/png"
	"os"
)

func main() {
	pixelgl.Run(launch)
}

func launch() {
	ingredientRepository := initStorage()

	window := createWindow(1024, 768)
	runGame(window, ingredientRepository)
}

func runGame(window *pixelgl.Window, storage IngredientFinder) {
	mortar := new(Mortar)

	var interactiveElements []Interactive

	alchemyWindow := loadSprite("mortar-interface")
	addIngredientButtonSprite := loadSprite("btn.add-ingredient")
	createPotionButtonSprite := loadSprite("btn.create-potion")
	exitButtonSprite := loadSprite("btn.exit")

	ingredientSelectors := []*Button{
		placeButton(addIngredientButtonSprite, window, 187, 180),
		placeButton(addIngredientButtonSprite, window, 187, 250),
		placeButton(addIngredientButtonSprite, window, 187, 320),
		placeButton(addIngredientButtonSprite, window, 187, 390),
	}

	for _, ingredientButton := range ingredientSelectors {
		ingredientButton.onclickfn = func () {}
	}

	createPotionButton := placeButton(createPotionButtonSprite, window, 253, 116)
	createPotionButton.onclickfn = func() { createPotion(mortar) }

	exitButton := placeButton(exitButtonSprite, window, 646, 115)
	exitButton.onclickfn = func() { os.Exit(0) }

	for _, button := range ingredientSelectors {
		interactiveElements = append(interactiveElements, button)
	}
	interactiveElements = append(interactiveElements, createPotionButton)
	interactiveElements = append(interactiveElements, exitButton)

	needRedraw := true
	for !window.Closed() {
		for _, interactiveElement := range interactiveElements {
			if window.JustPressed(pixelgl.MouseButtonLeft) && window.MouseInsideWindow() {
				if interactiveElement.IsUnderPosition(window.MousePosition()) {
					interactiveElement.Click()
				}
			}

			if !needRedraw && interactiveElement.NeedsRedraw() {
				needRedraw = true
			}
		}

		if needRedraw {
			window.Clear(colornames.White)
			placeSprite(alchemyWindow, window, 0, 0)

			for _, interactiveElement := range interactiveElements {
				interactiveElement.Draw()
			}

			needRedraw = false
		}

		window.Update()
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
