package game

import (
	"log"

	"github.com/gookit/event"

	"github.com/TemirkhanN/alchemist/pkg/alchemy/alchemist"
	"github.com/TemirkhanN/alchemist/pkg/alchemy/ingredient"
	"github.com/TemirkhanN/alchemist/pkg/gui"
)

type primaryLayout struct {
	initialized bool
	activeSlot  alchemist.Slot
	graphics    *gui.Layer

	background         *gui.SpriteCanvas
	effectsPreview     *gui.Layer
	statusText         *gui.TextCanvas
	ingredientSlots    map[alchemist.Slot]gui.Canvas
	ingredients        map[alchemist.Slot]*ingredient.Ingredient
	createPotionButton *gui.Button
	exitButton         *gui.Button
}

func newPrimaryLayout(window *gui.Window, player *alchemist.Alchemist) *primaryLayout {
	layout := new(primaryLayout)
	if layout.initialized {
		log.Fatal("can not initialize layout more than one time")
	}

	layout.initialized = true
	layout.graphics = gui.CreateLayer(window.Width(), window.Height(), true)

	backgroundSprite := gameAssets.GetSprite("interface.alchemy")
	addIngredientBtnSprite := gameAssets.GetSprite("btn.add-ingredient")
	createPotionBtnSprite := gameAssets.GetSprite("btn.create-potion")
	exitBtnSprite := gameAssets.GetSprite("btn.exit")

	button1 := gui.CreateButton(addIngredientBtnSprite)
	button1.SetClickHandler(func() {
		layout.activeSlot = alchemist.FirstSlot
		newAddIngredientButtonClickedEvent(layout.activeSlot)
	})

	button2 := gui.CreateButton(addIngredientBtnSprite)
	button2.SetClickHandler(func() {
		layout.activeSlot = alchemist.SecondSlot
		newAddIngredientButtonClickedEvent(layout.activeSlot)
	})

	button3 := gui.CreateButton(addIngredientBtnSprite)
	button3.SetClickHandler(func() {
		layout.activeSlot = alchemist.ThirdSlot
		newAddIngredientButtonClickedEvent(layout.activeSlot)
	})

	button4 := gui.CreateButton(addIngredientBtnSprite)
	button4.SetClickHandler(func() {
		layout.activeSlot = alchemist.FourthSlot
		newAddIngredientButtonClickedEvent(layout.activeSlot)
	})

	layout.background = gui.CreateSpriteCanvas(backgroundSprite)

	defaultSlots := func() map[alchemist.Slot]gui.Canvas {
		return map[alchemist.Slot]gui.Canvas{
			alchemist.FirstSlot:  button1,
			alchemist.SecondSlot: button2,
			alchemist.ThirdSlot:  button3,
			alchemist.FourthSlot: button4,
		}
	}
	layout.ingredientSlots = defaultSlots()
	layout.ingredients = make(map[alchemist.Slot]*ingredient.Ingredient, 4)

	layout.createPotionButton = gui.CreateButton(createPotionBtnSprite)
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

	layout.exitButton = gui.CreateButton(exitBtnSprite)
	layout.exitButton.SetClickHandler(func() { window.Close() })

	layout.effectsPreview = gui.CreateLayer(300, 270, true)
	layout.statusText = gui.CreateTextCanvas("", tesOblivion24Font, 200)

	layout.registerEventHandlers(player)

	layout.render()

	return layout
}

func (layout *primaryLayout) registerEventHandlers(player *alchemist.Alchemist) {
	event.On(eventIngredientSelected, event.ListenerFunc(func(e event.Event) error {
		actualEvent := e.(*ingredientSelected)
		layout.ingredients[layout.activeSlot] = actualEvent.ingredient

		selectedIngredientButton := gui.CreateButton(getIngredientSprite(*actualEvent.ingredient))
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
				effectPreviewLayout := gui.CreateLayer(260, 50, true)

				effectCanvas := gui.CreateSpriteCanvas(gameAssets.GetSprite(effect.Name()).Frame(potionEffectFrameSize))
				effectPreviewLayout.AddElement(
					effectCanvas,
					gui.NewPosition(0, (effectPreviewLayout.Height()-effectCanvas.Height())/2),
				)

				descriptionCanvas := gui.CreateTextCanvas(effect.Description(), tesOblivion24Font, 225)
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

func (layout *primaryLayout) render() {
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

func newAddIngredientButtonClickedEvent(inSlot alchemist.Slot) {
	err := event.TriggerEvent(&addIngredientButtonClicked{slot: inSlot, BasicEvent: event.BasicEvent{}})
	if err != nil {
		log.Fatal(err.Error())
	}
}
