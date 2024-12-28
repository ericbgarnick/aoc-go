package day12

import (
	"fmt"

	"github.com/ericbgarnick/aoc-go/util"
)

func Part1() {
	garden := util.ScanAOCDataFile(2024, 12)
	fmt.Printf("Part 1: %d\n", PriceFencing(garden))
}

func Part2() {
	fmt.Printf("Part 2: %s\n", "TODO")
}

type Position struct {
	row, col int
}

type Region struct {
	cropType  uint8
	positions map[Position]bool
	perimeter int
}

func NewRegion() *Region {
	r := Region{}
	r.positions = make(map[Position]bool)
	return &r
}

func PriceFencing(garden []string) int {
	crops := findRegions(garden)
	var fencePrice int
	for _, regions := range crops {
		for _, region := range regions {
			fencePrice += len(region.positions) * region.perimeter
		}
	}
	return fencePrice
}

func findRegions(garden []string) map[uint8][]*Region {
	var crops = make(map[uint8][]*Region)
	for r, row := range garden {
		for c := range row {
			cropType := row[c]
			if regions, ok := crops[cropType]; ok {
				var mapped bool
				for _, region := range regions {
					if region.positions[Position{r, c}] {
						mapped = true
						break
					}
				}
				if mapped {
					continue
				}
			}
			newRegion := buildRegion(garden, Position{r, c})
			crops[cropType] = append(crops[cropType], newRegion)
		}
	}
	return crops
}

func buildRegion(garden []string, start Position) *Region {
	region := NewRegion()
	region.cropType = garden[start.row][start.col]
	var toMap = []*Position{&start}
	regionHelper(garden, region, toMap)
	return region
}

func regionHelper(garden []string, region *Region, toMap []*Position) {
	if len(toMap) == 0 {
		return
	}
	nextPos := toMap[0]
	toMap = toMap[1:]
	region.positions[*nextPos] = true
	neighbors := getNeighbors(nextPos)
	for _, n := range neighbors {
		var neighborCrop uint8
		if n.row < 0 || n.row >= len(garden) || n.col < 0 || n.col >= len(garden[0]) {
			neighborCrop = '!'
		} else {
			neighborCrop = garden[n.row][n.col]
		}
		if _, ok := region.positions[*n]; !ok && neighborCrop == region.cropType {
			region.positions[*n] = true
			toMap = append(toMap, n)
		} else if neighborCrop != region.cropType {
			region.perimeter++
		}
	}
	regionHelper(garden, region, toMap)
}

func getNeighbors(p *Position) []*Position {
	return []*Position{
		{p.row - 1, p.col},
		{p.row, p.col - 1},
		{p.row + 1, p.col},
		{p.row, p.col + 1},
	}
}
