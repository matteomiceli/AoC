package main

import (
	"fmt"
	"matteomiceli/aoc/2024/utils"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func Five() {
	var input = utils.FetchInput("5")

	// parse input
	lines := utils.ReadLines(input)
	rules := []string{}
	updates := []string{}
	for _, line := range lines {
		if strings.Contains(line, "|") {
			rules = append(rules, line)
		} else if strings.Contains(line, ",") {
			updates = append(updates, line)
		}
	}

	rulesMap := map[string][]string{}
	for _, rule := range rules {
		splitRule := strings.Split(rule, "|")
		_, exists := rulesMap[splitRule[0]]
		if exists {
			rulesMap[splitRule[0]] = append(rulesMap[splitRule[0]], splitRule[1])
		} else {
			rulesMap[splitRule[0]] = []string{splitRule[1]}
		}
	}

	correctlyOrdered, notOrdered := findCorrectlyOrderedUpdates(rulesMap, updates)

	fmt.Println("sum of ordered middle nums", sumMiddle(correctlyOrdered))

	sorted := sortUnordered(rulesMap, notOrdered)

	fmt.Println("sum of sorted middle nums", sumMiddle(sorted))
}

func sortUnordered(rulesMap map[string][]string, unordered [][]string) [][]string {
	sorted := [][]string{}
	for _, list := range unordered {
		sort.Slice(list, func(i int, j int) bool {
			num := list[j]
			comesAfter := rulesMap[num]

			if slices.Contains(comesAfter, list[i]) {
				return false
			}
			return true
		})
		sorted = append(sorted, list)
	}
	return sorted
}

func findCorrectlyOrderedUpdates(rulesMap map[string][]string, updates []string) ([][]string, [][]string) {
	notOrdered := [][]string{}
	correctlyOrdered := [][]string{}
	for _, update := range updates {
		list := strings.Split(update, ",")

		isCompliant := true
		for i, num := range list {
			comesAfter := rulesMap[num]

			for _, after := range comesAfter {
				if slices.Contains(list[:i], after) {
					// number seen before it's supposed to appear
					isCompliant = false
					break
				}
			}
		}
		if isCompliant {
			correctlyOrdered = append(correctlyOrdered, list)
		} else {
			notOrdered = append(notOrdered, list)
		}
	}
	return correctlyOrdered, notOrdered
}

func sumMiddle(correctlyOrdered [][]string) int {
	sum := 0
	for _, list := range correctlyOrdered {
		mid := len(list) / 2
		v, _ := strconv.Atoi(list[mid])
		sum += v
	}
	return sum
}
