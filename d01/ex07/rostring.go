package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) >= 2 {
		words := strings.Fields(os.Args[1])
		l := len(words)
		if l > 1 {
			words = append(words[1:], words[0])
		}
		fmt.Println(strings.Join(words, " "))
	}
}
