package day12_test

import (
	"fmt"
	"testing"

	"github.com/ericbgarnick/aoc-go/y2024/day12"
	"github.com/stretchr/testify/assert"
)

func Test_PriceFencing(t *testing.T) {
	tests := map[int][]string{
		140: {
			"AAAA",
			"BBCD",
			"BBCC",
			"EEEC",
		},
		772: {
			"OOOOO",
			"OXOXO",
			"OOOOO",
			"OXOXO",
			"OOOOO",
		},
		1930: {
			"RRRRIICCFF",
			"RRRRIICCCF",
			"VVRRRCCFFF",
			"VVRCCCJFFF",
			"VVVVCJJCFE",
			"VVIVCCJJEE",
			"VVIIICJJEE",
			"MIIIIIJJEE",
			"MIIISIJEEE",
			"MMMISSJEEE",
		},
	}
	for price, garden := range tests {
		t.Run(fmt.Sprintf("%d: %dX%d", price, len(garden), len(garden[0])), func(t *testing.T) {
			got := day12.PriceFencing(garden)
			assert.Equal(t, price, got)
		})
	}
}
