package day06

import (
	"fmt"
	"math"
)

type Race struct {
	Time   int
	Record int
}

var SampleRacesPt1 = []Race{
	{Time: 7, Record: 9},
	{Time: 15, Record: 40},
	{Time: 30, Record: 200},
}

var ChallengeRacesPt1 = []Race{
	{Time: 50, Record: 242},
	{Time: 74, Record: 1017},
	{Time: 86, Record: 1691},
	{Time: 85, Record: 1252},
}

var SampleRacesPt2 = []Race{
	{Time: 71530, Record: 940200},
}

var ChallengeRacesPt2 = []Race{
	{Time: 50748685, Record: 242101716911252},
}

// Part1 finds the solution for races in ChallengeRacesPt1
func Part1() {
	fmt.Printf("PART 1: %d\n", optimalChargeTimes(ChallengeRacesPt1))
}

// Part2 finds the solution for races in ChallengeRacesPt2
func Part2() {
	fmt.Printf("PART 2: %d\n", optimalChargeTimes(ChallengeRacesPt2))
}

// optimalChargeTimes returns the product of counts of possible vehicle charge times that would beat
// the existing record(s). Charge times for existing records can be derived using the quadratic formula.
func optimalChargeTimes(races []Race) int {
	var (
		lowInt  int
		highInt int
	)
	solution := 1
	for _, r := range races {
		low, high := quadraticSolution(-1, float64(r.Time), float64(-1*r.Record))
		lowInt = int(low)
		highInt = int(high)
		if float64(highInt) == high {
			highInt -= 1
		}
		solution *= highInt - lowInt
	}
	return solution
}

func quadraticSolution(a, b, c float64) (float64, float64) {
	low := (-1*b + math.Sqrt(b*b-4*a*c)) / (2 * a)
	high := (-1*b - math.Sqrt(b*b-4*a*c)) / (2 * a)
	return low, high
}
