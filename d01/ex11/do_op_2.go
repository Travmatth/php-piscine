package main

import (
	"fmt"
	"os"
	"errors"
	"regexp"
	"strconv"
)

var r = regexp.MustCompile("^\\s*([0-9]*)\\s*([%*+-/])\\s*([0-9]*)\\s*$")

func splitString(arg string) (lhs int, op string, rhs int, err error) {
	matches := r.FindAllStringSubmatchIndex(arg, -1)
	if len(matches) == 0 {
		return 0, "", 0, errors.New("Syntax Error")
	}
	m := matches[0]
	lhs, _ = strconv.Atoi(arg[m[2]:m[3]])
	op = arg[m[4]:m[5]]
	rhs, _ = strconv.Atoi(arg[m[6]:m[7]])
	return
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Syntax Error")
	}
	lhs, op, rhs, err := splitString(os.Args[1])
	if err != nil {
		fmt.Println("Syntax Error")
	}
	switch op {
	case "+":
		fmt.Println(lhs + rhs)
	case "-":
		fmt.Println(lhs - rhs)
	case "*":
		fmt.Println(lhs * rhs)
	case "/":
		fmt.Println(lhs / rhs)
	case "%":
		fmt.Println(lhs % rhs)
	}
}
