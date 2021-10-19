package domain_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/TemirkhanN/alchemist/domain"
)

var ingredients = []struct {
	name    string
	effects []string
}{
	{
		name:    "Alkanet Flower",
		effects: []string{"Restore Intelligence", "Resist Poison", "Light", "Damage Fatigue"},
	},
	{
		name:    "Aloe Vera Leaves",
		effects: []string{"Restore Fatigue", "Restore Health", "Damage Magicka", "Invisibility"},
	},
	{
		name:    "Arrowroot",
		effects: []string{"Restore Agility", "Damage Luck", "Fortify Strength", "Burden"},
	},
	{
		name:    "Beef",
		effects: []string{"Restore Fatigue", "Shield", "Fortify Agility", "Dispel"},
	},
	{
		name:    "Bergamot Seeds",
		effects: []string{"Resist Disease", "Dispel", "Damage Magicka", "Silence"},
	},
	{
		name:    "Blackberry",
		effects: []string{"Restore Fatigue", "Resist Shock", "Fortify Endurance", "Restore Magicka"},
	},
	{
		name:    "Bloodgrass",
		effects: []string{"Chameleon", "Resist Paralysis", "Burden", "Fortify Health"},
	},
	{
		name:    "Boar Meat",
		effects: []string{"Restore Health", "Damage Speed", "Fortify Health", "Burden"},
	},
	{
		name:    "Bog Beacon Asco Cap",
		effects: []string{"Restore Magicka", "Shield", "Damage Personality", "Damage Endurance"},
	},
	{
		name:    "Bonemeal",
		effects: []string{"Damage Fatigue", "Resist Fire", "Fortify Luck", "Night-Eye"},
	},
	{
		name:    "Bread Loaf",
		effects: []string{"Restore Fatigue", "Detect Life", "Damage Agility", "Damage Strength"},
	},
	{
		name:    "Cairn Bolete Cap",
		effects: []string{"Restore Health", "Damage Intelligence", "Resist Paralysis", "Shock Damage"},
	},
	{
		name:    "Carrot",
		effects: []string{"Restore Fatigue", "Night-Eye", "Fortify Intelligence", "Damage Endurance"},
	},
	{
		name:    "Cheese Wheel",
		effects: []string{"Restore Fatigue", "Resist Paralysis", "Damage Luck", "Fortify Willpower"},
	},
	{
		name:    "Cinnabar Polypore Red Cap",
		effects: []string{"Restore Agility", "Shield", "Damage Personality", "Damage Endurance"},
	},
	{
		name: "Cinnabar Polypore Yellow Cap",
		effects: []string{
			"Restore Endurance", "Fortify Endurance", "Damage Personality",
			"Reflect Spell",
		},
	},
	{
		name:    "Clannfear Claws",
		effects: []string{"Cure Disease", "Resist Disease", "Paralyze", "Damage Health"},
	},
	{
		name: "Clouded Funnel Cap",
		effects: []string{
			"Restore Intelligence", "Fortify Intelligence", "Damage Endurance",
			"Damage Magicka",
		},
	},
	{
		name:    "Columbine Root Pulp",
		effects: []string{"Restore Personality", "Resist Frost", "Fortify Magicka", "Chameleon"},
	},
	{
		name:    "Corn",
		effects: []string{"Restore Fatigue", "Restore Intelligence", "Damage Agility", "Shock Shield"},
	},
	{
		name:    "Crab Meat",
		effects: []string{"Restore Endurance", "Resist Shock", "Damage Fatigue", "Fire Shield"},
	},
	{
		name:    "Daedra Heart",
		effects: []string{"Restore Health", "Shock Shield", "Damage Magicka", "Silence"},
	},
	{
		name:    "Daedra Silk",
		effects: []string{"Burden", "Night-Eye", "Chameleon", "Damage Endurance"},
	},
	{
		name:    "Daedroth Teeth",
		effects: []string{"Night-Eye", "Frost Shield", "Burden", "Light"},
	},
	{
		name:    "Dreugh Wax",
		effects: []string{"Damage Fatigue", "Resist Poison", "Water Breathing", "Damage Health"},
	},
	{
		name:    "Dryad Saddle Polypore Cap",
		effects: []string{"Restore Luck", "Resist Frost", "Damage Speed", "Frost Damage"},
	},
	{
		name:    "Ectoplasm",
		effects: []string{"Shock Damage", "Dispel", "Fortify Magicka", "Damage Health"},
	},
	{
		name:    "Elf Cup Cap",
		effects: []string{"Damage Willpower", "Cure Disease", "Fortify Strength", "Damage Intelligence"},
	},
	{
		name:    "Emetic Russula Cap",
		effects: []string{"Restore Agility", "Shield", "Damage Personality", "Damage Endurance"},
	},
	{
		name:    "Fennel Seeds",
		effects: []string{"Restore Fatigue", "Damage Intelligence", "Damage Magicka", "Paralyze"},
	},
	{
		name:    "Fire Salts",
		effects: []string{"Fire Damage", "Resist Frost", "Restore Magicka", "Fire Shield"},
	},
	{
		name:    "Flax Seeds",
		effects: []string{"Restore Magicka", "Feather", "Shield", "Damage Health"},
	},
	{
		name:    "Flour",
		effects: []string{"Restore Fatigue", "Damage Personality", "Fortify Fatigue", "Reflect Damage"},
	},
	{
		name:    "Fly Amanita Cap",
		effects: []string{"Restore Agility", "Burden", "Restore Health", "Shock Damage"},
	},
	{
		name:    "Foxglove Nectar",
		effects: []string{"Resist Poison", "Resist Paralysis", "Restore Luck", "Resist Disease"},
	},
	{
		name:    "Frost Salts",
		effects: []string{"Frost Damage", "Resist Fire", "Silence", "Frost Shield"},
	},
	{
		name:    "Garlic",
		effects: []string{"Resist Disease", "Damage Agility", "Frost Shield", "Fortify Strength"},
	},
	{
		name:    "Ginkgo Leaf",
		effects: []string{"Restore Speed", "Fortify Magicka", "Damage Luck", "Shock Damage"},
	},
	{
		name:    "Ginseng",
		effects: []string{"Damage Luck", "Cure Poison", "Burden", "Fortify Magicka"},
	},
	{
		name:    "Glow Dust",
		effects: []string{"Restore Speed", "Light", "Reflect Spell", "Damage Health"},
	},
	{
		name:    "Grapes",
		effects: []string{"Restore Fatigue", "Water Walking", "Dispel", "Damage Health"},
	},
	{
		name:    "Green Stain Cup Cap",
		effects: []string{"Restore Fatigue", "Damage Speed", "Reflect Damage", "Damage Health"},
	},
	{
		name:    "Green Stain Shelf Cap",
		effects: []string{"Restore Luck", "Fortify Luck", "Damage Fatigue", "Restore Health"},
	},
	{
		name:    "Ham",
		effects: []string{"Restore Fatigue", "Restore Health", "Damage Magicka", "Damage Luck"},
	},
	{
		name:    "Harrada",
		effects: []string{"Damage Health", "Damage Magicka", "Silence", "Paralyze"},
	},
	{
		name:    "Imp Gall",
		effects: []string{"Fortify Personality", "Cure Paralysis", "Damage Health", "Fire Damage"},
	},
	{
		name:    "Ironwood Nut",
		effects: []string{"Restore Intelligence", "Resist Fire", "Damage Fatigue", "Fortify Health"},
	},
	{
		name:    "Lady's Mantle Leaves",
		effects: []string{"Restore Health", "Damage Endurance", "Night-Eye", "Feather"},
	},
	{
		name:    "Lavender Sprig",
		effects: []string{"Restore Personality", "Fortify Willpower", "Restore Health", "Damage Luck"},
	},
	{
		name:    "Leek",
		effects: []string{"Restore Fatigue", "Fortify Agility", "Damage Personality", "Damage Strength"},
	},
	{
		name:    "Lettuce",
		effects: []string{"Restore Fatigue", "Restore Luck", "Fire Shield", "Damage Personality"},
	},
	{
		name:    "Lichor",
		effects: []string{"Restore Magicka"},
	},
	{
		name:    "Milk Thistle Seeds",
		effects: []string{"Light", "Frost Damage", "Cure Paralysis", "Paralyze"},
	},
	{
		name:    "Minotaur Horn",
		effects: []string{"Restore Willpower", "Burden", "Fortify Endurance", "Resist Paralysis"},
	},
	{
		name:    "Monkshood Root Pulp",
		effects: []string{"Restore Strength", "Damage Intelligence", "Fortify Endurance", "Burden"},
	},
	{
		name:    "Morning Glory Root Pulp",
		effects: []string{"Burden", "Damage Willpower", "Frost Shield", "Damage Magicka"},
	},
	{
		name:    "Mort Flesh",
		effects: []string{"Damage Fatigue", "Damage Luck", "Fortify Health", "Silence"},
	},
	{
		name:    "Motherwort Sprig",
		effects: []string{"Resist Poison", "Damage Fatigue", "Silence", "Invisibility"},
	},
	{
		name:    "Nightshade",
		effects: []string{"Damage Health", "Burden", "Damage Luck", "Fortify Magicka"},
	},
	{
		name:    "Ogre's Teeth",
		effects: []string{"Damage Intelligence", "Resist Paralysis", "Shock Damage", "Fortify Strength"},
	},
	{
		name:    "Onion",
		effects: []string{"Restore Fatigue", "Water Breathing", "Detect Life", "Damage Health"},
	},
	{
		name:    "Orange",
		effects: []string{"Restore Fatigue", "Detect Life", "Burden", "Shield"},
	},
	{
		name:    "Pear",
		effects: []string{"Restore Fatigue", "Damage Speed", "Fortify Speed", "Damage Health"},
	},
	{
		name:    "Peony Seeds",
		effects: []string{"Restore Strength", "Damage Health", "Damage Speed", "Restore Fatigue"},
	},
	{
		name:    "Potato",
		effects: []string{"Restore Fatigue", "Shield", "Burden", "Frost Shield"},
	},
	{
		name:    "Primrose Leaves",
		effects: []string{"Restore Willpower", "Restore Personality", "Fortify Luck", "Damage Strength"},
	},
	{
		name:    "Pumpkin",
		effects: []string{"Restore Fatigue", "Damage Agility", "Damage Personality", "Detect Life"},
	},
	{
		name:    "Purgeblood Salts",
		effects: []string{"Restore Magicka", "Damage Health", "Fortify Magicka", "Dispel"},
	},
	{
		name:    "Radish",
		effects: []string{"Restore Fatigue", "Damage Endurance", "Chameleon", "Burden"},
	},
	{
		name:    "Rat Meat",
		effects: []string{"Damage Fatigue", "Detect Life", "Damage Magicka", "Silence"},
	},
	{
		name:    "Redwort Flower",
		effects: []string{"Resist Frost", "Cure Poison", "Damage Health", "Invisibility"},
	},
	{
		name:    "Rice",
		effects: []string{"Restore Fatigue", "Silence", "Shock Shield", "Damage Agility"},
	},
	{
		name:    "Root Pulp",
		effects: []string{"Cure Disease", "Damage Willpower", "Fortify Strength", "Damage Intelligence"},
	},
	{
		name:    "Sacred Lotus Seeds",
		effects: []string{"Resist Frost", "Damage Health", "Feather", "Dispel"},
	},
	{
		name:    "Scales",
		effects: []string{"Damage Willpower", "Water Breathing", "Damage Health", "Water Walking"},
	},
	{
		name:    "Scamp Skin",
		effects: []string{"Damage Magicka", "Resist Shock", "Reflect Damage", "Damage Health"},
	},
	{
		name:    "Shepherd's Pie",
		effects: []string{"Cure Disease", "Shield", "Fortify Agility", "Dispel"},
	},
	{
		name:    "Somnalius Frond",
		effects: []string{"Restore Speed", "Damage Endurance", "Fortify Health", "Feather"},
	},
	{
		name:    "Spiddal Stick",
		effects: []string{"Damage Health", "Damage Magicka", "Fire Damage", "Restore Fatigue"},
	},
	{
		name:    "St. Jahn's Wort Nectar",
		effects: []string{"Resist Shock", "Damage Health", "Cure Poison", "Chameleon"},
	},
	{
		name:    "Steel-Blue Entoloma Cap",
		effects: []string{"Restore Magicka", "Fire Damage", "Resist Frost", "Burden"},
	},
	{
		name:    "Stinkhorn Cap",
		effects: []string{"Damage Health", "Restore Magicka", "Water Walking", "Invisibility"},
	},
	{
		name:    "Strawberry",
		effects: []string{"Restore Fatigue", "Cure Poison", "Damage Health", "Reflect Damage"},
	},
	{
		name:    "Summer Bolete Cap",
		effects: []string{"Restore Agility", "Shield", "Damage Personality", "Damage Endurance"},
	},
	{
		name:    "Sweetcake",
		effects: []string{"Restore Fatigue", "Feather", "Restore Health", "Burden"},
	},
	{
		name:    "Sweetroll",
		effects: []string{"Restore Fatigue", "Resist Disease", "Damage Personality", "Fortify Health"},
	},
	{
		name:    "Taproot",
		effects: []string{"Restore Luck", "Damage Endurance", "Resist Poison", "Shock Shield"},
	},
	{
		name:    "Tiger Lily Nectar",
		effects: []string{"Restore Endurance", "Damage Strength", "Water Walking", "Damage Willpower"},
	},
	{
		name:    "Tinder Polypore Cap",
		effects: []string{"Restore Willpower", "Resist Disease", "Invisibility", "Damage Magicka"},
	},
	{
		name:    "Tobacco",
		effects: []string{"Restore Fatigue", "Resist Paralysis", "Damage Magicka", "Dispel"},
	},
	{
		name:    "Tomato",
		effects: []string{"Restore Fatigue", "Detect Life", "Burden", "Shield"},
	},
	{
		name:    "Troll Fat",
		effects: []string{"Damage Agility", "Fortify Personality", "Damage Willpower", "Damage Health"},
	},
	{
		name:    "Vampire Dust",
		effects: []string{"Silence", "Resist Disease", "Frost Damage", "Invisibility"},
	},
	{
		name:    "Venison",
		effects: []string{"Restore Health", "Feather", "Damage Health", "Chameleon"},
	},
	{
		name:    "Viper's Bugloss Leaves",
		effects: []string{"Resist Paralysis", "Night-Eye", "Burden", "Cure Paralysis"},
	},
	{
		name:    "Water Hyacinth Nectar",
		effects: []string{"Damage Luck", "Damage Fatigue", "Restore Magicka", "Fortify Magicka"},
	},
	{
		name:    "Watermelon",
		effects: []string{"Restore Fatigue", "Light", "Burden", "Damage Health"},
	},
	{
		name:    "Wheat Grain",
		effects: []string{"Restore Fatigue", "Damage Magicka", "Fortify Health", "Damage Personality"},
	},
	{
		name:    "White Seed Pod",
		effects: []string{"Restore Strength", "Water Breathing", "Silence", "Light"},
	},
	{
		name:    "Wisp Stalk Caps",
		effects: []string{"Damage Health", "Damage Willpower", "Damage Intelligence", "Fortify Speed"},
	},
	{
		name:    "Wormwood Leaves",
		effects: []string{"Fortify Fatigue", "Invisibility", "Damage Health", "Damage Magicka"},
	},
}

