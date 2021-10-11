package domain

import (
	"errors"
	"strings"
)

type IngredientRepository struct {
	ingredients []Ingredient
}

type IngredientFinder interface {
	FindByName(name string) (Ingredient, error)
	FindByNames(names []string) ([]Ingredient, error)
	All() []Ingredient
}

func (r IngredientRepository) FindByName(name string) (Ingredient, error) {
	for _, ingredient := range r.ingredients {
		if strings.EqualFold(ingredient.name, name) {
			return ingredient, nil
		}
	}

	return Ingredient{}, errors.New("ingredient \"" + name + "\" not found")
}

func (r IngredientRepository) FindByNames(names []string) ([]Ingredient, error) {
	ingredients := make([]Ingredient, len(names))

	for key, name := range names {
		ingredient, err := r.FindByName(name)
		if err != nil {
			return []Ingredient{}, err
		}

		ingredients[key] = ingredient
	}

	return ingredients, nil
}

func (r IngredientRepository) All() []Ingredient {
	return r.ingredients
}

var (
	allKnownIngredients = []Ingredient{
		{
			"Alkanet Flower",
			[]Effect{
				restoreIntelligenceEffect,
				resistPoisonEffect,
				lightEffect,
				damageFatigueEffect,
			},
		}, {
			"Aloe Vera Leaves",
			[]Effect{
				restoreFatigueEffect,
				restoreHealthEffect,
				damageMagickaEffect,
				invisibilityEffect,
			},
		}, {
			"Ambrosia",
			[]Effect{
				restoreFatigueEffect,
				damageLuckEffect,
				fortifyWillpowerEffect,
				damageHealthEffect,
			},
		}, {
			"Arrowroot",
			[]Effect{
				restoreAgilityEffect,
				damageLuckEffect,
				fortifyStrengthEffect,
				burdenEffect,
			},
		}, {
			"Beef",
			[]Effect{
				restoreFatigueEffect,
				shieldEffect,
				fortifyAgilityEffect,
				dispelEffect,
			},
		}, {
			"Bergamot Seeds",
			[]Effect{
				resistDiseaseEffect,
				dispelEffect,
				damageMagickaEffect,
				silenceEffect,
			},
		}, {
			"Blackberry",
			[]Effect{
				restoreFatigueEffect,
				resistShockEffect,
				fortifyEnduranceEffect,
				restoreMagickaEffect,
			},
		}, {
			"Bloodgrass",
			[]Effect{
				chameleonEffect,
				resistParalysisEffect,
				burdenEffect,
				fortifyHealthEffect,
			},
		}, {
			"Boar Meat",
			[]Effect{
				restoreHealthEffect,
				damageSpeedEffect,
				fortifyHealthEffect,
				burdenEffect,
			},
		}, {
			"Bog Beacon Asco Cap",
			[]Effect{
				restoreMagickaEffect,
				shieldEffect,
				damagePersonalityEffect,
				damageEnduranceEffect,
			},
		}, {
			"Bonemeal",
			[]Effect{
				restoreFatigueEffect,
				detectLifeEffect,
				damageAgilityEffect,
				damageStrengthEffect,
			},
		}, {
			"Cairn Bolete Cap",
			[]Effect{
				restoreHealthEffect,
				damageIntelligenceEffect,
				resistParalysisEffect,
				shockDamageEffect,
			},
		}, {
			"Carrot",
			[]Effect{
				restoreFatigueEffect,
				resistFireEffect, // todo night eye
				fortifyIntelligenceEffect,
				damageEnduranceEffect,
			},
		}, {
			"Cheese Wheel",
			[]Effect{
				restoreFatigueEffect,
				resistParalysisEffect,
				damageLuckEffect,
				fortifyWillpowerEffect,
			},
		}, {
			"Cinnabar Polypore Red Cap",
			[]Effect{
				restoreAgilityEffect,
				shieldEffect,
				damagePersonalityEffect,
				damageEnduranceEffect,
			},
		}, {
			"Cinnabar Polypore Yellow Cap",
			[]Effect{
				restoreEnduranceEffect,
				fortifyEnduranceEffect,
				damagePersonalityEffect,
				reflectSpellEffect,
			},
		}, {
			"Clannfear Claws",
			[]Effect{
				cureDiseaseEffect,
				resistDiseaseEffect,
				paralyzeEffect,
				damageHealthEffect,
			},
		}, {
			"Clouded Funnel Cap",
			[]Effect{
				restoreIntelligenceEffect,
				fortifyIntelligenceEffect,
				damageEnduranceEffect,
				damageMagickaEffect,
			},
		}, {
			"Columbine Root Pulp",
			[]Effect{
				restorePersonalityEffect,
				resistFrostEffect,
				fortifyMagickaEffect,
				chameleonEffect,
			},
		}, {
			"Corn",
			[]Effect{
				restoreFatigueEffect,
				restoreIntelligenceEffect,
				damageAgilityEffect,
				shockShieldEffect,
			},
		}, {
			"Crab Meat",
			[]Effect{
				restoreEnduranceEffect,
				resistShockEffect,
				damageFatigueEffect,
				fireShieldEffect,
			},
		}, {
			"Daedra Heart",
			[]Effect{
				restoreHealthEffect,
				shockShieldEffect,
				damageMagickaEffect,
				silenceEffect,
			},
		}, {
			"Daedra Silk",
			[]Effect{
				paralyzeEffect,
				restoreFatigueEffect,
				damageHealthEffect,
				reflectDamageEffect,
			},
		}, {
			"Daedroth Teeth",
			[]Effect{
				resistFireEffect,
				damageHealthEffect,
				restoreHealthEffect,
				fireShieldEffect,
			},
		}, {
			"Dreugh Wax",
			[]Effect{
				damageFatigueEffect,
				resistPoisonEffect,
				waterBreathingEffect,
				damageHealthEffect,
			},
		}, {
			"Dryad Saddle Polypore Cap",
			[]Effect{
				restoreLuckEffect,
				resistFrostEffect,
				damageSpeedEffect,
				frostDamageEffect,
			},
		}, {
			"Ectoplasm",
			[]Effect{
				shockDamageEffect,
				dispelEffect,
				fortifyMagickaEffect,
				damageHealthEffect,
			},
		}, {
			"Elf Cup Cap",
			[]Effect{
				damageWillpowerEffect,
				cureDiseaseEffect,
				fortifyStrengthEffect,
				damageIntelligenceEffect,
			},
		}, {
			"Emetic Russula Cap",
			[]Effect{
				restoreAgilityEffect,
				shieldEffect,
				damagePersonalityEffect,
				damageEnduranceEffect,
			},
		}, {
			"Fennel Seeds",
			[]Effect{
				restoreFatigueEffect,
				damageIntelligenceEffect,
				damageMagickaEffect,
				paralyzeEffect,
			},
		}, {
			"Fire Salts",
			[]Effect{
				fireDamageEffect,
				resistFrostEffect,
				restoreMagickaEffect,
				fireShieldEffect,
			},
		}, {
			"Flax Seeds",
			[]Effect{
				restoreMagickaEffect,
				featherEffect,
				shieldEffect,
				damageHealthEffect,
			},
		}, {
			"Flour",
			[]Effect{
				restoreFatigueEffect,
				damagePersonalityEffect,
				fortifyFatigueEffect,
				reflectDamageEffect,
			},
		}, {
			"Fly Amanita Cap",
			[]Effect{
				restoreAgilityEffect,
				burdenEffect,
				restoreHealthEffect,
				shockDamageEffect,
			},
		}, {
			"Foxglove Nectar",
			[]Effect{
				resistPoisonEffect,
				resistParalysisEffect,
				restoreLuckEffect,
				resistDiseaseEffect,
			},
		}, {
			"Frost Salts",
			[]Effect{
				frostDamageEffect,
				resistFireEffect,
				silenceEffect,
				frostShieldEffect,
			},
		}, {
			"Garlic",
			[]Effect{
				resistDiseaseEffect,
				damageAgilityEffect,
				frostShieldEffect,
				fortifyStrengthEffect,
			},
		}, {
			"Ginkgo Leaf",
			[]Effect{
				restoreSpeedEffect,
				fortifyMagickaEffect,
				damageLuckEffect,
				shockDamageEffect,
			},
		}, {
			"Ginseng",
			[]Effect{
				damageLuckEffect,
				curePoisonEffect,
				burdenEffect,
				fortifyMagickaEffect,
			},
		}, {
			"Glow Dust",
			[]Effect{
				restoreSpeedEffect,
				lightEffect,
				reflectSpellEffect,
				damageHealthEffect,
			},
		}, {
			"Grapes",
			[]Effect{
				restoreFatigueEffect,
				waterWalkingEffect,
				dispelEffect,
				damageHealthEffect,
			},
		}, {
			"Green Stain Cup Cap",
			[]Effect{
				restoreFatigueEffect,
				damageSpeedEffect,
				reflectDamageEffect,
				damageHealthEffect,
			},
		}, {
			"Green Stain Shelf Cap",
			[]Effect{
				restoreLuckEffect,
				fortifyLuckEffect,
				damageFatigueEffect,
				restoreHealthEffect,
			},
		}, {
			"Ham",
			[]Effect{
				restoreFatigueEffect,
				restoreHealthEffect,
				damageMagickaEffect,
				damageLuckEffect,
			},
		}, {
			"Harrada",
			[]Effect{
				damageHealthEffect,
				damageMagickaEffect,
				silenceEffect,
				paralyzeEffect,
			},
		}, {
			"Imp Gall",
			[]Effect{
				fortifyPersonalityEffect,
				cureParalysisEffect,
				damageHealthEffect,
				fireDamageEffect,
			},
		}, {
			"Ironwood Nut",
			[]Effect{
				restoreIntelligenceEffect,
				resistFireEffect,
				damageFatigueEffect,
				fortifyHealthEffect,
			},
		}, {
			"Lady's Mantle Leaves",
			[]Effect{
				restoreIntelligenceEffect,
				resistFireEffect,
				damageFatigueEffect,
				fortifyHealthEffect,
			},
		}, {
			"Lavender Sprig",
			[]Effect{
				restorePersonalityEffect,
				fortifyWillpowerEffect,
				restoreHealthEffect,
				damageLuckEffect,
			},
		}, {
			"Leek",
			[]Effect{
				restoreFatigueEffect,
				fortifyAgilityEffect,
				damagePersonalityEffect,
				damageStrengthEffect,
			},
		}, {
			"Lettuce",
			[]Effect{
				restoreFatigueEffect,
				restoreLuckEffect,
				fireShieldEffect,
				damagePersonalityEffect,
			},
		}, {
			"Lichor",
			[]Effect{
				cureDiseaseEffect,
				resistPoisonEffect,
				damageAgilityEffect,
				fortifyWillpowerEffect,
			},
		}, {
			"Milk Thistle Seeds",
			[]Effect{
				lightEffect,
				frostDamageEffect,
				cureParalysisEffect,
				paralyzeEffect,
			},
		}, {
			"Minotaur Horn",
			[]Effect{
				restoreWillpowerEffect,
				burdenEffect,
				fortifyEnduranceEffect,
				resistParalysisEffect,
			},
		}, {
			"Monkshood Root Pulp",
			[]Effect{
				restoreStrengthEffect,
				damageIntelligenceEffect,
				fortifyEnduranceEffect,
				burdenEffect,
			},
		}, {
			"Morning Glory Root Pulp",
			[]Effect{
				burdenEffect,
				damageWillpowerEffect,
				frostShieldEffect,
				damageMagickaEffect,
			},
		}, {
			"Mort Flesh",
			[]Effect{
				damageFatigueEffect,
				damageLuckEffect,
				fortifyHealthEffect,
				silenceEffect,
			},
		}, {
			"Motherwort Sprig",
			[]Effect{
				resistPoisonEffect,
				damageFatigueEffect,
				silenceEffect,
				invisibilityEffect,
			},
		}, {
			"Mugwort Seeds",
			[]Effect{
				fortifyHealthEffect,
				damageFatigueEffect,
				dispelEffect,
				damageMagickaEffect,
			},
		}, {
			"Nightshade",
			[]Effect{
				damageHealthEffect,
				burdenEffect,
				damageLuckEffect,
				fortifyMagickaEffect,
			},
		}, {
			"Ogre's Teeth",
			[]Effect{
				damageIntelligenceEffect,
				resistParalysisEffect,
				shockDamageEffect,
				fortifyStrengthEffect,
			},
		}, {
			"Onion",
			[]Effect{
				restoreFatigueEffect,
				waterBreathingEffect,
				detectLifeEffect,
				damageHealthEffect,
			},
		}, {
			"Orange",
			[]Effect{
				restoreFatigueEffect,
				detectLifeEffect,
				burdenEffect,
				shieldEffect,
			},
		}, {
			"Pear",
			[]Effect{
				restoreFatigueEffect,
				damageSpeedEffect,
				fortifySpeedEffect,
				damageHealthEffect,
			},
		}, {
			"Peony Seeds",
			[]Effect{
				restoreStrengthEffect,
				damageHealthEffect,
				damageSpeedEffect,
				restoreFatigueEffect,
			},
		}, {
			"Potato",
			[]Effect{
				restoreFatigueEffect,
				shieldEffect,
				burdenEffect,
				frostShieldEffect,
			},
		}, {
			"Primrose Leaves",
			[]Effect{
				restoreWillpowerEffect,
				restorePersonalityEffect,
				fortifyLuckEffect,
				damageStrengthEffect,
			},
		}, {
			"Pumpkin",
			[]Effect{
				restoreFatigueEffect,
				damageAgilityEffect,
				damagePersonalityEffect,
				detectLifeEffect,
			},
		}, {
			"Purgeblood Salts",
			[]Effect{
				restoreMagickaEffect,
				damageHealthEffect,
				fortifyMagickaEffect,
				dispelEffect,
			},
		}, {
			"Radish",
			[]Effect{
				restoreFatigueEffect,
				damageEnduranceEffect,
				chameleonEffect,
				burdenEffect,
			},
		}, {
			"Rat Meat",
			[]Effect{
				damageFatigueEffect,
				detectLifeEffect,
				damageMagickaEffect,
				silenceEffect,
			},
		}, {
			"Redwort Flower",
			[]Effect{
				resistFrostEffect,
				curePoisonEffect,
				damageHealthEffect,
				invisibilityEffect,
			},
		}, {
			"Rice",
			[]Effect{
				restoreFatigueEffect,
				silenceEffect,
				shockShieldEffect,
				damageAgilityEffect,
			},
		}, {
			"Root Pulp",
			[]Effect{
				cureDiseaseEffect,
				damageWillpowerEffect,
				fortifyStrengthEffect,
				damageIntelligenceEffect,
			},
		}, {
			"Sacred Lotus Seeds",
			[]Effect{
				resistFrostEffect,
				damageHealthEffect,
				featherEffect,
				dispelEffect,
			},
		}, {
			"Scales",
			[]Effect{
				damageWillpowerEffect,
				waterBreathingEffect,
				damageHealthEffect,
				waterWalkingEffect,
			},
		}, {
			"Scamp Skin",
			[]Effect{
				damageMagickaEffect,
				resistShockEffect,
				reflectDamageEffect,
				damageHealthEffect,
			},
		}, {
			"Shepherd's Pie",
			[]Effect{
				cureDiseaseEffect,
				shieldEffect,
				fortifyAgilityEffect,
				dispelEffect,
			},
		}, {
			"S'jirra's Famous Potato Bread",
			[]Effect{
				detectLifeEffect,
				restoreHealthEffect,
				damageAgilityEffect,
				damageStrengthEffect,
			},
		}, {
			"Somnalius Frond",
			[]Effect{
				restoreSpeedEffect,
				damageEnduranceEffect,
				fortifyHealthEffect,
				featherEffect,
			},
		}, {
			"Spiddal Stick",
			[]Effect{
				damageHealthEffect,
				damageMagickaEffect,
				fireDamageEffect,
				restoreFatigueEffect,
			},
		}, {
			"St. Jahn's Wort Nectar",
			[]Effect{
				resistShockEffect,
				damageHealthEffect,
				curePoisonEffect,
				chameleonEffect,
			},
		}, {
			"Steel-Blue Entoloma Cap",
			[]Effect{
				restoreMagickaEffect,
				fireDamageEffect,
				resistFrostEffect,
				burdenEffect,
			},
		}, {
			"Stinkhorn Cap",
			[]Effect{
				damageHealthEffect,
				restoreMagickaEffect,
				waterWalkingEffect,
				invisibilityEffect,
			},
		}, {
			"Strawberry",
			[]Effect{
				restoreFatigueEffect,
				curePoisonEffect,
				damageHealthEffect,
				reflectDamageEffect,
			},
		}, {
			"Summer Bolete Cap",
			[]Effect{
				restoreAgilityEffect,
				shieldEffect,
				damagePersonalityEffect,
				damageEnduranceEffect,
			},
		}, {
			"Sweetcake",
			[]Effect{
				restoreFatigueEffect,
				featherEffect,
				restoreHealthEffect,
				burdenEffect,
			},
		}, {
			"Sweetroll",
			[]Effect{
				restoreFatigueEffect,
				resistDiseaseEffect,
				damagePersonalityEffect,
				fortifyHealthEffect,
			},
		}, {
			"Taproot",
			[]Effect{
				restoreLuckEffect,
				damageEnduranceEffect,
				resistPoisonEffect,
				shockShieldEffect,
			},
		}, {
			"Tiger Lily Nectar",
			[]Effect{
				restoreEnduranceEffect,
				damageStrengthEffect,
				waterWalkingEffect,
				damageWillpowerEffect,
			},
		}, {
			"Tinder Polypore Cap",
			[]Effect{
				restoreWillpowerEffect,
				resistDiseaseEffect,
				invisibilityEffect,
				damageMagickaEffect,
			},
		}, {
			"Tobacco",
			[]Effect{
				restoreFatigueEffect,
				resistParalysisEffect,
				damageMagickaEffect,
				dispelEffect,
			},
		}, {
			"Tomato",
			[]Effect{
				restoreFatigueEffect,
				detectLifeEffect,
				burdenEffect,
				shieldEffect,
			},
		}, {
			"Troll Fat",
			[]Effect{
				damageAgilityEffect,
				fortifyPersonalityEffect,
				damageWillpowerEffect,
				damageHealthEffect,
			},
		}, {
			"Vampire Dust",
			[]Effect{
				silenceEffect,
				resistDiseaseEffect,
				frostDamageEffect,
				invisibilityEffect,
			},
		}, {
			"Venison",
			[]Effect{
				restoreHealthEffect,
				featherEffect,
				damageHealthEffect,
				chameleonEffect,
			},
		}, {
			"Viper's Bugloss Leaves",
			[]Effect{
				restoreMagickaEffect,
				damageHealthEffect,
				fortifyMagickaEffect,
				dispelEffect,
			},
		}, {
			"Water Hyacinth Nectar",
			[]Effect{
				damageLuckEffect,
				damageFatigueEffect,
				restoreMagickaEffect,
				fortifyMagickaEffect,
			},
		}, {
			"Watermelon",
			[]Effect{
				restoreFatigueEffect,
				lightEffect,
				burdenEffect,
				damageHealthEffect,
			},
		}, {
			"Wheat Grain",
			[]Effect{
				restoreFatigueEffect,
				damageMagickaEffect,
				fortifyHealthEffect,
				damagePersonalityEffect,
			},
		}, {
			"White Seed Pod",
			[]Effect{
				restoreStrengthEffect,
				waterBreathingEffect,
				silenceEffect,
				lightEffect,
			},
		}, {
			"Wisp Stalk Caps",
			[]Effect{
				damageHealthEffect,
				damageWillpowerEffect,
				damageIntelligenceEffect,
				fortifySpeedEffect,
			},
		}, {
			"Wormwood Leaves",
			[]Effect{
				fortifyFatigueEffect,
				invisibilityEffect,
				damageHealthEffect,
				damageMagickaEffect,
			},
		}, {
			"Ashes of Hindaril",
			[]Effect{
				silenceEffect,
				resistDiseaseEffect,
				frostDamageEffect,
				invisibilityEffect,
			},
		},
	}

	IngredientsDatabase = IngredientRepository{ingredients: allKnownIngredients}
)
