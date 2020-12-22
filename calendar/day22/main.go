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
	_, winningDeck := playSpaceCards(0, false, d1, d2)
	score := calculateScore(winningDeck)
	return score
}

func solveB(d1 *list.List, d2 *list.List) int {
	_, winningDeck := playSpaceCards(0, true, d1, d2)
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

func copyDeck(deck *list.List) *list.List {
	newDeck := new(list.List)
	for item := deck.Front(); item != nil; {
		newDeck.PushBack(item.Value)
		item = item.Next()
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

func playSpaceCards(gameNr int, recursive bool, d1 *list.List, d2 *list.List) (int, *list.List) {
	gameNr++
	fmt.Printf("\n=== Game %d ===\n", gameNr)

	rounds := map[int][]string{}

	roundNr := 1
	// Play game
	for d1.Len() > 0 && d2.Len() > 0 {

		if recursive {
			// Before either player deals a card, if there was a previous round in this game that had exactly the same cards in
			// the same order in the same players' decks, the game instantly ends in a win for player 1. Previous rounds from
			// other games are not considered. (This prevents infinite games of Recursive Combat, which everyone agrees is a bad idea.)
			d1h := getDeckHash(d1)
			d2h := getDeckHash(d2)

			if funk.Contains(rounds[1], d1h) {
				fmt.Printf("player 1 deck %s has already been used in this game\n", d1h)
				return 1, d1
			}
			if funk.Contains(rounds[2], d2h) {
				fmt.Printf("player 2 deck %s has already been used in this game\n", d2h)
				return 1, d1
			}
		}

		fmt.Printf("-- Round %d (Game %d)\n", roundNr, gameNr)
		printDeck(1, d1)
		printDeck(2, d2)

		p1e := d1.Front()
		p2e := d2.Front()

		p1i := p1e.Value.(int)
		p2i := p2e.Value.(int)

		rounds[1] = append(rounds[1], getDeckHash(d1))
		rounds[2] = append(rounds[2], getDeckHash(d2))

		d1.Remove(p1e)
		d2.Remove(p2e)

		fmt.Printf("Player 1 plays: %d\n", p1i)
		fmt.Printf("Player 2 plays: %d\n", p2i)

		p1won := p1i > p2i

		if recursive {
			if p1i <= d1.Len() && p2i <= d2.Len() {
				fmt.Println("Playing a sub game to determine the winner...")
				d1copy := copyDeck(d1)
				d2copy := copyDeck(d2)
				winner, _ := playSpaceCards(gameNr, true, d1copy, d2copy)
				fmt.Printf("The winner of game %d is player %d\n", gameNr, winner)
				p1won = winner == 1
			}
		}

		if p1won {
			d1.PushBack(p1i)
			d1.PushBack(p2i)
			fmt.Printf("Player 1 wins round %d of game %d\n", roundNr, gameNr)
		} else {
			d2.PushBack(p2i)
			d2.PushBack(p1i)
			fmt.Printf("Player 2 wins round %d of game %d\n", roundNr, gameNr)
		}

		roundNr++
		fmt.Println()
		fmt.Println()
	}

	// Decide winner
	if d1.Len() == 0 {
		return 2, d2
	} else {
		return 1, d1
	}
}
