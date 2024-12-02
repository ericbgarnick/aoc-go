package day07

import (
	"fmt"
	"github.com/ericbgarnick/aoc-go/util"
	"slices"
	"sort"
	"strings"
)

type camelHandRank int

const (
	highCard camelHandRank = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

const camelCardBasicOrder = "23456789TJQKA"
const camelCardJokerOrder = "J23456789TQKA"

type camelHand struct {
	cards  string
	bid    int
	rank   camelHandRank
	jokers bool
}

// rankHand sets the rank property on c, using jokers if c.jokers is true.
func (c *camelHand) rankHand() {
	counts := countCards(*c)
	if c.jokers {
		counts = useJokers(counts)
	}
	switch len(counts) {
	case 1:
		c.rank = fiveOfAKind
	case 2:
		for _, count := range counts {
			if count == 1 || count == 4 {
				c.rank = fourOfAKind
			} else {
				c.rank = fullHouse
			}
			return
		}
	case 3:
		for _, count := range counts {
			if count == 3 {
				c.rank = threeOfAKind
				return
			}
		}
		c.rank = twoPair
	case 4:
		c.rank = onePair
	default:
		c.rank = highCard
	}
}

func countCards(c camelHand) map[rune]int {
	var counts = map[rune]int{}
	for _, card := range c.cards {
		count, ok := counts[card]
		if ok {
			counts[card] = count + 1
		} else {
			counts[card] = 1
		}
	}
	return counts
}

// useJokers increases the most abundant non-joker card by the number of jokers in counts.
func useJokers(counts map[rune]int) map[rune]int {
	jokerCount, ok := counts['J']
	if ok {
		var (
			maxCount int
			maxCard  rune
		)
		delete(counts, 'J')
		for card, count := range counts {
			if count > maxCount {
				maxCount = count
				maxCard = card
			}
		}
		counts[maxCard] += jokerCount
	}
	return counts
}

func (c *camelHand) lt(other camelHand) bool {
	if c.rank == other.rank {
		var (
			thisStr  []byte
			otherStr []byte
			s1Pos    int
			s2Pos    int
		)
		for i := 0; i < len(c.cards); i++ {
			thisStr = append(thisStr, c.cards[i])
			otherStr = append(otherStr, other.cards[i])
		}
		compareResult := slices.CompareFunc(thisStr, otherStr, func(s1, s2 byte) int {
			if s1 == s2 {
				return 0
			}
			if c.jokers {
				s1Pos = strings.Index(camelCardJokerOrder, string(s1))
				s2Pos = strings.Index(camelCardJokerOrder, string(s2))
			} else {
				s1Pos = strings.Index(camelCardBasicOrder, string(s1))
				s2Pos = strings.Index(camelCardBasicOrder, string(s2))
			}
			if s1Pos < s2Pos {
				return -1
			} else {
				return 1
			}
		})
		return compareResult < 0
	}
	return c.rank < other.rank
}

// Part1 ranks hands and produces the sum of rank*bid for each hand.
func Part1() {
	hands := loadHands(false)
	sort.Slice(hands, func(i, j int) bool {
		return hands[i].lt(hands[j])
	})
	var total int
	for i, h := range hands {
		total += (i + 1) * h.bid
	}
	fmt.Printf("PART 1: %d\n", total)
}

// Part2 behaves like Part1 but uses jokers.
func Part2() {
	hands := loadHands(true)
	sort.Slice(hands, func(i, j int) bool {
		return hands[i].lt(hands[j])
	})
	var total int
	for i, h := range hands {
		total += (i + 1) * h.bid
	}
	fmt.Printf("PART 2: %d\n", total)
}

func loadHands(jokers bool) []camelHand {
	var hands []camelHand
	fileLines := util.ScanFile("y2023/day07/data.txt")
	for _, line := range fileLines {
		hands = append(hands, parseHand(line, jokers))
	}
	return hands
}

func parseHand(line string, jokers bool) camelHand {
	hand := camelHand{jokers: jokers}
	parts := strings.Split(line, " ")
	hand.cards = parts[0]
	hand.rankHand()
	hand.bid = util.MustParseInt(parts[1])
	return hand
}
