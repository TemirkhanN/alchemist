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
var ingredientsDatabase = initStorage()

const (
	INGREDIENT_SELECTION_BUTTON_PRESSED = 1
)

func main() {
	pixelgl.Run(launch)
}

func launch() {
	assets := new(GUI.Assets)
	// todo shall filesystem be passed by reference or not?
	err := assets.RegisterAssets("assets/sprites", &spritesFs)

	if err != nil {
		log.Fatal(err)
	}

	window := GUI.CreateWindow(1024, 768)
	runGame(window, assets)
}

func runGame(window *GUI.Window, assets *GUI.Assets) {
	eventDispatcher := make(chan int)

	mortar := new(Mortar)
	mortar.alchemyLevel = MortarLevel(APPRENTICE)

	mainLayout := new(MainLayout)
	mainLayout.init(window, assets, mortar, eventDispatcher)

	backpackLayout := new(BackpackLayout)
	backpackLayout.init(window, assets, mortar, eventDispatcher)

	for !window.Closed() {
		window.Refresh()
	}
}

// todo
func addIngredient(ingredient Ingredient, mortar *Mortar, layer *GUI.Layer, dispatcher chan int) func() {
	return func() {
		err := mortar.AddIngredient(ingredient)
		// TODO remove after creating structured GUI interface
		if err != nil {
			log.Fatal(err)
		}
		layer.Hide()
	}
}

type MainLayout struct {
	initialized bool

	background *GUI.SpriteCanvas
	textBlock  *GUI.TextCanvas
	firstSlot  *GUI.Button
	secondSlot *GUI.Button
	thirdSlot  *GUI.Button
	fourthSlot *GUI.Button
	createPotionButton *GUI.Button
	exitButton *GUI.Button
}

type BackpackLayout struct {
	initialized bool

	background *GUI.SpriteCanvas
	ingredients []*Ingredient
	ingredientsBtns []*GUI.Button
}

func (layout *MainLayout) init(window *GUI.Window, assets *GUI.Assets, mortar *Mortar, dispatcher chan int) {
	if layout.initialized {
		log.Fatal("can not initialize layout more than one time")
	}

	backgroundSprite := assets.GetSprite("mortar-interface")
	addIngredientBtnSprite := assets.GetSprite("btn.add-ingredient")
	createPotionBtnSprite := assets.GetSprite("btn.create-potion")
	exitBtnSprite := assets.GetSprite("btn.exit")

	layout.background = window.CreateSpriteCanvas(backgroundSprite, GUI.Position{})

	layout.firstSlot = window.CreateButton(addIngredientBtnSprite, GUI.Position{X: 187, Y: 390})
	layout.firstSlot.SetClickHandler(func () {dispatcher <- INGREDIENT_SELECTION_BUTTON_PRESSED})
	layout.secondSlot = window.CreateButton(addIngredientBtnSprite, GUI.Position{X: 187, Y: 320})
	layout.secondSlot.SetClickHandler(func () {dispatcher <- INGREDIENT_SELECTION_BUTTON_PRESSED})
	layout.thirdSlot = window.CreateButton(addIngredientBtnSprite, GUI.Position{X: 187, Y: 250})
	layout.thirdSlot.SetClickHandler(func () {dispatcher <- INGREDIENT_SELECTION_BUTTON_PRESSED})
	layout.fourthSlot = window.CreateButton(addIngredientBtnSprite, GUI.Position{X: 187, Y: 180})
	layout.fourthSlot.SetClickHandler(func () {dispatcher <- INGREDIENT_SELECTION_BUTTON_PRESSED})

	layout.createPotionButton = window.CreateButton(createPotionBtnSprite, GUI.Position{X: 253, Y: 116})
	layout.createPotionButton.SetClickHandler(func() {
		potion, err := mortar.Pestle()
		if err != nil {
			log.Fatal(err)
		}
		layout.textBlock.ChangeText(potion.Description())
	})

	layout.exitButton = window.CreateButton(exitBtnSprite, GUI.Position{X: 646, Y: 115})
	layout.exitButton.SetClickHandler(func() { os.Exit(0) })

	layout.textBlock = window.CreateTextCanvas("Description here", GUI.Position{X: 555, Y: 430})

	firstLayer := new(GUI.Layer)

	firstLayer.AddCanvas(layout.background)
	firstLayer.AddCanvas(layout.textBlock)
	firstLayer.AddCanvas(layout.firstSlot)
	firstLayer.AddCanvas(layout.secondSlot)
	firstLayer.AddCanvas(layout.thirdSlot)
	firstLayer.AddCanvas(layout.fourthSlot)
	firstLayer.AddCanvas(layout.createPotionButton)
	firstLayer.AddCanvas(layout.exitButton)
	firstLayer.Show()

	window.AddLayer(firstLayer)
}

// todo rename repo to backpack
func (layout *BackpackLayout) init(window *GUI.Window, assets *GUI.Assets, mortar *Mortar, dispatcher chan int) {
	if layout.initialized {
		log.Fatal("can not initialize layout more than one time")
	}

	closeButtonSprite := assets.GetSprite("btn.exit")
	ingredientsLayoutSprite := assets.GetSprite("ingredients-interface")

	firstLayer := new(GUI.Layer)
	layout.background = window.CreateSpriteCanvas(ingredientsLayoutSprite, GUI.Position{})

	closeBackpackBtn := window.CreateButton(closeButtonSprite, GUI.Position{X: 410, Y: 65})
	closeBackpackBtn.SetClickHandler(func() { firstLayer.Hide() })

	lastIngredientPosition := GUI.Position{X: 50, Y: 500}
	for _, ingredient := range ingredientsDatabase.All() {
		ingredientBtn := window.CreateButton(assets.GetSprite(ingredient.sprite), lastIngredientPosition)
		ingredientBtn.SetClickHandler(addIngredient(ingredient, mortar, firstLayer, dispatcher))
		lastIngredientPosition.Y -= 64
		layout.ingredientsBtns = append(layout.ingredientsBtns, ingredientBtn)
		// Backpack shall contain
		layout.ingredients = append(layout.ingredients, &ingredient)
	}

	firstLayer.AddCanvas(layout.background)
	for _, ingredientBtn := range layout.ingredientsBtns {
		firstLayer.AddCanvas(ingredientBtn)
	}

	firstLayer.AddCanvas(closeBackpackBtn)
	firstLayer.Hide()

	go func () {
		for {
			signal := <-dispatcher
			if signal == INGREDIENT_SELECTION_BUTTON_PRESSED {
				firstLayer.Show()
			}
		}
	}()

	window.AddLayer(firstLayer)
}
