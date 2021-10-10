package main

import (
	"embed"
	"github.com/TemirkhanN/alchemist/GUI"
	"github.com/TemirkhanN/alchemist/domain"
	"github.com/faiface/pixel/pixelgl"
	"github.com/gookit/event"
	"log"
	"os"
	"strings"
)

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
	initialized                      bool
	graphics                         *GUI.Layer
	window                           *GUI.Window
	background                       *GUI.SpriteCanvas
	ingredientsBtns                  []*GUI.Button
	closeButton                      *GUI.Button
	ingredientsVerticalOffset        float64
	ingredientsVerticalDefaultOffset float64

	ingredients []*domain.Ingredient
	alchemist   *domain.Alchemist
}

//go:embed assets/sprites
var spritesFs embed.FS
var assets = func() *GUI.Assets {
	loadedAssets := new(GUI.Assets)
	// todo shall filesystem be passed by reference or not?
	err := loadedAssets.RegisterAssets("assets", &spritesFs)
	if err != nil {
		log.Fatal(err)
	}

	return loadedAssets
}()

func main() {
	pixelgl.Run(func() {
		launch(1024, 768)
	})
}

func launch(windowWidth float64, windowHeight float64) {
	window := GUI.CreateWindow(windowWidth, windowHeight)

	alchemyLevelHardcoded := 26
	luckLevelHardcoded := 5
	alchemist := domain.NewAlchemist(alchemyLevelHardcoded, luckLevelHardcoded, domain.NewNoviceMortar())

	mainLayout := NewMainLayout(window, alchemist)
	backpackLayout := NewBackpackLayout(window, alchemist)

	window.AddLayer(mainLayout.graphics)
	window.AddLayer(backpackLayout.graphics)

	for !window.Closed() {
		window.Refresh()
	}
}

func NewMainLayout(window *GUI.Window, alchemist *domain.Alchemist) *MainLayout {
	layout := new(MainLayout)
	if layout.initialized {
		log.Fatal("can not initialize layout more than one time")
	}
	layout.initialized = true
	layout.graphics = GUI.NewLayer(0, 0, true)

	backgroundSprite := assets.GetSprite("interface.alchemy")
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
		if len(alchemist.UsedIngredients()) < 2 {
			layout.textBlock.ChangeText("You need at least 2 ingredients to make potion")
			return
		}

		_, err := alchemist.BrewPotion("Some hardcoded potion name")
		if err != nil {
			log.Fatal(err)
		}
		layout.textBlock.ChangeText("You have created a potion. todo description here")
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

		ingredientIcon := GetIngredientSprite(*actualEvent.ingredient)
		slotPosition := layout.ingredientSlots[layout.activeSlot].Position()
		layout.ingredientSlots[layout.activeSlot] = window.CreateSpriteCanvas(ingredientIcon, slotPosition)

		if alchemist.CanUseIngredient(actualEvent.ingredient) {
			alchemist.UseIngredient(actualEvent.ingredient)
		}

		layout.activeSlot = Slot(None)

		layout.render()
		return nil
	}))

	layout.render()

	return layout
}

func (layout *MainLayout) render() {
	// if it is not initialized, then it is an empty layout. nothing to show
	if !layout.initialized {
		return
	}

	layout.graphics.Clear()
	layout.graphics.AddElement(layout.background)
	layout.graphics.AddElement(layout.textBlock)
	for _, slotCanvas := range layout.ingredientSlots {
		layout.graphics.AddElement(slotCanvas)
	}
	layout.graphics.AddElement(layout.createPotionButton)
	layout.graphics.AddElement(layout.exitButton)
	layout.graphics.Show()
}

