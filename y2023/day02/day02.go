package day02

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	gameIDPattern = regexp.MustCompile(`Game (\d+)`)
	cubeLimits    = map[string]int{"red": 12, "green": 13, "blue": 14}
)

// Part1 sums the game IDs from games having valid sets of cubes.
func Part1() {
	total := solution(possibleGameID)
	fmt.Printf("PART 1: %d\n", total)
}

// Part2 sums the power from all games where power is the product of
// the minimum possible number of cubes of each color in that game.
func Part2() {
	total := solution(gamePower)
	fmt.Printf("PART 2: %d\n", total)
}

func solution(valueFunc func(string) (int, error)) int {
	readFile, err := os.Open("y2023/day02/data.txt")
	if err != nil {
		panic(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var total int
	for fileScanner.Scan() {
		gameID, err := valueFunc(fileScanner.Text())
		if err != nil {
			panic(err)
		}
		total += gameID
	}
	return total
}

// possibleGameID returns the game ID for the given line
// if it is a possible game, otherwise returns 0.
func possibleGameID(game string) (int, error) {
	split := strings.Split(game, ": ")
	gameInfo := split[0]
	sets := split[1]
	gameID := gameIDPattern.FindStringSubmatch(gameInfo)
	for _, set := range strings.Split(sets, "; ") {
		if !isPossibleSet(set) {
			return 0, nil
		}
	}
	return strconv.Atoi(gameID[1])
}

// isPossibleSet returns true if counts for cubes of each color are below values in cubeLimits.
func isPossibleSet(set string) bool {
	for _, selection := range strings.Split(set, ", ") {
		split := strings.Split(selection, " ")
		count, err := strconv.Atoi(split[0])
		color := split[1]
		if err != nil {
			panic(err)
		}
		if count > cubeLimits[color] {
			return false
		}
	}
	return true
}

// gamePower returns the power value for a game by multiplying together
// the minimum possible number of cubes of each color for the given game.
func gamePower(game string) (int, error) {
	var minimums = &map[string]int{}
	split := strings.Split(game, ": ")
	sets := split[1]
	for _, set := range strings.Split(sets, "; ") {
		minimumsForSet(set, minimums)
	}
	product := 1
	for _, minCount := range *minimums {
		product *= minCount
	}
	return product, nil
}

// minimumsForSet updates the minimums map with the lowest
// possible count of cubes of each color for the given set.
func minimumsForSet(set string, minimums *map[string]int) {
	for _, selection := range strings.Split(set, ", ") {
		split := strings.Split(selection, " ")
		count, err := strconv.Atoi(split[0])
		color := split[1]
		if err != nil {
			panic(err)
		}
		if count > (*minimums)[color] {
			(*minimums)[color] = count
		}
	}
}
