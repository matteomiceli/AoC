package main

import (
	"fmt"
	"matteomiceli/aoc/2024/utils"
	"sort"
	"strconv"
	"strings"
)

func Four() {
	var input = utils.FetchInput("4")

	directions := [][]int{
		// x, y
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	instancesOfXMAS := 0

	// part 1
	rows := utils.ReadLines(input)
	for y := range rows {
		column := rows[y]
		for x := range column {
			char := string(column[x])
			if char == "X" {
				// check all directions for the word XMAS
				for _, direction := range directions {
					if wordFoundInDirection("XMAS", x, y, rows, direction) {
						instancesOfXMAS++
					}
				}
			}
		}
	}

	fmt.Println("Instances of XMAS:", instancesOfXMAS)

	// part 2
	diagonals := [][]int{
		// x, y
		{-1, -1},
		{-1, 1},
		{1, -1},
		{1, 1},
	}
	masCross := map[string]int{}

	for y := range rows {
		column := rows[y]
		for x := range column {
			char := string(column[x])
			if char == "M" {
				for _, d := range diagonals {
					if wordFoundInDirection("MAS", x, y, rows, d) {
						// add range to list
						xRange := []int{x, x + (d[0] * 2)}
						yRange := []int{y, y + (d[1] * 2)}
						sort.Ints(xRange)
						sort.Ints(yRange)

						//convert range to key
						key := ""
						for i := range len(xRange) {
							key += strconv.Itoa(xRange[i])
							key += strconv.Itoa(yRange[i])
						}
						_, exists := masCross[key]
						if exists {
							masCross[key] += 1
						} else {
							masCross[key] = 1
						}
					}
				}
			}
		}
	}

	masCrosses := 0
	for entry := range masCross {
		if masCross[entry] > 1 {
			masCrosses++
		}
	}

	fmt.Println("Instances of MAS crossed:", masCrosses)
}

func wordFoundInDirection(word string, x int, y int, rows []string, xyMod []int) bool {
	current := ""
	for range word {
		current += string(rows[y][x])
		if !strings.Contains(word, current) {
			return false
		}

		// apply directional modifications
		x += xyMod[0]
		y += xyMod[1]
		if !isInRange(x, y, rows) {
			break
		}
	}
	if current == word {
		return true
	}
	return false
}

func isInRange(x int, y int, rows []string) bool {
	if len(rows) <= y || y < 0 {
		return false
	}
	if len(rows[y]) <= x || x < 0 {
		return false
	}
	return true
}
