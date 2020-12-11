package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

// Seat ...
type Seat struct {
	x        int
	y        int
	occupied bool
}

func (s *Seat) id() string {
	return fmt.Sprintf("(%d, %d)", s.x, s.y)
}

func solution1(seats map[string]*Seat, adjacentMap, _ map[string][]*Seat) (answer int) {
	for {
		changes := make([]string, 0)
		for id, adjacents := range adjacentMap {
			occupiedAdjacents := 0
			for _, a := range adjacents {
				if a.occupied {
					occupiedAdjacents++
				}
			}
			if seats[id].occupied && occupiedAdjacents >= 4 {
				changes = append(changes, id)
			}
			if !seats[id].occupied && occupiedAdjacents == 0 {
				changes = append(changes, id)
			}
		}

		for _, id := range changes {
			seats[id].occupied = !seats[id].occupied
		}

		if len(changes) == 0 {
			for _, seat := range seats {
				if seat.occupied {
					answer++
				}
			}
			break
		}
	}

	return
}

func solution2(seats map[string]*Seat, _, visibleMap map[string][]*Seat) (answer int) {
	for {
		changes := make([]string, 0)
		for id, visibles := range visibleMap {
			occupiedVisibles := 0
			for _, a := range visibles {
				if a.occupied {
					occupiedVisibles++
				}
			}
			if seats[id].occupied && occupiedVisibles >= 5 {
				changes = append(changes, id)
			}
			if !seats[id].occupied && occupiedVisibles == 0 {
				changes = append(changes, id)
			}
		}

		for _, id := range changes {
			seats[id].occupied = !seats[id].occupied
		}

		if len(changes) == 0 {
			for _, seat := range seats {
				if seat.occupied {
					answer++
				}
			}
			break
		}
	}

	return
}

func parseSeats(input string) (map[string]*Seat, map[string][]*Seat, map[string][]*Seat) {
	rows := strings.Split(input, "\n")

	seats := make(map[string]*Seat)
	adjacentMap := make(map[string][]*Seat)
	visibleMap := make(map[string][]*Seat)

	for y, row := range rows {
		for x, col := range row {
			if col == 'L' {
				seat := &Seat{x: x, y: y}
				seats[seat.id()] = seat
			}
		}
	}

	for _, seat := range seats {
		adjacents := make([]*Seat, 0)
		for _, candidate := range seats {
			if candidate.id() == seat.id() {
				continue
			}
			if math.Abs(float64(seat.y-candidate.y)) < 2 && math.Abs(float64(seat.x-candidate.x)) < 2 {
				adjacents = append(adjacents, candidate)
			}
		}
		adjacentMap[seat.id()] = adjacents

		visibles := make([]*Seat, 0)
		// find top
		for dy := 1; seat.y+dy < len(rows); dy++ {
			if []rune(rows[seat.y+dy])[seat.x] == 'L' {
				visibles = append(visibles, seats[fmt.Sprintf("(%d, %d)", seat.x, seat.y+dy)])
				break
			}
		}
		// find bottom
		for dy := -1; seat.y+dy >= 0; dy-- {
			if []rune(rows[seat.y+dy])[seat.x] == 'L' {
				visibles = append(visibles, seats[fmt.Sprintf("(%d, %d)", seat.x, seat.y+dy)])
				break
			}
		}
		// find right
		for dx := 1; seat.x+dx < len([]rune(rows[seat.y])); dx++ {
			if []rune(rows[seat.y])[seat.x+dx] == 'L' {
				visibles = append(visibles, seats[fmt.Sprintf("(%d, %d)", seat.x+dx, seat.y)])
				break
			}
		}
		// find left
		for dx := -1; seat.x+dx >= 0; dx-- {
			if []rune(rows[seat.y])[seat.x+dx] == 'L' {
				visibles = append(visibles, seats[fmt.Sprintf("(%d, %d)", seat.x+dx, seat.y)])
				break
			}
		}
		// find top right
		for dy, dx := -1, 1; seat.y+dy >= 0 && seat.x+dx < len([]rune(rows[seat.y+dy])); dy, dx = dy-1, dx+1 {
			if []rune(rows[seat.y+dy])[seat.x+dx] == 'L' {
				visibles = append(visibles, seats[fmt.Sprintf("(%d, %d)", seat.x+dx, seat.y+dy)])
				break
			}
		}
		// find top left
		for dy, dx := -1, -1; seat.y+dy >= 0 && seat.x+dx >= 0; dy, dx = dy-1, dx-1 {
			if []rune(rows[seat.y+dy])[seat.x+dx] == 'L' {
				visibles = append(visibles, seats[fmt.Sprintf("(%d, %d)", seat.x+dx, seat.y+dy)])
				break
			}
		}
		// find bottom right
		for dy, dx := 1, 1; seat.y+dy < len(rows) && seat.x+dx < len([]rune(rows[seat.y+dy])); dy, dx = dy+1, dx+1 {
			if []rune(rows[seat.y+dy])[seat.x+dx] == 'L' {
				visibles = append(visibles, seats[fmt.Sprintf("(%d, %d)", seat.x+dx, seat.y+dy)])
				break
			}
		}
		// find bottom left
		for dy, dx := 1, -1; seat.y+dy < len(rows) && seat.x+dx >= 0; dy, dx = dy+1, dx-1 {
			if []rune(rows[seat.y+dy])[seat.x+dx] == 'L' {
				visibles = append(visibles, seats[fmt.Sprintf("(%d, %d)", seat.x+dx, seat.y+dy)])
				break
			}
		}
		visibleMap[seat.id()] = visibles
	}

	return seats, adjacentMap, visibleMap
}

func main() {
	bytes, _ := ioutil.ReadFile("input.txt")
	input := string(bytes)
	fmt.Println(solution1(parseSeats(input)))
	fmt.Println(solution2(parseSeats(input)))
}
