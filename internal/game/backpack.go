package game

import (
	"log"

	"github.com/gookit/event"

	"github.com/TemirkhanN/alchemist/assets"
	"github.com/TemirkhanN/alchemist/pkg/alchemy/alchemist"
	"github.com/TemirkhanN/alchemist/pkg/alchemy/ingredient"
	"github.com/TemirkhanN/alchemist/pkg/gui"
)

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

// newBackpackLayout todo rename repo to backpack.
func newBackpackLayout(window *gui.Window, player *alchemist.Alchemist) *BackpackLayout {
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

		ingredientBtn := layout.window.CreateButton(getIngredientSprite(*ingr))
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
