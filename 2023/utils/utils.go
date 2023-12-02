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
