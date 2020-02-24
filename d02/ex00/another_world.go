package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) >= 1 {
		fields := strings.Fields(args[0])
		str := strings.Join(fields, " ")
		fmt.Println(str)
	}
}
