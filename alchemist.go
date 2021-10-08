package main

import (
	"embed"
	"github.com/TemirkhanN/alchemist/GUI"
	"github.com/faiface/pixel/pixelgl"
	"github.com/gookit/event"
	_ "image/png"
	"log"
	"os"
)

//go:embed assets/sprites/*.png
var spritesFs embed.FS
var ingredientsDatabase = initStorage()

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
	mortar := new(Mortar)
	mortar.alchemyLevel = MortarLevel(APPRENTICE)

	mainLayout := new(MainLayout)
	mainLayout.init(window, assets, mortar)

	backpackLayout := new(BackpackLayout)
	backpackLayout.init(window, assets, mortar)

	for !window.Closed() {
		window.Refresh()
	}
}

type Slot uint8

const (
	None Slot = iota
	First
	Second
	Third
	Fourth
)

type MainLayout struct {
	initialized bool
	activeSlot  Slot
	graphics    *GUI.Layer

	background         *GUI.SpriteCanvas
	textBlock          *GUI.TextCanvas
	ingredientSlots    map[Slot]GUI.Canvas
	createPotionButton *GUI.Button
	exitButton         *GUI.Button
}

type BackpackLayout struct {
	initialized bool

	background      *GUI.SpriteCanvas
	ingredients     []*Ingredient
	ingredientsBtns []*GUI.Button
}

func (layout *MainLayout) init(window *GUI.Window, assets *GUI.Assets, mortar *Mortar) {
	if layout.initialized {
		log.Fatal("can not initialize layout more than one time")
	}
	layout.initialized = true
	layout.graphics = new(GUI.Layer)

	backgroundSprite := assets.GetSprite("mortar-interface")
	addIngredientBtnSprite := assets.GetSprite("btn.add-ingredient")
	createPotionBtnSprite := assets.GetSprite("btn.create-potion")
	exitBtnSprite := assets.GetSprite("btn.exit")

	button1 := window.CreateButton(addIngredientBtnSprite, GUI.Position{X: 187, Y: 390})
	button1.SetClickHandler(func() {
		layout.activeSlot = Slot(First)
		event.TriggerEvent(&AddIngredientButtonClicked{slot: layout.activeSlot})
	})
	button2 := window.CreateButton(addIngredientBtnSprite, GUI.Position{X: 187, Y: 320})
	button2.SetClickHandler(func() {
		layout.activeSlot = Slot(Second)
		event.TriggerEvent(&AddIngredientButtonClicked{slot: layout.activeSlot})
	})
	button3 := window.CreateButton(addIngredientBtnSprite, GUI.Position{X: 187, Y: 250})
	button3.SetClickHandler(func() {
		layout.activeSlot = Slot(Third)
		event.TriggerEvent(&AddIngredientButtonClicked{slot: layout.activeSlot})
	})
	button4 := window.CreateButton(addIngredientBtnSprite, GUI.Position{X: 187, Y: 180})
	button4.SetClickHandler(func() {
		layout.activeSlot = Slot(Fourth)
		event.TriggerEvent(&AddIngredientButtonClicked{slot: layout.activeSlot})
	})

	layout.background = window.CreateSpriteCanvas(backgroundSprite, GUI.Position{})

	layout.ingredientSlots = map[Slot]GUI.Canvas{
		Slot(First):  button1,
		Slot(Second): button2,
		Slot(Third):  button3,
		Slot(Fourth): button4,
	}

	layout.createPotionButton = window.CreateButton(createPotionBtnSprite, GUI.Position{X: 253, Y: 116})
	layout.createPotionButton.SetClickHandler(func() {
		potion, err := mortar.Pestle()
		if err != nil {
			log.Fatal(err)
		}
		layout.textBlock.ChangeText(potion.Description())
		layout.ingredientSlots = map[Slot]GUI.Canvas{
			Slot(First):  button1,
			Slot(Second): button2,
			Slot(Third):  button3,
			Slot(Fourth): button4,
		}
		layout.render()
	})

	layout.exitButton = window.CreateButton(exitBtnSprite, GUI.Position{X: 646, Y: 115})
	layout.exitButton.SetClickHandler(func() { os.Exit(0) })

	layout.textBlock = window.CreateTextCanvas("Description here", GUI.Position{X: 555, Y: 430})

	event.On(EventIngredientSelected, event.ListenerFunc(func(e event.Event) error {
		actualEvent := e.(*IngredientSelected)

		ingredientIcon := assets.GetSprite(actualEvent.ingredient.sprite)
		slotPosition := layout.ingredientSlots[layout.activeSlot].Position()
		layout.ingredientSlots[layout.activeSlot] = window.CreateSpriteCanvas(ingredientIcon, slotPosition)

		mortar.AddIngredient(*actualEvent.ingredient)

		layout.activeSlot = Slot(None)

		layout.render()
		return nil
	}))

	layout.render()

	window.AddLayer(layout.graphics)
}

func (layout *MainLayout) render() {
	// if it is not initialized, then it is an empty layout. nothing to show
	if !layout.initialized {
		return
	}

	layout.graphics.Clear()
	layout.graphics.AddCanvas(layout.background)
	layout.graphics.AddCanvas(layout.textBlock)
	for _, slotCanvas := range layout.ingredientSlots {
		layout.graphics.AddCanvas(slotCanvas)
	}
	layout.graphics.AddCanvas(layout.createPotionButton)
	layout.graphics.AddCanvas(layout.exitButton)
	layout.graphics.Show()
}

// todo rename repo to backpack
func (layout *BackpackLayout) init(window *GUI.Window, assets *GUI.Assets, mortar *Mortar) {
	if layout.initialized {
		log.Fatal("can not initialize layout more than one time")
	}
	layout.initialized = true

	for _, ingredient := range ingredientsDatabase.All() {
		deref := ingredient
		layout.ingredients = append(layout.ingredients, &deref)
	}

	closeButtonSprite := assets.GetSprite("btn.exit")
	ingredientsLayoutSprite := assets.GetSprite("ingredients-interface")

	firstLayer := new(GUI.Layer)
	layout.background = window.CreateSpriteCanvas(ingredientsLayoutSprite, GUI.Position{})

	closeBackpackBtn := window.CreateButton(closeButtonSprite, GUI.Position{X: 410, Y: 65})
	closeBackpackBtn.SetClickHandler(func() { firstLayer.Hide() })

	lastIngredientPosition := GUI.Position{X: 50, Y: 500}
	for _, ingredient := range layout.ingredients {
		ingredientBtn := window.CreateButton(assets.GetSprite(ingredient.sprite), lastIngredientPosition)
		ingredientBtn.SetClickHandler(func(selected *Ingredient) func() {
			return func() {
				firstLayer.Hide()
				event.FireEvent(&IngredientSelected{ingredient: selected})
			}
		}(ingredient))

		lastIngredientPosition.Y -= 64
		layout.ingredientsBtns = append(layout.ingredientsBtns, ingredientBtn)
	}

	firstLayer.AddCanvas(layout.background)
	for _, ingredientBtn := range layout.ingredientsBtns {
		firstLayer.AddCanvas(ingredientBtn)
	}

	firstLayer.AddCanvas(closeBackpackBtn)
	firstLayer.Hide()

	event.On(EventAddIngredientButtonClicked, event.ListenerFunc(func(e event.Event) error {
		firstLayer.Show()
		return nil
	}))

	window.AddLayer(firstLayer)
}
