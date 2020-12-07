package main

import "strings"

type bag struct {
	name     string
	children []*bag
}

type bagInput struct {
	nameStr     string
	childrenStr string
}

func getParentBags(inputValues []string) []*bag {
	var parentBags []*bag

	// Create list of bagInputs so we have access to ALL parent names:
	var bagInputs []bagInput
	for _, line := range inputValues {
		temp := strings.Split(line, " contain ")

		bagInputs = append(bagInputs, bagInput{
			nameStr:     temp[0],
			childrenStr: temp[1],
		})
		parentBags = append(parentBags, &bag{
			name:     temp[0],
			children: make([]*bag, 0),
		})
	}

	// Actually append all child bags to bag.children
	for _, bagInput := range bagInputs {
		bag := findBagByName(parentBags, bagInput.nameStr)
		for _, childStr := range strings.Split(bagInput.childrenStr, ", ") {
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
