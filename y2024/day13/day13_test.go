package day13_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/ericbgarnick/aoc-go/y2024/day13"
	"github.com/stretchr/testify/assert"
)

func TestSolveButtons(t *testing.T) {
	type testCase struct {
		machine *day13.ClawMachine
		cost    int
	}
	tests := []testCase{
		{
			machine: &day13.ClawMachine{
				ButtonA: day13.NewButton(3, 3, 1),
				ButtonB: day13.NewButton(1, 2, 4),
				Prize:   day13.NewPosition(16, 12),
			},
			cost: 14,
		},
		{
			machine: &day13.ClawMachine{
				ButtonA: day13.NewButton(3, 3, 1),
				ButtonB: day13.NewButton(1, 2, 4),
				Prize:   day13.NewPosition(16, 11),
			},
			cost: 0,
		},
		{
			machine: &day13.ClawMachine{
				ButtonA: day13.NewButton(3, 5, 4),
				ButtonB: day13.NewButton(1, 3, 6),
				Prize:   day13.NewPosition(25, 38),
			},
			cost: 11,
		},
		{
			machine: &day13.ClawMachine{
				ButtonA: day13.NewButton(3, 94, 34),
				ButtonB: day13.NewButton(1, 22, 67),
				Prize:   day13.NewPosition(8400, 5400),
			},
			cost: 280,
		},
		{
			machine: &day13.ClawMachine{
				ButtonA: day13.NewButton(3, 26, 66),
				ButtonB: day13.NewButton(1, 67, 21),
				Prize:   day13.NewPosition(12748, 12176),
			},
			cost: 0,
		},
		{
			machine: &day13.ClawMachine{
				ButtonA: day13.NewButton(3, 17, 86),
				ButtonB: day13.NewButton(1, 84, 37),
				Prize:   day13.NewPosition(7870, 6450),
			},
			cost: 200,
		},
		{
			machine: &day13.ClawMachine{
				ButtonA: day13.NewButton(3, 69, 23),
				ButtonB: day13.NewButton(1, 27, 71),
				Prize:   day13.NewPosition(18641, 10279),
			},
			cost: 0,
		},
		{
			machine: &day13.ClawMachine{
				ButtonA: day13.NewButton(3, 94, 34),
				ButtonB: day13.NewButton(1, 22, 67),
				Prize:   day13.NewPosition(10000000008400, 10000000005400),
			},
			cost: 0,
		},
		{
			machine: &day13.ClawMachine{
				ButtonA: day13.NewButton(3, 26, 66),
				ButtonB: day13.NewButton(1, 67, 21),
				Prize:   day13.NewPosition(10000000012748, 10000000012176),
			},
			cost: 459236326669,
		},
		{
			machine: &day13.ClawMachine{
				ButtonA: day13.NewButton(3, 17, 86),
				ButtonB: day13.NewButton(1, 84, 37),
				Prize:   day13.NewPosition(10000000007870, 10000000006450),
			},
			cost: 0,
		},
		{
			machine: &day13.ClawMachine{
				ButtonA: day13.NewButton(3, 69, 23),
				ButtonB: day13.NewButton(1, 27, 71),
				Prize:   day13.NewPosition(10000000018641, 10000000010279),
			},
			cost: 416082282239,
		},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v @ %d", tc.machine.Prize, tc.cost), func(t *testing.T) {
			gotCost := day13.SolveButtons(tc.machine, math.MaxInt)
			assert.Equal(t, tc.cost, gotCost)
		})
	}
}
