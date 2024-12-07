package day05

import (
	"fmt"
	"slices"
	"sort"
	"strings"

	"github.com/ericbgarnick/aoc-go/util"
)

func Part1() {
	rules, updates := parseInput()
	var result int
	for _, update := range updates {
		if isValidUpdate(update, rules) {
			middleIdx := (len(update) - 1) / 2
			result += update[middleIdx]
		}
	}
	fmt.Printf("PART 1: %d\n", result)
}

func Part2() {
	rules, updates := parseInput()
	var result int
	for _, update := range updates {
		if !isValidUpdate(update, rules) {
			update = orderUpdate(update, rules)
			middleIdx := (len(update) - 1) / 2
			result += update[middleIdx]
		}
	}
	fmt.Printf("PART 2: %d\n", result)
}

func parseInput() (map[int][]int, [][]int) {
	var (
		rawRules, rawUpdates []string
		parsingUpdates       bool
	)
	for _, line := range util.ScanAOCDataFile(2024, 5) {
		if line == "" {
			parsingUpdates = true
			continue
		}
		if parsingUpdates {
			rawUpdates = append(rawUpdates, line)
		} else {
			rawRules = append(rawRules, line)
		}
	}
	return parseRules(rawRules), parseUpdates(rawUpdates)
}

func parseRules(rawRules []string) map[int][]int {
	rules := map[int][]int{}
	for _, rule := range rawRules {
		parts := strings.Split(rule, "|")
		from := util.MustParseInt(parts[0])
		to := util.MustParseInt(parts[1])
		rules[from] = append(rules[from], to)
	}
	return rules
}

func parseUpdates(rawUpdates []string) [][]int {
	var updates [][]int
	for _, rawUpdate := range rawUpdates {
		var update []int
		for _, step := range strings.Split(rawUpdate, ",") {
			update = append(update, util.MustParseInt(step))
		}
		updates = append(updates, update)
	}
	return updates
}

func isValidUpdate(update []int, rules map[int][]int) bool {
	for i := 1; i < len(update); i++ {
		curStep := update[i-1]
		nextStep := update[i]
		validNextSteps := rules[curStep]
		if !slices.Contains(validNextSteps, nextStep) {
			return false
		}
	}
	return true
}

func orderUpdate(update []int, rules map[int][]int) []int {
	sort.Slice(update, func(i, j int) bool {
		validSteps := rules[update[i]]
		return slices.Contains(validSteps, update[j])
	})
	return update
}
