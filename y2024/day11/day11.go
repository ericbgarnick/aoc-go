package day11

import (
	"fmt"

	"github.com/ericbgarnick/aoc-go/util"
)

var (
	stones = []int{0, 4, 4979, 24, 4356119, 914, 85734, 698829}
	memo   = make(map[int][]int)
)

func Part1() {
	memo = make(map[int][]int)
	var total int
	for _, s := range stones {
		total += Expand(s, 25, 25)
	}
	fmt.Printf("Part 1: %d\n", total)
}

func Part2() {
	memo = make(map[int][]int)
	var total int
	for _, s := range stones {
		total += Expand(s, 75, 75)
	}
	fmt.Printf("Part 2: %d\n", total)
}

func Expand(stone int, totalBlinks, blinksRemaining int) int {
	if blinksRemaining == 0 {
		return 1
	}

	if result, ok := memo[stone]; ok {
		if total := result[blinksRemaining-1]; total > 0 {
			return total
		}
	} else {
		memo[stone] = make([]int, totalBlinks)
	}

	var total int
	for _, s := range Blink(stone) {
		total += Expand(s, totalBlinks, blinksRemaining-1)
	}
	memo[stone][blinksRemaining-1] = total

	return total
}

func Blink(stone int) []int {
	if stone == 0 {
		return []int{1}
	}
	if stoneStr := fmt.Sprintf("%d", stone); len(stoneStr)%2 == 0 {
		return []int{
			util.MustParseInt(stoneStr[:len(stoneStr)/2]),
			util.MustParseInt(stoneStr[len(stoneStr)/2:]),
		}
	}
	return []int{stone * 2024}
}
