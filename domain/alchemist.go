package domain

import "errors"

type Mastery int

const (
	MasteryNovice Mastery = iota
	MasteryApprentice
	MasteryJourneyman
	MasteryExpert
	MasteryMaster
)

type Alchemist struct {
	luckLevel    int
	alchemyLevel int

	mortar                   *Mortar
	currentlyUsedIngredients []*Ingredient
}

func NewAlchemist(level int, luckLevel int, mortar *Mortar) *Alchemist {
	return &Alchemist{
		alchemyLevel: level,
		luckLevel:    luckLevel,
		mortar:       mortar,
	}
}

func (a *Alchemist) CanUseIngredient(ingredient *Ingredient) bool {
	if len(a.currentlyUsedIngredients) == 0 {
		return true
	}

	// todo this will produce logic error on attempt to switch one ingredient for another
	if len(a.currentlyUsedIngredients) == 4 {
		return false
	}

	canUse := false
	for _, usedIngredient := range a.currentlyUsedIngredients {
		if usedIngredient.Name() == ingredient.Name() {
			return false
		}

		if !canUse && a.CanCombineIngredients(usedIngredient, ingredient) {
			canUse = true
		}
	}

	return canUse
}

func (a *Alchemist) UseIngredient(newIngredient *Ingredient) error {
	if len(a.currentlyUsedIngredients) == 0 || a.CanUseIngredient(newIngredient) {
		a.currentlyUsedIngredients = append(a.currentlyUsedIngredients, newIngredient)
		return nil
	}

	return errors.New("ingredients must have similar effects to be combined")
}

func (a *Alchemist) UsedIngredients() []*Ingredient {
	return a.currentlyUsedIngredients
}

func (a *Alchemist) CanCombineIngredients(ingredient1 *Ingredient, ingredient2 *Ingredient) bool {
	for _, effect1 := range a.DetermineEffects(ingredient1) {
		if effect1.IsUnknown() {
			continue
		}
		for _, effect2 := range a.DetermineEffects(ingredient2) {
			if effect2.IsUnknown() {
				continue
			}
			if effect1.name == effect2.name {
				return true
			}
		}
	}

	return false
}

func (a *Alchemist) DetermineEffects(ingredient *Ingredient) []Effect {
	identifiableAmountOfEffects := a.IdentifiableAmountOfEffects()

	var effects []Effect
	for i := 0; i < 4; i++ {
		effect := ingredient.effects[i]
		if i+1 > identifiableAmountOfEffects {
			effect = HideEffect(effect)
		}

		effects = append(effects, effect)
	}

	return effects
}

func (a *Alchemist) DiscardIngredients() {
	a.currentlyUsedIngredients = nil
}

func (a *Alchemist) BrewPotion(potionName string) (Potion, error) {
	usedIngredientsAmount := len(a.UsedIngredients())
	if usedIngredientsAmount < 2 && a.Mastery() != MasteryMaster {
		return Potion{}, errors.New("there are not enough ingredients to create a potion")
	}

	potionEffects := make(map[string]Effect)
	for _, ingredient := range a.currentlyUsedIngredients {
		for _, effect := range a.DetermineEffects(ingredient) {
			if effect.IsUnknown() {
				continue
			}
			_, effectExists := potionEffects[effect.name]
			if effectExists {
				// todo type overflow
				effect.increased = true
			}

			if a.Mastery() == MasteryMaster && usedIngredientsAmount == 1 {
				effect.increased = true
			}
			potionEffects[effect.name] = effect
		}
	}
	a.DiscardIngredients()

	list := make([]Effect, 0, len(potionEffects))
	for _, potionEffect := range potionEffects {
		// remove effects that didn't match between multiple ingredients
		if !potionEffect.increased {
			continue
		}
		potionEffect.power = a.effectStrength()
		list = append(list, potionEffect)
	}

	return Potion{
		name:    potionName,
		effects: list,
	}, nil
}

func (a *Alchemist) Mastery() Mastery {
	switch {
	case a.alchemyLevel < 25:
		return MasteryNovice
	case a.alchemyLevel < 50:
		return MasteryApprentice
	case a.alchemyLevel < 75:
		return MasteryJourneyman
	case a.alchemyLevel < 100:
		return MasteryExpert
	case a.alchemyLevel == 100:
		return MasteryMaster
	default:
		panic("alchemist has unknown mastery level. Probably wrong level set somehow")
	}
}

func (a *Alchemist) IdentifiableAmountOfEffects() int {
	switch a.Mastery() {
	case MasteryNovice:
		return 1
	case MasteryApprentice:
		return 2
	case MasteryJourneyman:
		return 3
	case MasteryExpert:
		return 4
	case MasteryMaster:
		return 4
	default:
		panic("alchemist has unknown mastery level. Probably wrong level set somehow")
	}
}

func (a *Alchemist) effectStrength() int {
	effectiveLevel := a.alchemyLevel + int(0.4*float64(a.luckLevel-50))
	if effectiveLevel < 0 {
		effectiveLevel = 0
	}
	if effectiveLevel > 100 {
		effectiveLevel = 100
	}
	return effectiveLevel + int(a.mortar.Strength())
}
