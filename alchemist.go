package main

import (
	"fmt"
)

func main() {
	bread := Ingredient{
		"Bread",
		[]Effect{
			{"Stamina+", "Restores stamina", 10},
			{"Hunger-", "Decreases hunger", -20},
		},
	}

	salmon := Ingredient{
		"Salmon",
		[]Effect{
			{"Stamina+", "Restores stamina", 30},
			{"Hunger-", "Decreases hunger", -50},
		},
	}

	ingredients := [2]Ingredient{
		salmon,
		bread,
	}

	mortar := new(Mortar)

	for _, ingredient := range ingredients {
		err := mortar.AddIngredient(ingredient)
		if err != nil {
			mortar.Clear()
			fmt.Println(err)
			return
		}
	}

	potion, err := mortar.Pestle()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Congratulations! You've created a potion!")
	fmt.Println(potion.Description())
}
