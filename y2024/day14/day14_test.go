package day14_test

import (
	"fmt"
	"testing"

	"github.com/ericbgarnick/aoc-go/y2024/day14"
	"github.com/stretchr/testify/assert"
)

/*
 floor: 5 x 7
 guard: (1, 5) => [-1, +2]
 . . . . . . .
 . . . . . . .
 . . . . . . .
 . . . . . . .
 . . . . . . .

 ticks:
 0 - (1, 5) / 2
 1 - (0, 0) / 1
 2 - (4, 2) / 3
 3 - (3, 4) / 4
 4 - (2, 6) / -
 5 - (1, 1) / 1
 6 - (0, 3) / -
*/

func TestSimulateGuard(t *testing.T) {
	var (
		floorRows = 5
		floorCols = 7
	)
	tests := map[int]int{
		0: 2,
		1: 1,
		2: 3,
		3: 4,
		4: 0,
		5: 1,
		6: 0,
	}
	for numTicks, wantQuadrant := range tests {
		t.Run(fmt.Sprintf("%d ticks @ %d quadrant", numTicks, wantQuadrant), func(t *testing.T) {
			g := day14.NewGuard(day14.NewPosition(1, 5), -1, 2)
			day14.SimulateGuard(g, 5, 7, numTicks)
			gotQuadrant := day14.GetQuadrant(g.GetPosition(), floorRows, floorCols)
			assert.Equal(t, wantQuadrant, gotQuadrant)
		})
	}
}

func TestDay14Run(t *testing.T) {
	var (
		floorRows = 7
		floorCols = 11
		numTicks  = 100
	)
	guards := []*day14.Guard{
		day14.NewGuard(day14.NewPosition(4, 0), -3, 3),
		day14.NewGuard(day14.NewPosition(3, 6), -3, -1),
		day14.NewGuard(day14.NewPosition(3, 10), 2, -1),
		day14.NewGuard(day14.NewPosition(0, 2), -1, 2),
		day14.NewGuard(day14.NewPosition(0, 0), 3, 1),
		day14.NewGuard(day14.NewPosition(0, 3), -2, -2),
		day14.NewGuard(day14.NewPosition(6, 7), -3, -1),
		day14.NewGuard(day14.NewPosition(0, 3), -2, -1),
		day14.NewGuard(day14.NewPosition(3, 9), 3, 2),
		day14.NewGuard(day14.NewPosition(3, 7), 2, -1),
		day14.NewGuard(day14.NewPosition(4, 2), -3, 2),
		day14.NewGuard(day14.NewPosition(5, 9), -3, -3),
	}
	want := 12
	got := day14.Run(guards, floorRows, floorCols, numTicks)

	assert.Equal(t, want, got)
}

func TestDay14ParseData(t *testing.T) {
	testData := []string{
		"p=0,4 v=3,-3",
		"p=6,3 v=-1,-3",
		"p=10,3 v=-1,2",
		"p=0,0 v=1,3",
	}
	wantGuards := []*day14.Guard{
		day14.NewGuard(day14.NewPosition(4, 0), -3, 3),
		day14.NewGuard(day14.NewPosition(3, 6), -3, -1),
		day14.NewGuard(day14.NewPosition(3, 10), 2, -1),
		day14.NewGuard(day14.NewPosition(0, 0), 3, 1),
	}
	gotGuards := day14.ParseData(testData)
	for i, guard := range gotGuards {
		assert.Equal(t, *wantGuards[i], *guard)
	}
}
