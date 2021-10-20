package game

import (
	"log"
	"os"
	"strings"

	"github.com/gookit/event"

	"github.com/TemirkhanN/alchemist/assets"
	"github.com/TemirkhanN/alchemist/pkg/alchemy/alchemist"
	"github.com/TemirkhanN/alchemist/pkg/alchemy/ingredient"
	"github.com/TemirkhanN/alchemist/pkg/gui"
)

type PrimaryLayout struct {
	initialized bool
	activeSlot  alchemist.Slot
	graphics    *gui.Layer

	background         *gui.SpriteCanvas
	effectsPreview     *gui.Layer
	statusText         *gui.TextCanvas
	ingredientSlots    map[alchemist.Slot]gui.Canvas
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

	ingredients []*ingredient.Ingredient
	alchemist   *alchemist.Alchemist
}

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

	mainLayout := NewMainLayout(window, player)
	backpackLayout := NewBackpackLayout(window, player)

	window.AddLayer(mainLayout.graphics, gui.ZeroPosition)
	window.AddLayer(backpackLayout.graphics, gui.ZeroPosition)

	for !window.Closed() {
		window.Refresh()
	}
}

func NewMainLayout(window *gui.Window, player *alchemist.Alchemist) *PrimaryLayout {
	layout := new(PrimaryLayout)
	if layout.initialized {
		log.Fatal("can not initialize layout more than one time")
	}

	layout.initialized = true
	layout.graphics = window.CreateLayer(window.Width(), window.Height(), true)

	backgroundSprite := assets.TESAssets.GetSprite("interface.alchemy")
	addIngredientBtnSprite := assets.TESAssets.GetSprite("btn.add-ingredient")
	createPotionBtnSprite := assets.TESAssets.GetSprite("btn.create-potion")
	exitBtnSprite := assets.TESAssets.GetSprite("btn.exit")

	button1 := window.CreateButton(addIngredientBtnSprite)
	button1.SetClickHandler(func() {
		layout.activeSlot = alchemist.FirstSlot
		newAddIngredientButtonClickedEvent(layout.activeSlot)
	})

	button2 := window.CreateButton(addIngredientBtnSprite)
	button2.SetClickHandler(func() {
		layout.activeSlot = alchemist.SecondSlot
		newAddIngredientButtonClickedEvent(layout.activeSlot)
	})

	button3 := window.CreateButton(addIngredientBtnSprite)
	button3.SetClickHandler(func() {
		layout.activeSlot = alchemist.ThirdSlot
		newAddIngredientButtonClickedEvent(layout.activeSlot)
	})

	button4 := window.CreateButton(addIngredientBtnSprite)
	button4.SetClickHandler(func() {
		layout.activeSlot = alchemist.FourthSlot
		newAddIngredientButtonClickedEvent(layout.activeSlot)
	})

	layout.background = window.CreateSpriteCanvas(backgroundSprite)

	defaultSlots := func() map[alchemist.Slot]gui.Canvas {
		return map[alchemist.Slot]gui.Canvas{
			alchemist.FirstSlot:  button1,
			alchemist.SecondSlot: button2,
			alchemist.ThirdSlot:  button3,
			alchemist.FourthSlot: button4,
		}
	}
	layout.ingredientSlots = defaultSlots()

	layout.createPotionButton = window.CreateButton(createPotionBtnSprite)
	layout.createPotionButton.SetClickHandler(func() {
		if !player.CanStartBrewing() {
			return
		}

		_, err := player.BrewPotion("Some hardcoded potion name")
		if err != nil {
			log.Fatal(err)
		}
		layout.statusText.ChangeText("You have created a potion!")
		layout.ingredientSlots = defaultSlots()
		layout.effectsPreview.Clear()
		layout.render()
	})

	layout.exitButton = window.CreateButton(exitBtnSprite)
	layout.exitButton.SetClickHandler(func() { os.Exit(0) })

	layout.effectsPreview = window.CreateLayer(300, 270, true)
	layout.statusText = window.CreateTextCanvas("", tesOblivion24Font, 200)

	layout.registerEventHandlers(player, window)

	layout.render()

	return layout
}

