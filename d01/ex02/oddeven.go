package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a number: ")
	for scanner.Scan() {
		cmd := strings.ToLower(scanner.Text())
		if val, err := strconv.Atoi(cmd); err != nil {
			fmt.Printf("'%s' is not a number\n", cmd)
		} else {
			if val%2 == 0 {
				fmt.Printf("The number %d is even\n", val)
			} else {
				fmt.Printf("The number %d is odd\n", val)
			}
		}
		fmt.Print("Enter a number: ")
	}
	fmt.Print("\n")
	if err := scanner.Err(); err != nil {
		if err != io.EOF {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
