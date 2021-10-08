package main

import (
	"github.com/TemirkhanN/alchemist/domain"
	"github.com/gookit/event"
)

type IngredientSelected struct {
	ingredient *domain.Ingredient
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
