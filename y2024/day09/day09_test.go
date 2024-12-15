package day09_test

import (
	"fmt"
	"testing"

	"github.com/ericbgarnick/aoc-go/y2024/day09"
	"github.com/stretchr/testify/assert"
)

func TestCalculateBlockChecksum(t *testing.T) {
	tests := map[string]int{
		"12345":               60,
		"2333133121414131402": 1928,
	}
	for diskMap, checksum := range tests {
		t.Run(fmt.Sprintf("block checksum %s: %d", diskMap, checksum), func(t *testing.T) {
			got := day09.CalculateBlockChecksum(diskMap)
			assert.Equal(t, checksum, got)
		})
	}
}

/*
 *  blockIdx 0   0   1   1   1   2   2   2   3   3   3   4   5   5   5   6   6   6   7   8   8   9   10  10  10  10  11  12  12  12  12  13  14  14  14  15  16  16  16  161718  18
 *  block id 0   0   9   9   2   1   1   1   7   7   7   .   4   4   .   3   3   3   .   .   .   .   5   5   5   5   .   6   6   6   6   .   .   .   .   .   8   8   8   8   .   .
 *  expanded 0   1   2   3   4   5   6   7   8   9   10  11  12  13  14  15  16  17  18  19  20  21  22  23  24  25  26  27  28  29  30  31  32  33  34  35  36  37  38  39  40  41
 *  checksum 0   0   18  45  53  58  64  71  127 190 260     308 360     405 453 504                 614 729 849 974     1136130414781658                    1946224225462858
 */

/*
 *  blockIdx 0   1   2   3   4
 *  block id 0   2   1   .   .
 *  expanded 0   1   2   3   4
 *  checksum 0   2   4
 */

/*
 *  blockIdx 0   0   0   0   1   1   1   2   2   3   4   4   5   5   5   6   6   6   6
 *  block id 0   0   0   0   2   2   .   1   1   .   .   .   .   .   .   3   3   3   3
 *  expanded 0   1   2   3   4   5   6   7   8   9   10  11  12  13  14  15  16  17  18
 *  checksum 0   0   0   0   8   18      25  33                          78  126 177 231
 */

func TestCalculateFileChecksum(t *testing.T) {
	tests := map[string]int{
		//"54321":               31,   // 00000....111..2 -> 000002111......
		//"2333133121414131402": 2858, // 00...111...2...333.44.5555.6666.777.888899 -> 00992111777.44.333....5555.6666.....8888..
		//"11111":               4,    // 0.1.2 -> 021..
		//"1234321":             57,   // 0..111....222..3 -> 03.111222.......
		"4321234": 231, // 0000...11.22...3333 -> 000022.11......3333
	}
	for diskMap, checksum := range tests {
		t.Run(fmt.Sprintf("file checksum %s: %d", diskMap, checksum), func(t *testing.T) {
			got := day09.CalculateFileChecksum(diskMap)
			assert.Equal(t, checksum, got)
		})
	}
}
