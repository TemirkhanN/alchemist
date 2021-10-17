package main

import (
	"embed"
	_ "image/png"
	"log"
	"os"
	"strings"

	"github.com/faiface/pixel/pixelgl"
	"github.com/gookit/event"

	"github.com/TemirkhanN/alchemist/domain"
	"github.com/TemirkhanN/alchemist/gui"
)

type PrimaryLayout struct {
	initialized bool
	activeSlot  domain.Slot
	graphics    *gui.Layer

	background         *gui.SpriteCanvas
	textBlock          *gui.TextCanvas
	ingredientSlots    map[domain.Slot]gui.Canvas
	createPotionButton *gui.Button
	exitButton         *gui.Button
}

type BackpackLayout struct {
	initialized     bool
	graphics        *gui.Layer
	window          *gui.Window
	background      *gui.SpriteCanvas
	ingredientsBtns []*gui.Button
	closeButton     *gui.Button

	ingredients []*domain.Ingredient
	alchemist   *domain.Alchemist
}

func main() {
	pixelgl.Run(func() {
		launch(1024, 768)
	})
}

func launch(windowWidth float64, windowHeight float64) {
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
	alchemist := domain.NewAlchemist(alchemyLevelHardcoded, luckLevelHardcoded, domain.NewNoviceMortar())

	mainLayout := NewMainLayout(window, alchemist)
	backpackLayout := NewBackpackLayout(window, alchemist)

	window.AddLayer(mainLayout.graphics, gui.ZeroPosition)
	window.AddLayer(backpackLayout.graphics, gui.ZeroPosition)

	for !window.Closed() {
		window.Refresh()
	}
}

func NewMainLayout(window *gui.Window, alchemist *domain.Alchemist) *PrimaryLayout {
	layout := new(PrimaryLayout)
	if layout.initialized {
		log.Fatal("can not initialize layout more than one time")
	}

	layout.initialized = true
	layout.graphics = gui.NewLayer(window.Width(), window.Height(), true)

	backgroundSprite := assets.GetSprite("interface.alchemy")
	addIngredientBtnSprite := assets.GetSprite("btn.add-ingredient")
	createPotionBtnSprite := assets.GetSprite("btn.create-potion")
	exitBtnSprite := assets.GetSprite("btn.exit")

	button1 := window.CreateButton(addIngredientBtnSprite)
	button1.SetClickHandler(func() {
		layout.activeSlot = domain.FirstSlot
		newAddIngredientButtonClickedEvent(layout.activeSlot)
	})

	button2 := window.CreateButton(addIngredientBtnSprite)
	button2.SetClickHandler(func() {
		layout.activeSlot = domain.SecondSlot
		newAddIngredientButtonClickedEvent(layout.activeSlot)
	})

	button3 := window.CreateButton(addIngredientBtnSprite)
	button3.SetClickHandler(func() {
		layout.activeSlot = domain.ThirdSlot
		newAddIngredientButtonClickedEvent(layout.activeSlot)
	})

	button4 := window.CreateButton(addIngredientBtnSprite)
	button4.SetClickHandler(func() {
		layout.activeSlot = domain.FourthSlot
		newAddIngredientButtonClickedEvent(layout.activeSlot)
	})

	layout.background = window.CreateSpriteCanvas(backgroundSprite)

	layout.ingredientSlots = map[domain.Slot]gui.Canvas{
		domain.FirstSlot:  button1,
		domain.SecondSlot: button2,
		domain.ThirdSlot:  button3,
		domain.FourthSlot: button4,
	}

	layout.createPotionButton = window.CreateButton(createPotionBtnSprite)
	layout.createPotionButton.SetClickHandler(func() {
		if !alchemist.CanStartBrewing() {
			return
		}

		potion, err := alchemist.BrewPotion("Some hardcoded potion name")
		if err != nil {
			log.Fatal(err)
		}
		layout.textBlock.ChangeText(potion.Description())
		layout.ingredientSlots = map[domain.Slot]gui.Canvas{
			domain.FirstSlot:  button1,
			domain.SecondSlot: button2,
			domain.ThirdSlot:  button3,
			domain.FourthSlot: button4,
		}
		layout.render()
	})

	layout.exitButton = window.CreateButton(exitBtnSprite)
	layout.exitButton.SetClickHandler(func() { os.Exit(0) })

	layout.textBlock = window.CreateTextCanvas("Description here")

	event.On(EventIngredientSelected, event.ListenerFunc(func(e event.Event) error {
		actualEvent := e.(*IngredientSelected)

		ingredientIcon := GetIngredientSprite(*actualEvent.ingredient)
		layout.ingredientSlots[layout.activeSlot] = window.CreateSpriteCanvas(ingredientIcon)

		if alchemist.CanUseIngredient(actualEvent.ingredient) {
			err := alchemist.UseIngredient(actualEvent.ingredient)
			if err != nil {
				layout.textBlock.ChangeText(err.Error())
			}
		}

		layout.activeSlot = domain.EmptySlot

		layout.render()

		return nil
	}))

	layout.render()

	return layout
}

