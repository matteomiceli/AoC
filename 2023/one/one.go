package one

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() int {
	d, _ := os.ReadFile("./one/data.txt")
	data := string(d)

	lines := strings.Split(data, "\n")

	var total int

	for _, l := range lines {
		num := []string{}
		maybeWordNum := ""

		// for number chars
		for i := 0; i < len(l); i++ {
			c := string(l[i])

			if isNum(c) {
				num = append(num, c)
			} else {
				maybeWordNum += c
				if isWordNum(maybeWordNum) != "" {
					num = append(num, isWordNum(maybeWordNum))
					maybeWordNum = ""
				}
			}
		}

		if len(num) != 0 {
			n, _ := strconv.Atoi(num[0] + num[len(num)-1])
			total += n
		}
	}
	fmt.Println(total)
	return total
}

func isNum(s string) bool {
	nums := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

	for _, n := range nums {
		if n == s {
			return true
		}
	}
	return false
}

func isWordNum(s string) string {
	spelledNums := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for i, w := range spelledNums {
		if strings.Contains(s, w) {
			return strconv.Itoa(i + 1) // returns "1", "2", "3", etc.
		}
	}
	return ""
}
