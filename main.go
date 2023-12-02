package main

import (
	"flag"
	"fmt"
	"github.com/ericbgarnick/aoc-go/day01"
	"github.com/ericbgarnick/aoc-go/day02"
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
	default:
		fmt.Printf("No solution for day %d\n", *dayNum)
	}
}
