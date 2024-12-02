package y2024

import (
	"fmt"
	"github.com/ericbgarnick/aoc-go/y2024/day01"
)

func Run(dayNum int) {
	switch dayNum {
	case 1:
		fmt.Println("Day 1")
		day01.Part1()
		day01.Part2()
	default:
		fmt.Printf("No solution for day %d\n", dayNum)
	}
}
