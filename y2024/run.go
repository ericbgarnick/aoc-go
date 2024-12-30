package y2024

import (
	"fmt"

	"github.com/ericbgarnick/aoc-go/y2024/day01"
	"github.com/ericbgarnick/aoc-go/y2024/day02"
	"github.com/ericbgarnick/aoc-go/y2024/day03"
	"github.com/ericbgarnick/aoc-go/y2024/day04"
	"github.com/ericbgarnick/aoc-go/y2024/day05"
	"github.com/ericbgarnick/aoc-go/y2024/day06"
	"github.com/ericbgarnick/aoc-go/y2024/day07"
	"github.com/ericbgarnick/aoc-go/y2024/day08"
	"github.com/ericbgarnick/aoc-go/y2024/day09"
	"github.com/ericbgarnick/aoc-go/y2024/day10"
	"github.com/ericbgarnick/aoc-go/y2024/day11"
	"github.com/ericbgarnick/aoc-go/y2024/day12"
	"github.com/ericbgarnick/aoc-go/y2024/day13"
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
	case 3:
		fmt.Println("Day 3")
		day03.Part1()
		day03.Part2()
	case 4:
		fmt.Println("Day 4")
		day04.Part1()
		day04.Part2()
	case 5:
		fmt.Println("Day 5")
		day05.Part1()
		day05.Part2()
	case 6:
		fmt.Println("Day 6")
		day06.Part1()
		day06.Part2()
	case 7:
		fmt.Println("Day 7")
		day07.Part1()
		day07.Part2()
	case 8:
		fmt.Println("Day 8")
		day08.Part1()
		day08.Part2()
	case 9:
		fmt.Println("Day 9")
		day09.Part1()
		day09.Part2()
	case 10:
		fmt.Println("Day 10")
		day10.Part1()
		day10.Part2()
	case 11:
		fmt.Println("Day 11")
		day11.Part1()
		day11.Part2()
	case 12:
		fmt.Println("Day 12")
		day12.Part1()
		day12.Part2()
	case 13:
		fmt.Println("Day 13")
		day13.Part1()
		day13.Part2()
	default:
		fmt.Printf("No solution for day %d\n", dayNum)
	}
}
