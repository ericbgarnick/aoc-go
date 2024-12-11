package day09_test

import (
	"fmt"
	"testing"

	"github.com/ericbgarnick/aoc-go/y2024/day09"
	"github.com/stretchr/testify/assert"
)

func TestDay09(t *testing.T) {
	tests := map[string]int{
		"12345":               60,
		"2333133121414131402": 1928,
	}
	for diskMap, checksum := range tests {
		t.Run(fmt.Sprintf("%s: %d", diskMap, checksum), func(t *testing.T) {
			got := day09.CalculateChecksum(diskMap)
			assert.Equal(t, checksum, got)
		})

	}
}