func (layout *PrimaryLayout) render() {
	// if it is not initialized, then it is an empty layout. nothing to show
	if !layout.initialized {
		return
	}

	layout.graphics.Clear()
	layout.graphics.AddElement(layout.background, gui.ZeroPosition)
	layout.graphics.AddElement(layout.textBlock, gui.NewPosition(555, 430))

	layout.graphics.AddElement(layout.ingredientSlots[domain.FirstSlot], gui.NewPosition(187, 390))
	layout.graphics.AddElement(layout.ingredientSlots[domain.SecondSlot], gui.NewPosition(187, 320))
	layout.graphics.AddElement(layout.ingredientSlots[domain.ThirdSlot], gui.NewPosition(187, 250))
	layout.graphics.AddElement(layout.ingredientSlots[domain.FourthSlot], gui.NewPosition(187, 180))

	layout.graphics.AddElement(layout.createPotionButton, gui.NewPosition(253, 116))
	layout.graphics.AddElement(layout.exitButton, gui.NewPosition(646, 115))
	layout.graphics.Show()
}

// NewBackpackLayout todo rename repo to backpack.
func NewBackpackLayout(window *gui.Window, alchemist *domain.Alchemist) *BackpackLayout {
	layout := new(BackpackLayout)
	if layout.initialized {
		log.Fatal("can not initialize layout more than one time")
	}

	layout.initialized = true
	layout.window = window
	layout.alchemist = alchemist

	for _, ingredient := range domain.IngredientsDatabase.All() {
		deref := ingredient
		layout.ingredients = append(layout.ingredients, &deref)
	}

	closeButtonSprite := assets.GetSprite("btn.exit")
	ingredientsLayoutSprite := assets.GetSprite("interface.ingredients")

	layout.graphics = gui.NewLayer(window.Width(), window.Height(), false)
	layout.background = window.CreateSpriteCanvas(ingredientsLayoutSprite)

	layout.closeButton = window.CreateButton(closeButtonSprite)
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
	layout.graphics.AddElement(layout.background, gui.ZeroPosition)

	ingredientsLayer := gui.NewLayer(480, 465, true, true)
	ingredientEffectsLayer := gui.NewLayer(238, 220, false)
	ingredientsEffectsLayerBackground := layout.window.CreateSpriteCanvas(assets.GetSprite("interface.effects"))

	layout.ingredientsBtns = nil
	offset := ingredientsLayer.Height()

	for _, ingredient := range layout.ingredients {
		if !layout.alchemist.CanUseIngredient(ingredient) {
			continue
		}

		ingredientBtn := layout.window.CreateButton(GetIngredientSprite(*ingredient))
		ingredientBtn.SetClickHandler(func(selected *domain.Ingredient) func() {
			return func() {
				// todo potentially vulnerable for mistake on main(mortar) side
				layout.graphics.Hide()
				err := event.FireEvent(&IngredientSelected{ingredient: selected, BasicEvent: event.BasicEvent{}})
				if err != nil {
					log.Fatal(err)
				}
			}
		}(ingredient))

		var lastHoveredIngredient string
		ingredientBtn.SetMouseOverHandler(func(hovered *domain.Ingredient) func() {
			return func() {
				lastHoveredIngredient = hovered.Name()
				ingredientEffectsLayer.Clear()
				ingredientEffectsLayer.AddElement(ingredientsEffectsLayerBackground, gui.ZeroPosition)
				posY := ingredientEffectsLayer.Height()
				for _, effect := range layout.alchemist.DetermineEffects(hovered) {
					posY -= 55
					effectPreview := layout.window.CreateSpriteCanvas(assets.GetSprite(effect.Name()))
					ingredientEffectsLayer.AddElement(effectPreview, gui.NewPosition(0, posY))
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

		layout.ingredientsBtns = append(layout.ingredientsBtns, ingredientBtn)

		offset -= 64

		ingredientsLayer.AddElement(ingredientBtn, gui.NewPosition(0, offset))
	}

	layout.graphics.Show()

	layout.graphics.AddElement(ingredientsLayer, gui.NewPosition(50, 115))
	layout.graphics.AddElement(ingredientEffectsLayer, gui.NewPosition(605, 200))
	layout.graphics.AddElement(layout.closeButton, gui.NewPosition(410, 65))
}

func GetIngredientSprite(ingredient domain.Ingredient) *gui.Sprite {
	spriteName := "ingr." + strings.ReplaceAll(strings.ToLower(ingredient.Name()), "'", "")

	return assets.GetSprite(spriteName)
}

//go:embed assets/sprites
var spritesFs embed.FS

var assets = func() *gui.Assets {
	loadedAssets := new(gui.Assets)
	// todo shall filesystem be passed by reference or not?
	err := loadedAssets.RegisterAssets("assets", &spritesFs)
	if err != nil {
		log.Fatal(err)
	}

	return loadedAssets
}()
