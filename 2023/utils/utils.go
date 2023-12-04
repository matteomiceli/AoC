package utils

import (
	"log"
	"os"
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
