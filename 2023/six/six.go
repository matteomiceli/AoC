package six

import (
	"fmt"
	"matteomiceli/aoc/utils"
	"strconv"
	"strings"
)

func Run() {
	lines := utils.ReadLines("six/data.txt")
	races := parseRaceData(lines)

	// part 1
	fmt.Println(raceResults(races))

	// part 2
	race := parseOneRace(lines)
	fmt.Println(raceResults(race))
}

// returns: margin of error, number of ways to beat record
func raceResults(races [][]int) (int, int) {
	beatRecord := []int{}

	for _, r := range races {
		beats := 0
		t := r[0]
		d := r[1]

		mpsMap := generateMpsMap(t)
		// check how many stategies beat record
		for _, v := range mpsMap {
			if v > d {
				beats++
			}
		}
		beatRecord = append(beatRecord, beats)
	}

	return utils.Multiply(beatRecord), utils.Sum(beatRecord)
}

func generateMpsMap(largest int) map[int]int {
	mpsMap := map[int]int{}

	// button hold down times
	for i := 0; i <= largest; i++ {
		mpsMap[i] = i * (largest - i)
	}

	return mpsMap
}

func parseRaceData(lines []string) [][]int {
	races := [][]int{}

	tString := strings.Split(lines[0], "Time:")[1]
	dString := strings.Split(lines[1], "Distance:")[1]

	times := removeEmptyEntriesAndConvert(strings.Split(tString, " "))
	distances := removeEmptyEntriesAndConvert(strings.Split(dString, " "))

	for i := 0; i < len(times); i++ {
		intTime := times[i]
		intDistance := distances[i]
		races = append(races, []int{intTime, intDistance})
	}

	return races
}

func parseOneRace(lines []string) [][]int {
	race := [][]int{}

	tString := strings.Split(lines[0], "Time:")[1]
	dString := strings.Split(lines[1], "Distance:")[1]

	time := strings.Join(strings.Split(tString, " "), "")
	distance := strings.Join(strings.Split(dString, " "), "")

	intTime, _ := strconv.Atoi(time)
	intDistance, _ := strconv.Atoi(distance)
	race = append(race, []int{intTime, intDistance})

	return race
}

func removeEmptyEntriesAndConvert(slice []string) []int {
	ints := []int{}
	for _, s := range slice {
		if s == "" {
			continue
		}
		int, _ := strconv.Atoi(s)
		ints = append(ints, int)
	}
	return ints
}
