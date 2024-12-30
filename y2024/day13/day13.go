package day13

import (
	"fmt"
	"math"

	"github.com/ericbgarnick/aoc-go/util"
)

const (
	maxButtonPushes = 100
	chunkSize       = 4
	ButtonACost     = 3
	ButtonBCost     = 1
)

func Part1() {
	maxCost := maxButtonPushes * (ButtonACost + ButtonBCost)
	machines := processInput(0)
	var totalCost int
	for _, m := range machines {
		totalCost += SolveButtons(m, maxCost)
	}
	fmt.Printf("Part 1: %d\n", totalCost)
}

func Part2() {
	maxCost := math.MaxInt
	machines := processInput(10000000000000)
	var totalCost int
	for _, m := range machines {
		totalCost += SolveButtons(m, maxCost)
	}
	fmt.Printf("Part 2: %d\n", totalCost)
}

type Position struct {
	x, y int
}

func NewPosition(x, y int) *Position {
	return &Position{
		x: x,
		y: y,
	}
}

type Button struct {
	cost                   int
	xIncrement, yIncrement int
}

func NewButton(cost int, xIncrement, yIncrement int) *Button {
	return &Button{cost: cost, xIncrement: xIncrement, yIncrement: yIncrement}
}

type ClawMachine struct {
	ButtonA *Button
	ButtonB *Button
	Prize   *Position
}

func processInput(prizeOffset int) []*ClawMachine {
	rawInput := util.ScanAOCDataFile(2024, 13)
	var (
		machines   []*ClawMachine
		buttonAIdx = 0
		buttonBIdx = 1
		prizeIdx   = 2
	)
	for prizeIdx < len(rawInput) {
		rawButtonA := rawInput[buttonAIdx]
		rawButtonB := rawInput[buttonBIdx]
		rawPrize := rawInput[prizeIdx]

		buttonAIncrements := util.IntPattern.FindAllString(rawButtonA, 2)
		buttonBIncrements := util.IntPattern.FindAllString(rawButtonB, 2)
		prizeCoordinates := util.IntPattern.FindAllString(rawPrize, 2)

		buttonA := NewButton(
			ButtonACost,
			util.MustParseInt(buttonAIncrements[0]),
			util.MustParseInt(buttonAIncrements[1]),
		)
		buttonB := NewButton(
			ButtonBCost,
			util.MustParseInt(buttonBIncrements[0]),
			util.MustParseInt(buttonBIncrements[1]),
		)
		prize := NewPosition(
			util.MustParseInt(prizeCoordinates[0])+prizeOffset,
			util.MustParseInt(prizeCoordinates[1])+prizeOffset,
		)
		machines = append(machines, &ClawMachine{ButtonA: buttonA, ButtonB: buttonB, Prize: prize})

		buttonAIdx += chunkSize
		buttonBIdx += chunkSize
		prizeIdx += chunkSize
	}
	return machines
}

func SolveButtons(m *ClawMachine, maxCost int) int {
	var costA, costB int
	incrementABx := m.ButtonA.xIncrement + m.ButtonB.xIncrement
	incrementABy := m.ButtonA.yIncrement + m.ButtonB.yIncrement

	costAB := m.ButtonA.cost + m.ButtonB.cost

	// calculate cost for pressing button A more
	combinedAx := ((m.ButtonA.xIncrement * m.Prize.y) - (m.ButtonA.yIncrement * m.Prize.x)) /
		((m.ButtonA.xIncrement * incrementABy) - (m.ButtonA.yIncrement * incrementABx))
	combinedAy := ((m.ButtonA.yIncrement * m.Prize.x) - (m.ButtonA.xIncrement * m.Prize.y)) /
		((m.ButtonA.yIncrement * incrementABx) - (m.ButtonA.xIncrement * incrementABy))

	if combinedAx == combinedAy {
		remainingAx := m.Prize.x - (combinedAx * incrementABx)
		remainingAy := m.Prize.y - (combinedAy * incrementABy)
		countAx := remainingAx / m.ButtonA.xIncrement
		countAy := remainingAy / m.ButtonA.yIncrement

		if countAx == countAy && countAx > 0 && countAx*m.ButtonA.xIncrement == remainingAx {
			costA = countAx*m.ButtonA.cost + combinedAx*costAB
		}
	}

	// calculate cost for pressing button B more
	combinedBx := ((m.ButtonB.xIncrement * m.Prize.y) - (m.ButtonB.yIncrement * m.Prize.x)) /
		((m.ButtonB.xIncrement * incrementABy) - (m.ButtonB.yIncrement * incrementABx))
	combinedBy := ((m.ButtonB.yIncrement * m.Prize.x) - (m.ButtonB.xIncrement * m.Prize.y)) /
		((m.ButtonB.yIncrement * incrementABx) - (m.ButtonB.xIncrement * incrementABy))

	if combinedBx == combinedBy {
		remainingBx := m.Prize.x - (combinedBx * incrementABx)
		remainingBy := m.Prize.y - (combinedBy * incrementABy)
		countBx := remainingBx / m.ButtonB.xIncrement
		countBy := remainingBy / m.ButtonB.yIncrement

		if countBx == countBy && countBx > 0 && countBx*m.ButtonB.xIncrement == remainingBx {
			costB = countBx*m.ButtonB.cost + combinedBx*costAB
		}
	}
	if (costA > 0 && costA < maxCost) && (costA <= costB || costB <= 0) {
		return costA
	} else if (costB > 0 && costB < maxCost) && (costB < costA || costA <= 0) {
		return costB
	} else {
		return 0
	}
}
