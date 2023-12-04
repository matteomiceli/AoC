package four

import (
	"fmt"
	"math"
	"matteomiceli/aoc/utils"
	"strings"
)

func Run() {
	cards := utils.ReadLines("four/data.txt")

	totalScratchValue := 0
	winners := map[int][]string{}

	for i, card := range cards {
		if len(card) == 0 {
			continue
		}

		WIN_MULTIPLIER := 2

		numbers := strings.Split(card, ":")[1]
		winningNums := strings.Split(strings.Split(numbers, "|")[0], " ")
		cardNums := strings.Split(strings.Split(numbers, "|")[1], " ")

		for _, cn := range cardNums {
			cardNum := strings.Trim(cn, " ")
			if isWinningNum(cardNum, winningNums) {
				winners[i] = append(winners[i], cardNum)
			}
		}
		if len(winners[i]) < 1 {
			winners[i] = []string{}
		}

		totalScratchValue += int(math.Floor(math.Pow(float64(WIN_MULTIPLIER), float64(len(winners[i])-1))))
	}
	fmt.Println("Total scratch val ", totalScratchValue)
	// part 2
	fmt.Println("Total cards ", calculateWinMoreScratchCards(winners))
}

func calculateWinMoreScratchCards(winners map[int][]string) int {
	totalCards := 0
	cardsCount := []int{}
	cardsEnd := len(winners)

	for i := 0; i < len(winners); i++ {
		cardsCount = append(cardsCount, 1) // we start with one copy of each card
	}

	for cardIdx := 0; cardIdx < len(winners); cardIdx++ {
		results := winners[cardIdx]
		cardWins := len(results)
		inc := 1

		for inc <= cardWins {
			nextIdx := inc + cardIdx
			if nextIdx < cardsEnd {
				cardsCount[nextIdx] += cardsCount[cardIdx] * 1
			}
			inc++
		}
	}

	for _, v := range cardsCount {
		totalCards += v
	}
	return totalCards
}

func isWinningNum(num string, winningNums []string) bool {
	if num == "" {
		return false
	}
	for _, winningNum := range winningNums {
		if num == strings.Trim(winningNum, " ") {
			return true
		}
	}
	return false
}
