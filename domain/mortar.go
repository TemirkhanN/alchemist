package domain

import (
	"errors"
)

type MortarLevel float32

const (
	Novice MortarLevel = 0.1
	//Apprentice = 0.25
	//JOURNEYMAN = 0.5
	//EXPERT = 0.75
	//MASTER = 1
)

type Mortar struct {
	alchemistLevel int
	level          MortarLevel
	ingredients    []*Ingredient
}

func (m *Mortar) AddIngredient(newIngredient *Ingredient) error {
	if m.IsEmpty() {
		m.ingredients = append(m.ingredients, newIngredient)
		return nil
	}

	if !m.IngredientAllowed(newIngredient) {
		return errors.New("can not add this ingredient")
	}

	for _, existingIngredient := range m.ingredients {
		if m.HaveSimilarEffects(existingIngredient, newIngredient) {
			m.ingredients = append(m.ingredients, newIngredient)
			return nil
		}
	}

	return errors.New("ingredients must have similar effects to be combined")
}

func (m *Mortar) IngredientAllowed(ingredient *Ingredient) bool {
	amount := len(m.ingredients)
	if amount == 0 {
		return true
	}

	if amount == 4 {
		return false
	}

	// todo is it allowed to mix ingredients that can be mixed but has no identified effects yet?
	for _, existingIngredient := range m.ingredients {
		if m.HaveSimilarEffects(existingIngredient, ingredient) {
			return true
		}
	}

	return false
}

func (m *Mortar) Ingredients() []*Ingredient {
	return m.ingredients
}

func (m *Mortar) Pestle() (Potion, error) {
	if len(m.ingredients) < 2 {
		return Potion{}, errors.New("there are not enough ingredients to create a potion")
	}

	potionEffects := make(map[string]Effect)
	for _, ingredient := range m.ingredients {
		for _, effect := range m.DetermineEffects(ingredient) {
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
		potionEffect.power = potionEffect.power * int16(m.level*25)
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

func NewNoviceMortar(alchemistLevel int) *Mortar {
	if alchemistLevel > 100 || alchemistLevel < 1 {
		panic("todo runtime")
	}

	return &Mortar{
		alchemistLevel: alchemistLevel,
		level:          MortarLevel(Novice),
		ingredients:    nil,
	}
}

func (m *Mortar) IsEmpty() bool {
	return len(m.ingredients) == 0
}

func (m *Mortar) HaveSimilarEffects(ingredient1 *Ingredient, ingredient2 *Ingredient) bool {
	for _, effect1 := range m.DetermineEffects(ingredient1) {
		for _, effect2 := range m.DetermineEffects(ingredient2) {
			if effect1.name == effect2.name {
				return true
			}
		}
	}

	return false
}

func (m *Mortar) DetermineEffects(ingredient *Ingredient) []Effect {
	identifiableAmountOfEffects := m.getIdentifiableAmountOfEffects()

	return ingredient.effects[:identifiableAmountOfEffects]
}

func (m *Mortar) getIdentifiableAmountOfEffects() int {
	switch true {
	case m.alchemistLevel < 25 :
		return 1
	case m.alchemistLevel < 50 :
		return 2
	case m.alchemistLevel < 75 :
		return 3
	case m.alchemistLevel < 100 :
		return 4
	case m.alchemistLevel == 100 :
		return 4
	default:
		return 1

	}
}
