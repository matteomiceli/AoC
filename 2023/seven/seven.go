package seven

import (
	"fmt"
	"matteomiceli/aoc/utils"
	"strings"
)

func Run() {
	lines := utils.ReadLines("seven/data.txt")
	sets := parseSets(lines)

	sortHandStrength(sets)
}

func sortHandStrength(sets [][]string) {
	cardMap := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
	//  handMap := []string{"high", "pair", "two-pair", "trips", "full", "four", "five"}
	for i := 0; i < len(sets); i++ {
		fmt.Println(getHand(sets[i], cardMap))
	}

}

func getHand(set []string, cardMap []string) string {
	// number of cards in hand
	collapsePairs := map[string]int{}

	for _, c := range set[0] {
		card := string(c)
		_, ok := collapsePairs[card]
		// if card is not in map
		if !ok {
			collapsePairs[card] = 1
		} else {
			collapsePairs[card] += 1
		}
	}

	switch len(collapsePairs) {
	case 5:
		return "high"
	case 4:
		return "pair"
	case 3:
		// can be two-pair or trips
		trips := false
		for _, v := range collapsePairs {
			if v == 3 {
				trips = true
			}
		}
		if trips {
			return "trips"
		}
		return "two-pair"
	case 2:
		// can be four or full
		four := false
		for _, v := range collapsePairs {
			if v == 4 {
				four = true
			}
		}
		if four {
			return "four"
		}
		return "full"
	}
	return ""
}

func parseSets(lines []string) [][]string {
	sets := [][]string{}
	for _, line := range lines {
		if line == "" {
			continue
		}

		fmt.Println(line)
		splitSet := strings.Split(line, " ")
		sets = append(sets, []string{splitSet[0], splitSet[1]}) // hand, bet
	}

	return sets
}
