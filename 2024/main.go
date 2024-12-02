package main

import "os"

var days = map[string]func(){
	"one": One,
}

func main() {
	days[os.Args[1]]()
}
