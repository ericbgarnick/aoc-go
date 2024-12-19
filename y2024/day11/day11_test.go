package day11_test

import (
	"fmt"
	"testing"

	"github.com/ericbgarnick/aoc-go/y2024/day11"
	"github.com/stretchr/testify/assert"
)

func Test_Blink(t *testing.T) {
	tests := map[int][]int{
		0:    {1},
		10:   {1, 0},
		1000: {10, 0},
		1001: {10, 1},
		1:    {2024},
	}
	for stone, result := range tests {
		t.Run(fmt.Sprintf("%d: %v", stone, result), func(t *testing.T) {
			got := day11.Blink(stone)
			assert.Equal(t, result, got)
		})
	}
}

func Test_Expand(t *testing.T) {
	t.Run("test case", func(t *testing.T) {
		var got int
		for _, s := range []int{125, 17} {
			got += day11.Expand(s, 25, 25)
		}
		want := 55312
		assert.Equal(t, want, got)
	})
}
