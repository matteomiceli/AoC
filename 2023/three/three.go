package three

import (
	"fmt"
	"matteomiceli/aoc/utils"
	"strconv"
)

func Run() {
	lines := utils.ReadLines("three/data.txt")
	validPartNums := []int{}
	maybeGears := map[string][]int{} // {"12x11y": [234, 467]} key is xy id

	for y, line := range lines {
		var currNum string
		validPart := false
		gearId := "" // empty if no gear neighbor

		for x, r := range line {
			char := string(r)

			// if we've got a numeric character start recording the num
			if isNumeric(char) {
				currNum += char

				// check symbol proximity
				has, _ := peakNeighbours(x, y, lines, isSymbol)
				if has {
					validPart = true
				}

				// check gears (part 2)
				has, id := peakNeighbours(x, y, lines, isGear)
				if has {
					gearId = id
				}

			} else {
				// if we've been parsing a num up until this point, save if valid
				if currNum != "" {
					saveValidNum(&currNum, &validPart, &validPartNums)
					saveGears(&currNum, &gearId, maybeGears)
					currNum = ""
				}
			}
			// if we get to end of line save num if exists
			if currNum != "" && x == len(line)-1 {
				saveValidNum(&currNum, &validPart, &validPartNums)
				saveGears(&currNum, &gearId, maybeGears)
				currNum = ""
			}
		}
	}

	fmt.Println("Sum of valid part nums ", utils.Sum(validPartNums))
	fmt.Println("Sum of gear ratios ", getGearRatios(maybeGears)) // part 2
}

func saveValidNum(currNum *string, validPart *bool, validPartNums *[]int) {
	intNum, _ := strconv.Atoi(*currNum)
	if *validPart {
		*validPartNums = append(*validPartNums, intNum)
		*validPart = false
	}
}

func saveGears(currNum *string, gearId *string, mbGears map[string][]int) {
	intNum, _ := strconv.Atoi(*currNum)
	if *gearId != "" {
		mbGears[*gearId] = append(mbGears[*gearId], intNum)
		*gearId = ""
	}
}

func getGearRatios(maybeGears map[string][]int) int {
	total := 0
	for _, v := range maybeGears {
		if len(v) == 2 { // eactly 2 parts connected to gear
			total += v[0] * v[1]
		}
	}
	return total
}

func peakNeighbours(x int, y int, yRange []string, checkFunc func(string) bool) (bool, string) { // returns xy id of neighbor
	sequence := []int{-1, 0, 1}

	for _, sy := range sequence {
		for _, sx := range sequence {
			if checkFunc(checkPoint(x+sx, y+sy, yRange)) {
				return true, fmt.Sprintf("%dx%dy", x+sx, y+sy)
			}
		}
	}

	return false, ""
}

func checkPoint(x int, y int, yRange []string) string {
	if x >= 0 && y >= 0 {
		xRange := yRange[y]

		maxX := len(xRange)
		maxY := len(yRange)

		if x < maxX && y < maxY {
			return string(xRange[x])
		}
	}

	return ""
}

func isNumeric(c string) bool {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	for _, num := range nums {
		if c == strconv.Itoa(num) {
			return true
		}
	}
	return false
}

func isSymbol(c string) bool {
	if c == "" {
		return false
	}

	if !isNumeric(c) && c != "." {
		return true
	}
	return false
}

func isGear(c string) bool {
	if c == "*" {
		return true
	}
	return false
}
