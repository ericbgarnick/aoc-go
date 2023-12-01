package main

import (
	"flag"
	"fmt"
	"github.com/ericbgarnick/aoc-go/day01"
)

func main() {
	dayNum := flag.Int("day", 1, "Day number to run")
	flag.Parse()

	switch *dayNum {
	case 1:
		fmt.Println("Day 1")
		day01.Part1()
		day01.Part2()
	}
}
