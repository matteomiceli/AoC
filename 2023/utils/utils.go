package utils

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadLines(path string) []string {
	d, e := os.ReadFile(path)
	if e != nil {
		log.Fatal(e)
	}

	data := string(d)

	return strings.Split(data, "\n")
}

func Sum(n []int) int {
	total := 0

	for _, v := range n {
		total += v
	}
	return total
}

func IsNumeric(c string) bool {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	for _, num := range nums {
		if c == strconv.Itoa(num) {
			return true
		}
	}
	return false
}
