package main

import "os"

var days = map[string]func(){
	"one": One,
	"two": Two,
}

func main() {
	days[os.Args[1]]()
}
