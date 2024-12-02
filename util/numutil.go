package util

func AbsInt(value int) int {
	if value >= 0 {
		return value
	}
	return -1 * value
}
