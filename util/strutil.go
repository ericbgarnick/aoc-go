package util

import "strconv"

func MustParseInt(numStr string) int {
	numInt, err := strconv.Atoi(numStr)
	if err != nil {
		panic(err)
	}
	return numInt
}
