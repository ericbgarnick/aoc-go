package day06

import (
	"fmt"

	"github.com/ericbgarnick/aoc-go/util"
)

const (
	guardStart = '^'
	obstacle   = '#'
)

type position struct {
	row, col int
}

type history map[position]map[rune]bool

type guard struct {
	*position
	direction rune
	startPos  position
	history
}

func newGuard(p *position, d rune) *guard {
	g := guard{
		position:  p,
		direction: d,
		history:   history{},
	}
	g.startPos = *g.position
	return &g
}

func Part1() {
	floorMap := util.ScanAOCDataFile(2024, 6)
	g := newGuard(findGuard(floorMap), guardStart)
	g, _ = simulate(floorMap, g)
	fmt.Printf("Part 1: %d\n", len(g.history))
}

func Part2() {
	var (
		loops   = make(map[position]bool)
		hasLoop bool
	)
	floorMap := util.ScanAOCDataFile(2024, 6)
	g := newGuard(findGuard(floorMap), guardStart)
	g, _ = simulate(floorMap, g)
	for p, dirs := range g.history {
		for d := range dirs {
			if !(p == g.startPos && d == guardStart) {
				newFloor, obstacleP := addObstacle(floorMap, &p, d)
				if newFloor != nil {
					if _, hasLoop = simulate(newFloor, newGuard(&g.startPos, guardStart)); hasLoop {
						loops[*obstacleP] = true
					}
				}
			}
		}
	}
	fmt.Printf("Part 2: %d\n", len(loops))
}

func simulate(floorMap []string, g *guard) (*guard, bool) {
	for isInBounds(floorMap, g.position) {
		prevDirections, ok := g.history[*g.position]
		if prevDirections[g.direction] {
			return g, true
		}
		if ok {
			g.history[*g.position][g.direction] = true
		} else {
			g.history[*g.position] = map[rune]bool{g.direction: true}
		}

		g = moveGuard(floorMap, g)
	}
	return g, false
}

func findGuard(floorMap []string) *position {
	for r, row := range floorMap {
		for c := range row {
			if floorMap[r][c] == guardStart {
				return &position{row: r, col: c}
			}
		}
	}
	panic("no guard found")
}

func isInBounds(floorMap []string, p *position) bool {
	return p.row >= 0 && p.col >= 0 && p.row < len(floorMap) && p.col < len(floorMap[0])
}

func moveGuard(floorMap []string, g *guard) *guard {
	nextD := g.direction

	nextP := getNextPosition(floorMap, g.position, g.direction)
	if nextP == g.position {
		nextD = turn(g.direction)
	}

	return &guard{
		position:  nextP,
		direction: nextD,
		startPos:  g.startPos,
		history:   g.history,
	}
}

func getNextPosition(floorMap []string, p *position, d rune) *position {
	curP := p
	nextP := p

	switch d {
	case '^':
		nextP = &position{row: curP.row - 1, col: curP.col}
		if isInBounds(floorMap, nextP) && floorMap[nextP.row][nextP.col] == obstacle {
			nextP = p
		}
	case '>':
		nextP = &position{row: curP.row, col: curP.col + 1}
		if isInBounds(floorMap, nextP) && floorMap[nextP.row][nextP.col] == obstacle {
			nextP = p
		}
	case 'v':
		nextP = &position{row: curP.row + 1, col: curP.col}
		if isInBounds(floorMap, nextP) && floorMap[nextP.row][nextP.col] == obstacle {
			nextP = p
		}
	case '<':
		nextP = &position{row: curP.row, col: curP.col - 1}
		if isInBounds(floorMap, nextP) && floorMap[nextP.row][nextP.col] == obstacle {
			nextP = p
		}
	}
	return nextP
}

func turn(direction rune) rune {
	switch direction {
	case '^':
		return '>'
	case '>':
		return 'v'
	case 'v':
		return '<'
	case '<':
		return '^'
	}
	panic(fmt.Sprintf("unknown position %s", string(direction)))
}

func addObstacle(floorMap []string, p *position, d rune) ([]string, *position) {
	nextP := getNextPosition(floorMap, p, d)
	if !isInBounds(floorMap, nextP) {
		return nil, nil
	}
	var newFloor []string
	for r, row := range floorMap {
		if r == nextP.row {
			row = row[:nextP.col] + string(obstacle) + row[nextP.col+1:]
		}
		newFloor = append(newFloor, row)
	}
	return newFloor, nextP
}

func printFloor(floorMap []string) {
	for _, row := range floorMap {
		fmt.Println(row)
	}
}
