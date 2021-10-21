package game

import (
	"github.com/gookit/event"

	"github.com/TemirkhanN/alchemist/pkg/alchemy/alchemist"
	"github.com/TemirkhanN/alchemist/pkg/alchemy/ingredient"
)

type ingredientSelected struct {
	ingredient *ingredient.Ingredient
	event.BasicEvent
}

type addIngredientButtonClicked struct {
	slot alchemist.Slot
	event.BasicEvent
}

func (e ingredientSelected) Name() string {
	return eventIngredientSelected
}

func (e addIngredientButtonClicked) Name() string {
	return eventAddIngredientButtonClicked
}

const (
	eventIngredientSelected         = "ingredientSelected"
	eventAddIngredientButtonClicked = "addIngredientButtonClicked"
)
