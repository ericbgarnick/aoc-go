package day08

import (
	"fmt"
	"github.com/ericbgarnick/aoc-go/util"
	"math"
	"regexp"
)

type node map[string]string

var nodePattern = regexp.MustCompile(`[0-9A-Z]{3}`)

type network map[string]node

// Part1 returns the number of steps taken from the starting node to the end node.
func Part1() {
	instructions, nw := loadMap()
	numSteps := traverseMapSingle(instructions, nw, "AAA", func(s string) bool {
		return s == "ZZZ"
	})
	fmt.Printf("PART 1: %d\n", numSteps)
}

// Part2 returns the number of steps required to get from all
// starting nodes to simultaneously be on all ending nodes.
func Part2() {
	instructions, nw := loadMap()
	numSteps := traverseMapMulti(instructions, nw)
	fmt.Printf("PART 2: %d\n", numSteps)
}

func loadMap() (string, network) {
	var nw = network{}
	fileLines := util.ScanFile("y2023/day08/data.txt")
	for _, line := range fileLines[2:] {
		labels := nodePattern.FindAllString(line, -1)
		newNode := node{"L": labels[1], "R": labels[2]}
		nw[labels[0]] = newNode
	}
	return fileLines[0], nw
}

// traverseMapSingle returns the number of steps taken ( > 0 ) going from startLabel until isDone returns true.
func traverseMapSingle(instructions string, nw network, startLabel string, isDone func(string) bool) int {
	var (
		numSteps       int
		curInstruction string
	)
	curNodeLabel := startLabel
	for numSteps == 0 || !isDone(curNodeLabel) {
		curInstruction = string(instructions[numSteps%len(instructions)])
		curNodeLabel = nw[curNodeLabel][curInstruction]
		numSteps += 1
	}
	return numSteps
}

// traverseMapMulti returns the number of steps taken to get from
// all nodes ending with A to arrive simultaneously at all nodes ending with Z.
func traverseMapMulti(instructions string, nw network) int {
	var (
		periods []int
	)
	for label := range nw {
		if label[2] == 'A' {
			endSteps := traverseMapSingle(instructions, nw, label, func(s string) bool {
				return s[2] == 'Z'
			})
			periods = append(periods, endSteps)
		}
	}

	return lcm(periods)
}

func lcm(values []int) int {
	var (
		combinedFactors = map[int]int{}
		product         = 1
	)
	for _, v := range values {
		factors := factorize(v)
		for f, count := range factors {
			if count > combinedFactors[f] {
				combinedFactors[f] = count
			}
		}
	}
	for f, count := range combinedFactors {
		product *= int(math.Pow(float64(f), float64(count)))
	}
	return product
}

func factorize(num int) map[int]int {
	var factors = map[int]int{}
	for candidate := 2; candidate <= num; candidate++ {
		if num%candidate == 0 {
			factors[candidate] += 1
			num /= candidate
			candidate = 2
		}
	}
	return factors
}
