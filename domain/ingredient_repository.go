package domain

import (
	"errors"
	"strings"
)

type IngredientRepository struct {
	ingredients []Ingredient
}

type IngredientFinder interface {
	FindByName(name string) (Ingredient, error)
	FindByNames(names []string) ([]Ingredient, error)
	All() []Ingredient
}

func initStorage() IngredientFinder {
	return IngredientRepository{
		ingredients: []Ingredient{
			{
				"Bread",
				[]Effect{
					{name: "Fatigue+", description: "Restores fatigue", power: 10},
					{name: "Hunger-", description: "Decreases hunger", power: -20},
				},
			},
			{
				"Cheese",
				[]Effect{
					{name: "Fatigue+", description: "Restores fatigue", power: 30},
					{name: "Resist Fire+", description: "Increases fire resistance", power: 10},
				},
			},
		},
	}
}

func (r IngredientRepository) FindByName(name string) (Ingredient, error) {
	for _, ingredient := range r.ingredients {
		if strings.EqualFold(ingredient.name, name) {
			return ingredient, nil
		}
	}

	return Ingredient{}, errors.New("ingredient \"" + name + "\" not found")
}

func (r IngredientRepository) FindByNames(names []string) ([]Ingredient, error) {
	var ingredients []Ingredient

	for _, name := range names {
		ingredient, err := r.FindByName(name)
		if err != nil {
			return []Ingredient{}, err
		}
		ingredients = append(ingredients, ingredient)
	}

	return ingredients, nil
}

func (r IngredientRepository) All() []Ingredient {
	return r.ingredients
}


var IngredientsDatabase = initStorage()
