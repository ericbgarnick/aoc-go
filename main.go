package main

import (
	"flag"
	"fmt"
	"github.com/ericbgarnick/aoc-go/day01"
	"github.com/ericbgarnick/aoc-go/day02"
	"github.com/ericbgarnick/aoc-go/day03"
	"github.com/ericbgarnick/aoc-go/day04"
	"github.com/ericbgarnick/aoc-go/day05"
	"github.com/ericbgarnick/aoc-go/day06"
	"github.com/ericbgarnick/aoc-go/day07"
	"github.com/ericbgarnick/aoc-go/day08"
	"github.com/ericbgarnick/aoc-go/day09"
)

func main() {
	dayNum := flag.Int("day", 1, "Day number to run")
	flag.Parse()

	switch *dayNum {
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
	default:
		fmt.Printf("No solution for day %d\n", *dayNum)
	}
}
