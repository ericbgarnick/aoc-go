package day15_test

import (
	"fmt"
	"testing"

	"github.com/ericbgarnick/aoc-go/y2024/day15"
	"github.com/stretchr/testify/assert"
)

func TestNextPosition(t *testing.T) {
	p := day15.NewPosition(3, 4)
	fTests := map[rune]day15.Position{
		'^': day15.NewPosition(2, 4),
		'v': day15.NewPosition(4, 4),
		'<': day15.NewPosition(3, 3),
		'>': day15.NewPosition(3, 5),
	}
	for d, wantP := range fTests {
		t.Run(fmt.Sprintf("forward %c", d), func(t *testing.T) {
			gotP := day15.NextPosition(p, d, false)
			assert.Equal(t, wantP, gotP)
		})
	}
	rTests := map[rune]day15.Position{
		'v': day15.NewPosition(2, 4),
		'^': day15.NewPosition(4, 4),
		'>': day15.NewPosition(3, 3),
		'<': day15.NewPosition(3, 5),
	}
	for d, wantP := range rTests {
		t.Run(fmt.Sprintf("reverse %c", d), func(t *testing.T) {
			gotP := day15.NextPosition(p, d, true)
			assert.Equal(t, wantP, gotP)
		})
	}
}

func TestWarehouse_MoveNarrow(t *testing.T) {
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
		wh.MoveNarrow('>')
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
		wh.MoveNarrow('>')
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
		wh.MoveNarrow('>')
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
		wh.MoveNarrow('>')
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
		wh.MoveNarrow('v')
		assert.Equal(t, wantFloorPlan, wh.GetFloorPlan())
	})
}

func TestWarehouse_MoveWide(t *testing.T) {
	t.Run("move horizontal", func(t *testing.T) {
		floorPlan := []string{
			"#.@[][].#",
		}
		wh := day15.NewWarehouse(floorPlan)
		wantFloorPlan := [][]rune{
			{
				'#', '.', '.', '@', '[', ']', '[', ']', '#',
			},
		}
		wh.MoveNarrow('>')
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

func TestWarehouse_PushWideBoxes(t *testing.T) {
	wh := day15.NewWarehouse([]string{
		"#.....#",
		"#.[]..#",
		"#..[].#",
	})
	boxPositions := []day15.Position{
		day15.NewPosition(1, 2),
		day15.NewPosition(1, 3),
		day15.NewPosition(2, 3),
		day15.NewPosition(2, 4),
	}
	toMove := []*day15.Position{
		&boxPositions[0],
		&boxPositions[1],
		&boxPositions[2],
		&boxPositions[3],
	}
	wantFloorPlan := [][]rune{
		{'#', '.', '[', ']', '.', '.', '#'},
		{'#', '.', '.', '[', ']', '.', '#'},
		{'#', '.', '.', '.', '.', '.', '#'},
	}
	wh.PushWideBoxes(toMove, '^')
	assert.Equal(t, wantFloorPlan, wh.GetFloorPlan())
}
