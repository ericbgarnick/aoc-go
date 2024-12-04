package day03

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ericbgarnick/aoc-go/util"
)

// Part1 sums the product of all integers inside a mul() expression
func Part1() {
	fmt.Printf("PART 1: %d\n", scanMemory(false))
}

// Part2 sums the product of all integers inside a mul() expression following a do() expression
func Part2() {
	fmt.Printf("PART 2: %d\n", scanMemory(true))
}

var (
	mulPattern  = `mul\((\d+),(\d+)\)`
	doPattern   = `do()`
	dontPattern = `don't()`
	mulRE       = regexp.MustCompile(mulPattern)
	toggleRE    = regexp.MustCompile(strings.Join([]string{
		regexp.QuoteMeta(doPattern),
		regexp.QuoteMeta(dontPattern),
		mulPattern,
	}, "|"))
)

func scanMemory(toggle bool) int {
	input := util.ScanFile("y2024/day03/data.txt")
	var (
		total   int
		execute = true
	)
	for _, line := range input {
		if toggle {
			for _, match := range toggleRE.FindAllStringSubmatch(line, -1) {
				switch match[0] {
				case doPattern:
					execute = true
				case dontPattern:
					execute = false
				default:
					if execute {
						total += util.MustParseInt(match[1]) * util.MustParseInt(match[2])
					}
				}
			}
		} else {
			for _, match := range mulRE.FindAllStringSubmatch(line, -1) {
				total += util.MustParseInt(match[1]) * util.MustParseInt(match[2])
			}
		}
	}
	return total
}
