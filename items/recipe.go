package items

import (
	"encoding/json"
	"errors"
	"fmt"
)

var allRecipes map[int]Recipe

func AllRecipes() *map[int]Recipe {
	if allRecipes == nil {
		if err := json.Unmarshal(allRecipes_json, &allRecipes); err != nil {
			panic(err)
		}
	}
	return &allRecipes
}

type Recipe struct {
	Id          int          `json:"id"`
	Ingredients []Ingredient `json:"ingredients"`
	Results     []Ingredient `json:"results"`
}

type Ingredient struct {
	Item   int `json:"item"`
	Amount int `json:"amount"`
}

func validateRecipes() error {
	recipes := AllRecipes()
	fmt.Println(fmt.Sprintf("validating %d recipes", len(*recipes)))
	for _, recipe := range *recipes {
		if err := validateRecipe(&recipe); err != nil {
			return err
		}
	}
	return nil
}

func validateRecipe(recipe *Recipe) error {
	allItems := AllItems()
	for _, i := range recipe.Ingredients {
		if _, ok := (*allItems)[i.Item]; !ok {
			return errors.New(fmt.Sprintf("Ingredient with id %d not found", i.Item))
		}
	}
	return nil
}