func (layout *PrimaryLayout) registerEventHandlers(player *alchemist.Alchemist, window *gui.Window) {
	event.On(EventIngredientSelected, event.ListenerFunc(func(e event.Event) error {
		actualEvent := e.(*IngredientSelected)

		ingredientIcon := GetIngredientSprite(*actualEvent.ingredient)
		layout.ingredientSlots[layout.activeSlot] = window.CreateSpriteCanvas(ingredientIcon)

		if player.CanUseIngredient(actualEvent.ingredient) {
			err := player.UseIngredient(actualEvent.ingredient)
			if err != nil {
				layout.statusText.ChangeText(err.Error())
			}
		}

		if player.CanStartBrewing() {
			potion, err := player.PredictPotion()
			if err != nil {
				log.Fatal(err)
			}
			layout.effectsPreview.Clear()

			maximumAvailableAmountOfEffects := 4
			for order, effect := range potion.Effects() {
				effectPreviewLayout := window.CreateLayer(260, 50, true)

				effectCanvas := window.CreateSpriteCanvas(assets.TESAssets.GetSprite(effect.Name()).Frame(potionEffectFrameSize))
				effectPreviewLayout.AddElement(
					effectCanvas,
					gui.NewPosition(0, (effectPreviewLayout.Height()-effectCanvas.Height())/2),
				)

				descriptionCanvas := window.CreateTextCanvas(effect.Description(), tesOblivion24Font, 225)
				effectPreviewLayout.AddElement(
					descriptionCanvas,
					gui.NewPosition(
						potionEffectFrameSize.Width()+5,
						(effectPreviewLayout.Height()-descriptionCanvas.Height())/2,
					),
				)

				layout.effectsPreview.AddElement(
					effectPreviewLayout,
					gui.NewPosition(
						10+gui.ZeroPosition.X(),
						effectPreviewLayout.Height()*float64(maximumAvailableAmountOfEffects-order),
					),
				)
			}
		}

		layout.activeSlot = alchemist.EmptySlot

		layout.render()

		return nil
	}))
}

func (layout *PrimaryLayout) render() {
	// if it is not initialized, then it is an empty layout. nothing to show
	if !layout.initialized {
		return
	}

	layout.graphics.Clear()
	layout.graphics.AddElement(layout.background, gui.ZeroPosition)
	layout.graphics.AddElement(layout.effectsPreview, gui.NewPosition(550, 180))
	layout.graphics.AddElement(layout.statusText, gui.NewPosition(180, 600))

	layout.graphics.AddElement(layout.ingredientSlots[alchemist.FirstSlot], gui.NewPosition(187, 390))
	layout.graphics.AddElement(layout.ingredientSlots[alchemist.SecondSlot], gui.NewPosition(187, 320))
	layout.graphics.AddElement(layout.ingredientSlots[alchemist.ThirdSlot], gui.NewPosition(187, 250))
	layout.graphics.AddElement(layout.ingredientSlots[alchemist.FourthSlot], gui.NewPosition(187, 180))

	layout.graphics.AddElement(layout.createPotionButton, gui.NewPosition(253, 116))
	layout.graphics.AddElement(layout.exitButton, gui.NewPosition(646, 115))
	layout.graphics.Show()
}

// NewBackpackLayout todo rename repo to backpack.
func NewBackpackLayout(window *gui.Window, player *alchemist.Alchemist) *BackpackLayout {
	layout := new(BackpackLayout)
	if layout.initialized {
		log.Fatal("can not initialize layout more than one time")
	}

	layout.initialized = true
	layout.window = window
	layout.alchemist = player

	for _, ingr := range ingredient.IngredientsDatabase.All() {
		deref := ingr
		layout.ingredients = append(layout.ingredients, &deref)
	}

	closeButtonSprite := assets.TESAssets.GetSprite("btn.exit")
	ingredientsLayoutSprite := assets.TESAssets.GetSprite("interface.ingredients")

	layout.graphics = window.CreateLayer(window.Width(), window.Height(), false)
	layout.background = window.CreateSpriteCanvas(ingredientsLayoutSprite)

	layout.closeButton = window.CreateButton(closeButtonSprite)
	layout.closeButton.SetClickHandler(func() { layout.graphics.Hide() })

	event.On(EventAddIngredientButtonClicked, event.ListenerFunc(func(e event.Event) error {
		layout.render()

		return nil
	}))

	return layout
}

