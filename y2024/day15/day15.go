package day15

import (
	"fmt"
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
	Box   = 'O'
	Wall  = '#'
	Robot = '@'
	Floor = '.'
)

type Position struct {
	row, col int
}

func NewPosition(row, col int) *Position {
	return &Position{row: row, col: col}
}

type Warehouse struct {
	robot     *Position
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
		wh.Move(d)
	}
	return wh.SumBoxCoords()
}

func (wh *Warehouse) Move(d rune) {
	// handle non-box movement
	nextP := NextPosition(wh.robot, d)
	if nextObject := wh.floorPlan[nextP.row][nextP.col]; nextObject == Wall {
		return
	} else if nextObject == Floor {
		wh.floorPlan[wh.robot.row][wh.robot.col] = Floor
		wh.robot = NewPosition(nextP.row, nextP.col)
		wh.floorPlan[nextP.row][nextP.col] = Robot
		return
	}

	// handle box movement
	nextRobotP := *nextP
	for wh.floorPlan[nextP.row][nextP.col] == Box {
		nextP = NextPosition(nextP, d)
	}

	// boxes against a wall
	if wh.floorPlan[nextP.row][nextP.col] != Floor {
		return
	}

	// shift boxes
	wh.floorPlan[wh.robot.row][wh.robot.col] = Floor
	wh.floorPlan[nextRobotP.row][nextRobotP.col] = Robot
	wh.robot = NewPosition(nextRobotP.row, nextRobotP.col)
	wh.floorPlan[nextP.row][nextP.col] = Box
}

func NextPosition(p *Position, d rune) *Position {
	if d == '>' {
		return NewPosition(p.row, p.col+1)
	} else if d == '<' {
		return NewPosition(p.row, p.col-1)
	} else if d == '^' {
		return NewPosition(p.row-1, p.col)
	} else if d == 'v' {
		return NewPosition(p.row+1, p.col)
	} else {
		panic(fmt.Sprintf("unknown direction %c", d))
	}
}

func (wh *Warehouse) SumBoxCoords() int {
	var total int
	for r, row := range wh.floorPlan {
		for c := range row {
			if wh.floorPlan[r][c] == Box {
				total += 100*r + c
			}
		}
	}
	return total
}
