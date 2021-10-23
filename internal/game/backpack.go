package game

import (
	"log"

	"github.com/gookit/event"

	"github.com/TemirkhanN/alchemist/pkg/alchemy/alchemist"
	"github.com/TemirkhanN/alchemist/pkg/alchemy/ingredient"
	"github.com/TemirkhanN/alchemist/pkg/gui"
)

type backpackLayout struct {
	initialized     bool
	graphics        *gui.Layer
	window          *gui.Window
	background      *gui.SpriteCanvas
	ingredientsBtns []*gui.Button
	closeButton     *gui.Button

	ingredients []*ingredient.Ingredient
	alchemist   *alchemist.Alchemist
}

// newBackpackLayout todo rename repo to backpack.
func newBackpackLayout(window *gui.Window, player *alchemist.Alchemist) *backpackLayout {
	layout := new(backpackLayout)
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

	closeButtonSprite := gameAssets.GetSprite("btn.exit")
	ingredientsLayoutSprite := gameAssets.GetSprite("interface.ingredients")

	layout.graphics = gui.CreateLayer(window.Width(), window.Height(), false)
	layout.background = window.CreateSpriteCanvas(ingredientsLayoutSprite)

	layout.closeButton = window.CreateButton(closeButtonSprite)
	layout.closeButton.SetClickHandler(func() { layout.graphics.Hide() })

	event.On(eventAddIngredientButtonClicked, event.ListenerFunc(func(e event.Event) error {
		layout.render()

		return nil
	}))

	return layout
}

func (layout *backpackLayout) render() {
	layout.graphics.Clear()
	layout.graphics.AddElement(layout.background, gui.ZeroPosition)

	ingredientsLayer := gui.CreateLayer(480, 465, true, true)
	ingredientEffectsLayer := gui.CreateLayer(238, 220, false)
	ingredientsEffectsLayerBackground := layout.window.CreateSpriteCanvas(gameAssets.GetSprite("interface.effects"))

	layout.ingredientsBtns = nil
	offset := ingredientsLayer.Height()

	for _, ingr := range layout.ingredients {
		if !layout.alchemist.CanUseIngredient(ingr) {
			continue
		}

		ingredientBtn := layout.window.CreateButton(getIngredientSprite(*ingr))
		ingredientBtn.SetClickHandler(func(selected *ingredient.Ingredient) func() {
			return func() {
				// todo potentially vulnerable for mistake on main(mortar) side
				layout.graphics.Hide()
				err := event.TriggerEvent(&ingredientSelected{ingredient: selected, BasicEvent: event.BasicEvent{}})
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
					effectPreview := layout.window.CreateSpriteCanvas(gameAssets.GetSprite(effect.Name()))
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
