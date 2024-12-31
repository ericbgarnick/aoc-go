package util

import "regexp"

var (
	IntPattern       = regexp.MustCompile(`\d+`)
	SignedIntPattern = regexp.MustCompile(`-?\d+`)
)

func AbsInt(value int) int {
	if value >= 0 {
		return value
	}
	return -1 * value
}
