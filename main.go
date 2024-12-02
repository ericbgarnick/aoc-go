package main

import (
	"flag"
	"github.com/ericbgarnick/aoc-go/y2023"
	"github.com/ericbgarnick/aoc-go/y2024"
)

func main() {
	year := flag.Int("year", 2024, "Advent of Code challenge year")
	dayNum := flag.Int("day", 1, "Day number to run")
	flag.Parse()

	switch *year {
	case 2023:
		y2023.Run(*dayNum)
	case 2024:
		y2024.Run(*dayNum)
	}
}
