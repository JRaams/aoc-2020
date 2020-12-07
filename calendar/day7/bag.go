package main

import (
	"strconv"
	"strings"
)

type bag struct {
	name        string
	children    []childBag
	childrenStr string
}

type childBag struct {
	amount int
	bag    *bag
}

func getParentBags(inputValues []string) []*bag {
	var parentBags []*bag

	// Setup parent bags
	for _, line := range inputValues { // vibrant bronze bags contain 3 dim olive bags.
		temp := strings.Split(line, " contain ") // ['vibrant bronze bags', '3 dim olive bags.']

		parentBags = append(parentBags, &bag{
			name:        temp[0], // 'vibrant bronze bags'
			children:    make([]childBag, 0),
			childrenStr: temp[1], // '3 dim olive bags.'
		})
	}

	// Add child bags (now that all parent bags have been set up)
	for _, parentBag := range parentBags {
		for _, childStr := range strings.Split(parentBag.childrenStr, ", ") {
			amount, _ := strconv.Atoi(strings.Split(childStr, " ")[0])
			bagName := strings.Join(strings.Split(childStr, " ")[1:], " ")
			if bagName == "other" {
				continue
			}
			bag := findBagByName(parentBags, bagName)
			parentBag.children = append(parentBag.children, childBag{
				amount: amount,
				bag:    bag,
			})
		}
		parentBag.childrenStr = ""
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
