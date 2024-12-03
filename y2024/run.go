package y2024

import (
	"fmt"

	"github.com/ericbgarnick/aoc-go/y2024/day01"
	"github.com/ericbgarnick/aoc-go/y2024/day02"
)

func Run(dayNum int) {
	switch dayNum {
	case 1:
		fmt.Println("Day 1")
		day01.Part1()
		day01.Part2()
	case 2:
		fmt.Println("Day 2")
		day02.Part1()
		day02.Part2()
	default:
		fmt.Printf("No solution for day %d\n", dayNum)
	}
}
