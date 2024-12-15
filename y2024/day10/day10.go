package day10

import (
	"fmt"

	"github.com/ericbgarnick/aoc-go/util"
)

type position struct {
	row, col int
}

func Part1() {
	trailMap := loadTrailMap()
	fmt.Printf("Part 1: %d\n", ScoreMap(trailMap))
}

func Part2() {
	trailMap := loadTrailMap()
	fmt.Printf("Part 2: %d\n", RateMap(trailMap))
}

func loadTrailMap() [][]int {
	rawMap := util.ScanAOCDataFile(2024, 10)
	var trailMap [][]int
	for _, row := range rawMap {
		newRow := make([]int, len(row))
		for c, val := range row {
			newRow[c] = util.MustParseInt(string(val))
		}
		trailMap = append(trailMap, newRow)
	}
	return trailMap
}

func ScoreMap(trailMap [][]int) int {
	var score int
	for r, row := range trailMap {
		for c, v := range row {
			if v == 0 {
				found := scorePosition(position{row: r, col: c}, trailMap)
				score += len(found)
			}
		}
	}
	return score
}

func scorePosition(start position, trailMap [][]int) map[position]bool {
	var peaks = make(map[position]bool)
	curVal := trailMap[start.row][start.col]
	if curVal == 9 {
		return map[position]bool{start: true}
	}
	if start.row > 0 && trailMap[start.row-1][start.col] == curVal+1 {
		found := scorePosition(position{row: start.row - 1, col: start.col}, trailMap)
		for peak := range found {
			peaks[peak] = true
		}
	}
	if start.row+1 < len(trailMap) && trailMap[start.row+1][start.col] == curVal+1 {
		found := scorePosition(position{row: start.row + 1, col: start.col}, trailMap)
		for peak := range found {
			peaks[peak] = true
		}
	}
	if start.col > 0 && trailMap[start.row][start.col-1] == curVal+1 {
		found := scorePosition(position{row: start.row, col: start.col - 1}, trailMap)
		for peak := range found {
			peaks[peak] = true
		}
	}
	if start.col+1 < len(trailMap[0]) && trailMap[start.row][start.col+1] == curVal+1 {
		found := scorePosition(position{row: start.row, col: start.col + 1}, trailMap)
		for peak := range found {
			peaks[peak] = true
		}
	}
	return peaks
}

func RateMap(trailMap [][]int) int {
	var rating int
	for r, row := range trailMap {
		for c, v := range row {
			if v == 0 {
				rating += ratePosition(position{row: r, col: c}, trailMap)
			}
		}
	}
	return rating
}

func ratePosition(start position, trailMap [][]int) int {
	var rating int
	curVal := trailMap[start.row][start.col]
	if curVal == 9 {
		return 1
	}
	if start.row > 0 && trailMap[start.row-1][start.col] == curVal+1 {
		rating += ratePosition(position{row: start.row - 1, col: start.col}, trailMap)
	}
	if start.row+1 < len(trailMap) && trailMap[start.row+1][start.col] == curVal+1 {
		rating += ratePosition(position{row: start.row + 1, col: start.col}, trailMap)
	}
	if start.col > 0 && trailMap[start.row][start.col-1] == curVal+1 {
		rating += ratePosition(position{row: start.row, col: start.col - 1}, trailMap)
	}
	if start.col+1 < len(trailMap[0]) && trailMap[start.row][start.col+1] == curVal+1 {
		rating += ratePosition(position{row: start.row, col: start.col + 1}, trailMap)
	}
	return rating
}
