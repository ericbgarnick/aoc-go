package day15

import (
	"fmt"
	"slices"
	"strings"

	"github.com/ericbgarnick/aoc-go/util"
)

func Part1() {
	wh, directions := loadData()
	fmt.Printf("Part 1: %d\n", Run(wh, directions))
}

func Part2() {
	fmt.Printf("Part 2: %s\n", "TODO")
}

const (
	SingleBox      = 'O'
	DoubleBoxLeft  = '['
	DoubleBoxRight = ']'
	Wall           = '#'
	Robot          = '@'
	Floor          = '.'
)

var boxShapes = []rune{SingleBox, DoubleBoxLeft, DoubleBoxRight}

type Position struct {
	row, col int
}

func NewPosition(row, col int) Position {
	return Position{row: row, col: col}
}

type Warehouse struct {
	robot     Position
	floorPlan [][]rune
}

func NewWarehouse(rawFloorPlan []string) *Warehouse {
	wh := Warehouse{}
	for r, row := range rawFloorPlan {
		floorRow := []rune(row)
		for c := range floorRow {
			if floorRow[c] == Robot {
				wh.robot = NewPosition(r, c)
			}
		}
		wh.floorPlan = append(wh.floorPlan, floorRow)
	}
	return &wh
}

func (wh *Warehouse) Print() {
	for _, row := range wh.floorPlan {
		fmt.Println(string(row))
	}
}

func (wh *Warehouse) IsBox(p Position) bool {
	return slices.Contains(boxShapes, wh.floorPlan[p.row][p.col])
}

func (wh *Warehouse) GetFloorPlan() [][]rune {
	return wh.floorPlan
}

func loadData() (*Warehouse, []rune) {
	rawData := util.ScanAOCDataFile(2024, 15)
	var (
		rawFloorPlan      []string
		directions        []rune
		readingDirections bool
	)
	for _, row := range rawData {
		row = strings.TrimSpace(row)
		if len(row) == 0 {
			readingDirections = true
		}
		if readingDirections {
			directions = append(directions, []rune(row)...)
		} else {
			rawFloorPlan = append(rawFloorPlan, row)
		}
	}
	return NewWarehouse(rawFloorPlan), directions
}

func Run(wh *Warehouse, directions []rune) int {
	for _, d := range directions {
		wh.MoveNarrow(d)
	}
	return wh.SumBoxCoordsNarrow()
}

func (wh *Warehouse) MoveNarrow(d rune) {
	// handle non-box movement
	nextP := NextPositionNarrow(wh.robot, d, false)
	if nextObject := wh.floorPlan[nextP.row][nextP.col]; nextObject == Wall {
		return
	} else if nextObject == Floor {
		wh.floorPlan[wh.robot.row][wh.robot.col] = Floor
		wh.robot = NewPosition(nextP.row, nextP.col)
		wh.floorPlan[nextP.row][nextP.col] = Robot
		return
	}

	// handle box movement
	for wh.IsBox(nextP) {
		nextP = NextPositionNarrow(nextP, d, false)
	}

	// boxes against a wall
	if wh.floorPlan[nextP.row][nextP.col] != Floor {
		return
	}

	// shift boxes
	var lastP Position
	for nextP != wh.robot {
		lastP = nextP
		nextP = NextPositionNarrow(nextP, d, true)
		wh.floorPlan[lastP.row][lastP.col] = wh.floorPlan[nextP.row][nextP.col]
	}
	wh.floorPlan[lastP.row][lastP.col] = Robot
	wh.robot = NewPosition(lastP.row, lastP.col)
	wh.floorPlan[nextP.row][nextP.col] = Floor
}

func NextPositionNarrow(p Position, d rune, reverse bool) Position {
	if (d == '>' && !reverse) || (d == '<' && reverse) {
		return NewPosition(p.row, p.col+1)
	} else if (d == '<' && !reverse) || (d == '>' && reverse) {
		return NewPosition(p.row, p.col-1)
	} else if (d == '^' && !reverse) || (d == 'v' && reverse) {
		return NewPosition(p.row-1, p.col)
	} else if (d == 'v' && !reverse) || (d == '^' && reverse) {
		return NewPosition(p.row+1, p.col)
	} else {
		panic(fmt.Sprintf("unknown direction %c", d))
	}
}

func (wh *Warehouse) SumBoxCoordsNarrow() int {
	var total int
	for r, row := range wh.floorPlan {
		for c := range row {
			if wh.floorPlan[r][c] == SingleBox {
				total += 100*r + c
			}
		}
	}
	return total
}

func (wh *Warehouse) MoveWide(d rune) {
	if d == '<' || d == '>' {

	}
}
