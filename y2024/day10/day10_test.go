package day10_test

import (
	"testing"

	"github.com/ericbgarnick/aoc-go/y2024/day10"
	"github.com/stretchr/testify/assert"
)

func Test_Day10(t *testing.T) {
	trailMap := [][]int{
		{8, 9, 0, 1, 0, 1, 2, 3},
		{7, 8, 1, 2, 1, 8, 7, 4},
		{8, 7, 4, 3, 0, 9, 6, 5},
		{9, 6, 5, 4, 9, 8, 7, 4},
		{4, 5, 6, 7, 8, 9, 0, 3},
		{3, 2, 0, 1, 9, 0, 1, 2},
		{0, 1, 3, 2, 9, 8, 0, 1},
		{1, 0, 4, 5, 6, 7, 3, 2},
	}
	t.Run("ScoreMap", func(t *testing.T) {
		want := 36
		got := day10.ScoreMap(trailMap)

		assert.Equal(t, want, got)
	})
	t.Run("RateMap", func(t *testing.T) {
		want := 81
		got := day10.RateMap(trailMap)

		assert.Equal(t, want, got)
	})
}
