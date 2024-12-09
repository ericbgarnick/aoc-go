package day07

import (
	"fmt"
	"slices"

	"github.com/ericbgarnick/aoc-go/util"
)

func Part1() {
	fmt.Printf("Part 1: %d\n", processData(false))
}

func Part2() {
	fmt.Printf("Part 2: %d\n", processData(true))
}

func processData(concat bool) int {
	var total int
	for _, line := range util.ScanAOCDataFile(2024, 7) {
		rawValues := util.IntPattern.FindAllString(line, -1)
		var values = make([]int, len(rawValues))
		for i, v := range rawValues {
			values[i] = util.MustParseInt(v)
		}
		if evaluate(values[1:], values[0], concat) {
			total += values[0]
		}
	}
	return total
}

func evaluate(values []int, target int, concat bool) bool {
	var (
		curTotals  = []int{values[0]}
		prevTotals []int
	)
	for i := 1; i < len(values); i++ {
		prevTotals = []int{}
		for _, t := range curTotals {
			prevTotals = append(prevTotals, t)
		}
		curTotals = []int{}
		for _, t := range prevTotals {
			if newSum := values[i] + t; newSum <= target {
				curTotals = append(curTotals, newSum)
			}
			if newProd := values[i] * t; newProd <= target {
				curTotals = append(curTotals, newProd)
			}
			if !concat {
				continue
			}
			if newConcat := concatenate(t, values[i]); newConcat <= target {
				curTotals = append(curTotals, newConcat)
			}
		}
	}
	return slices.Contains(curTotals, target)
}

func concatenate(v1, v2 int) int {
	v2Copy := float32(v2)
	for v2Copy >= 1 {
		v1 *= 10
		v2Copy /= 10
	}
	return v1 + v2
}
