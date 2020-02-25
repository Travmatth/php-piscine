package main

import (
	"fmt"
)

func main() {
	str := "Hello         World 	AAA"
	arr := Sorted(str)
	fmt.Print("Array\n(\n")
	for i, word := range arr {
		fmt.Printf("\t[%d] => %s\n", i, word)
	}
	fmt.Print(")\n")
}
