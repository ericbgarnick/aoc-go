package day01

import (
	"bufio"
	"fmt"
	"github.com/ericbgarnick/aoc-go/util"
	"os"
	"regexp"
	"sort"
	"strconv"
)

// Part1 sums the different between each pair of values
// after sorting the two input lists.
func Part1() {
	total := solution(sumDifferences)
	fmt.Printf("PART 1: %d\n", total)
}

// Part2 sums the product of values in list 1
// multiplied by the number of times that value occurs in list 2.
func Part2() {
	total := solution(findSimilarity)
	fmt.Printf("PART 2: %d\n", total)
}

func solution(f func(scanner *bufio.Scanner) int) int {
	readFile, err := os.Open("y2024/day01/data.txt")
	if err != nil {
		panic(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	return f(fileScanner)
}

func sumDifferences(fileScanner *bufio.Scanner) int {
	var l1, l2 []int
	for fileScanner.Scan() {
		pair := fileScanner.Text()
		values := util.IntPattern.FindAllString(pair, -1)
		v1 := util.MustParseInt(values[0])
		l1 = append(l1, v1)
		v2 := util.MustParseInt(values[1])
		l2 = append(l2, v2)
	}
	sort.Ints(l1)
	sort.Ints(l2)
	var total int
	for i, v1 := range l1 {
		v2 := l2[i]
		total += util.AbsInt(v2 - v1)
	}
	return total
}

func findSimilarity(fileScanner *bufio.Scanner) int {
	numPattern := regexp.MustCompile(`\d+`)
	var (
		m1 = make(map[int]int)
		m2 = make(map[int]int)
	)
	for fileScanner.Scan() {
		pair := fileScanner.Text()
		values := numPattern.FindAllString(pair, -1)
		v1, err := strconv.Atoi(values[0])
		if err != nil {
			panic(err)
		}
		m1[v1] += 1
		v2, err := strconv.Atoi(values[1])
		if err != nil {
			panic(err)
		}
		m2[v2] += 1
	}
	var total int
	for v := range m1 {
		c := m2[v]
		total += v * c
	}
	return total
}
