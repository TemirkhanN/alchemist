package main

import (
	"errors"
)

type Mortar struct {
	ingredients []Ingredient
}

func (m *Mortar) AddIngredient(newIngredient Ingredient) error {
	if len(m.ingredients) == 0 {
		m.ingredients = append(m.ingredients, newIngredient)
		return nil
	}
	if len(m.ingredients) == 4 {
		return errors.New("there can not be more than 4 ingredients")
	}

	for _, existingIngredient := range m.ingredients {
		if existingIngredient.hasSimilarEffects(newIngredient) {
			m.ingredients = append(m.ingredients, newIngredient)
			return nil
		}
	}

	return errors.New("ingredients must have similar effects to be combined")
}

func (m *Mortar) Pestle() (Potion, error) {
	if len(m.ingredients) < 2 {
		return Potion{}, errors.New("there are not enough ingredients to create a potion")
	}

	potionEffects := make(map[string]Effect)
	for _, ingredient := range m.ingredients {
		for _, effect := range ingredient.effects {
			existingEffect, effectExists := potionEffects[effect.name]
			if effectExists {
				// todo type overflow
				effect.power += existingEffect.power
			}

			potionEffects[effect.name] = effect
		}
	}
	m.Clear()

	list := make([]Effect, 0, len(potionEffects))
	for _, potionEffect := range potionEffects {
		list = append(list, potionEffect)
	}

	return Potion{
		name:    "Some random potion name",
		effects: list,
	}, nil
}

func (m *Mortar) Clear() {
	m.ingredients = nil
}
