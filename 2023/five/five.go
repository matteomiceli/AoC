package five

import (
	"fmt"
	"matteomiceli/aoc/utils"
	"strconv"
	"strings"
)

func Run() {
	lines := utils.ReadLines("five/data.txt")
	seeds, maps := parseAlmanac(lines, parseSeeds)

	// part 2
	rngSeeds := parseSeedRange(seeds)
	fmt.Println(rngSeeds)

	fmt.Println("Part 1 ----")
	findClosestSeedLocation(seeds, maps)
	fmt.Println("Part 2 ----")
	findLowestLocationFromSeedRange(rngSeeds, maps)
}

func findClosestSeedLocation(seeds []int, maps Map) {
	closestLocation := 0

	for i, s := range seeds {
		// lookup chain
		soil := lookUpDestination(s, maps["seed-to-soil"])
		fertilizer := lookUpDestination(soil, maps["soil-to-fertilizer"])
		water := lookUpDestination(fertilizer, maps["fertilizer-to-water"])
		light := lookUpDestination(water, maps["water-to-light"])
		temperature := lookUpDestination(light, maps["light-to-temperature"])
		humidity := lookUpDestination(temperature, maps["temperature-to-humidity"])
		location := lookUpDestination(humidity, maps["humidity-to-location"])

		if location < closestLocation || i == 0 {
			closestLocation = location
		}
	}

	fmt.Println(closestLocation)
}

// walk backward from lowest location value
func findLowestLocationFromSeedRange(seedRng []Range, maps Map) {
	lowest := 0
	// find lowest location val
	for i, v := range maps["humidity-to-location"] {
		if i == 0 {
			lowest = v.destination.min
			continue
		}

		if v.destination.min < lowest {
			lowest = v.destination.min
		}

		// iterate from lowest and track lowest seed
	}
}

func parseAlmanac(lines []string, parseSeedFunc func(string) []int) ([]int, Map) {
	seeds := []int{}
	maps := Map{}

	currMapName := ""

	for _, line := range lines {
		// ignore empty lines
		if line == "" {
			continue
		}

		// parse seeds
		if strings.Contains(line, "seeds:") {
			seeds = parseSeedFunc(line)
			continue
		}

		// parse map titles
		if !utils.IsNumeric(string(line[0])) {
			title := strings.Split(line, " map:")[0]
			currMapName = title
			maps[title] = []MapEntry{}
			continue
		}

		// parse numbers
		if utils.IsNumeric(string(line[0])) {
			maps[currMapName] = append(maps[currMapName], parseMapEntry(line))
		}

	}
	return seeds, maps
}

// part 1
func parseSeeds(seedString string) []int {
	seeds := []int{}
	splitNums := strings.Trim(strings.Split(seedString, ":")[1], " ")
	stringNums := strings.Split(splitNums, " ")
	for _, v := range stringNums {
		n, _ := strconv.Atoi(v)
		seeds = append(seeds, n)
	}
	return seeds
}

// part 2
func parseSeedRange(seedInts []int) []Range {
	seeds := []Range{}
	for i := 0; i < len(seedInts); i += 2 {
		start := seedInts[i]
		rng := seedInts[i+1]

		seedRng := Range{}
		seedRng.min = start
		seedRng.max = start + rng - 1
		seeds = append(seeds, seedRng)
	}
	return seeds
}

// parses the map number ranges
func parseMapEntry(line string) MapEntry {
	entry := MapEntry{}

	vals := strings.Split(line, " ")
	intVals := convertStringsToInt(vals)

	entry.destination.min = intVals[0]
	entry.destination.max = intVals[0] + intVals[2] - 1 // -1 because inclusive
	entry.source.min = intVals[1]
	entry.source.max = intVals[1] + intVals[2] - 1 // -1 because inclusive

	return entry
}

func convertStringsToInt(strings []string) []int {
	ints := []int{}
	for _, s := range strings {
		n, _ := strconv.Atoi(s)
		ints = append(ints, n)
	}
	return ints
}

type Map = map[string][]MapEntry

type MapEntry struct {
	source      Range
	destination Range
}

type Range struct {
	min int
	max int
}

func lookUpDestination(source int, entry []MapEntry) int {
	for _, m := range entry {
		offset := 0
		// in range
		if m.source.min <= source && source <= m.source.max {
			offset = source - m.source.min

			return m.destination.min + offset
		}
	}
	return source
}

func lookUpSource(destination int, entry []MapEntry) int {
	for _, m := range entry {
		offset := 0
		// in range
		if m.destination.min <= destination && destination <= m.destination.max {
			offset = destination - m.destination.min

			return m.source.min + offset
		}
	}
	return destination
}
