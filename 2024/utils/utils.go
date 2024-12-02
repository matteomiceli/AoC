package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func ReadLines(data string) []string {
	data = strings.TrimSpace(data)

	return strings.Split(data, "\n")
}

func Sum(n []int) int {
	total := 0

	for _, v := range n {
		total += v
	}
	return total
}

func Multiply(n []int) int {
	total := 1

	for _, v := range n {
		total *= v
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

func FetchInput(day string) string {
	url := fmt.Sprintf("https://adventofcode.com/2024/day/%s/input", day)
	path := fmt.Sprintf("data/day%s.txt", day)

	// if file already exists, retrieve
	_, err := os.Stat(path)
	if err == nil {
		d, err := os.ReadFile(path)
		if err != nil {
			log.Fatal(err)
		}

		return string(d)
	}

	// Doesn't exist, fetch input
	req, e := http.NewRequest(http.MethodGet, url, nil)
	if e != nil {
		log.Fatal(e)
	}

	req.Header.Set("Cookie", fmt.Sprintf("session=%s", os.Getenv("AOC_TOKEN")))

	res, e := http.DefaultClient.Do(req)
	if e != nil {
		log.Fatal(e)
	}

	defer res.Body.Close()
	d, e := io.ReadAll(res.Body)
	if e != nil {
		log.Fatal(e)
	}

	err = os.WriteFile(path, d, 0644)
	if err != nil {
		log.Fatal(err)
	}

	return string(d)
}
