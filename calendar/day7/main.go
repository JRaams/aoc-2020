package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/jraams/aoc-2020/helpers"
)

type bag struct {
	name     string
	children []*bag
}

type bagInput struct {
	nameStr     string
	childrenStr string
}

func main() {
	// Load input from file
	inputPath, _ := filepath.Abs("input")
	inputValues := helpers.GetInputValues(inputPath)
	parentBags := getParentBags(inputValues)

	// fmt.Println(parentBags)
	// Part 1
	amountOfBagsWithShinyGoldChild := getAmountBagsWithName(parentBags, "shiny gold")
	amountOfBagsWithShinyGoldChild--
	fmt.Printf("Solution part 1: %d bag colors eventually contain at least one shiny gold bad", amountOfBagsWithShinyGoldChild)
	fmt.Println()
}

func getParentBags(inputValues []string) []*bag {
	var parentBags []*bag

	// Create list of bagInputs so we have access to ALL parent names:
	var bagInputs []bagInput
	for _, line := range inputValues {
		temp := strings.Split(line, " contain ")
		bagName := strings.Replace(temp[0], " bags", "", 1)

		bagInputs = append(bagInputs, bagInput{
			nameStr:     bagName,
			childrenStr: temp[1],
		})
		parentBags = append(parentBags, &bag{
			name:     bagName,
			children: make([]*bag, 0),
		})
	}

	// Actually append all child bags to bag.children
	for _, bagInput := range bagInputs {
		bag := findBagByName(parentBags, bagInput.nameStr)
		childrenStr := bagInput.childrenStr[:len(bagInput.childrenStr)-1]
		childrenStr = strings.ReplaceAll(strings.ReplaceAll(childrenStr, " bags", ""), " bag", "")
		for _, childStr := range strings.Split(childrenStr, ", ") {
			bagName := strings.Join(strings.Split(childStr, " ")[1:], " ")
			if bagName == "other" {
				continue
			}
			childBag := findBagByName(parentBags, bagName)

			bag.children = append(bag.children, childBag)
		}
	}

	return parentBags
}

func findBagByName(bags []*bag, name string) *bag {
	for _, bag := range bags {
		if bag.name == name {
			return bag
		}
	}
	return nil
}

func getAmountBagsWithName(parentBags []*bag, name string) int {
	amount := 0

	for _, parentBag := range parentBags {
		if bagHasChildWithName(*parentBag, name) {
			amount++
		}
	}

	return amount
}

func bagHasChildWithName(bag bag, name string) bool {
	if bag.name == name {
		return true
	}

	for _, childBag := range bag.children {
		if bagHasChildWithName(*childBag, name) {
			return true
		}
	}

	return false
}
