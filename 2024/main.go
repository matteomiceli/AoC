package main

import "os"

var days = map[string]func(){
	"one":   One,
	"two":   Two,
	"three": Three,
}

func main() {
	days[os.Args[1]]()
}
