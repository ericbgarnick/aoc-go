package day02

import (
	"fmt"

	"github.com/ericbgarnick/aoc-go/util"
)

// Part1 counts the input data rows that are monotonically
// increasing or decreasing within the given min/max differences.
func Part1() {
	total := countSafeReports(false)
	fmt.Printf("PART 1: %d\n", total)
}

// Part2 counts the input data rows that are monotonically increasing or decreasing
// within the given min/max differences after removing one level value from the row
func Part2() {
	total := countSafeReports(true)
	fmt.Printf("PART 2: %d\n", total)
}

const (
	minDiff = 1
	maxDiff = 3
)

func countSafeReports(withDampener bool) int {
	data := util.ScanFile("y2024/day02/data.txt")
	var numSafe int
	for _, line := range data {
		rawReport := util.IntPattern.FindAllString(line, -1)
		var report = make([]int, len(rawReport))
		for i, level := range rawReport {
			report[i] = util.MustParseInt(level)
		}
		if withDampener {
			for i := range report {
				dampened := applyDampener(report, i)
				if isSafeReport(dampened) {
					numSafe++
					break
				}
			}
		} else {
			if isSafeReport(report) {
				numSafe++
			}
		}
	}
	return numSafe
}

func applyDampener(report []int, i int) []int {
	var (
		dampened = make([]int, len(report)-1)
		j, k     int
	)
	for j < len(report) && k < len(dampened) {
		if i == j {
			j++
		}
		dampened[k] = report[j]
		j++
		k++
	}
	return dampened
}

func isSafeReport(report []int) bool {
	prev := report[0]
	cur := report[1]
	increasing := levelIncreasing(prev, cur)
	if increasing == nil {
		return false
	}
	for i := 2; i < len(report); i++ {
		prev = report[i-1]
		cur = report[i]
		if nextStep := levelIncreasing(prev, cur); nextStep == nil || *nextStep != *increasing {
			return false
		}
	}
	return true
}

func levelIncreasing(prev, cur int) *bool {
	absDiff := util.AbsInt(cur - prev)
	if absDiff < minDiff || absDiff > maxDiff {
		return nil
	}
	increasing := cur > prev
	return &increasing
}
