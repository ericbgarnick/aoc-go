package day09

import (
	"fmt"

	"github.com/ericbgarnick/aoc-go/util"
)

func Part1() {
	diskMap := util.ScanAOCDataFile(2024, 9)[0]
	fmt.Printf("Part 1: %d\n", CalculateBlockChecksum(diskMap))
}

func Part2() {
	diskMap := util.ScanAOCDataFile(2024, 9)[0]
	fmt.Printf("Part 2: %d\n", CalculateFileChecksum(diskMap))
}

func CalculateBlockChecksum(diskMap string) int {
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

func CalculateFileChecksum(diskMap string) int {
	var (
		isFile                          bool
		checksum, expandedIdx, blockIdx int
		revBlockIdx                     = len(diskMap) - 1
		revBlockLen                     = util.MustParseInt(diskMap[revBlockIdx : revBlockIdx+1])
		movedIdx                        = make(map[int]bool)
	)
	for blockIdx < len(diskMap) {
		isFile = !isFile
		blockLen := util.MustParseInt(diskMap[blockIdx : blockIdx+1])
		if isFile && !movedIdx[blockIdx] {
			blockID := blockIdx / 2
			for range blockLen {
				checksum += blockID * expandedIdx
				expandedIdx++
			}
		} else if !movedIdx[blockIdx] {
			revBlockIdx = len(diskMap) - 1
			revBlockLen = util.MustParseInt(diskMap[revBlockIdx : revBlockIdx+1])
			for blockLen > 0 {
				for movedIdx[revBlockIdx] || (revBlockIdx > blockIdx && revBlockLen > blockLen) {
					revBlockIdx -= 2
					revBlockLen = util.MustParseInt(diskMap[revBlockIdx : revBlockIdx+1])
				}
				if revBlockLen <= blockLen && revBlockIdx > blockIdx {
					blockID := revBlockIdx / 2
					for range revBlockLen {
						checksum += blockID * expandedIdx
						expandedIdx++
					}
					movedIdx[revBlockIdx] = true
					blockLen -= revBlockLen
					revBlockIdx -= 2
					revBlockLen = util.MustParseInt(diskMap[revBlockIdx : revBlockIdx+1])
				} else {
					expandedIdx += blockLen
					blockLen = 0
				}
			}
		} else {
			expandedIdx += blockLen
		}
		blockIdx++
	}
	return checksum
}
