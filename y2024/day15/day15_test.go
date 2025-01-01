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
	t.Run("shift robot", func(t *testing.T) {
		floorPlan := []string{
			"#.@.OOO.#",
		}
		wh := day15.NewWarehouse(floorPlan)
		wantFloorPlan := [][]rune{
			{
				'#', '.', '.', '@', 'O', 'O', 'O', '.', '#',
			},
		}
		wh.Move('>')
		assert.Equal(t, wantFloorPlan, wh.GetFloorPlan())
	})
	t.Run("robot against a wall", func(t *testing.T) {
		floorPlan := []string{
			"#.@#OOO.#",
		}
		wh := day15.NewWarehouse(floorPlan)
		wantFloorPlan := [][]rune{
			{
				'#', '.', '@', '#', 'O', 'O', 'O', '.', '#',
			},
		}
		wh.Move('>')
		assert.Equal(t, wantFloorPlan, wh.GetFloorPlan())
	})
	t.Run("push boxes", func(t *testing.T) {
		floorPlan := []string{
			"#.@OOO..#",
		}
		wh := day15.NewWarehouse(floorPlan)
		wantFloorPlan := [][]rune{
			{
				'#', '.', '.', '@', 'O', 'O', 'O', '.', '#',
			},
		}
		wh.Move('>')
		assert.Equal(t, wantFloorPlan, wh.GetFloorPlan())
	})
	t.Run("boxes against a wall", func(t *testing.T) {
		floorPlan := []string{
			"#.@OOO#.#",
		}
		wh := day15.NewWarehouse(floorPlan)
		wantFloorPlan := [][]rune{
			{
				'#', '.', '@', 'O', 'O', 'O', '#', '.', '#',
			},
		}
		wh.Move('>')
		assert.Equal(t, wantFloorPlan, wh.GetFloorPlan())
	})
	t.Run("move vertically", func(t *testing.T) {
		floorPlan := []string{
			"#.@#",
			"#..#",
		}
		wh := day15.NewWarehouse(floorPlan)
		wantFloorPlan := [][]rune{
			{
				'#', '.', '.', '#',
			},
			{
				'#', '.', '@', '#',
			},
		}
		wh.Move('v')
		assert.Equal(t, wantFloorPlan, wh.GetFloorPlan())
	})
}

func TestDay15Run(t *testing.T) {
	t.Run("small sample", func(t *testing.T) {
		wh := day15.NewWarehouse([]string{
			"########",
			"#..O.O.#",
			"##@.O..#",
			"#...O..#",
			"#.#.O..#",
			"#...O..#",
			"#......#",
			"########",
		})
		directions := []rune("<^^>>>vv<v>>v<<")
		want := 2028
		got := day15.Run(wh, directions)
		assert.Equal(t, want, got)
	})
	t.Run("large sample", func(t *testing.T) {
		wh := day15.NewWarehouse([]string{
			"##########",
			"#..O..O.O#",
			"#......O.#",
			"#.OO..O.O#",
			"#..O@..O.#",
			"#O#..O...#",
			"#O..O..O.#",
			"#.OO.O.OO#",
			"#....O...#",
			"##########",
		})
		directions := []rune(
			"<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^vvv<<^>^v^^><<>>><>^<<><^vv^^<" +
				">vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v" +
				">v^^<^^vv<<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^^><^><>>><>^^<<^^v>>" +
				"><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>" +
				"^<v^><<<^>>^v<v^v<v^>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^<><^^>^^^<" +
				"><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<" +
				">>>^<^>>>>>^<<^v>^vvv<>^<><<v>v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^")
		want := 10092
		got := day15.Run(wh, directions)
		assert.Equal(t, want, got)
	})
}
