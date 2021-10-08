package domain

import (
	"errors"
)

type MortarLevel float32

const (
	APPRENTICE MortarLevel = 0.1
	//NOVICE = 0.25
	//JOURNEYMAN = 0.5
	//EXPERT = 0.75
	//MASTER = 1
)

type Mortar struct {
	alchemyLevel MortarLevel
	ingredients  []Ingredient
}

func (m *Mortar) AddIngredient(newIngredient Ingredient) error {
	if len(m.ingredients) == 0 {
		m.ingredients = append(m.ingredients, newIngredient)
		return nil
	}

	if !m.IngredientAllowed(newIngredient) {
		return errors.New("can not add this ingredient")
	}

	for _, existingIngredient := range m.ingredients {
		if existingIngredient.hasSimilarEffects(newIngredient) {
			m.ingredients = append(m.ingredients, newIngredient)
			return nil
		}
	}

	return errors.New("ingredients must have similar effects to be combined")
}

func (m Mortar) IngredientAllowed(ingredient Ingredient) bool {
	amount := len(m.ingredients)
	if amount == 0 {
		return true
	}

	if amount == 4 {
		return false
	}

	for _, existingIngredient := range m.ingredients {
		if existingIngredient.hasSimilarEffects(ingredient) {
			return true
		}
	}

	return false
}

func (m Mortar) Ingredients() []Ingredient {
	return m.ingredients
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
				effect.increased = true
			}

			potionEffects[effect.name] = effect
		}
	}
	m.Clear()

	list := make([]Effect, 0, len(potionEffects))
	for _, potionEffect := range potionEffects {
		// remove effects that didn't match between multiple ingredients
		if !potionEffect.increased {
			continue
		}
		potionEffect.power = potionEffect.power * int16(m.alchemyLevel * 25)
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

func NewApprenticeMortar() *Mortar {
	return &Mortar{
		alchemyLevel: MortarLevel(APPRENTICE),
		ingredients:  nil,
	}
}
