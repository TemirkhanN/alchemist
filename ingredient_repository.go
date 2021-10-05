package main

import (
	"errors"
	"strings"
)

type IngredientRepository struct {
	ingredients []Ingredient
}

type IngredientFinder interface {
	FindByName(name string) Ingredient
	FindByNames(names []string) []Ingredient
}

func initStorage() IngredientRepository {
	return IngredientRepository{
		ingredients: []Ingredient{
			{
				"Bread",
				[]Effect{
					{"Stamina+", "Restores stamina", 10},
					{"Hunger-", "Decreases hunger", -20},
				},
			},
			{
				"Salmon",
				[]Effect{
					{"Stamina+", "Restores stamina", 30},
					{"Hunger-", "Decreases hunger", -50},
				},
			},
		},
	}
}

func (r IngredientRepository) FindByName(name string) (Ingredient, error) {
	for _, ingredient := range r.ingredients {
		if strings.ToLower(ingredient.name) == strings.ToLower(name) {
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
