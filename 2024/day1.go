package main

import (
	"fmt"
	"math"
	"matteomiceli/aoc/2024/utils"
	"sort"
	"strconv"
	"strings"
)

func One() {
	var input = utils.FetchInput("1")

	lines := (strings.Split(input, "\n"))
	list1 := []int{}
	list2 := []int{}

	for _, line := range lines {
		if line == "" {
			continue
		}
		nums := strings.Split(line, "   ")
		num1, _ := strconv.Atoi(string(nums[0]))
		num2, _ := strconv.Atoi(string(nums[1]))
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	sumOfDifferences(list1, list2)
	similarity(list1, list2)
}

func sumOfDifferences(list1 []int, list2 []int) {
	diffs := []int{}

	sort.Slice(list1, func(i int, j int) bool {
		return list1[j] > list1[i]
	})
	sort.Slice(list2, func(i int, j int) bool {
		return list2[j] > list2[i]
	})

	for i := range list1 {
		// some stupid type coversion to use math Abs int -> float64 -> int
		diffs = append(diffs, int(math.Abs(float64(list1[i])-float64(list2[i]))))
	}

	fmt.Println("part 1:", utils.Sum(diffs))
}

func similarity(list1 []int, list2 []int) {
	list2NumFreq := map[int]int{}
	similarity := []int{}

	for _, num := range list2 {
		_, exists := list2NumFreq[num]
		if exists {
			list2NumFreq[num]++
		} else {
			list2NumFreq[num] = 1
		}
	}

	for _, num := range list1 {
		freq, exists := list2NumFreq[num]
		if exists {
			similarity = append(similarity, freq*num)
		} else {
			similarity = append(similarity, 0)
		}
	}

	fmt.Println("part 2:", utils.Sum(similarity))
}
