package day04

import (
	"fmt"

	"github.com/ericbgarnick/aoc-go/util"
)

const fileName = "y2024/day04/data.txt"

// Part1 counts the occurrences of the string "XMAS" in the grid, going in any direction.
func Part1() {
	grid := util.ScanFile(fileName)
	word := "XMAS"
	fmt.Printf("Part 1: %d\n", countOccurrences(grid, word))
}

// Part2 counts the occurrences in the grid of the string "MAS" crossing diagonally on "A".
func Part2() {
	grid := util.ScanFile(fileName)
	fmt.Printf("Part 2: %d\n", countMASCrosses(grid))
}

type position struct {
	row, col int
}

func countOccurrences(grid []string, word string) int {
	var total int
	for r, row := range grid {
		for c := range row {
			if row[c] == word[0] {
				p := position{row: r, col: c}
				if searchNorth(grid, word, p) {
					total++
				}
				if searchSouth(grid, word, p) {
					total++
				}
				if searchEast(grid, word, p) {
					total++
				}
				if searchWest(grid, word, p) {
					total++
				}
				if searchNortheast(grid, word, p) {
					total++
				}
				if searchNorthwest(grid, word, p) {
					total++
				}
				if searchSoutheast(grid, word, p) {
					total++
				}
				if searchSouthwest(grid, word, p) {
					total++
				}
			}
		}
	}
	return total
}

func searchNorth(grid []string, word string, start position) bool {
	if len(word)-start.row > 1 {
		return false
	}
	var curLetter int
	for curLetter < len(word) {
		if grid[start.row-curLetter][start.col] != word[curLetter] {
			return false
		}
		curLetter++
	}
	return true
}

func searchSouth(grid []string, word string, start position) bool {
	if start.row+len(word) > len(grid) {
		return false
	}
	var curLetter int
	for curLetter < len(word) {
		if grid[start.row+curLetter][start.col] != word[curLetter] {
			return false
		}
		curLetter++
	}
	return true
}

func searchEast(grid []string, word string, start position) bool {
	if start.col+len(word) > len(grid[0]) {
		return false
	}
	var curLetter int
	for curLetter < len(word) {
		if grid[start.row][start.col+curLetter] != word[curLetter] {
			return false
		}
		curLetter++
	}
	return true
}

func searchWest(grid []string, word string, start position) bool {
	if len(word)-start.col > 1 {
		return false
	}
	var curLetter int
	for curLetter < len(word) {
		if grid[start.row][start.col-curLetter] != word[curLetter] {
			return false
		}
		curLetter++
	}
	return true
}

func searchNortheast(grid []string, word string, start position) bool {
	if len(word)-start.row > 1 || start.col+len(word) > len(grid[0]) {
		return false
	}
	var curLetter int
	for curLetter < len(word) {
		if grid[start.row-curLetter][start.col+curLetter] != word[curLetter] {
			return false
		}
		curLetter++
	}
	return true
}

func searchNorthwest(grid []string, word string, start position) bool {
	if len(word)-start.row > 1 || len(word)-start.col > 1 {
		return false
	}
	var curLetter int
	for curLetter < len(word) {
		if grid[start.row-curLetter][start.col-curLetter] != word[curLetter] {
			return false
		}
		curLetter++
	}
	return true
}

func searchSoutheast(grid []string, word string, start position) bool {
	if start.row+len(word) > len(grid) || start.col+len(word) > len(grid[0]) {
		return false
	}
	var curLetter int
	for curLetter < len(word) {
		if grid[start.row+curLetter][start.col+curLetter] != word[curLetter] {
			return false
		}
		curLetter++
	}
	return true
}

func searchSouthwest(grid []string, word string, start position) bool {
	if start.row+len(word) > len(grid) || len(word)-start.col > 1 {
		return false
	}
	var curLetter int
	for curLetter < len(word) {
		if grid[start.row+curLetter][start.col-curLetter] != word[curLetter] {
			return false
		}
		curLetter++
	}
	return true
}

func countMASCrosses(grid []string) int {
	var total int
	for r, row := range grid {
		for c := range row {
			if row[c] == 'A' && hasMSDiagonals(grid, position{row: r, col: c}) {
				total++
			}
		}
	}
	return total
}

func hasMSDiagonals(grid []string, middle position) bool {
	if middle.row == 0 || middle.col == 0 {
		return false
	}
	if middle.row+1 >= len(grid) || middle.col+1 >= len(grid[0]) {
		return false
	}
	nw := grid[middle.row-1][middle.col-1]
	ne := grid[middle.row-1][middle.col+1]
	sw := grid[middle.row+1][middle.col-1]
	se := grid[middle.row+1][middle.col+1]
	if ((nw == 'M' || nw == 'S') && (se == 'M' || se == 'S') && nw != se) &&
		((ne == 'M' || ne == 'S') && (sw == 'M' || sw == 'S') && ne != sw) {
		return true
	}
	return false
}
