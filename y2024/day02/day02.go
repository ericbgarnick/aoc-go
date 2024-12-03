package day02

import (
	"fmt"
	"github.com/ericbgarnick/aoc-go/util"
)

// Part1 sums the different between each pair of values
// after sorting the two input lists.
func Part1() {
	total := countSafeReports()
	fmt.Printf("PART 1: %d\n", total)
}

const (
	minDiff = 1
	maxDiff = 3
)

func countSafeReports() int {
	data := util.ScanFile("y2024/day02/data.txt")
	var numSafe int
	for _, line := range data {
		safe := true
		values := util.IntPattern.FindAllString(line, -1)
		prev := util.MustParseInt(values[0])
		cur := util.MustParseInt(values[1])
		increasing := levelIncreasing(prev, cur)
		if increasing == nil {
			continue
		}
		for i := 2; i < len(values); i++ {
			prev = util.MustParseInt(values[i-1])
			cur = util.MustParseInt(values[i])
			if nextStep := levelIncreasing(prev, cur); nextStep == nil || *nextStep != *increasing {
				safe = false
			}
		}
		if safe {
			numSafe++
		}
	}
	return numSafe
}

func levelIncreasing(prev, cur int) *bool {
	absDiff := util.AbsInt(cur - prev)
	if absDiff < minDiff || absDiff > maxDiff {
		return nil
	}
	increasing := cur > prev
	return &increasing
}
