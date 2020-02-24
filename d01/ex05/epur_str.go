package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		words := strings.Fields(args[0])
		fmt.Println(strings.Join(words, " "))
	}
}
