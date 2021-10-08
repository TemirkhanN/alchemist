package main

import "github.com/gookit/event"

type IngredientSelected struct {
	ingredient *Ingredient
	event.BasicEvent
}

type AddIngredientButtonClicked struct {
	slot Slot
	event.BasicEvent
}

func (e IngredientSelected) Name() string{
	return EventIngredientSelected
}

func (e AddIngredientButtonClicked) Name() string {
	return EventAddIngredientButtonClicked
}

const (
	EventIngredientSelected = "ingredientSelected"
	EventAddIngredientButtonClicked = "addIngredientButtonClicked"
)
