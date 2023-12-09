package day05

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

var (
	dataSectionIdPattern = regexp.MustCompile(`[a-z]+:`)
	digitsPattern        = regexp.MustCompile(`\d+`)
	mapTypesPattern      = regexp.MustCompile(`([a-z]+)-to-([a-z]+)`)
)

type seedRange struct {
	startNumber int
	length      int
}

type resource struct {
	resourceType   string
	resourceNumber int
}

type conversionRange struct {
	sourceStart      int
	destinationStart int
	length           int
}

type conversionMap struct {
	sourceType      string
	destinationType string
	ranges          []conversionRange
}

func (cm *conversionMap) convert(source resource) resource {
	var result = resource{resourceType: cm.destinationType}
	if source.resourceType != cm.sourceType {
		panic(fmt.Errorf("invalid source resource type %s", source.resourceType))
	}
	for _, r := range cm.ranges {
		if source.resourceNumber >= r.sourceStart && source.resourceNumber < r.sourceStart+r.length {
			diff := source.resourceNumber - r.sourceStart
			result.resourceNumber = r.destinationStart + diff
			return result
		}
	}
	result.resourceNumber = source.resourceNumber
	return result
}

// Part1 returns the lowest location number for all input seed values.
func Part1() {
	fmt.Printf("PART 1: %d\n", findClosestLocation(false))
}

// Part2 returns the lowest location number for all input seed ranges.
// NOTE: This takes a couple of minutes to run.
func Part2() {
	fmt.Printf("PART 2: %d\n", findClosestLocation(true))
}

func findClosestLocation(seedRanges bool) int {
	seeds, maps := readInput(seedRanges)
	var wg sync.WaitGroup

	closestLocation := math.MaxInt
	c := make(chan int, len(seeds))

	for _, sr := range seeds {
		wg.Add(1)
		go closestForSeedRange(sr, maps, c, &wg)

	}

	wg.Wait()
	close(c)

	for loc := range c {
		if loc < closestLocation {
			closestLocation = loc
		}
	}
	return closestLocation
}

func readInput(seedRanges bool) ([]seedRange, map[string]conversionMap) {
	var (
		seeds []seedRange
		maps  = map[string]conversionMap{}
		cm    conversionMap
	)
	readFile, err := os.Open("day05/data.txt")
	if err != nil {
		panic(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := strings.Trim(fileScanner.Text(), "\n")
		if len(line) == 0 {
			continue
		}
		switch dataSectionIdPattern.FindString(line) {
		case "seeds:":
			seedData := digitsPattern.FindAllString(line, -1)
			if seedRanges {
				for startNumIdx := 0; startNumIdx < len(seedData); startNumIdx += 2 {
					sr := seedRange{
						startNumber: mustParseInt(seedData[startNumIdx]),
						length:      mustParseInt(seedData[startNumIdx+1]),
					}
					seeds = append(seeds, sr)
				}
			} else {
				for _, seedNum := range seedData {
					sr := seedRange{
						startNumber: mustParseInt(seedNum),
						length:      1,
					}
					seeds = append(seeds, sr)
				}
			}
		case "map:":
			if len(cm.ranges) != 0 {
				maps[cm.sourceType] = cm
			}
			cm = conversionMap{}
			types := mapTypesPattern.FindStringSubmatch(line)
			cm.sourceType = types[1]
			cm.destinationType = types[2]
		default:
			conversionData := digitsPattern.FindAllString(line, -1)
			r := conversionRange{
				destinationStart: mustParseInt(conversionData[0]),
				sourceStart:      mustParseInt(conversionData[1]),
				length:           mustParseInt(conversionData[2]),
			}
			cm.ranges = append(cm.ranges, r)
		}
	}
	maps[cm.sourceType] = cm
	return seeds, maps
}

func closestForSeedRange(sr seedRange, maps map[string]conversionMap, c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	closestLocation := math.MaxInt
	for diff := 0; diff < sr.length; diff++ {
		newLocation := getSeedLocation(sr.startNumber+diff, maps)
		if newLocation < closestLocation {
			closestLocation = newLocation
		}
	}
	c <- closestLocation
}

func getSeedLocation(seedNumber int, conversionMaps map[string]conversionMap) int {
	r := resource{
		resourceType:   "seed",
		resourceNumber: seedNumber,
	}
	cm, ok := conversionMaps[r.resourceType]
	for ok {
		r = cm.convert(r)
		cm, ok = conversionMaps[r.resourceType]
	}
	return r.resourceNumber
}

func mustParseInt(numStr string) int {
	numInt, err := strconv.Atoi(numStr)
	if err != nil {
		panic(err)
	}
	return numInt
}
