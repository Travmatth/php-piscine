package main

import (
	"fmt"
	"sort"
	"strings"
)

func Sorted(str string) {
	arr := strings.Fields(str)
	sort.Strings(arr)
	fmt.Print("Array\n(\n")
	for i, word := range arr {
		fmt.Printf("\t[%d] => %s\n", i, word)
	}
	fmt.Print(")\n")
}
