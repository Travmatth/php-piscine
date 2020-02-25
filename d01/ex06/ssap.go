package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		words := make([]string, 0)
		for _, arg := range os.Args[1:] {
			words = append(words, strings.Fields(arg)...)
		}
		sort.Strings(words)
		fmt.Println(strings.Join(words, "\n"))
	}
}
