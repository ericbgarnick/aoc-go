package day04

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
)

type ScratchCard struct {
	winningNumbers     []string
	playerNumbers      []string
	resultingCardCount int
}

var (
	numbersPattern        = regexp.MustCompile(`\d+`)
	winningNumbersPattern = regexp.MustCompile(`: ([\d ]+) \|`)
	playerNumbersPattern  = regexp.MustCompile(`\| ([\d ]+)$`)
)

// Part1 does something.
func Part1() {
	cards := readCards()
	fmt.Printf("PART 1: %d\n", sumCardValues(cards))
}

// Part2 sums the gear ratios for all gears in the schematic.
func Part2() {
	cards := readCards()
	fmt.Printf("PART 2: %d\n", sumCardCounts(cards))
}

func readCards() []ScratchCard {
	var cards []ScratchCard
	readFile, err := os.Open("day04/data.txt")
	if err != nil {
		panic(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		cards = append(cards, parseCard(fileScanner.Text()))
	}
	return cards
}

// sumCardValues returns the sum of values for each ScratchCard, where a card's value is
// 2^x such that x is the number of playerNumbers that match winningNumbers.
func sumCardValues(cards []ScratchCard) int {
	var total int
	for _, sc := range cards {
		matches := sumCardMatches(sc)
		if matches == 0 {
			continue
		}
		total += int(math.Pow(2, float64(matches-1)))
	}
	return total
}

// sumCardCounts returns the total number of cards the given collection of ScratchCards is worth.
func sumCardCounts(cards []ScratchCard) int {
	var total int
	memo := memoizeCardCounts(cards)
	for _, cardCount := range memo {
		total += cardCount
	}
	return total
}

// memoizeCardCounts returns an in-order list of counts indicating
// how many cards each ScratchCard is worth (including itself).
func memoizeCardCounts(cards []ScratchCard) []int {
	var (
		memo     = make([]int, len(cards))
		numCards int
	)
	for i := len(cards) - 1; i >= 0; i-- {
		numCards = 1
		moreCards := sumCardMatches(cards[i])
		for j := i + 1; moreCards > 0 && j < len(cards); j++ {
			numCards += memo[j]
			moreCards--
		}
		memo[i] = numCards
	}
	return memo
}

// sumCardMatches returns the number of playerNumbers that match winningNumbers for the given ScratchCard.
func sumCardMatches(sc ScratchCard) int {
	var matches int

	var winningNumbersSet = map[string]bool{}
	for _, wn := range sc.winningNumbers {
		winningNumbersSet[wn] = true
	}

	for _, pn := range sc.playerNumbers {
		if _, ok := winningNumbersSet[pn]; ok {
			matches += 1
		}
	}
	return matches
}

func parseCard(card string) ScratchCard {
	var sc ScratchCard
	winningNumbersStr := winningNumbersPattern.FindStringSubmatch(card)[1]
	playerNumbersStr := playerNumbersPattern.FindStringSubmatch(card)[1]

	sc.winningNumbers = numbersPattern.FindAllString(winningNumbersStr, -1)
	sc.playerNumbers = numbersPattern.FindAllString(playerNumbersStr, -1)

	return sc
}
