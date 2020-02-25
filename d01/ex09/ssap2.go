package main

import (
	"os"
	"unicode"
	"sort"
	"fmt"
	"strings"
)

func determineOrder(runes [][]rune) func(i, j int) bool {
	return  func(i, j int) bool {
		min, n, m := 0, len(runes[i]), len(runes[j])
		if n < m {
			min = m
		} else {
			min = n
		}
		less := true
		for k := 0; k < min; k++ {
			first, second := runes[i], runes[j]
			firstIsLetter := unicode.IsLetter(first[k])
			firstIsNumber := unicode.IsNumber(first[k])
			secondIsLetter := unicode.IsLetter(second[k])
			secondIsNumber := unicode.IsNumber(second[k])
			switch {
			case k == n:
				return true
			case k == m:
				return false
			case firstIsLetter && secondIsLetter:
				firstLower := unicode.ToLower(first[k])
				secondLower := unicode.ToLower(rune(second[k]))
				return firstLower < secondLower
			case firstIsLetter:
				return true
			case secondIsLetter:
				return false
			case firstIsNumber && secondIsNumber:
				return int(first[k]) < int(second[k])
			case firstIsNumber:
				return true
			case secondIsNumber:
				return false
			default:
				return first[k] < second[k]
			}
		}
		return less
	}
}

func sort_strings(arr []string) (sorted []string) {
	runes := [][]rune{}
	for i := 0; i < len(arr); i++ {
		items := strings.Fields(arr[i])
		for _, item := range items {
			runes = append(runes, []rune(item))
		}
	}
	sort.SliceStable(runes, determineOrder(runes))
	sorted = []string{}
	for _, str := range runes {
		sorted = append(sorted, string(str))
	}
	return sorted
}

func main() {
	if len(os.Args) > 1 {
		array := os.Args[1:]
		runes := sort_strings(array)
		for i := 0; i < len(runes); i++ {
			fmt.Println(runes[i])
		}
	}
}
