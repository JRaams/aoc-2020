package main

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	"github.com/jraams/aoc-2020/helpers"
	"github.com/thoas/go-funk"
)

func main() {
	inputPath, _ := filepath.Abs("input")
	lines := helpers.GetInputValues(inputPath)
	foods := loadFoods(lines)
	allIngMap := getAllIngMap(foods)

	// Part a
	a := solveA(foods, allIngMap)
	fmt.Printf("Solution day 21 part a: %d\n", a)

	// Part b
	b := solveB(allIngMap)
	fmt.Printf("Solution day 21 part b: %s\n", b)
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

func getAllIngMap(allFoods []food) map[string][]string {
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
		for all, ings := range allIngMap {
			diffSuspectsAndKnown := helpers.StringArrDifference(ings, dangerousIngs)
			if len(diffSuspectsAndKnown) == 1 {
				allIngMap[all] = diffSuspectsAndKnown
				dangerousIngs = append(dangerousIngs, ings...)
			}
		}
	}

	return allIngMap
}

func solveA(allFoods []food, allIngMap map[string][]string) int {
	var dangerousIngs []string
	for _, ings := range allIngMap {
		dangerousIngs = append(dangerousIngs, ings...)
	}

	a := 0
	for _, food := range allFoods {
		d := helpers.StringArrDifference(food.ingredients, dangerousIngs)
		a += len(d)
	}
	return a
}

func solveB(allIngMap map[string][]string) string {
	keys := make([]string, 0)
	for k := range allIngMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	b := ""
	for _, k := range keys {
		b += allIngMap[k][0] + ","
	}
	b = b[:len(b)-1]

	return b
}
