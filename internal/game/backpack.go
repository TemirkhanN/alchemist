package game

import (
	"log"

	"github.com/gookit/event"

	"github.com/TemirkhanN/alchemist/pkg/alchemy/alchemist"
	"github.com/TemirkhanN/alchemist/pkg/alchemy/ingredient"
	"github.com/TemirkhanN/alchemist/pkg/gui/geometry"
	"github.com/TemirkhanN/alchemist/pkg/gui/graphics"
)

type backpackLayout struct {
	initialized     bool
	graphics        *graphics.Layout
	background      *graphics.SpriteCanvas
	ingredientsBtns []*graphics.Button
	closeButton     *graphics.Button

	ingredients []*ingredient.Ingredient
	alchemist   *alchemist.Alchemist
}

func newBackpackLayout(window *graphics.Window, player *alchemist.Alchemist) *backpackLayout {
	layout := new(backpackLayout)
	layout.initialized = true
	layout.alchemist = player

	for _, ingr := range ingredient.IngredientsDatabase.All() {
		deref := ingr
		layout.ingredients = append(layout.ingredients, &deref)
	}

	closeButtonSprite := gameAssets.GetSprite("btn.exit")
	ingredientsLayoutSprite := gameAssets.GetSprite("interface.ingredients")

	layout.graphics = graphics.NewLayout(window.Width(), window.Height(), false)
	layout.background = graphics.NewSpriteCanvas(ingredientsLayoutSprite)

	layout.closeButton = graphics.NewButton(closeButtonSprite)
	layout.closeButton.SetClickHandler(func(geometry.Position) { layout.graphics.Hide() })

	event.On(eventAddIngredientButtonClicked, event.ListenerFunc(func(e event.Event) error {
		layout.render()

		return nil
	}))

	return layout
}

func (layout *backpackLayout) render() {
	layout.graphics.Clear()
	layout.graphics.AddElement(layout.background)

	ingredientsLayer := graphics.NewLayout(480, 465, true, true)
	ingredientEffectsLayer := graphics.NewLayout(238, 220, false)
	ingredientsEffectsLayerBackground := graphics.NewSpriteCanvas(gameAssets.GetSprite("interface.effects"))

	layout.ingredientsBtns = nil
	offset := ingredientsLayer.Height()

	for _, ingr := range layout.ingredients {
		if !layout.alchemist.CanUseIngredient(ingr) {
			continue
		}

		ingredientBtn := graphics.NewButton(getIngredientSprite(*ingr))
		ingredientBtn.SetClickHandler(func(selected *ingredient.Ingredient) func(geometry.Position) {
			return func(geometry.Position) {
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
				ingredientEffectsLayer.AddElement(ingredientsEffectsLayerBackground)
				posY := ingredientEffectsLayer.Height()
				for _, effect := range layout.alchemist.DetermineEffects(hovered) {
					posY -= 55
					effectPreview := graphics.NewSpriteCanvas(gameAssets.GetSprite(effect.Name()))
					ingredientEffectsLayer.AddElement(effectPreview, geometry.NewPosition(0, posY))
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

		ingredientsLayer.AddElement(ingredientBtn, geometry.NewPosition(0, offset))
	}

	layout.graphics.AddElement(ingredientsLayer, geometry.NewPosition(50, 115))
	layout.graphics.AddElement(ingredientEffectsLayer, geometry.NewPosition(605, 200))
	layout.graphics.AddElement(layout.closeButton, geometry.NewPosition(410, 65))

	layout.graphics.Show()
}
