package main

import (
	"container/list"
	"math"
	"strconv"
	"strings"

	"github.com/jraams/aoc-2020/helpers"
	"github.com/thoas/go-funk"
)

var gameCounter = 1

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

// Once the game ends, you can calculate the winning player's score. The bottom card in their deck is worth the value of
// the card multiplied by 1, the second-from-the-bottom card is worth the value of the card multiplied by 2, and so on.
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
	return copyDeckUpTo(deck, math.MaxInt32)
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

func playSpaceCards(gameNr int, recursive bool, rounds map[int]*map[int][]string, d1 *list.List, d2 *list.List) (int, *list.List) {
	if !funk.Contains(rounds, gameNr) {
		rounds[gameNr] = &map[int][]string{}
	}

	roundNr := 1
	// he game consists of a series of rounds
	for d1.Len() > 0 && d2.Len() > 0 {

		if recursive {
			// Before either player deals a card, if there was a previous round in this game that had exactly the same cards in
			// the same order in the same players' decks, the game instantly ends in a win for player 1. Previous rounds from
			// other games are not considered. (This prevents infinite games of Recursive Combat, which everyone agrees is a bad idea.)
			d1h := getDeckHash(d1)
			d2h := getDeckHash(d2)

			// Previous rounds from other games are not considered.
			roundMap := *rounds[gameNr]
			if funk.Contains(roundMap[1], d1h) || funk.Contains(roundMap[2], d2h) {
				return 1, d1
			}
		}

		// both players draw their top card,
		p1e := d1.Front()
		p2e := d2.Front()
		p1i := p1e.Value.(int)
		p2i := p2e.Value.(int)

		roundMap := *rounds[gameNr]
		roundMap[1] = append(roundMap[1], getDeckHash(d1))
		roundMap[2] = append(roundMap[2], getDeckHash(d2))

		d1.Remove(p1e)
		d2.Remove(p2e)

		// the player with the higher-valued card wins the round.
		p1won := p1i > p2i

		if recursive {
			// If both players have at least as many cards remaining in their deck as the value of the card they just drew,
			// the winner of the round is determined by playing a new game of Recursive Combat
			if p1i <= d1.Len() && p2i <= d2.Len() {
				d1copy := copyDeckUpTo(d1, p1i)
				d2copy := copyDeckUpTo(d2, p2i)
				gameCounter++
				winner, _ := playSpaceCards(gameCounter, true, rounds, d1copy, d2copy)
				p1won = winner == 1
			}
		}

		// The winner keeps both cards, placing them on the bottom of their own deck so that the winner's card is above the
		// other card. If this causes a player to have all of the cards, they win, and the game ends.
		if p1won {
			d1.PushBack(p1i)
			d1.PushBack(p2i)
		} else {
			d2.PushBack(p2i)
			d2.PushBack(p1i)
		}

		roundNr++
	}

	// Decide winner
	if d1.Len() > 0 {
		return 1, d1
	}

	return 2, d2
}
