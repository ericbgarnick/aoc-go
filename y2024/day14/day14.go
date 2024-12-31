package day14

import (
	"fmt"
	"strings"
	"sync"

	"github.com/ericbgarnick/aoc-go/util"
)

const (
	floorRows = 103
	floorCols = 101
)

func Part1() {
	rawData := util.ScanAOCDataFile(2024, 14)
	guards := ParseData(rawData)
	fmt.Printf("Part 1: %d\n", Run(guards, floorRows, floorCols, 100))
}

func Part2() {
	rawData := util.ScanAOCDataFile(2024, 14)
	guards := ParseData(rawData)
	fmt.Print("Part 2: ")
	FindShape(guards, floorRows, floorCols)
}

type Position struct {
	row, col int
}

func NewPosition(r, c int) *Position {
	return &Position{row: r, col: c}
}

type Guard struct {
	position    *Position
	rowVelocity int
	colVelocity int
}

func (g *Guard) GetPosition() *Position {
	return g.position
}

func NewGuard(p *Position, rVel, cVel int) *Guard {
	return &Guard{
		position:    p,
		rowVelocity: rVel,
		colVelocity: cVel,
	}
}

func ParseData(lines []string) []*Guard {
	var guards = make([]*Guard, len(lines))
	for i, line := range lines {
		rawValues := util.SignedIntPattern.FindAllString(line, 4)
		guards[i] = NewGuard(
			NewPosition(util.MustParseInt(rawValues[1]), util.MustParseInt(rawValues[0])),
			util.MustParseInt(rawValues[3]),
			util.MustParseInt(rawValues[2]),
		)
	}
	return guards
}

func FindShape(guards []*Guard, floorRows, floorCols int) {
	wg := sync.WaitGroup{}
	var numTicks int
	for {
		for _, g := range guards {
			wg.Add(1)
			go func() {
				SimulateGuard(g, floorRows, floorCols, 1)
				wg.Done()
			}()
		}
		wg.Wait()
		numTicks++
		if (numTicks-39)%101 == 0 {
			//display(guards, floorRows, floorCols, numTicks)
		}
		if numTicks == 7412 {
			display(guards, floorRows, floorCols, numTicks)
			break
		}
	}
}

func Run(guards []*Guard, floorRows, floorCols, totalTicks int) int {
	var (
		quadrants [4]int
		result    = 1
		m         sync.Mutex
	)
	wg := sync.WaitGroup{}
	for _, g := range guards {
		go func() {
			SimulateGuard(g, floorRows, floorCols, totalTicks)
			q := GetQuadrant(g.GetPosition(), floorRows, floorCols)
			if q != 0 {
				m.Lock()
				quadrants[q-1]++
				m.Unlock()
			}
			wg.Done()
		}()
		wg.Add(1)
	}
	wg.Wait()
	for _, qCount := range quadrants {
		result *= qCount
	}
	return result
}

func SimulateGuard(g *Guard, floorRows, floorCols, totalTicks int) {
	var newRow, newCol int
	for t := 0; t < totalTicks; t++ {
		newRow = (g.position.row + g.rowVelocity) % floorRows
		if newRow < 0 {
			g.position.row = floorRows + newRow
		} else {
			g.position.row = newRow
		}

		newCol = (g.position.col + g.colVelocity) % floorCols
		if newCol < 0 {
			g.position.col = floorCols + newCol
		} else {
			g.position.col = newCol
		}
	}
}

// GetQuadrant returns the quadrant (1-4) of the given position, going in read-order from top-left:
//
//	1 2
//	3 4
func GetQuadrant(p *Position, floorRows, floorCols int) int {
	var (
		rowHalf = 1
		colHalf = 0
	)
	if p.row == (floorRows/2) || p.col == (floorCols/2) {
		return 0
	}
	if p.row > (floorRows / 2) {
		rowHalf += 2
	}
	if p.col > (floorCols / 2) {
		colHalf += 1
	}
	return rowHalf + colHalf
}

func display(guards []*Guard, floorRows, floorCols, numTicks int) {
	var floor = make([][]string, floorRows)
	for i := range floor {
		newRow := make([]string, floorCols)
		for j := range newRow {
			newRow[j] = "."
		}
		floor[i] = newRow
	}
	for _, g := range guards {
		floor[g.position.row][g.position.col] = "#"
	}
	fmt.Printf("%d SECONDS\n", numTicks)
	for _, row := range floor {
		fmt.Println(strings.Join(row, ""))
	}
}
