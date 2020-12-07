package main

import (
	"strconv"
	"strings"
)

type bag struct {
	name     string
	children []childBag
}

type childBag struct {
	amount int
	bag    *bag
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
			children: make([]childBag, 0),
		})
	}

	// Actually append all child bags to bag.children
	for _, bagInput := range bagInputs {
		bag := findBagByName(parentBags, bagInput.nameStr)
		for _, childStr := range strings.Split(bagInput.childrenStr, ", ") {
			amount, _ := strconv.Atoi(strings.Split(childStr, " ")[0])
			bagName := strings.Join(strings.Split(childStr, " ")[1:], " ")
			if bagName == "other" {
				continue
			}
			cb := findBagByName(parentBags, bagName)

			bag.children = append(bag.children, childBag{
				amount: amount,
				bag:    cb,
			})
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
		if bagHasChildWithName(*childBag.bag, name) {
			return true
		}
	}
	return false
}

func getTotalBagSum(parentBag bag) int {
	totalAmount := 1
	for _, childBag := range parentBag.children {
		totalAmount += getTotalBagSum(*childBag.bag) * childBag.amount
	}
	return totalAmount
}
