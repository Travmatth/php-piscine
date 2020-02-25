package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 2 {
		args := os.Args[1]
		words := strings.Fields(args)
		fmt.Println(strings.Join(words, " "))
	}
}
