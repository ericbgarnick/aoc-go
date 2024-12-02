package day03

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type schematic [][]rune
type numSpan struct {
	start int
	end   int
}

var s schematic

func init() {
	s = loadSchematic()
}

// Part1 sums the part numbers in the schematic.
func Part1() {
	fmt.Printf("PART 1: %d\n", sumPartNums())
}

// Part2 sums the gear ratios for all gears in the schematic.
func Part2() {
	fmt.Printf("PART 2: %d\n", sumGearRatios())
}

func loadSchematic() schematic {
	readFile, err := os.Open("y2023/day03/data.txt")
	if err != nil {
		panic(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var newSchematic schematic
	for fileScanner.Scan() {
		var newRow []rune
		for _, symbol := range strings.Trim(fileScanner.Text(), "\n") {
			newRow = append(newRow, symbol)
		}
		newSchematic = append(newSchematic, newRow)
	}
	return newSchematic
}

// sumPartNums returns the sum of all numbers in the
// schematic that are adjacent to any non-period symbol.
func sumPartNums() int {
	var symbol rune
	total := 0
	for r, row := range s {
		for c := 0; c < len(row); c++ {
			symbol = row[c]
			if isDigit(symbol) {
				var number []rune
				for i := c; i < len(row) && isDigit(row[i]); i++ {
					number = append(number, row[i])
				}
				if hasAdjacentSymbol(r, c, len(number)) {
					total += mustParseInt(string(number))
				}
				c += len(number)
			}
		}
	}
	return total
}

// hasAdjacentSymbol returns true if there is a non-period symbol vertically, horizontally
// or diagonally adjacent to the number appearing at the given position in the schematic.
func hasAdjacentSymbol(numRow, numStartIdx, numLength int) bool {
	checkRowStartIdx := numStartIdx - 1
	if numStartIdx == 0 {
		checkRowStartIdx += 1
	}
	checkRowEndIdx := numStartIdx + numLength
	if checkRowEndIdx == len(s[0]) {
		checkRowEndIdx -= 1
	}
	if numRow > 0 {
		if segmentHasPartSymbol(s[numRow-1][checkRowStartIdx : checkRowEndIdx+1]) {
			return true
		}
	}
	if numRow < len(s)-1 {
		if segmentHasPartSymbol(s[numRow+1][checkRowStartIdx : checkRowEndIdx+1]) {
			return true
		}
	}
	if checkRowStartIdx < numStartIdx && isPartSymbol(s[numRow][checkRowStartIdx]) {
		return true
	}
	if checkRowEndIdx == numStartIdx+numLength && isPartSymbol(s[numRow][numStartIdx+numLength]) {
		return true
	}
	return false
}

func segmentHasPartSymbol(segment []rune) bool {
	for _, symbol := range segment {
		if isPartSymbol(symbol) {
			return true
		}
	}
	return false
}

// sumGearRatios returns the sum of all gear ratios in the schematic.
func sumGearRatios() int {
	total := 0
	for r, row := range s {
		for c := 0; c < len(row); c++ {
			if row[c] == '*' {
				adjacentPartNums := findAdjacentPartNums(r, c)
				if len(adjacentPartNums) == 2 {
					total += adjacentPartNums[0] * adjacentPartNums[1]
				}
			}
		}
	}
	return total
}

// findAdjacentPartNums returns a slice of all part numbers that are vertically,
// horizontally or diagonally adjacent to the position at row r, column c.
func findAdjacentPartNums(r, c int) []int {
	var (
		partNums []int
		ns       numSpan
	)
	if r > 0 {
		partNums = append(partNums, partNumsForRow(r-1, c)...)
	}
	if r < len(s)-1 {
		partNums = append(partNums, partNumsForRow(r+1, c)...)
	}
	if isDigit(s[r][c+1]) {
		ns = findNumSpan(r, c+1)
		partNums = append(partNums, mustParseInt(string(s[r][ns.start:ns.end+1])))
	}
	if isDigit(s[r][c-1]) {
		ns = findNumSpan(r, c-1)
		partNums = append(partNums, mustParseInt(string(s[r][ns.start:ns.end+1])))
	}
	return partNums
}

// partNumsForRow returns part numbers that
// overlap columns c-1, c, c+1 for the given row r.
func partNumsForRow(r, c int) []int {
	var (
		partNums []int
		ns       numSpan
	)
	cIdx := c - 1
	if cIdx < 0 {
		cIdx = c
	}
	if isDigit(s[r][cIdx]) {
		ns = findNumSpan(r, cIdx)
		partNums = append(partNums, mustParseInt(string(s[r][ns.start:ns.end+1])))
		cIdx = ns.end + 2
	} else if isDigit(s[r][cIdx+1]) {
		ns = findNumSpan(r, cIdx+1)
		partNums = append(partNums, mustParseInt(string(s[r][ns.start:ns.end+1])))
		cIdx = ns.end + 2
	} else {
		cIdx += 2
	}
	if cIdx == c+1 && isDigit(s[r][cIdx]) {
		ns = findNumSpan(r, cIdx)
		partNums = append(partNums, mustParseInt(string(s[r][ns.start:ns.end+1])))
	}
	return partNums
}

// findNumSpan returns a numSpan indicating the first and last index
// of the number overlapping the position at row r and column c
func findNumSpan(r, c int) numSpan {
	var ns = numSpan{start: c, end: c}
	for i := c + 1; i < len(s[r]) && isDigit(s[r][i]); i++ {
		ns.end += 1
	}
	for i := c - 1; i >= 0 && isDigit(s[r][i]); i-- {
		ns.start -= 1
	}
	return ns
}

func isDigit(symbol rune) bool {
	return symbol >= '0' && symbol <= '9'
}

func isPartSymbol(symbol rune) bool {
	return !(isDigit(symbol) || symbol == '.')
}

func mustParseInt(numStr string) int {
	numInt, err := strconv.Atoi(numStr)
	if err != nil {
		panic(err)
	}
	return numInt
}