// NewBackpackLayout todo rename repo to backpack
func NewBackpackLayout(window *GUI.Window, alchemist *domain.Alchemist) *BackpackLayout {
	layout := new(BackpackLayout)
	if layout.initialized {
		log.Fatal("can not initialize layout more than one time")
	}
	layout.initialized = true
	layout.window = window
	layout.alchemist = alchemist
	// Allows scrolling ingredient list
	layout.ingredientsVerticalDefaultOffset = 500
	layout.ingredientsVerticalOffset = layout.ingredientsVerticalDefaultOffset

	for _, ingredient := range domain.IngredientsDatabase.All() {
		deref := ingredient
		layout.ingredients = append(layout.ingredients, &deref)
	}

	closeButtonSprite := assets.GetSprite("btn.exit")
	ingredientsLayoutSprite := assets.GetSprite("interface.ingredients")

	layout.graphics = GUI.NewLayer(0, 0, false)
	layout.background = window.CreateSpriteCanvas(ingredientsLayoutSprite, GUI.Position{})

	layout.closeButton = window.CreateButton(closeButtonSprite, GUI.Position{X: 410, Y: 65})
	layout.closeButton.SetClickHandler(func() { layout.graphics.Hide() })

	event.On(EventAddIngredientButtonClicked, event.ListenerFunc(func(e event.Event) error {
		layout.graphics.Show()
		layout.render()

		return nil
	}))

	return layout
}

func (layout *BackpackLayout) render() {
	layout.graphics.Clear()
	layout.graphics.AddElement(layout.background)

	ingredientsLayer := GUI.NewLayer(0, 448, true)
	ingredientEffectsLayer := GUI.NewLayer(0, 0, false)
	ingredientsLayerBackgroundPosition := GUI.Position{X: 605, Y: 200}
	ingredientsLayerBackground := layout.window.CreateSpriteCanvas(assets.GetSprite("interface.effects"), ingredientsLayerBackgroundPosition)

	layout.ingredientsBtns = nil
	offset := layout.ingredientsVerticalDefaultOffset
	for _, ingredient := range layout.ingredients {
		if !layout.alchemist.CanUseIngredient(ingredient) {
			continue
		}
		ingredientBtn := layout.window.CreateButton(
			GetIngredientSprite(*ingredient),
			GUI.Position{X: 50, Y: offset},
		)
		ingredientBtn.SetClickHandler(func(selected *domain.Ingredient) func() {
			return func() {
				// todo potentially vulnerable for mistake on main(mortar) side
				layout.graphics.Hide()
				event.FireEvent(&IngredientSelected{ingredient: selected})
			}
		}(ingredient))

		var lastHoveredIngredient string
		ingredientBtn.SetMouseOverHandler(func(hovered *domain.Ingredient) func() {
			return func() {
				lastHoveredIngredient = hovered.Name()
				ingredientEffectsLayer.Clear()
				ingredientEffectsLayer.AddElement(ingredientsLayerBackground)
				initialPosition := GUI.Position{X: 610, Y: 370}
				for _, effect := range layout.alchemist.DetermineEffects(hovered) {
					effectPreview := layout.window.CreateSpriteCanvas(assets.GetSprite(effect.Name()), initialPosition)
					ingredientEffectsLayer.AddElement(effectPreview)

					initialPosition.Y -= 55
				}
				ingredientEffectsLayer.Show()
			}
		}(ingredient))
		ingredientBtn.SetMouseOutHandler(func(hovered *domain.Ingredient) func() {
			return func() {
				// If mouse is out, and it is not different ingredient hovered, then hide effect layer
				if lastHoveredIngredient == hovered.Name() {
					ingredientEffectsLayer.Clear()
					ingredientEffectsLayer.Hide()
				}
			}
		}(ingredient))

		offset -= 64
		layout.ingredientsBtns = append(layout.ingredientsBtns, ingredientBtn)
		ingredientsLayer.AddElement(ingredientBtn)
	}
	ingredientsLayer.Show()

	layout.graphics.AddElement(ingredientsLayer)
	layout.graphics.AddElement(ingredientEffectsLayer)

	layout.graphics.AddElement(layout.closeButton)
}

func GetIngredientSprite(ingredient domain.Ingredient) *GUI.Sprite {
	spriteName := "ingr." + strings.ReplaceAll(strings.ToLower(ingredient.Name()), "'", "")

	return assets.GetSprite(spriteName)
}
