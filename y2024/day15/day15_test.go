package day15_test

import (
	"testing"

	"github.com/ericbgarnick/aoc-go/y2024/day15"
	"github.com/stretchr/testify/assert"
)

func TestNextPosition(t *testing.T) {
	p := day15.NewPosition(3, 4)
	tests := map[rune]*day15.Position{
		'^': day15.NewPosition(2, 4),
		'v': day15.NewPosition(4, 4),
		'<': day15.NewPosition(3, 3),
		'>': day15.NewPosition(3, 5),
	}
	for d, wantP := range tests {
		t.Run(string(d), func(t *testing.T) {
			gotP := day15.NextPosition(p, d)
			assert.Equal(t, *wantP, *gotP)
		})
	}
}

func TestWarehouse_Move(t *testing.T) {
	t.Run("shift robot right", func(t *testing.T) {
		floopPlan := []string{
			"#.@.OOO..#",
		}
		wh := day15.NewWarehouse(floopPlan)
		wantFloorPlan := [][]rune{
			{
				'#', '.', '.', '@', 'O', 'O', 'O', '.', '.', '#',
			},
		}
		wh.Move(day15.NewPosition(0, 2), '>')
		assert.Equal(t, wantFloorPlan, wh.GetFloorPlan())
	})
	t.Run("shift boxes right", func(t *testing.T) {
		floopPlan := []string{
			"#..@OOO..#",
		}
		wh := day15.NewWarehouse(floopPlan)
		wantFloorPlan := [][]rune{
			{
				'#', '.', '.', '.', '.', '@', 'O', 'O', 'O', '#',
			},
		}
		wh.Move(day15.NewPosition(0, 2), '>')
		assert.Equal(t, wantFloorPlan, wh.GetFloorPlan())
	})
}
