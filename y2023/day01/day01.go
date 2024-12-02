package day01

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// Part1 sums the 2-digit number formed from the first and last digit
// from each line in the input data.
func Part1() {
	total := solution(findCalibrationValueDigitsOnly)
	fmt.Printf("PART 1: %d\n", total)
}

// Part2 sums the 2-digit number formed from the first and last
// digit or number word from each line in the input data.
func Part2() {
	total := solution(findCalibrationValueWithWords)
	fmt.Printf("PART 2: %d\n", total)
}

func solution(searchFunc func(string) (int, error)) int {
	readFile, err := os.Open("y2023/day01/data.txt")
	if err != nil {
		panic(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var total int
	for fileScanner.Scan() {
		cVal, err := searchFunc(fileScanner.Text())
		if err != nil {
			panic(err)
		}
		total += cVal
	}
	return total
}

func findCalibrationValueDigitsOnly(line string) (int, error) {
	digitPattern := regexp.MustCompile(`\d`)
	firstDigit := digitPattern.FindString(line)
	lastDigit := digitPattern.FindString(reverseString(line))
	return strconv.Atoi(fmt.Sprintf("%s%s", firstDigit, lastDigit))
}

// findCalibrationValueWithWords returns a 2-digit number formed by the first and last digit line,
// using numeric digits and word numbers.
//
// Note: number words may overlap e.g. "sevenine", "oneight", etc.
func findCalibrationValueWithWords(line string) (int, error) {
	rawPattern := `one|two|three|four|five|six|seven|eight|nine`
	frontPattern := regexp.MustCompile(rawPattern + `|\d`)
	backPattern := regexp.MustCompile(reverseString(rawPattern) + `|\d`)
	firstDigit := frontPattern.FindString(line)
	lastDigit := reverseString(backPattern.FindString(reverseString(line)))
	return strconv.Atoi(fmt.Sprintf("%s%s", parseValue(firstDigit), parseValue(lastDigit)))
}

func reverseString(original string) string {
	var reversed string
	for _, v := range original {
		reversed = string(v) + reversed
	}
	return reversed
}

func parseValue(rawValue string) string {
	switch rawValue {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	case "zero":
		return "0"
	default:
		return rawValue
	}
}
