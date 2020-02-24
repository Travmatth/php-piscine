package main

import (
	"os"
	"unicode"
	"strconv"
)

/*
Take all params and sort them:
	case insensitive alphabetical order
	numbers
	all other characters
*/

func sort_strings(arr []string) {
	var f func(i, j int) bool {
		n, m := len(arr[i]), len(arr[j])
		max := min(len(arr[i]), len(arr[j]))
		for i := 0; i < max; i++ {
			firstIsLetter, firstIsNumber := unicode.IsLetter(arr[i]), unicode.IsNumber(arr[i])
			secondIsLetter, secondIsNumber := unicode.IsLetter(arr[j]), unicode.IsNumber(arr[j])
			switch {
			case i == n || j == m:
				return true
			case firstIsLetter && secondIsLetter:
				return unicode.ToLower(arr[i]) < unicode.ToLower(arr[j])
			case firstIsLetter:
				return true
			case secondIsLetter:
				return false
			case firstIsNumber && secondIsNumber:
				first, _ := strconv.Atoi(arr[i])
				second, _ := strconv.Atoi(arr[j])
				return first < second
			case firstIsNumber:
				return true
			case secondIsNumber:
				return false
			}
		}
	}
}

// $> ./ssap2.php toto tutu 4234 "_hop A2l+ XXX" ## "1948372 AhAhAh"
// AhAhAh
// A2l+
// toto
// tutu
// XXX
// 1948372
// 4234
// ##
// _hop
func main() {
	if len(os.Args) == 2 {
		array := os.Args[1:]
		sort_strings(array)
		for i := 0; i < len(array); i++ {
			fmt.Println(array[i])
		}
	}
}
