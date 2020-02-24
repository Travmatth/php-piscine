package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
This PHP program will take 3 arguments. The second is an arithmetic operation
amongst : ’+’, ’-’, ’*’, ’/’, ’%’. The first and the third are numbers. You need to make
this operation and display the result. The program doesn’t manage errors, except the
number of arguments given. Spaces and tabulations can be presented in all 3 arguments.
*/

func toNum(arg string) (num int) {
	trimmed := strings.TrimSpace(arg)
	num, _ = strconv.Atoi(trimmed)
	return
}

func toOp(arg string) (op string) {
	op = strings.TrimSpace(arg)
	return
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Incorrect Parameters")
		os.Exit(1)
	}
	lhs, op, rhs := toNum(os.Args[1]), toOp(os.Args[2]), toNum(os.Args[3])
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