func (layout *BackpackLayout) render() {
	layout.graphics.Clear()
	layout.graphics.AddElement(layout.background, gui.ZeroPosition)

	ingredientsLayer := layout.window.CreateLayer(480, 465, true, true)
	ingredientEffectsLayer := layout.window.CreateLayer(238, 220, false)
	ingredientsEffectsLayerBackground := layout.window.CreateSpriteCanvas(assets.TESAssets.GetSprite("interface.effects"))

	layout.ingredientsBtns = nil
	offset := ingredientsLayer.Height()

	for _, ingr := range layout.ingredients {
		if !layout.alchemist.CanUseIngredient(ingr) {
			continue
		}

		ingredientBtn := layout.window.CreateButton(GetIngredientSprite(*ingr))
		ingredientBtn.SetClickHandler(func(selected *ingredient.Ingredient) func() {
			return func() {
				// todo potentially vulnerable for mistake on main(mortar) side
				layout.graphics.Hide()
				err := event.FireEvent(&IngredientSelected{ingredient: selected, BasicEvent: event.BasicEvent{}})
				if err != nil {
					log.Fatal(err)
				}
			}
		}(ingr))

		var lastHoveredIngredient string
		ingredientBtn.SetMouseOverHandler(func(hovered *ingredient.Ingredient) func() {
			return func() {
				lastHoveredIngredient = hovered.Name()
				ingredientEffectsLayer.Clear()
				ingredientEffectsLayer.AddElement(ingredientsEffectsLayerBackground, gui.ZeroPosition)
				posY := ingredientEffectsLayer.Height()
				for _, effect := range layout.alchemist.DetermineEffects(hovered) {
					posY -= 55
					effectPreview := layout.window.CreateSpriteCanvas(assets.TESAssets.GetSprite(effect.Name()))
					ingredientEffectsLayer.AddElement(effectPreview, gui.NewPosition(0, posY))
				}
				ingredientEffectsLayer.Show()
			}
		}(ingr))
		ingredientBtn.SetMouseOutHandler(func(hovered *ingredient.Ingredient) func() {
			return func() {
				// If mouse is out, and it is not different ingredient hovered, then hide effect layer
				if lastHoveredIngredient == hovered.Name() {
					ingredientEffectsLayer.Clear()
					ingredientEffectsLayer.Hide()
				}
			}
		}(ingr))

		layout.ingredientsBtns = append(layout.ingredientsBtns, ingredientBtn)

		offset -= 64

		ingredientsLayer.AddElement(ingredientBtn, gui.NewPosition(0, offset))
	}

	layout.graphics.AddElement(ingredientsLayer, gui.NewPosition(50, 115))
	layout.graphics.AddElement(ingredientEffectsLayer, gui.NewPosition(605, 200))
	layout.graphics.AddElement(layout.closeButton, gui.NewPosition(410, 65))

	layout.graphics.Show()
}

func GetIngredientSprite(ingr ingredient.Ingredient) *gui.Sprite {
	spriteName := "ingr." + strings.ReplaceAll(strings.ToLower(ingr.Name()), "'", "")

	return assets.TESAssets.GetSprite(spriteName)
}

var potionEffectFrameSize = gui.FrameSize{
	LeftBottom: gui.NewPosition(10, 15),
	RightTop:   gui.NewPosition(35, 40),
}

var tesOblivion24Font = gui.LoadFont("TESOblivionFont", "assets/font/Kingthings Petrock.ttf", 24)
