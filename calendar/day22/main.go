package main

import (
	"container/list"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/jraams/aoc-2020/helpers"
	"github.com/thoas/go-funk"
)

const debug = true

func main() {
	inputPath, _ := filepath.Abs("input")
	lines := helpers.GetInputValues(inputPath)
	d1, d2 := loadDecks(lines)

	// // Part a
	// a := solveA(d1, d2)
	// fmt.Printf("Solution day 22 part a: %d\n", a)

	// Part b
	b := solveB(d1, d2)
	fmt.Printf("Solution day 22 part b: %d\n", b)
}

func loadDecks(lines []string) (*list.List, *list.List) {
	d1, d2 := list.New(), list.New()

	loadingDeck1 := true
	for _, line := range lines {
		if strings.Contains(line, "Player") {
			continue
		}
		if len(line) == 0 {
			loadingDeck1 = false
			continue
		}
		intVal := helpers.MustAtoi(line)
		if loadingDeck1 {
			d1.PushBack(intVal)
		} else {
			d2.PushBack(intVal)
		}
	}

	return d1, d2
}

func solveA(d1 *list.List, d2 *list.List) int {
	gameRoundDeckMap := map[int]*map[int][]string{}
	_, winningDeck := playSpaceCards(1, false, gameRoundDeckMap, d1, d2)
	score := calculateScore(winningDeck)
	return score
}

func solveB(d1 *list.List, d2 *list.List) int {
	gameRoundDeckMap := map[int]*map[int][]string{}
	_, winningDeck := playSpaceCards(1, true, gameRoundDeckMap, d1, d2)
	score := calculateScore(winningDeck)
	return score
}

func calculateScore(deck *list.List) int {
	score := 0
	for mult := deck.Len(); mult > 0; {
		val := deck.Front()
		score += mult * val.Value.(int)
		deck.Remove(val)
		mult--
	}
	return score
}

func getDeckHash(deck *list.List) string {
	hash := ""
	for item := deck.Front(); item != nil; {
		hash += strconv.Itoa(item.Value.(int))
		hash += "-"
		item = item.Next()
	}
	return hash
}

func copyDeckUpTo(deck *list.List, maxItems int) *list.List {
	newDeck := new(list.List)
	count := 0
	for item := deck.Front(); item != nil && count < maxItems; {
		newDeck.PushBack(item.Value)
		item = item.Next()
		count++
	}
	return newDeck
}

func printDeck(player int, deck *list.List) {
	fmt.Printf("Player %d's deck: ", player)
	for item := deck.Front(); item != nil; {
		intVal := item.Value.(int)
		fmt.Printf("%d, ", intVal)
		item = item.Next()
	}
	fmt.Println()
}

var gameCounter = 1

func playSpaceCards(gameNr int, recursive bool, rounds map[int]*map[int][]string, d1 *list.List, d2 *list.List) (int, *list.List) {
	if debug {
		fmt.Printf("\n=== Game %d ===\n", gameNr)
	}

	if !funk.Contains(rounds, gameNr) {
		rounds[gameNr] = &map[int][]string{}
	}

	roundNr := 1
	// Play game
	for d1.Len() > 0 && d2.Len() > 0 {

		if recursive {
			// Before either player deals a card, if there was a previous round in this game that had exactly the same cards in
			// the same order in the same players' decks, the game instantly ends in a win for player 1. Previous rounds from
			// other games are not considered. (This prevents infinite games of Recursive Combat, which everyone agrees is a bad idea.)
			d1h := getDeckHash(d1)
			d2h := getDeckHash(d2)

			roundMap := *rounds[gameNr]
			if funk.Contains(roundMap[1], d1h) {
				if debug {
					fmt.Printf("player 1 deck %s has already been used in this game\n", d1h)
				}
				return 1, d1
			}
			if funk.Contains(roundMap[2], d2h) {
				if debug {
					fmt.Printf("player 2 deck %s has already been used in this game\n", d2h)
				}
				return 1, d1
			}
		}

		if debug {
			fmt.Printf("-- Round %d (Game %d)\n", roundNr, gameNr)
			printDeck(1, d1)
			printDeck(2, d2)
		}

		p1e := d1.Front()
		p2e := d2.Front()

		p1i := p1e.Value.(int)
		p2i := p2e.Value.(int)

		roundMap := *rounds[gameNr]
		roundMap[1] = append(roundMap[1], getDeckHash(d1))
		roundMap[2] = append(roundMap[2], getDeckHash(d2))

		d1.Remove(p1e)
		d2.Remove(p2e)

		if debug {
			fmt.Printf("Player 1 plays: %d\n", p1i)
			fmt.Printf("Player 2 plays: %d\n", p2i)
		}

		p1won := p1i > p2i

		if recursive {
			if p1i <= d1.Len() && p2i <= d2.Len() {
				if debug {
					fmt.Println("Playing a sub game to determine the winner...")
				}
				d1copy := copyDeckUpTo(d1, p1i)
				d2copy := copyDeckUpTo(d2, p2i)
				gameCounter++
				winner, _ := playSpaceCards(gameCounter, true, rounds, d1copy, d2copy)
				if debug {
					fmt.Printf("The winner of game %d is player %d\n", gameNr, winner)
				}
				p1won = winner == 1
			}
		}

		if p1won {
			d1.PushBack(p1i)
			d1.PushBack(p2i)
			if debug {
				fmt.Printf("Player 1 wins round %d of game %d\n", roundNr, gameNr)
			}
		} else {
			d2.PushBack(p2i)
			d2.PushBack(p1i)
			if debug {
				fmt.Printf("Player 2 wins round %d of game %d\n", roundNr, gameNr)
			}
		}

		roundNr++
		if debug {
			fmt.Println()
			fmt.Println()
		}
	}

	// Decide winner
	if d1.Len() == 0 {
		return 2, d2
	} else {
		return 1, d1
	}
}
