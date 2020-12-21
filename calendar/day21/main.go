package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/jraams/aoc-2020/helpers"
	"github.com/thoas/go-funk"
)

func main() {
	inputPath, _ := filepath.Abs("input")
	lines := helpers.GetInputValues(inputPath)
	foods := loadFoods(lines)
	dangerousIngs := getDangerousIngs(foods)

	// Part 1
	a := solveA(foods, dangerousIngs)
	fmt.Printf("Solution day 21 part a: %d\n", a)
}

type food struct {
	ingredients []string
	allergens   []string
}

func loadFoods(lines []string) []food {
	var foods []food

	junkReplacer := strings.NewReplacer("(", "", ")", "", ",", "")
	for id := 0; id < len(lines); id++ {
		line := junkReplacer.Replace(lines[id])
		words := strings.Split(line, " ")
		food := food{
			ingredients: []string{},
			allergens:   []string{},
		}

		addingIngredients := true
		for _, word := range words {
			if word == "contains" {
				addingIngredients = false
				continue
			}

			if addingIngredients {
				food.ingredients = append(food.ingredients, word)
			} else {
				food.allergens = append(food.allergens, word)
			}
		}

		foods = append(foods, food)
	}

	return foods
}

func getDangerousIngs(allFoods []food) []string {
	allIngMap := map[string][]string{}

	for _, food := range allFoods {
		for _, a := range food.allergens {
			if !funk.Contains(allIngMap, a) {
				allIngMap[a] = food.ingredients
			} else {
				newIngs := []string{}
				for _, ing := range food.ingredients {
					if funk.Contains(allIngMap[a], ing) {
						newIngs = append(newIngs, ing)
					}
				}
				allIngMap[a] = newIngs
			}
		}
	}

	dangerousIngs := []string{}
	for len(dangerousIngs) < len(allIngMap) {
		for _, suspects := range allIngMap {
			diffSuspectsAndKnown := helpers.StringArrDifference(suspects, dangerousIngs)
			if len(diffSuspectsAndKnown) == 1 {
				suspects = diffSuspectsAndKnown
				dangerousIngs = append(dangerousIngs, suspects...)
			}
		}
	}
	return dangerousIngs
}

func solveA(allFoods []food, dangerousIngs []string) int {
	a := 0
	for _, food := range allFoods {
		d := helpers.StringArrDifference(food.ingredients, dangerousIngs)
		a += len(d)
	}
	return a
}
