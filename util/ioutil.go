package util

import (
	"bufio"
	"os"
	"strings"
)

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