func TestIngredientRepository_FindByName(t *testing.T) {
	for i := 0; i < len(ingredients); i++ {
		ingredientDetails := ingredients[i]
		t.Run(ingredientDetails.name, func(sub *testing.T) {
			sub.Parallel()
			assert := assert.New(sub)
			ingredient, err := domain.IngredientsDatabase.FindByName(ingredientDetails.name)

			assert.NoError(err)
			assert.Equal(ingredientDetails.name, ingredient.Name())

			assert.Len(ingredient.Effects(), len(ingredientDetails.effects))
			for effectPosition, effectName := range ingredientDetails.effects {
				assert.Equal(effectName, ingredient.Effects()[effectPosition].Name())
			}
		},
		)
	}
}

func TestIngredientRepository_FindByNames(t *testing.T) {
	chunkSize := 10
	totalAmountOfIngredients := len(ingredients)
	from := 0
	to := chunkSize

	for {
		chunk := ingredients[from:to]

		t.Run(fmt.Sprintf("Testing ingredients between %d %d", from, to), func(sub *testing.T) {
			sub.Parallel()
			assert := assert.New(sub)
			names := make([]string, len(chunk))
			for key, ingredientDetails := range chunk {
				names[key] = ingredientDetails.name
			}
			result, err := domain.IngredientsDatabase.FindByNames(names)
			assert.NoError(err)
			assert.Len(result, len(names))

			for key, ingredient := range result {
				ingredientDetails := chunk[key]
				assert.Equal(ingredientDetails.name, ingredient.Name())

				assert.Len(ingredient.Effects(), len(ingredientDetails.effects))
				for effectPosition, effectName := range ingredientDetails.effects {
					assert.Equal(effectName, ingredient.Effects()[effectPosition].Name())
				}
			}
		})

		from = to

		to += chunkSize
		if to >= totalAmountOfIngredients {
			to = totalAmountOfIngredients
		}

		if from >= totalAmountOfIngredients {
			break
		}
	}
}

func TestIngredientRepository_All(t *testing.T) {
	assert := assert.New(t)

	all := domain.IngredientsDatabase.All()

	assert.Len(all, len(ingredients))

	for key, ingredient := range all {
		ingredientDetails := ingredients[key]

		assert.Equal(ingredientDetails.name, ingredient.Name())

		assert.Len(ingredient.Effects(), len(ingredientDetails.effects))

		for effectPosition, effectName := range ingredientDetails.effects {
			assert.Equal(effectName, ingredient.Effects()[effectPosition].Name())
		}
	}
}
