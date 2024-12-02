package day09

import (
	"fmt"
	"github.com/ericbgarnick/aoc-go/util"
	"regexp"
)

type direction string

const (
	fwd direction = "FWD"
	bwd direction = "BWD"
)

var digitPattern = regexp.MustCompile(`-?\d+`)

// Part1 sums the next values for the end of each line in the report.
func Part1() {
	var total int
	report := loadReport(fwd)
	for _, h := range report {
		total += extendHistory(h, fwd)
	}
	fmt.Printf("PART 1: %d\n", total)
}

// Part2 sums the previous values for the beginning of each line in the report.
func Part2() {
	var total int
	report := loadReport(bwd)
	for _, h := range report {
		total += extendHistory(h, bwd)
	}
	fmt.Printf("PART 2: %d\n", total)
}

func loadReport(d direction) [][]int {
	fileLines := util.ScanFile("y2023/day09/data.txt")
	var report [][]int
	for _, line := range fileLines {
		var historyLine []int
		for _, pt := range digitPattern.FindAllString(line, -1) {
			val := util.MustParseInt(pt)
			if d == fwd {
				historyLine = append(historyLine, val)
			} else {
				historyLine = append([]int{val}, historyLine...)
			}
		}
		report = append(report, historyLine)
	}
	return report
}

// extendHistory returns the next value in the history sequence going in the direction indicated.
func extendHistory(history []int, d direction) int {
	var (
		newValue            int
		derivedValue        int
		derivedHistory      []int
		derivedHistoryTotal int
	)
	for i := 1; i < len(history); i++ {
		if d == fwd {
			derivedValue = history[i] - history[i-1]
		} else {
			derivedValue = history[i-1] - history[i]
		}
		derivedHistoryTotal += derivedValue
		derivedHistory = append(derivedHistory, derivedValue)
	}
	if derivedHistoryTotal == 0 {
		return history[0]
	}
	if d == fwd {
		newValue = history[len(history)-1] + extendHistory(derivedHistory, d)
	} else {
		newValue = history[len(history)-1] - extendHistory(derivedHistory, d)
	}
	return newValue
}
