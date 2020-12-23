package main

import (
	"github.com/jraams/aoc-2020/helpers"
	"github.com/thoas/go-funk"
)

type valueCupMap map[int]*cup

type cup struct {
	value int
	next  *cup
}

func load(labels string, extendCups bool) (valueCupMap, *cup) {
	// Your labeling is still correct for the first few cups;
	var intValues []int
	for _, b := range []byte(labels) {
		intValues = append(intValues, helpers.MustAtoi(string(b)))
	}
	// after that, the remaining cups are just numbered in an increasing fashion starting from the number after the
	// highest number in your list and proceeding one by one until one million is reached.
	if extendCups {
		for i := len(labels) + 1; i <= 1000000; i++ {
			intValues = append(intValues, i)
		}
	}

	vc := valueCupMap{}
	var prev *cup

	for _, i := range intValues {
		c := &cup{
			value: i,
		}
		vc[i] = c
		if prev != nil {
			prev.next = c
		}
		prev = c
	}
	prev.next = vc[intValues[0]]

	// Before the crab starts, it will designate the first cup in your list as the current cup.
	return vc, prev
}

func playCrabGame(vc valueCupMap, currentCup *cup, moves int) {
	// The crab is then going to do N moves.
	for i := 1; i <= moves; i++ {
		// The crab selects a new current cup: the cup which is immediately clockwise of the current cup.
		currentCup = currentCup.next

		// The crab picks up the three cups that are immediately clockwise of the current cup.
		pickedUpCups := []*cup{
			currentCup.next,
			currentCup.next.next,
			currentCup.next.next.next,
		}
		// They are removed from the circle;
		currentCup.next = currentCup.next.next.next.next

		// The crab selects a destination cup
		destinationVal := currentCup.value
		searching_dest := true
		for searching_dest {
			destinationVal = destinationVal - 1
			if destinationVal <= 0 {
				destinationVal = len(vc)
			}
			if !funk.Contains(pickedUpCups, vc[destinationVal]) {
				searching_dest = false
			}
		}
		destination := vc[destinationVal]

		// The crab places the cups it just picked up so that they are immediately clockwise of the destination cup.
		pickedUpCups[2].next = destination.next
		destination.next = pickedUpCups[0]
	}
}
