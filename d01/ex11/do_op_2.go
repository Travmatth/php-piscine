package main

import (
	"fmt"
	"os"
	"unicode"
)

/*
This time, only 1 argument is on the menu. That one contains the whole calculation
that needs to be done. This calculation will always be under the following format number
operator number. A new error message “Syntax Error” will now complete the prior
message in case the syntax isn’t correct. There can be no space between the numbers
and the operator, or there can be many. The expected results is the same.
*/

func splitString(arg string) (lhs int, op string, rhs int) {
	first := true
	lhs, op, rhs = 0, "", 0
	for _, c := range arg {
		switch {
		case unicode.IsSpace(c):
			continue
		case first && unicode.IsNumber(c):
			next := int(c)
			lhs = (lhs * 10) + next
		case !first && unicode.IsNumber(c):
			next := int(c)
			rhs = (rhs * 10) + next
		default:
			first = false
			op = string(c)
		}
	}
	return
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("“Syntax Error")
		os.Exit(1)
	}
	lhs, op, rhs := splitString(os.Args[1])
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
