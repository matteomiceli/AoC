package main

import (
	"fmt"
	"math"
	"matteomiceli/aoc/2024/utils"
	"strconv"
	"strings"
)

func Two() {
	var input = utils.FetchInput("2")

	safeLevels := 0

	levels := utils.ReadLines(input)
	for _, level := range levels {
		reports := strings.Split(level, " ")

		if checkSafeReport(reports) {
			safeLevels++
		} else {
			// part 2
			// brute force all options (gross)
			for i := range reports {
				// reports without i
				newReports := []string{}
				if i == 0 {
					newReports = reports[i+1:]
				} else if i == len(reports)-1 {
					newReports = reports[:len(reports)-1]
				} else {
					newReports = append(newReports, reports[:i]...)
					newReports = append(newReports, reports[i+1:]...)
				}

				if checkSafeReport(newReports) {
					safeLevels++
					break
				}
			}

		}

	}
	fmt.Println("safe levels:", safeLevels)
}

func checkSafeReport(reports []string) bool {
	safe := true
	increasing := 0
	decreasing := 0

	for i := range reports {
		currReport, _ := strconv.Atoi(reports[i])
		// does a next index exist?
		if i+1 > len(reports)-1 {
			continue
		}
		nxtReport, _ := strconv.Atoi(reports[i+1])

		diff := nxtReport - currReport
		if math.Abs(float64(diff)) < 1 || math.Abs(float64(diff)) > 3 {
			safe = false
		}

		// check if decreasing, increasing
		if currReport == nxtReport {
			safe = false
		}
		if currReport < nxtReport {
			increasing++
		}
		if currReport > nxtReport {
			decreasing++
		}
	}

	if increasing > 0 && decreasing > 0 {
		safe = false
	}

	return safe
}
