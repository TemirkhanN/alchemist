package main

type Effect struct {
	name        string
	description string
	power       int16
}

type Ingredient struct {
	sprite  string
	name    string
	effects []Effect
}

func (ingredient1 Ingredient) hasSimilarEffects(ingredient2 Ingredient) bool {
	for _, effect1 := range ingredient1.effects {
		for _, effect2 := range ingredient2.effects {
			if effect1.name == effect2.name {
				return true
			}
		}
	}

	return false
}
