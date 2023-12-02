package two

import (
	"fmt"
	"matteomiceli/aoc/utils"
	"strconv"
	"strings"
)

func Run() {
	lines := utils.ReadLines("two/data.txt")

	validGameIds := []int{}
	powers := []int{}

	for i, gameString := range lines {
		if gameString == "" {
			continue
		}

		if validGame(gameString) {
			validGameIds = append(validGameIds, i+1)
		}

		r, g, b := fewestCubes(gameString)
		powers = append(powers, (r * g * b))
	}

	fmt.Println("Sum of valid games ", sum(validGameIds))
	fmt.Println("Summed powers of least required cubes ", sum(powers))
}

// part 1
func validGame(g string) bool {
	maxCubes := map[string]int{"red": 12, "green": 13, "blue": 14}

	gameData := strings.Split(g, ":")[1]
	handFulls := strings.Split(gameData, ";")

	for _, handFull := range handFulls {
		for _, cube := range strings.Split(handFull, ",") {
			parts := strings.Split(strings.Trim(cube, " "), " ")
			num, _ := strconv.Atoi(parts[0])
			color := parts[1]

			switch color {
			case "red":
				if num > maxCubes["red"] {
					return false
				}
			case "green":
				if num > maxCubes["green"] {
					return false
				}
			case "blue":
				if num > maxCubes["blue"] {
					return false
				}
			}
		}
	}
	return true
}

// part 2
func fewestCubes(g string) (int, int, int) {
	gameData := strings.Split(g, ":")[1]
	handFulls := strings.Split(gameData, ";")

	red := 0
	green := 0
	blue := 0

	for _, handFull := range handFulls {
		for _, cube := range strings.Split(handFull, ",") {
			parts := strings.Split(strings.Trim(cube, " "), " ")
			num, _ := strconv.Atoi(parts[0])
			color := parts[1]

			switch color {
			case "red":
				if red < num {
					red = num
				}
			case "green":
				if green < num {
					green = num
				}
			case "blue":
				if blue < num {
					blue = num
				}
			}
		}
	}
	return red, green, blue
}

func sum(n []int) int {
	total := 0

	for _, v := range n {
		total += v
	}
	return total
}
