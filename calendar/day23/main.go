package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"time"

	"github.com/jraams/aoc-2020/helpers"
	"github.com/thoas/go-funk"
)

func main() {
	inputPath, _ := filepath.Abs("input")
	lines := helpers.GetInputValues(inputPath)

	// Part a
	vca, firstValue := load(lines[0], false)
	a := solveA(vca, firstValue)
	fmt.Printf("Solution day 23 part a: %s\n", a)

	// // Part b
	// cupsB := load(lines[0], true)
	// b := solveB(cupsB)
	// fmt.Printf("Solution day 23 part b: %d\n", b)
}

type valueCupMap map[int]*cup

type cup struct {
	value int
	next  *cup
}

func load(labels string, extendCups bool) (valueCupMap, *cup) {
	var intValues []int
	for _, b := range []byte(labels) {
		intValues = append(intValues, helpers.MustAtoi(string(b)))
	}
	if extendCups {
		for i := len(labels); i <= 1000000; i++ {
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
	return vc, prev
}

func playCrabGame(vc valueCupMap, currentCup *cup, moves int) {
	// Before the crab starts, it will designate the first cup in your list as the current cup.
	start := time.Now()
	fmt.Println("Started at", start)

	// The crab is then going to do N moves.
	for i := 1; i <= moves; i++ {
		if i%1000000 == 0 {
			fmt.Println(i / 1000000)
			fmt.Println(time.Since(start))
		}
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

	fmt.Println("Done after ", time.Since(start))
}

func solveA(vc valueCupMap, currentCup *cup) string {
	playCrabGame(vc, currentCup, 100)
	str := ""
	for cup := vc[1].next; cup.value != 1; cup = cup.next {
		str += strconv.Itoa(cup.value)
	}
	return str
}

// func solveB(c cups) int {
// 	playCrabGame(&c, 10000)

// 	ix_1 := c.index(1)
// 	b := c.values[ix_1+1] * c.values[ix_1+2]
// 	return b
// }
