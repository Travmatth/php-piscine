package main

import (
	"fmt"
	"sort"
)

/*
You need to create a little function that will reply true or false
according to whether the array passed as argument is sorted or not.
*/

func ft_is_sort(array []string) bool {
	return sort.SliceIsSorted(array, func (i, j int) bool {
		return array[i] < array[j]
	})
}

func main() {
	// array := []string{"!/@#;^", "42", "Hello World", "hi", "zZzZzZz"}
	array := []string{"What", "are", "we", "doing", "now"}
	if ft_is_sort(array) == false {
		fmt.Println("The array is not sorted")
	} else {
		fmt.Println("The array is sorted")
	}
}