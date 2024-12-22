package day12

import (
	"fmt"

	"github.com/ericbgarnick/aoc-go/util"
)

func Part1() {
	garden := util.ScanAOCDataFile(2024, 12)
	fmt.Printf("Part 1: %d\n", PriceFencing(garden, false))
}

func Part2() {
	garden := util.ScanAOCDataFile(2024, 12)
	fmt.Printf("Part 2: %d\n", PriceFencing(garden, true))
}

const NoCrop = '!'

type Position struct {
	row, col int
}

type Region struct {
	cropType                       uint8
	positions                      map[Position]bool
	antiRegions                    []*Region
	perimeter                      int
	outerSides, innerSides         int
	minRow, maxRow, minCol, maxCol int
}

func (r *Region) addPosition(newPosition Position) {
	r.positions[newPosition] = true
	if newPosition.row > r.maxRow {
		r.maxRow = newPosition.row
	}
	if newPosition.row < r.minRow {
		r.minRow = newPosition.row
	}
	if newPosition.col > r.maxCol {
		r.maxCol = newPosition.col
	}
	if newPosition.col < r.minCol {
		r.minCol = newPosition.col
	}
}

func NewRegion(garden []string, start Position) *Region {
	r := Region{}
	r.positions = make(map[Position]bool)
	r.cropType = garden[start.row][start.col]
	r.minRow = len(garden) - 1
	r.minCol = len(garden[0]) - 1
	return &r
}

func PriceFencing(garden []string, priceBySides bool) int {
	crops := findRegions(garden)
	for _, regions := range crops {
		for _, region := range regions {
			region.antiRegions = findAntiRegions(garden, region)
		}
	}

	var fencePrice int
	if priceBySides {
		for _, crop := range crops {
			for _, r := range crop {
				rMatch := func(ct uint8) bool {
					return r.cropType == ct
				}
				CountOuterSides(garden, r, rMatch)
				for _, ar := range r.antiRegions {
					arMatch := func(ct uint8) bool {
						return ct != NoCrop && ct != r.cropType
					}
					CountOuterSides(garden, ar, arMatch)
				}
			}
		}
		for _, regions := range crops {
			for _, r := range regions {
				for _, otherR := range r.antiRegions {
					if regionContains(r, otherR) {
						r.innerSides += otherR.outerSides
					}
				}
			}
		}
	}
	for _, regions := range crops {
		for _, region := range regions {
			if priceBySides {
				fencePrice += len(region.positions) * (region.outerSides + region.innerSides)
			} else {
				fencePrice += len(region.positions) * region.perimeter
			}
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
			match := func(p *Position, ct uint8) bool {
				return ct == cropType
			}
			antiMatch := func(p *Position, ct uint8) bool {
				return ct != cropType
			}
			newRegion := buildRegion(garden, Position{r, c}, match, antiMatch)
			crops[cropType] = append(crops[cropType], newRegion)
		}
	}
	return crops
}

func findAntiRegions(garden []string, region *Region) []*Region {
	var antiRegions []*Region
	for r, row := range garden {
		for c := range row {
			mapped := false
			cropType := row[c]
			if region.cropType == cropType {
				continue
			}
			for _, ar := range antiRegions {
				if ar.positions[Position{r, c}] {
					mapped = true
				}
			}
			if mapped {
				continue
			}
			match := func(p *Position, ct uint8) bool {
				return ct != NoCrop && !region.positions[*p]
			}
			antiMatch := func(p *Position, ct uint8) bool {
				return ct != NoCrop && region.positions[*p]
			}
			newRegion := buildRegion(garden, Position{r, c}, match, antiMatch)
			antiRegions = append(antiRegions, newRegion)
		}
	}
	return antiRegions
}

type cropMatch func(cropType uint8) bool

type regionMatch func(p *Position, cropType uint8) bool

