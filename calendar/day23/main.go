package main

import (
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/jraams/aoc-2020/helpers"
	"github.com/thoas/go-funk"
)

func main() {
	inputPath, _ := filepath.Abs("input")
	lines := helpers.GetInputValues(inputPath)
	cups := load(lines[0])

	a := solveA(cups)
	fmt.Printf("Solution day 23 part a: %s\n", a)
}

func load(labels string) cups {
	c := cups{
		values: []int{},
	}

	for _, b := range []byte(labels) {
		c.values = append(c.values, helpers.MustAtoi(string(b)))
	}
	return c
}

type cups struct {
	values []int
}

func (c *cups) index(value int) int {
	for i := 0; i < len(c.values); i++ {
		if c.values[i] == value {
			return i
		}
	}
	return -1
}

func (c *cups) remove(value int) {
	valIx := c.index(value)
	temp := append([]int{}, c.values[:valIx]...)
	c.values = append(temp, c.values[valIx+1:]...)
}

func playCrapGame(c *cups, moves int) {
	current := -1
	cups_len := len(c.values)

	// The crab is then going to do N moves.
	for i := 1; i <= moves; i++ {
		// Before the crab starts, it will designate the first cup in your list as the current cup.
		current_ix := 0
		if current != -1 {
			current_ix = (c.index(current) + 1) % cups_len
		}
		// The crab selects a new current cup: the cup which is immediately clockwise of the current cup.
		current = c.values[current_ix]

		// 2. Pick up cups
		picked_up := []int{}
		for j := 0; j < 3; j++ {
			picked_up = append(picked_up, c.values[(current_ix+j+1)%cups_len])
		}
		for _, p := range picked_up {
			c.remove(p)
		}

		// The crab selects a destination cup
		searching_dest := true
		destination := current
		if current == -1 {
			destination = cups_len
		}
		for searching_dest {
			// the cup with a label equal to the current cup's label minus one.
			destination = destination - 1
			// If at any point in this process the value goes below the lowest value on any cup's label,
			// it wraps around to the highest value on any cup's label instead.
			if destination <= 0 {
				destination = cups_len
			}
			if !funk.Contains(picked_up, destination) {
				searching_dest = false
			}
		}
		destination_ix := c.index(destination)

		// The crab places the cups it just picked up so that they are immediately clockwise of the destination cup.
		temp := append([]int{}, c.values[:destination_ix+1]...)
		temp = append(temp, picked_up...)
		temp = append(temp, c.values[destination_ix+1:]...)
		c.values = temp
	}
}

func solveA(c cups) string {
	playCrapGame(&c, 100)

	ix_1 := c.index(1)
	result := append([]int{}, c.values[ix_1+1:]...)
	result = append(result, c.values[:ix_1]...)

	str := ""
	for _, val := range result {
		str += strconv.Itoa(val)
	}
	return str
}
