package day12_test

import (
	"fmt"
	"testing"

	"github.com/ericbgarnick/aoc-go/y2024/day12"
	"github.com/stretchr/testify/assert"
)

func Test_PriceFencingByPerimeter(t *testing.T) {
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
			got := day12.PriceFencing(garden, false)
			assert.Equal(t, price, got)
		})
	}
}

func Test_PriceFencingBySides(t *testing.T) {
	tests := map[int][]string{
		80: {
			"AAAA",
			"BBCD",
			"BBCC",
			"EEEC",
		},
		//36: {
		//	"XAX",
		//	"AXA",
		//	"XAX",
		//},
		68: {
			"OOO",
			"OXO",
			"OOO",
		},
		196: {
			"XXXXX",
			"XOOOX",
			"XOXOX",
			"XOOOX",
			"XXXXX",
		},
		436: {
			"OOOOO",
			"OXOXO",
			"OOOOO",
			"OXOXO",
			"OOOOO",
		},
		236: {
			"EEEEE",
			"EXXXX",
			"EEEEE",
			"EXXXX",
			"EEEEE",
		},
		368: {
			"AAAAAA",
			"AAABBA",
			"AAABBA",
			"ABBAAA",
			"ABBAAA",
			"AAAAAA",
		},
		1206: {
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
		148: { // X: 14 * (4+4), O: 3 * 6, I: 3 * 6 => 112 + 18 + 18
			"XXXXX",
			"XOIIX",
			"XOOIX",
			"XXXXX",
		},
		220: { // A: 15 * (4+8), D: 5 * 8 => 180 + 40 = 220
			"AAAAA",
			"ADDDA",
			"ADADA",
			"AAAAA",
		},
		388: { // 24 * (4+4), 196
			"OOOOOOO",
			"OXXXXXO",
			"OXOOOXO",
			"OXOXOXO",
			"OXOOOXO",
			"OXXXXXO",
			"OOOOOOO",
		},
	}
	for price, garden := range tests {
		t.Run(fmt.Sprintf("%d: %dX%d", price, len(garden), len(garden[0])), func(t *testing.T) {
			got := day12.PriceFencing(garden, true)
			assert.Equal(t, price, got)
		})
	}
}
