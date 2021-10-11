package domain

import "errors"

type Alchemist struct {
	luckLevel    int
	alchemyLevel int

	mortar                   *Mortar
	currentlyUsedIngredients []*Ingredient
}

func NewAlchemist(level int, luckLevel int, mortar *Mortar) *Alchemist {
	return &Alchemist{
		alchemyLevel:             level,
		luckLevel:                luckLevel,
		mortar:                   mortar,
		currentlyUsedIngredients: []*Ingredient{},
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

func (a *Alchemist) CanStartBrewing() bool {
	usedIngredientsAmount := len(a.UsedIngredients())
	if usedIngredientsAmount == 0 {
		return false
	}

	if usedIngredientsAmount == 1 && !a.IsMaster() {
		return false
	}

	return true
}

func (a *Alchemist) BrewPotion(potionName string) (Potion, error) {
	if !a.CanStartBrewing() {
		return Potion{}, errors.New("there are not enough ingredients to create a potion")
	}

	defer a.DiscardIngredients()

	usedIngredientsAmount := len(a.UsedIngredients())
	if usedIngredientsAmount == 1 && a.IsMaster() {
		theOnlyEffect := a.currentlyUsedIngredients[0].Effects()[0]
		theOnlyEffect.increased = true
		theOnlyEffect.power = a.effectStrength()

		return Potion{
			name:    potionName,
			effects: []Effect{theOnlyEffect},
		}, nil
	}

	allEffects := make(map[string]Effect)

	for _, ingredient := range a.currentlyUsedIngredients {
		for _, effect := range a.DetermineEffects(ingredient) {
			if effect.IsUnknown() {
				continue
			}

			_, effectExists := allEffects[effect.Name()]
			if effectExists {
				effect.increased = true
			}

			allEffects[effect.Name()] = effect
		}
	}

	potionEffects := make([]Effect, 0)

	for _, potionEffect := range allEffects {
		// remove effects that didn't match between multiple ingredients
		if !potionEffect.increased {
			continue
		}

		potionEffect.power = a.effectStrength()
		potionEffects = append(potionEffects, potionEffect)
	}

	return Potion{
		name:    potionName,
		effects: potionEffects,
	}, nil
}

func (a *Alchemist) IsNovice() bool {
	return a.alchemyLevel < 25
}

func (a *Alchemist) IsApprentice() bool {
	return a.alchemyLevel >= 25 && a.alchemyLevel < 50
}

func (a *Alchemist) IsJourneyMan() bool {
	return a.alchemyLevel >= 50 && a.alchemyLevel < 75
}

func (a *Alchemist) IsExpert() bool {
	return a.alchemyLevel >= 75 && a.alchemyLevel < 100
}

func (a *Alchemist) IsMaster() bool {
	return a.alchemyLevel == 100
}

func (a *Alchemist) IdentifiableAmountOfEffects() int {
	switch {
	case a.IsNovice():
		return 1
	case a.IsApprentice():
		return 2
	case a.IsJourneyMan():
		return 3
	case a.IsExpert():
		return 4
	case a.IsMaster():
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
