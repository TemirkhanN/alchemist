package main

import (
	"log"

	"github.com/gookit/event"

	"github.com/TemirkhanN/alchemist/domain"
)

type IngredientSelected struct {
	ingredient *domain.Ingredient
	event.BasicEvent
}

type AddIngredientButtonClicked struct {
	slot Slot
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

func newAddIngredientButtonClickedEvent(inSlot Slot) {
	err := event.TriggerEvent(&AddIngredientButtonClicked{slot: inSlot, BasicEvent: event.BasicEvent{}})
	if err != nil {
		log.Fatal(err.Error())
	}
}
