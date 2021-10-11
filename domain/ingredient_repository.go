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

func initStorage() IngredientFinder {
	return IngredientRepository{
		ingredients: []Ingredient{
			{
				"Alkanet Flower",
				[]Effect{
					{name: "Restore Intelligence", eType: "+", power: 1},
					{name: "Resist Poison", eType: "+", power: 1},
					{name: "Light", eType: "+", power: 1},
					{name: "Damage Fatigue", eType: "-", power: 1},
				},
			}, {
				"Aloe Vera Leaves",
				[]Effect{
					{name: "Restore Fatigue", eType: "+", power: 1},
					{name: "Restore Health", eType: "+", power: 1},
					{name: "Damage Magicka", eType: "-", power: 1},
					{name: "Invisibility", eType: "+", power: 1},
				},
			}, {
				"Ambrosia",
				[]Effect{
					{name: "Restore Fatigue", eType: "+", power: 1},
					{name: "Damage Luck", eType: "-", power: 1},
					{name: "Fortify Willpower", eType: "+", power: 1},
					{name: "Damage Health", eType: "-", power: 1},
				},
			}, {
				"Arrowroot",
				[]Effect{
					{name: "Restore Agility", eType: "+", power: 1},
					{name: "Damage Luck", eType: "-", power: 1},
					{name: "Fortify Strength", eType: "+", power: 1},
					{name: "Burden", eType: "-", power: 1},
				},
			}, {
				"Beef",
				[]Effect{
					{name: "Restore Fatigue", eType: "+", power: 1},
					{name: "Shield", eType: "+", power: 1},
					{name: "Fortify Agility", eType: "+", power: 1},
					{name: "Dispel", eType: "+", power: 1},
				},
			}, {
				"Bergamot Seeds",
				[]Effect{
					{name: "Resist Disease", eType: "+", power: 1},
					{name: "Dispel", eType: "+", power: 1},
					{name: "Damage Magicka", eType: "-", power: 1},
					{name: "Silence", eType: "-", power: 1},
				},
			}, {
				"Blackberry",
				[]Effect{
					{name: "Restore Fatigue", eType: "+", power: 1},
					{name: "Resist Shock", eType: "+", power: 1},
					{name: "Fortify Endurance", eType: "+", power: 1},
					{name: "Restore Magicka", eType: "+", power: 1},
				},
			}, {
				"Bloodgrass",
				[]Effect{
					{name: "Chameleon", eType: "+", power: 1},
					{name: "Resist Paralysis", eType: "+", power: 1},
					{name: "Burden", eType: "-", power: 1},
					{name: "Fortify Health", eType: "+", power: 1},
				},
			}, {
				"Boar Meat",
				[]Effect{
					{name: "Restore Health", eType: "+", power: 1},
					{name: "Damage Speed", eType: "-", power: 1},
					{name: "Fortify Health", eType: "+", power: 1},
					{name: "Burden", eType: "-", power: 1},
				},
			}, {
				"Bog Beacon Asco Cap",
				[]Effect{
					{name: "Restore Magicka", eType: "+", power: 1},
					{name: "Shield", eType: "+", power: 1},
					{name: "Damage Personality", eType: "-", power: 1},
					{name: "Damage Endurance", eType: "-", power: 1},
				},
			}, {
				"Bonemeal",
				[]Effect{
					{name: "Restore Fatigue", eType: "+", power: 1},
					{name: "Detect Life", eType: "+", power: 1},
					{name: "Damage Agility", eType: "-", power: 1},
					{name: "Damage Strength", eType: "-", power: 1},
				},
			}, {
				"Cairn Bolete Cap",
				[]Effect{
					{name: "Restore Health", eType: "+", power: 1},
					{name: "Damage Intelligence", eType: "-", power: 1},
					{name: "Resist Paralysis", eType: "+", power: 1},
					{name: "Shock Damage", eType: "-", power: 1},
				},
			}, {
				"Carrot",
				[]Effect{
					{name: "Restore Fatigue", eType: "+", power: 1},
					{name: "Resist Fire", eType: "+", power: 1},
					{name: "Fire Shield", eType: "+", power: 1},
					{name: "Damage Agility", eType: "-", power: 1},
				},
			}, {
				"Cheese Wheel",
				[]Effect{
					{name: "Restore Fatigue", eType: "+", power: 1},
					{name: "Resist Paralysis", eType: "+", power: 1},
					{name: "Damage Luck", eType: "-", power: 1},
					{name: "Fortify Willpower", eType: "+", power: 1},
				},
			}, {
				"Cinnabar Polypore Red Cap",
				[]Effect{
					{name: "Restore Agility", eType: "+", power: 1},
					{name: "Shield", eType: "+", power: 1},
					{name: "Damage Personality", eType: "-", power: 1},
					{name: "Damage Endurance", eType: "-", power: 1},
				},
			}, {
				"Cinnabar Polypore Yellow Cap",
				[]Effect{
					{name: "Restore Endurance", eType: "+", power: 1},
					{name: "Fortify Endurance", eType: "+", power: 1},
					{name: "Damage Personality", eType: "-", power: 1},
					{name: "Reflect Spell", eType: "+", power: 1},
				},
			}, {
				"Clannfear Claws",
				[]Effect{
					{name: "Cure Disease", eType: "+", power: 1},
					{name: "Resist Disease", eType: "+", power: 1},
					{name: "Paralyze", eType: "-", power: 1},
					{name: "Damage Health", eType: "-", power: 1},
				},
			}, {
				"Clouded Funnel Cap",
				[]Effect{
					{name: "Restore Intelligence", eType: "+", power: 1},
					{name: "Fortify Intelligence", eType: "+", power: 1},
					{name: "Damage Endurance", eType: "-", power: 1},
					{name: "Damage Magicka", eType: "-", power: 1},
				},
			}, {
				"Columbine Root Pulp",
				[]Effect{
					{name: "Restore Personality", eType: "+", power: 1},
					{name: "Resist Frost", eType: "+", power: 1},
					{name: "Fortify Magicka", eType: "+", power: 1},
					{name: "Chameleon", eType: "+", power: 1},
				},
			}, {
				"Corn",
				[]Effect{
					{name: "Restore Fatigue", eType: "+", power: 1},
					{name: "Restore Intelligence", eType: "+", power: 1},
					{name: "Damage Agility", eType: "-", power: 1},
					{name: "Shock Shield", eType: "+", power: 1},
				},
			}, {
				"Crab Meat",
				[]Effect{
					{name: "Restore Endurance", eType: "+", power: 1},
					{name: "Resist Shock", eType: "+", power: 1},
					{name: "Damage Fatigue", eType: "-", power: 1},
					{name: "Fire Shield", eType: "+", power: 1},
				},
			}, {
				"Daedra Heart",
				[]Effect{
					{name: "Restore Health", eType: "+", power: 1},
					{name: "Shock Shield", eType: "+", power: 1},
					{name: "Damage Magicka", eType: "-", power: 1},
					{name: "Silence", eType: "-", power: 1},
				},
			}, {
				"Daedra Silk",
				[]Effect{
					{name: "Paralyze", eType: "-", power: 1},
					{name: "Restore Fatigue", eType: "+", power: 1},
					{name: "Damage Health", eType: "-", power: 1},
					{name: "Reflect Damage", eType: "+", power: 1},
				},
			}, {
				"Daedroth Teeth",
				[]Effect{
					{name: "Resist Fire", eType: "+", power: 1},
					{name: "Damage Health", eType: "-", power: 1},
					{name: "Restore Health", eType: "+", power: 1},
					{name: "Fire Shield", eType: "+", power: 1},
				},
			}, {
				"Dreugh Wax",
				[]Effect{
					{name: "Damage Fatigue", eType: "-", power: 1},
					{name: "Resist Poison", eType: "+", power: 1},
					{name: "Water Breathing", eType: "+", power: 1},
					{name: "Damage Health", eType: "-", power: 1},
				},
			}, {
				"Dryad Saddle Polypore Cap",
				[]Effect{
					{name: "Restore Luck", eType: "+", power: 1},
					{name: "Resist Frost", eType: "+", power: 1},
					{name: "Damage Speed", eType: "-", power: 1},
					{name: "Frost Damage", eType: "-", power: 1},
				},
			}, {
				"Ectoplasm",
				[]Effect{
					{name: "Shock Damage", eType: "-", power: 1},
					{name: "Dispel", eType: "+", power: 1},
					{name: "Fortify Magicka", eType: "+", power: 1},
					{name: "Damage Health", eType: "-", power: 1},
				},
			}, {
				"Elf Cup Cap",
				[]Effect{
					{name: "Damage Willpower", eType: "-", power: 1},
					{name: "Cure Disease", eType: "+", power: 1},
					{name: "Fortify Strength", eType: "+", power: 1},
					{name: "Damage Intelligence", eType: "-", power: 1},
				},
			}, {
				"Emetic Russula Cap",
				[]Effect{
					{name: "Restore Agility", eType: "+", power: 1},
					{name: "Shield", eType: "+", power: 1},
					{name: "Damage Personality", eType: "-", power: 1},
					{name: "Damage Endurance", eType: "-", power: 1},
				},
			}, {
				"Fennel Seeds",
				[]Effect{
					{name: "Restore Fatigue", eType: "+", power: 1},
					{name: "Damage Intelligence", eType: "-", power: 1},
					{name: "Damage Magicka", eType: "-", power: 1},
					{name: "Paralyze", eType: "-", power: 1},
				},
			}, {
				"Fire Salts",
				[]Effect{
					{name: "Fire Damage", eType: "-", power: 1},
					{name: "Resist Frost", eType: "+", power: 1},
					{name: "Restore Magicka", eType: "+", power: 1},
					{name: "Fire Shield", eType: "+", power: 1},
				},
			}, {
				"Flax Seeds",
				[]Effect{
					{name: "Restore Magicka", eType: "+", power: 1},
					{name: "Feather", eType: "+", power: 1},
					{name: "Shield", eType: "+", power: 1},
					{name: "Damage Health", eType: "-", power: 1},
				},
			}, {
				"Flour",
				[]Effect{
					{name: "Restore Fatigue", eType: "+", power: 1},
					{name: "Damage Personality", eType: "-", power: 1},
					{name: "Fortify Fatigue", eType: "+", power: 1},
					{name: "Reflect Damage", eType: "+", power: 1},
				},
			}, {
				"Fly Amanita Cap",
				[]Effect{
					{name: "Restore Agility", eType: "+", power: 1},
					{name: "Burden", eType: "-", power: 1},
					{name: "Restore Health", eType: "+", power: 1},
					{name: "Shock Damage", eType: "-", power: 1},
				},
			}, {
				"Foxglove Nectar",
				[]Effect{
					{name: "Resist Poison", eType: "+", power: 1},
					{name: "Resist Paralysis", eType: "+", power: 1},
					{name: "Restore Luck", eType: "+", power: 1},
					{name: "Resist Disease", eType: "+", power: 1},
				},
			}, {
				"Frost Salts",
				[]Effect{
					{name: "Frost Damage", eType: "-", power: 1},
					{name: "Resist Fire", eType: "+", power: 1},
					{name: "Silence", eType: "-", power: 1},
					{name: "Frost Shield", eType: "+", power: 1},
				},
			}, {
				"Garlic",
				[]Effect{
					{name: "Resist Disease", eType: "+", power: 1},
					{name: "Damage Agility", eType: "-", power: 1},
					{name: "Frost Shield", eType: "+", power: 1},
					{name: "Fortify Strength", eType: "+", power: 1},
				},
			}, {
				"Ginkgo Leaf",
				[]Effect{
					{name: "Restore Speed", eType: "+", power: 1},
					{name: "Fortify Magicka", eType: "+", power: 1},
					{name: "Damage Luck", eType: "-", power: 1},
					{name: "Shock Damage", eType: "-", power: 1},
				},
			}, {
				"Ginseng",
				[]Effect{
					{name: "Damage Luck", eType: "-", power: 1},
					{name: "Cure Poison", eType: "+", power: 1},
					{name: "Burden", eType: "-", power: 1},
					{name: "Fortify Magicka", eType: "+", power: 1},
				},
			}, {
				"Glow Dust",
				[]Effect{
					{name: "Restore Speed", eType: "+", power: 1},
					{name: "Light", eType: "+", power: 1},
					{name: "Reflect Spell", eType: "+", power: 1},
					{name: "Damage Health", eType: "-", power: 1},
				},
			}, {
				"Grapes",
				[]Effect{
					{name: "Restore Fatigue", eType: "+", power: 1},
					{name: "Water Walking", eType: "+", power: 1},
					{name: "Dispel", eType: "+", power: 1},
					{name: "Damage Health", eType: "-", power: 1},
				},
			}, {
				"Green Stain Cup Cap",
				[]Effect{
					{name: "Restore Fatigue", eType: "+", power: 1},
					{name: "Damage Speed", eType: "-", power: 1},
					{name: "Reflect Damage", eType: "+", power: 1},
					{name: "Damage Health", eType: "-", power: 1},
				},
			}, {
				"Green Stain Shelf Cap",
				[]Effect{
					{name: "Restore Luck", eType: "+", power: 1},
					{name: "Fortify Luck", eType: "+", power: 1},
					{name: "Damage Fatigue", eType: "-", power: 1},
					{name: "Restore Health", eType: "+", power: 1},
				},
			}, {
				"Ham",
				[]Effect{
					{name: "Restore Fatigue", eType: "+", power: 1},
					{name: "Restore Health", eType: "+", power: 1},
					{name: "Damage Magicka", eType: "-", power: 1},
					{name: "Damage Luck", eType: "-", power: 1},
				},
			}, {
				"Harrada",
				[]Effect{
					{name: "Damage Health", eType: "-", power: 1},
					{name: "Damage Magicka", eType: "-", power: 1},
					{name: "Silence", eType: "-", power: 1},
					{name: "Paralyze", eType: "-", power: 1},
				},
			}, {
				"Imp Gall",
				[]Effect{
					{name: "Fortify Personality", eType: "+", power: 1},
					{name: "Cure Paralysis", eType: "+", power: 1},
					{name: "Damage Health", eType: "-", power: 1},
					{name: "Fire Damage", eType: "-", power: 1},
				},
			}, {
				"Ironwood Nut",
				[]Effect{
					{name: "Restore Intelligence", eType: "+", power: 1},
					{name: "Resist Fire", eType: "+", power: 1},
					{name: "Damage Fatigue", eType: "-", power: 1},
					{name: "Fortify Health", eType: "+", power: 1},
				},
			}, {
				"Lady's Mantle Leaves",
				[]Effect{
					{name: "Restore Intelligence", eType: "+", power: 1},
					{name: "Resist Fire", eType: "+", power: 1},
					{name: "Damage Fatigue", eType: "-", power: 1},
					{name: "Fortify Health", eType: "+", power: 1},
				},
			}, {
				"Lavender Sprig",
				[]Effect{
					{name: "Restore Personality", eType: "+", power: 1},
					{name: "Fortify Willpower", eType: "+", power: 1},
					{name: "Restore Health", eType: "+", power: 1},
					{name: "Damage Luck", eType: "-", power: 1},
				},
			}, {
				"Leek",
				[]Effect{
					{name: "Restore Fatigue", eType: "+", power: 1},
					{name: "Fortify Agility", eType: "+", power: 1},
					{name: "Damage Personality", eType: "-", power: 1},
					{name: "Damage Strength", eType: "-", power: 1},
				},
			}, {
				"Lettuce",
				[]Effect{
					{name: "Restore Fatigue", eType: "+", power: 1},
					{name: "Restore Luck", eType: "+", power: 1},
					{name: "Fire Shield", eType: "+", power: 1},
					{name: "Damage Personality", eType: "-", power: 1},
				},
			}, {
				"Lichor",
				[]Effect{
					{name: "Cure Disease", eType: "+", power: 1},
					{name: "Resist Poison", eType: "+", power: 1},
					{name: "Damage Agility", eType: "-", power: 1},
					{name: "Fortify Willpower", eType: "+", power: 1},
				},
			}, {
				"Milk Thistle Seeds",
				[]Effect{
					{name: "Light", eType: "+", power: 1},
					{name: "Frost Damage", eType: "-", power: 1},
					{name: "Cure Paralysis", eType: "+", power: 1},
					{name: "Paralyze", eType: "-", power: 1},
				},
			}, {
				"Minotaur Horn",
				[]Effect{
					{name: "Restore Willpower", eType: "+", power: 1},
					{name: "Burden", eType: "-", power: 1},
					{name: "Fortify Endurance", eType: "+", power: 1},
					{name: "Resist Paralysis", eType: "+", power: 1},
				},
			}, {
				"Monkshood Root Pulp",
				[]Effect{
					{name: "Restore Strength", eType: "+", power: 1},
					{name: "Damage Intelligence", eType: "-", power: 1},
					{name: "Fortify Endurance", eType: "+", power: 1},
					{name: "Burden", eType: "-", power: 1},
				},
			}, {
				"Morning Glory Root Pulp",
				[]Effect{
					{name: "Burden", eType: "-", power: 1},
					{name: "Damage Willpower", eType: "-", power: 1},
					{name: "Frost Shield", eType: "+", power: 1},
					{name: "Damage Magicka", eType: "-", power: 1},
				},
			}, {
				"Mort Flesh",
				[]Effect{
					{name: "Damage Fatigue", eType: "-", power: 1},
					{name: "Damage Luck", eType: "-", power: 1},
					{name: "Fortify Health", eType: "+", power: 1},
					{name: "Silence", eType: "-", power: 1},
				},
			}, {
				"Motherwort Sprig",
				[]Effect{
					{name: "Resist Poison", eType: "+", power: 1},
					{name: "Damage Fatigue", eType: "-", power: 1},
					{name: "Silence", eType: "-", power: 1},
					{name: "Invisibility", eType: "+", power: 1},
				},
			}, {
				"Mugwort Seeds",
				[]Effect{
					{name: "Fortify Health", eType: "+", power: 1},
					{name: "Damage Fatigue", eType: "-", power: 1},
					{name: "Dispel", eType: "+", power: 1},
					{name: "Damage Magicka", eType: "-", power: 1},
				},
			}, {
				"Nightshade",
				[]Effect{
					{name: "Damage Health", eType: "-", power: 1},
					{name: "Burden", eType: "-", power: 1},
					{name: "Damage Luck", eType: "-", power: 1},
					{name: "Fortify Magicka", eType: "+", power: 1},
				},
			}, {
				"Ogre's Teeth",
				[]Effect{
					{name: "Damage Intelligence", eType: "-", power: 1},
					{name: "Resist Paralysis", eType: "+", power: 1},
					{name: "Shock Damage", eType: "-", power: 1},
					{name: "Fortify Strength", eType: "+", power: 1},
				},
			}, {
				"Onion",
				[]Effect{
					{name: "Restore Fatigue", eType: "+", power: 1},
					{name: "Water Breathing", eType: "+", power: 1},
					{name: "Detect Life", eType: "+", power: 1},
					{name: "Damage Health", eType: "-", power: 1},
				},
			}, {
				"Orange",
				[]Effect{
					{name: "Restore Fatigue", eType: "+", power: 1},
					{name: "Detect Life", eType: "+", power: 1},
					{name: "Burden", eType: "-", power: 1},
					{name: "Shield", eType: "+", power: 1},
				},
			}, {
				"Pear",
				[]Effect{
					{name: "Restore Fatigue", eType: "+", power: 1},
					{name: "Damage Speed", eType: "-", power: 1},
					{name: "Fortify Speed", eType: "+", power: 1},
					{name: "Damage Health", eType: "-", power: 1},
				},
			}, {
				"Peony Seeds",
				[]Effect{
					{name: "Restore Strength", eType: "+", power: 1},
					{name: "Damage Health", eType: "-", power: 1},
					{name: "Damage Speed", eType: "-", power: 1},
					{name: "Restore Fatigue", eType: "+", power: 1},
				},
			}, {
				"Potato",
				[]Effect{
					{name: "Restore Fatigue", eType: "+", power: 1},
					{name: "Shield", eType: "+", power: 1},
					{name: "Burden", eType: "-", power: 1},
					{name: "Frost Shield", eType: "+", power: 1},
				},
			}, {
				"Primrose Leaves",
				[]Effect{
					{name: "Restore Willpower", eType: "+", power: 1},
					{name: "Restore Personality", eType: "+", power: 1},
					{name: "Fortify Luck", eType: "+", power: 1},
					{name: "Damage Strength", eType: "-", power: 1},
				},
			}, {
				"Pumpkin",
				[]Effect{
					{name: "Restore Fatigue", eType: "+", power: 1},
					{name: "Damage Agility", eType: "-", power: 1},
					{name: "Damage Personality", eType: "-", power: 1},
					{name: "Detect Life", eType: "+", power: 1},
				},
			}, {
				"Purgeblood Salts",
				[]Effect{
					{name: "Restore Magicka", eType: "+", power: 1},
					{name: "Damage Health", eType: "-", power: 1},
					{name: "Fortify Magicka", eType: "+", power: 1},
					{name: "Dispel", eType: "+", power: 1},
				},
			}, {
				"Radish",
				[]Effect{
					{name: "Restore Fatigue", eType: "+", power: 1},
					{name: "Damage Endurance", eType: "-", power: 1},
					{name: "Chameleon", eType: "+", power: 1},
					{name: "Burden", eType: "-", power: 1},
				},
			}, {
				"Rat Meat",
				[]Effect{
					{name: "Damage Fatigue", eType: "-", power: 1},
					{name: "Detect Life", eType: "+", power: 1},
					{name: "Damage Magicka", eType: "-", power: 1},
					{name: "Silence", eType: "-", power: 1},
				},
			}, {
				"Redwort Flower",
				[]Effect{
					{name: "Resist Frost", eType: "+", power: 1},
					{name: "Cure Poison", eType: "+", power: 1},
					{name: "Damage Health", eType: "-", power: 1},
					{name: "Invisibility", eType: "+", power: 1},
				},
			}, {
				"Rice",
				[]Effect{
					{name: "Restore Fatigue", eType: "+", power: 1},
					{name: "Silence", eType: "-", power: 1},
					{name: "Shock Shield", eType: "+", power: 1},
					{name: "Damage Agility", eType: "-", power: 1},
				},
			}, {
				"Root Pulp",
				[]Effect{
					{name: "Cure Disease", eType: "+", power: 1},
					{name: "Damage Willpower", eType: "-", power: 1},
					{name: "Fortify Strength", eType: "+", power: 1},
					{name: "Damage Intelligence", eType: "-", power: 1},
				},
			}, {
				"Sacred Lotus Seeds",
				[]Effect{
					{name: "Resist Frost", eType: "+", power: 1},
					{name: "Damage Health", eType: "-", power: 1},
					{name: "Feather", eType: "+", power: 1},
					{name: "Dispel", eType: "+", power: 1},
				},
			}, {
				"Scales",
				[]Effect{
					{name: "Damage Willpower", eType: "-", power: 1},
					{name: "Water Breathing", eType: "+", power: 1},
					{name: "Damage Health", eType: "-", power: 1},
					{name: "Water Walking", eType: "+", power: 1},
				},
			}, {
				"Scamp Skin",
				[]Effect{
					{name: "Damage Magicka", eType: "-", power: 1},
					{name: "Resist Shock", eType: "+", power: 1},
					{name: "Reflect Damage", eType: "+", power: 1},
					{name: "Damage Health", eType: "-", power: 1},
				},
			}, {
				"Shepherd's Pie",
				[]Effect{
					{name: "Cure Disease", eType: "+", power: 1},
					{name: "Shield", eType: "+", power: 1},
					{name: "Fortify Agility", eType: "+", power: 1},
					{name: "Dispel", eType: "+", power: 1},
				},
			}, {
				"S'jirra's Famous Potato Bread",
				[]Effect{
					{name: "Detect Life", eType: "+", power: 1},
					{name: "Restore Health", eType: "+", power: 1},
					{name: "Damage Agility", eType: "-", power: 1},
					{name: "Damage Strength", eType: "-", power: 1},
				},
			}, {
				"Somnalius Frond",
				[]Effect{
					{name: "Restore Speed", eType: "+", power: 1},
					{name: "Damage Endurance", eType: "-", power: 1},
					{name: "Fortify Health", eType: "+", power: 1},
					{name: "Feather", eType: "+", power: 1},
				},
			}, {
				"Spiddal Stick",
				[]Effect{
					{name: "Damage Health", eType: "-", power: 1},
					{name: "Damage Magicka", eType: "-", power: 1},
					{name: "Fire Damage", eType: "-", power: 1},
					{name: "Restore Fatigue", eType: "+", power: 1},
				},
			}, {
				"St. Jahn's Wort Nectar",
				[]Effect{
					{name: "Resist Shock", eType: "+", power: 1},
					{name: "Damage Health", eType: "-", power: 1},
					{name: "Cure Poison", eType: "+", power: 1},
					{name: "Chameleon", eType: "+", power: 1},
				},
			}, {
				"Steel-Blue Entoloma Cap",
				[]Effect{
					{name: "Restore Magicka", eType: "+", power: 1},
					{name: "Fire Damage", eType: "-", power: 1},
					{name: "Resist Frost", eType: "+", power: 1},
					{name: "Burden", eType: "-", power: 1},
				},
			}, {
				"Stinkhorn Cap",
				[]Effect{
					{name: "Damage Health", eType: "-", power: 1},
					{name: "Restore Magicka", eType: "+", power: 1},
					{name: "Water Walking", eType: "+", power: 1},
					{name: "Invisibility", eType: "+", power: 1},
				},
			}, {
				"Strawberry",
				[]Effect{
					{name: "Restore Fatigue", eType: "+", power: 1},
					{name: "Cure Poison", eType: "+", power: 1},
					{name: "Damage Health", eType: "-", power: 1},
					{name: "Reflect Damage", eType: "+", power: 1},
				},
			}, {
				"Summer Bolete Cap",
				[]Effect{
					{name: "Restore Agility", eType: "+", power: 1},
					{name: "Shield", eType: "+", power: 1},
					{name: "Damage Personality", eType: "-", power: 1},
					{name: "Damage Endurance", eType: "-", power: 1},
				},
			}, {
				"Sweetcake",
				[]Effect{
					{name: "Restore Fatigue", eType: "+", power: 1},
					{name: "Feather", eType: "+", power: 1},
					{name: "Restore Health", eType: "+", power: 1},
					{name: "Burden", eType: "-", power: 1},
				},
			}, {
				"Sweetroll",
				[]Effect{
					{name: "Restore Fatigue", eType: "+", power: 1},
					{name: "Resist Disease", eType: "+", power: 1},
					{name: "Damage Personality", eType: "-", power: 1},
					{name: "Fortify Health", eType: "+", power: 1},
				},
			}, {
				"Taproot",
				[]Effect{
					{name: "Restore Luck", eType: "+", power: 1},
					{name: "Damage Endurance", eType: "-", power: 1},
					{name: "Resist Poison", eType: "+", power: 1},
					{name: "Shock Shield", eType: "+", power: 1},
				},
			}, {
				"Tiger Lily Nectar",
				[]Effect{
					{name: "Restore Endurance", eType: "+", power: 1},
					{name: "Damage Strength", eType: "-", power: 1},
					{name: "Water Walking", eType: "+", power: 1},
					{name: "Damage Willpower", eType: "-", power: 1},
				},
			}, {
				"Tinder Polypore Cap",
				[]Effect{
					{name: "Restore Willpower", eType: "+", power: 1},
					{name: "Resist Disease", eType: "+", power: 1},
					{name: "Invisibility", eType: "+", power: 1},
					{name: "Damage Magicka", eType: "-", power: 1},
				},
			}, {
				"Tobacco",
				[]Effect{
					{name: "Restore Fatigue", eType: "+", power: 1},
					{name: "Resist Paralysis", eType: "+", power: 1},
					{name: "Damage Magicka", eType: "-", power: 1},
					{name: "Dispel", eType: "+", power: 1},
				},
			}, {
				"Tomato",
				[]Effect{
					{name: "Restore Fatigue", eType: "+", power: 1},
					{name: "Detect Life", eType: "+", power: 1},
					{name: "Burden", eType: "-", power: 1},
					{name: "Shield", eType: "+", power: 1},
				},
			}, {
				"Troll Fat",
				[]Effect{
					{name: "Damage Agility", eType: "-", power: 1},
					{name: "Fortify Personality", eType: "+", power: 1},
					{name: "Damage Willpower", eType: "-", power: 1},
					{name: "Damage Health", eType: "-", power: 1},
				},
			}, {
				"Vampire Dust",
				[]Effect{
					{name: "Silence", eType: "-", power: 1},
					{name: "Resist Disease", eType: "+", power: 1},
					{name: "Frost Damage", eType: "-", power: 1},
					{name: "Invisibility", eType: "+", power: 1},
				},
			}, {
				"Venison",
				[]Effect{
					{name: "Restore Health", eType: "+", power: 1},
					{name: "Feather", eType: "+", power: 1},
					{name: "Damage Health", eType: "-", power: 1},
					{name: "Chameleon", eType: "+", power: 1},
				},
			}, {
				"Viper's Bugloss Leaves",
				[]Effect{
					{name: "Restore Magicka", eType: "+", power: 1},
					{name: "Damage Health", eType: "-", power: 1},
					{name: "Fortify Magicka", eType: "+", power: 1},
					{name: "Dispel", eType: "+", power: 1},
				},
			}, {
				"Water Hyacinth Nectar",
				[]Effect{
					{name: "Damage Luck", eType: "-", power: 1},
					{name: "Damage Fatigue", eType: "-", power: 1},
					{name: "Restore Magicka", eType: "+", power: 1},
					{name: "Fortify Magicka", eType: "+", power: 1},
				},
			}, {
				"Watermelon",
				[]Effect{
					{name: "Restore Fatigue", eType: "+", power: 1},
					{name: "Light", eType: "+", power: 1},
					{name: "Burden", eType: "-", power: 1},
					{name: "Damage Health", eType: "-", power: 1},
				},
			}, {
				"Wheat Grain",
				[]Effect{
					{name: "Restore Fatigue", eType: "+", power: 1},
					{name: "Damage Magicka", eType: "-", power: 1},
					{name: "Fortify Health", eType: "+", power: 1},
					{name: "Damage Personality", eType: "-", power: 1},
				},
			}, {
				"White Seed Pod",
				[]Effect{
					{name: "Restore Strength", eType: "+", power: 1},
					{name: "Water Breathing", eType: "+", power: 1},
					{name: "Silence", eType: "-", power: 1},
					{name: "Light", eType: "+", power: 1},
				},
			}, {
				"Wisp Stalk Caps",
				[]Effect{
					{name: "Damage Health", eType: "-", power: 1},
					{name: "Damage Willpower", eType: "-", power: 1},
					{name: "Damage Intelligence", eType: "-", power: 1},
					{name: "Fortify Speed", eType: "+", power: 1},
				},
			}, {
				"Wormwood Leaves",
				[]Effect{
					{name: "Fortify Fatigue", eType: "+", power: 1},
					{name: "Invisibility", eType: "+", power: 1},
					{name: "Damage Health", eType: "-", power: 1},
					{name: "Damage Magicka", eType: "-", power: 1},
				},
			}, {
				"Ashes of Hindaril",
				[]Effect{
					{name: "Silence", eType: "-", power: 1},
					{name: "Resist Disease", eType: "+", power: 1},
					{name: "Frost Damage", eType: "-", power: 1},
					{name: "Invisibility", eType: "+", power: 1},
				},
			},
		},
	}
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

var IngredientsDatabase = initStorage()
