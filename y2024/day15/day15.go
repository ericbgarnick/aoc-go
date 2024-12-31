package day15

import (
	"fmt"
)

func Part1() {
	fmt.Printf("Part 1: %s\n", "TODO")
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
	numRows, numCols int
	floorPlan        [][]rune
}

func NewWarehouse(rawFloorPlan []string) *Warehouse {
	wh := Warehouse{
		numRows: len(rawFloorPlan),
		numCols: len(rawFloorPlan[0]),
	}
	for _, row := range rawFloorPlan {
		wh.floorPlan = append(wh.floorPlan, []rune(row))
	}
	return &wh
}

func (wh *Warehouse) GetFloorPlan() [][]rune {
	return wh.floorPlan
}

func (wh *Warehouse) Move(robotP *Position, d rune) {
	// sanity-check we are starting with a robot
	if wh.floorPlan[robotP.row][robotP.col] != Robot {
		panic(fmt.Sprintf(
			"non-robot start for shift boxes: %c at %v",
			wh.floorPlan[robotP.row][robotP.col], robotP),
		)
	}

	// handle non-box spaces
	nextP := NextPosition(robotP, d)
	if nextObject := wh.floorPlan[nextP.row][nextP.col]; nextObject == Wall {
		return
	} else if nextObject == Floor {
		wh.floorPlan[robotP.row][robotP.col] = Floor
		wh.floorPlan[nextP.row][nextP.col] = Robot
		return
	}

	// traverse boxes
	for wh.floorPlan[nextP.row][nextP.col] == Box {
		nextP = NextPosition(nextP, d)
	}

	// traverse floor spaces, adding boxes
	var numFloorTiles int
	for wh.floorPlan[nextP.row][nextP.col] == Floor {
		numFloorTiles++
		wh.floorPlan[nextP.row][nextP.col] = Box
		nextP = NextPosition(nextP, d)
	}

	// clean up moved boxes and move robot
	wh.floorPlan[robotP.row][robotP.col] = Floor
	nextP = NextPosition(robotP, d)
	for i := 1; i < numFloorTiles; i++ {
		wh.floorPlan[nextP.row][nextP.col] = Floor
		nextP = NextPosition(nextP, d)
	}
	wh.floorPlan[robotP.row][robotP.col] = Robot
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
