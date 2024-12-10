package day08

import (
	"fmt"

	"github.com/ericbgarnick/aoc-go/util"
)

type position struct {
	row, col int
}

const empty = '.'

func Part1() {
	fmt.Printf("Part 1: %d\n", solve(findSimpleAntinodes))
}

func Part2() {
	fmt.Printf("Part 2: %d\n", solve(findComplexAntinodes))
}

func solve(antinodesFunc func(position, position, int, int) []position) int {
	var (
		numRows, numCols int
		antennas         = make(map[rune][]position)
		antinodes        = make(map[position]bool)
	)
	for r, line := range util.ScanAOCDataFile(2024, 8) {
		numRows++
		numCols = len(line)
		for c, symbol := range line {
			if symbol != empty {
				antennas[symbol] = append(antennas[symbol], position{row: r, col: c})
			}
		}
	}
	for _, positions := range antennas {
		addAntinodes(positions, numRows, numCols, &antinodes, antinodesFunc)
	}
	return len(antinodes)
}

func addAntinodes(positions []position, maxRow, maxCol int, antinodes *map[position]bool, antinodesFunc func(position, position, int, int) []position) {
	for _, p1 := range positions {
		for _, p2 := range positions {
			if p1 != p2 {
				newAntinodes := antinodesFunc(p1, p2, maxRow, maxCol)
				for _, a := range newAntinodes {
					(*antinodes)[a] = true
				}
			}
		}
	}
}

func findSimpleAntinodes(p1, p2 position, maxRow, maxCol int) []position {
	var antinodes []position
	rowDiff := p1.row - p2.row
	colDiff := p1.col - p2.col
	a1 := position{row: p1.row + rowDiff, col: p1.col + colDiff}
	if positionInBounds(a1, maxRow, maxCol) {
		antinodes = append(antinodes, a1)
	}
	a2 := position{row: p2.row - rowDiff, col: p2.col - colDiff}
	if positionInBounds(a2, maxRow, maxCol) {
		antinodes = append(antinodes, a2)
	}
	return antinodes
}

func findComplexAntinodes(p1, p2 position, maxRow, maxCol int) []position {
	var antinodes []position
	rowDiff := p1.row - p2.row
	colDiff := p1.col - p2.col
	aRow := p2.row
	aCol := p2.col
	for positionInBounds(position{row: aRow, col: aCol}, maxRow, maxCol) {
		antinodes = append(antinodes, position{row: aRow, col: aCol})
		aRow -= rowDiff
		aCol -= colDiff
	}
	aRow = p1.row
	aCol = p1.col
	for positionInBounds(position{row: aRow, col: aCol}, maxRow, maxCol) {
		antinodes = append(antinodes, position{row: aRow, col: aCol})
		aRow += rowDiff
		aCol += colDiff
	}
	return antinodes
}

func positionInBounds(p position, maxRow, maxCol int) bool {
	return p.row >= 0 && p.col >= 0 && p.row < maxRow && p.col < maxCol
}
