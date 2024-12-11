package day09

import (
	"fmt"

	"github.com/ericbgarnick/aoc-go/util"
)

func Part1() {
	diskMap := util.ScanAOCDataFile(2024, 9)[0]
	fmt.Printf("Part 1: %d\n", CalculateChecksum(diskMap))
}

func Part2() {
	fmt.Printf("Part 2: %s\n", "TODO")
}

func CalculateChecksum(diskMap string) int {
	var (
		isFile                                               bool
		checksum, expandedIdx, blockIdx, expandedRevBlockIdx int
		revBlockIdx                                          = len(diskMap) - 1
		revBlockLen                                          = util.MustParseInt(diskMap[revBlockIdx : revBlockIdx+1])
	)
	for blockIdx <= revBlockIdx {
		isFile = !isFile
		blockLen := util.MustParseInt(diskMap[blockIdx : blockIdx+1])
		if isFile && blockIdx < revBlockIdx {
			blockID := blockIdx / 2
			for range blockLen {
				checksum += blockID * expandedIdx
				expandedIdx++
			}
		} else {
			for range blockLen {
				if expandedRevBlockIdx >= revBlockLen {
					expandedRevBlockIdx = 0
					revBlockIdx -= 2
					revBlockLen = util.MustParseInt(diskMap[revBlockIdx : revBlockIdx+1])
				}
				if blockIdx > revBlockIdx {
					return checksum
				}
				blockID := revBlockIdx / 2
				checksum += blockID * expandedIdx
				expandedRevBlockIdx++
				expandedIdx++
			}
		}
		blockIdx++
	}
	return checksum
}
