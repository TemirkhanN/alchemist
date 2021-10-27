package game

import (
	"log"

	"github.com/gookit/event"
	"golang.org/x/image/colornames"

	"github.com/TemirkhanN/alchemist/pkg/alchemy/alchemist"
	"github.com/TemirkhanN/alchemist/pkg/alchemy/ingredient"
	"github.com/TemirkhanN/alchemist/pkg/gui/geometry"
	"github.com/TemirkhanN/alchemist/pkg/gui/graphics"
)

type primaryLayout struct {
	initialized bool
	activeSlot  alchemist.Slot
	graphics    *graphics.Layout

	background         *graphics.SpriteCanvas
	effectsPreview     *graphics.Layout
	statusText         *graphics.TextCanvas
	ingredientSlots    map[alchemist.Slot]graphics.Canvas
	ingredients        map[alchemist.Slot]*ingredient.Ingredient
	createPotionButton *graphics.Button
	exitButton         *graphics.Button
}

func newPrimaryLayout(window *graphics.Window, player *alchemist.Alchemist) *primaryLayout {
	layout := new(primaryLayout)
	layout.initialized = true
	layout.graphics = graphics.NewLayer(window.Width(), window.Height(), true)

	backgroundSprite := gameAssets.GetSprite("interface.alchemy")
	addIngredientBtnSprite := gameAssets.GetSprite("btn.add-ingredient")
	createPotionBtnSprite := gameAssets.GetSprite("btn.create-potion")
	exitBtnSprite := gameAssets.GetSprite("btn.exit")

	button1 := graphics.NewButton(addIngredientBtnSprite)
	button1.SetClickHandler(func() {
		layout.activeSlot = alchemist.FirstSlot
		newAddIngredientButtonClickedEvent(layout.activeSlot)
	})

	button2 := graphics.NewButton(addIngredientBtnSprite)
	button2.SetClickHandler(func() {
		layout.activeSlot = alchemist.SecondSlot
		newAddIngredientButtonClickedEvent(layout.activeSlot)
	})

	button3 := graphics.NewButton(addIngredientBtnSprite)
	button3.SetClickHandler(func() {
		layout.activeSlot = alchemist.ThirdSlot
		newAddIngredientButtonClickedEvent(layout.activeSlot)
	})

	button4 := graphics.NewButton(addIngredientBtnSprite)
	button4.SetClickHandler(func() {
		layout.activeSlot = alchemist.FourthSlot
		newAddIngredientButtonClickedEvent(layout.activeSlot)
	})

	layout.background = graphics.NewSpriteCanvas(backgroundSprite)

	defaultSlots := func() map[alchemist.Slot]graphics.Canvas {
		return map[alchemist.Slot]graphics.Canvas{
			alchemist.FirstSlot:  button1,
			alchemist.SecondSlot: button2,
			alchemist.ThirdSlot:  button3,
			alchemist.FourthSlot: button4,
		}
	}
	layout.ingredientSlots = defaultSlots()
	layout.ingredients = make(map[alchemist.Slot]*ingredient.Ingredient, 4)

	layout.createPotionButton = graphics.NewButton(createPotionBtnSprite)
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
		layout.ingredients = make(map[alchemist.Slot]*ingredient.Ingredient, 4)
		layout.effectsPreview.Clear()
		layout.render()
	})

	layout.exitButton = graphics.NewButton(exitBtnSprite)
	layout.exitButton.SetClickHandler(func() { window.Close() })

	layout.effectsPreview = graphics.NewLayer(300, 270, true)
	layout.statusText = graphics.NewTextCanvas("", tesOblivion24Font, 200, colornames.Sienna)

	layout.registerEventHandlers(player)

	layout.render()

	return layout
}

func (layout *primaryLayout) registerEventHandlers(player *alchemist.Alchemist) {
	event.On(eventIngredientSelected, event.ListenerFunc(func(e event.Event) error {
		actualEvent := e.(*ingredientSelected)
		layout.ingredients[layout.activeSlot] = actualEvent.ingredient

		selectedIngredientButton := graphics.NewButton(getIngredientSprite(*actualEvent.ingredient))
		selectedSlot := layout.activeSlot
		selectedIngredientButton.SetClickHandler(func() {
			layout.activeSlot = selectedSlot

			player.DiscardIngredients()

			for slot, usedIngredient := range layout.ingredients {
				if slot == selectedSlot {
					continue
				}
				err := player.UseIngredient(usedIngredient)
				if err != nil {
					// This shall be unreachable statement
					layout.statusText.ChangeText(err.Error())
				}
			}
			newAddIngredientButtonClickedEvent(layout.activeSlot)
		})

		layout.ingredientSlots[layout.activeSlot] = selectedIngredientButton

		player.DiscardIngredients()
		for _, usedIngredient := range layout.ingredients {
			if player.CanUseIngredient(usedIngredient) {
				err := player.UseIngredient(usedIngredient)
				if err != nil {
					layout.statusText.ChangeText(err.Error())
				}
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
				effectPreviewLayout := graphics.NewLayer(260, 50, true)

				effectCanvas := graphics.NewSpriteCanvas(gameAssets.GetSprite(effect.Name()).Frame(potionEffectFrameSize))
				effectPreviewLayout.AddElement(
					effectCanvas,
					geometry.NewPosition(0, (effectPreviewLayout.Height()-effectCanvas.Height())/2),
				)

				descriptionCanvas := graphics.NewTextCanvas(
					effect.Description(),
					tesOblivion24Font,
					225,
					colornames.Sienna,
				)
				effectPreviewLayout.AddElement(
					descriptionCanvas,
					geometry.NewPosition(
						potionEffectFrameSize.Width()+5,
						(effectPreviewLayout.Height()-descriptionCanvas.Height())/2,
					),
				)

				layout.effectsPreview.AddElement(
					effectPreviewLayout,
					geometry.NewPosition(
						10+geometry.ZeroPosition.X(),
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

func (layout *primaryLayout) render() {
	// if it is not initialized, then it is an empty layout. nothing to show
	if !layout.initialized {
		return
	}

	layout.graphics.Clear()
	layout.graphics.AddElement(layout.background, geometry.ZeroPosition)
	layout.graphics.AddElement(layout.effectsPreview, geometry.NewPosition(550, 180))
	layout.graphics.AddElement(layout.statusText, geometry.NewPosition(180, 600))

	layout.graphics.AddElement(layout.ingredientSlots[alchemist.FirstSlot], geometry.NewPosition(187, 390))
	layout.graphics.AddElement(layout.ingredientSlots[alchemist.SecondSlot], geometry.NewPosition(187, 320))
	layout.graphics.AddElement(layout.ingredientSlots[alchemist.ThirdSlot], geometry.NewPosition(187, 250))
	layout.graphics.AddElement(layout.ingredientSlots[alchemist.FourthSlot], geometry.NewPosition(187, 180))

	layout.graphics.AddElement(layout.createPotionButton, geometry.NewPosition(253, 116))
	layout.graphics.AddElement(layout.exitButton, geometry.NewPosition(646, 115))
	layout.graphics.Show()
}

func newAddIngredientButtonClickedEvent(inSlot alchemist.Slot) {
	err := event.TriggerEvent(&addIngredientButtonClicked{slot: inSlot, BasicEvent: event.BasicEvent{}})
	if err != nil {
		log.Fatal(err.Error())
	}
}
