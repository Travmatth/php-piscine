package main

import (
	"fmt"
	"os"
	"strings"
)

/*
Your goal is to create a program that will display the corresponding value to a key
given as first argument amongst an unlimited number of couples formated like this:
“key:value” given as following arguments.
*/

func main() {
	if len(os.Args) == 1 {
		return
	}
	lookup := os.Args[1]
	arr := os.Args[2:]
	dict := make(map[string]string)
	for i := 0; i < len(arr); i++ {
		str := strings.Split(arr[i], ":")
		dict[str[0]] = str[1]
	}
	if val, ok := dict[lookup]; ok {
		fmt.Println(val)
	}
}