func buildRegion(garden []string, start Position, match regionMatch, antiMatch regionMatch) *Region {
	region := NewRegion(garden, start)
	var toMap = []*Position{&start}
	regionHelper(garden, region, toMap, match, antiMatch)
	return region
}

func regionHelper(garden []string, region *Region, toMap []*Position, match regionMatch, antiMatch regionMatch) {
	if len(toMap) == 0 {
		return
	}
	nextPos := toMap[0]
	toMap = toMap[1:]
	region.addPosition(*nextPos)
	neighbors := getNeighbors(nextPos)
	for _, n := range neighbors {
		var neighborCrop uint8
		if n.row < 0 || n.row >= len(garden) || n.col < 0 || n.col >= len(garden[0]) {
			neighborCrop = NoCrop
		} else {
			neighborCrop = garden[n.row][n.col]
		}
		if _, ok := region.positions[*n]; !ok && match(n, neighborCrop) {
			region.addPosition(*n)
			toMap = append(toMap, n)
		} else if antiMatch(n, neighborCrop) {
			region.perimeter++
		}
	}
	regionHelper(garden, region, toMap, match, antiMatch)
}

func getNeighbors(p *Position) []*Position {
	return []*Position{
		{p.row - 1, p.col},
		{p.row, p.col - 1},
		{p.row + 1, p.col},
		{p.row, p.col + 1},
	}
}

func CountOuterSides(garden []string, region *Region, match cropMatch) {
	var start Position
	for p := range region.positions {
		if p.row >= start.row && p.col >= start.col {
			start = p
		}
	}
	var direction = 'S'
	nextPos, nextDirection := traceSides(garden, start, region.cropType, direction, match)
	if nextDirection != direction {
		region.outerSides++
		direction = nextDirection
	}
	for !(nextPos == start && direction == 'S') {
		nextPos, nextDirection = traceSides(garden, nextPos, region.cropType, direction, match)
		if nextDirection != direction {
			region.outerSides++
			direction = nextDirection
		}
	}
}

// traceSides returns the next position and next direction
// always keeping the wall on the left-hand side
func traceSides(garden []string, p Position, cropType uint8, d rune, match cropMatch) (Position, rune) {
	if d == 'N' && (p.row == 0 || !match(garden[p.row-1][p.col])) {
		return p, 'E'
	} else if d == 'N' && p.col > 0 && match(garden[p.row-1][p.col-1]) {
		return Position{p.row - 1, p.col - 1}, 'W'
	} else if d == 'N' {
		return Position{p.row - 1, p.col}, d
	}

	if d == 'E' && (p.col == len(garden[0])-1 || !match(garden[p.row][p.col+1])) {
		return p, 'S'
	} else if d == 'E' && p.row > 0 && match(garden[p.row-1][p.col+1]) {
		return Position{p.row - 1, p.col + 1}, 'N'
	} else if d == 'E' {
		return Position{p.row, p.col + 1}, d
	}

	if d == 'S' && (p.row == len(garden)-1 || !match(garden[p.row+1][p.col])) {
		return p, 'W'
	} else if d == 'S' && p.col+1 < len(garden[p.row]) && match(garden[p.row+1][p.col+1]) {
		return Position{p.row + 1, p.col + 1}, 'E'
	} else if d == 'S' {
		return Position{p.row + 1, p.col}, d
	}

	if d == 'W' && (p.col == 0 || !match(garden[p.row][p.col-1])) {
		return p, 'N'
	} else if d == 'W' && p.row+1 < len(garden) && match(garden[p.row+1][p.col-1]) {
		return Position{p.row + 1, p.col - 1}, 'S'
	} else if d == 'W' {
		return Position{p.row, p.col - 1}, d
	}

	panic(fmt.Sprintf("unknown direction %s", string(d)))
}

func regionContains(r *Region, otherR *Region) bool {
	return r.minRow < otherR.minRow &&
		r.minCol < otherR.minCol &&
		r.maxRow > otherR.maxRow &&
		r.maxCol > otherR.maxCol
}
