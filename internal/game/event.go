package game

import (
	"log"

	"github.com/gookit/event"

	"github.com/TemirkhanN/alchemist/pkg/alchemy/alchemist"
	"github.com/TemirkhanN/alchemist/pkg/alchemy/ingredient"
)

type IngredientSelected struct {
	ingredient *ingredient.Ingredient
	event.BasicEvent
}

type AddIngredientButtonClicked struct {
	slot alchemist.Slot
	event.BasicEvent
}

func (e IngredientSelected) Name() string {
	return EventIngredientSelected
}

func (e AddIngredientButtonClicked) Name() string {
	return EventAddIngredientButtonClicked
}

const (
	EventIngredientSelected         = "ingredientSelected"
	EventAddIngredientButtonClicked = "addIngredientButtonClicked"
)

func newAddIngredientButtonClickedEvent(inSlot alchemist.Slot) {
	err := event.TriggerEvent(&AddIngredientButtonClicked{slot: inSlot, BasicEvent: event.BasicEvent{}})
	if err != nil {
		log.Fatal(err.Error())
	}
}
