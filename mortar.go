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
	for _, existingIngredient := range m.ingredients {
		if existingIngredient.hasSimilarEffects(newIngredient) {
			m.ingredients = append(m.ingredients, newIngredient)
			return nil
		}
	}

	return errors.New("ingredients must have similar effects to be combined")
}

func (m *Mortar) Pestle() (p Potion, e error) {
	if len(m.ingredients) < 2 {
		return p, errors.New("there are not enough ingredients to create a potion")
	}

	potionEffects := make(map[string]Effect)

	for _, ingredient := range m.ingredients {
		for _, effect := range ingredient.effects {
			existingEffect, effectExists := potionEffects[effect.name]
			var newEffect Effect
			if effectExists {
				newEffect = Effect{
					name: existingEffect.name,
					description: existingEffect.description,
					// todo type overflow
					power: existingEffect.power + effect.power,
				}
			} else {
				newEffect = effect
			}

			potionEffects[effect.name] = newEffect
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
