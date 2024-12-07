package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ScanAOCDataFile(year, day int) []string {
	fileName := fmt.Sprintf("y%d/day%02d/data.txt", year, day)
	return ScanFile(fileName)
}

func ScanFile(fileName string) []string {
	readFile, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer readFile.Close()

	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, strings.Trim(scanner.Text(), "\n"))
	}
	return lines
}
