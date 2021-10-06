package main

import (
	"bufio"
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"log"
	"os"
	"strings"
)

func main() {
	pixelgl.Run(launch)
}

func launch() {
	cfg := pixelgl.WindowConfig{
		Title:  "Alchemist",
		Bounds: pixel.R(0, 0, 1024, 768),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	ingredientRepository := initStorage()

	runGame(win, ingredientRepository)
}

func runGame(window *pixelgl.Window, storage IngredientFinder) {
	fmt.Println("Input comma-separated ingredients names(bread, salmon)")

	var usedIngredients []Ingredient
	reader := bufio.NewReader(os.Stdin)
	for {
		ingredientsNames, _ := reader.ReadString('\n')
		if len(ingredientsNames) > 1 {
			var ingredientNames []string
			for _, ingredientName := range strings.Split(ingredientsNames, ",") {
				ingredientName = strings.Trim(ingredientName, " ")
				ingredientName = strings.Trim(ingredientName, "\n")
				ingredientNames = append(ingredientNames, ingredientName)
			}
			ingredients, err := storage.FindByNames(ingredientNames)
			if err != nil {
				log.Fatal(err)
			}
			usedIngredients = ingredients
			break
		}
	}

	mortar := new(Mortar)

	for _, ingredient := range usedIngredients {
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

	for !window.Closed() {
		window.Update()
	}
}
