package main

import "errors"

type position struct {
	y              int
	x              int
	isFloor        bool
	isOccupied     bool
	tempIsOccupied bool
	adjacentSeats  []position
}

func (p *position) getChar() string {
	if p.isFloor {
		return "."
	}
	if p.isOccupied {
		return "#"
	}
	return "L"
}

func makePosition(s string, y int, x int) position {
	p := position{}
	p.y = y
	p.x = x
	p.isFloor = s == "."
	p.isOccupied = s == "#"
	p.adjacentSeats = []position{}
	return p
}

// If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
func (p *position) checkRule1(g grid) bool {
	if p.isOccupied {
		return false
	}

	for _, adjacentSeat := range p.adjacentSeats {
		if g[adjacentSeat.y][adjacentSeat.x].isOccupied {
			return false
		}
	}
	return true
}

// If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
func (p *position) checkRule2(g grid, seatTolerance int) bool {
	if !p.isOccupied {
		return false
	}

	occupiedNeighbourCount := 0
	for _, adjacentSeat := range p.adjacentSeats {
		if g[adjacentSeat.y][adjacentSeat.x].isOccupied {
			occupiedNeighbourCount++
		}
	}
	return occupiedNeighbourCount >= seatTolerance
}

func getNeighBourSeats(g grid, y int, x int, onlyCheckAdjacentSeats bool) []position {
	var seats []position
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dy == 0 && dx == 0 {
				continue
			}
			if position, err := getSeatInDir(g, y, x, dy, dx, onlyCheckAdjacentSeats); err == nil {
				seats = append(seats, *position)
			}
		}
	}
	return seats
}

func getSeatInDir(g grid, y int, x int, dy int, dx int, quitAfterFirst bool) (*position, error) {
	newY := y + dy
	newX := x + dx

	if newY < 0 || newX < 0 || newY >= len(g) || newX >= len(g[y]) {
		return nil, errors.New("")
	}

	position := g[newY][newX]
	if !position.isFloor {
		return &position, nil
	}

	if quitAfterFirst {
		return nil, errors.New("")
	}

	return getSeatInDir(g, newY, newX, dy, dx, quitAfterFirst)
}
