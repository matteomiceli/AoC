package main

import (
	"fmt"
	"matteomiceli/aoc/2024/utils"
	"regexp"
	"strings"
)

func Three() {
	var input = utils.FetchInput("3")

	total := 0
	do := true

	regex := "mul\\(\\d*,\\d*\\)"                     // part 1
	regex = "mul\\(\\d*,\\d*\\)|do\\(\\)|don't\\(\\)" // part 2
	r, _ := regexp.Compile(regex)

	matchedFactors := r.FindAll([]byte(input), -1)

	for _, factor := range matchedFactors {
		isMult := strings.Contains(string(factor), "mul")
		isDont := strings.Contains(string(factor), "don't")

		if isDont {
			do = false
		} else if isMult {
			if do {
				// parse numbers 'mul(x,y)' -> '[x,y]'
				partial := strings.Replace(string(factor), "mul(", "", -1)
				partial = string(partial[:len(partial)-1])
				nums := strings.Split(partial, ",")

				multiplied := utils.MultiplyStrings(nums)
				total += multiplied
			}
		} else {
			do = true
		}
	}

	fmt.Println("total:", total)
}
