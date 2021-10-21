package alchemist

import (
	"errors"
	"math"

	"github.com/TemirkhanN/alchemist/pkg/alchemy/ingredient"
)

type Slot struct {
	value uint8
}

type Alchemist struct {
	luckLevel    int
	alchemyLevel int

	mortar                   *Mortar
	currentlyUsedIngredients []*ingredient.Ingredient
}

func NewAlchemist(level int, luckLevel int, mortar *Mortar) *Alchemist {
	return &Alchemist{
		alchemyLevel:             level,
		luckLevel:                luckLevel,
		mortar:                   mortar,
		currentlyUsedIngredients: []*ingredient.Ingredient{},
	}
}

func (a *Alchemist) CanUseIngredient(ingredient *ingredient.Ingredient) bool {
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

func (a *Alchemist) UseIngredient(newIngredient *ingredient.Ingredient) error {
	if len(a.currentlyUsedIngredients) == 0 || a.CanUseIngredient(newIngredient) {
		a.currentlyUsedIngredients = append(a.currentlyUsedIngredients, newIngredient)

		return nil
	}

	return errors.New("ingredients must have similar effects to be combined")
}

func (a *Alchemist) UsedIngredients() []*ingredient.Ingredient {
	return a.currentlyUsedIngredients
}

func (a *Alchemist) CanCombineIngredients(ingredient1 *ingredient.Ingredient, ingredient2 *ingredient.Ingredient) bool {
	for _, effect1 := range a.DetermineEffects(ingredient1) {
		if effect1.IsUnknown() {
			continue
		}

		for _, effect2 := range a.DetermineEffects(ingredient2) {
			if effect2.IsUnknown() {
				continue
			}

			if effect1.Name() == effect2.Name() {
				return true
			}
		}
	}

	return false
}

func (a *Alchemist) DetermineEffects(ofIngredient *ingredient.Ingredient) []ingredient.Effect {
	identifiableAmountOfEffects := a.IdentifiableAmountOfEffects()
	ingredientEffectsAmount := len(ofIngredient.Effects())

	var effects []ingredient.Effect

	for i := 0; i < 4 && ingredientEffectsAmount > i; i++ {
		effect := ofIngredient.Effects()[i]
		if i+1 > identifiableAmountOfEffects {
			effect = effect.HideEffectDetails()
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

	potion, err := a.PredictPotion()
	if err != nil {
		return potion, err
	}

	a.DiscardIngredients()

	potion.name = potionName

	return potion, nil
}

func (a Alchemist) PredictPotion() (Potion, error) {
	if !a.CanStartBrewing() {
		return Potion{}, errors.New("there are not enough ingredients to create a potion")
	}

	usedIngredientsAmount := len(a.UsedIngredients())
	if usedIngredientsAmount == 1 && a.IsMaster() {
		theOnlyEffect := a.currentlyUsedIngredients[0].Effects()[0]

		return Potion{
			name:    "",
			effects: []PotionEffect{a.Refine(theOnlyEffect)},
		}, nil
	}

	potionEffects := make(map[string]PotionEffect)
	allEffects := make(map[string]ingredient.Effect)

	for _, usedIngredient := range a.currentlyUsedIngredients {
		for _, effect := range a.DetermineEffects(usedIngredient) {
			if effect.IsUnknown() {
				continue
			}

			_, effectExists := allEffects[effect.Name()]
			if !effectExists {
				allEffects[effect.Name()] = effect

				continue
			}

			potionEffects[effect.Name()] = a.Refine(effect)
		}
	}

	effects := make([]PotionEffect, 0)
	for _, effect := range potionEffects {
		effects = append(effects, effect)
	}

	return Potion{
		name:    "",
		effects: effects,
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

func (a *Alchemist) Refine(effect ingredient.Effect) PotionEffect {
	magnitude := math.Round(a.calculateMagnitude(effect))
	if magnitude < 1 {
		magnitude = 1
	}

	duration := math.Round(a.calculateDuration(effect))
	if duration < 1 {
		duration = 1
	}

	return PotionEffect{
		magnitude: magnitude,
		duration:  duration,
		effect:    effect,
	}
}

func (a *Alchemist) effectiveAlchemyLevel() float64 {
	effectiveLevel := float64(a.alchemyLevel) + (0.4 * float64(a.luckLevel-50))
	if effectiveLevel < 0 {
		return 0
	}

	if effectiveLevel > 100 {
		return 100
	}

	return effectiveLevel
}

func (a *Alchemist) calculateMagnitude(effect ingredient.Effect) float64 {
	if effect.IsDurationOnly() {
		return 1.0
	}

	delta := 4.0
	if effect.IsMagnitudeOnly() {
		delta = 1.0
	}

	return math.Pow((a.effectiveAlchemyLevel()+a.mortar.Strength())/(effect.BaseCost()/10*delta), 1/2.28)
}

func (a *Alchemist) calculateDuration(effect ingredient.Effect) float64 {
	if effect.IsMagnitudeOnly() {
		return 1.0
	}

	if effect.IsDurationOnly() {
		return (a.effectiveAlchemyLevel() + a.mortar.Strength()) / (effect.BaseCost() / 10)
	}

	return 4 * a.calculateMagnitude(effect)
}

var (
	EmptySlot  = Slot{value: 0}
	FirstSlot  = Slot{value: 1}
	SecondSlot = Slot{value: 2}
	ThirdSlot  = Slot{value: 3}
	FourthSlot = Slot{value: 4}
)
