package domain

type Effect struct {
	name        string
	description string
	power       int16
	increased   bool
}

type Ingredient struct {
	name    string
	effects []Effect
}

func (i Ingredient) hasSimilarEffects(withIngredient Ingredient) bool {
	for _, effect1 := range i.effects {
		for _, effect2 := range withIngredient.effects {
			if effect1.name == effect2.name {
				return true
			}
		}
	}

	return false
}

func (i Ingredient) Name() string {
	return i.name
}
