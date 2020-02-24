package main

import (
	"fmt"
	"strings"
)

/*
You need to create a little function that will reply true or false
according to whether the array passed as argument is sorted or not.
*/

func ft_is_sort(array []string) bool {
	n := len(array)
	for i := 0; i < n-1; i++ {
		if strings.Compare(array[i], array[i+1]) == 1 {
			return false
		}
	}
	return true
}

func main() {
	array := {"!/@#;^", "42", "Hello World", "hi", "zZzZzZz"}
	if ft_is_sort(array) == false {
		fmt.Println("The array is not sorted")
	} else {
		fmt.Println("The array is sorted")
	}
}