package main

import (
	"sort"
	"strings"
)

func Sorted(str string) (arr []string) {
	arr = strings.Fields(str)
	sort.Strings(arr)
	return
}
