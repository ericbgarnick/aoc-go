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

type guard struct {
	*position
	direction rune
}

func Part1() {
	var history = make(map[position]bool)
	floorMap := util.ScanAOCDataFile(2024, 6)
	g := guard{
		position:  findGuard(floorMap),
		direction: guardStart,
	}
	for isInBounds(floorMap, g.position) {
		history[*g.position] = true
		g = moveGuard(floorMap, g)
	}
	fmt.Printf("Part 1: %d\n", len(history))
}

func Part2() {
	fmt.Println("Part 2: TODO")
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

func moveGuard(floorMap []string, g guard) guard {
	curP := g.position
	nextP := g.position
	nextD := g.direction

	switch g.direction {
	case '^':
		nextP = &position{row: curP.row - 1, col: curP.col}
		if isInBounds(floorMap, nextP) && floorMap[nextP.row][nextP.col] == obstacle {
			nextP = g.position
			nextD = '>'
		}
	case '>':
		nextP = &position{row: curP.row, col: curP.col + 1}
		if isInBounds(floorMap, nextP) && floorMap[nextP.row][nextP.col] == obstacle {
			nextP = g.position
			nextD = 'v'
		}
	case 'v':
		nextP = &position{row: curP.row + 1, col: curP.col}
		if isInBounds(floorMap, nextP) && floorMap[nextP.row][nextP.col] == obstacle {
			nextP = g.position
			nextD = '<'
		}
	case '<':
		nextP = &position{row: curP.row, col: curP.col - 1}
		if isInBounds(floorMap, nextP) && floorMap[nextP.row][nextP.col] == obstacle {
			nextP = g.position
			nextD = '^'
		}
	}

	return guard{
		position:  nextP,
		direction: nextD,
	}
}
