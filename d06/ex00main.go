package main

import (
	"fmt"
	"d06/ex00"
)

func main() {
	doc, err := Color.Doc()
	if err == nil {
		fmt.Print(doc)
	}

	Color.Verbose = true
	red := Color.NewColor([]int{0xff, 0, 0})
	defer red.Release()
	green := Color.NewColor([]int{255 << 8})
	defer green.Release()
	blue := Color.NewColor([]int{0, 0, 0xff})
	defer blue.Release()
	yellow := red.Add(green)
	defer yellow.Release()
	cyan := green.Add(blue)
	defer cyan.Release()
	magenta := blue.Add(red)
	defer magenta.Release()
	white := red.Add(green).Add(blue)
	defer white.Release()

	fmt.Println(red)
	fmt.Println(green)
	fmt.Println(blue)
	fmt.Println(yellow)
	fmt.Println(cyan)
	fmt.Println(magenta)
	fmt.Println(white)

	Color.Verbose = false

	black := white.Sub(red).Sub(green).Sub(blue)
	defer black.Release()
	fmt.Println("Black:", black)

	Color.Verbose = true

	darkgrey := Color.NewColor([]int{(10 << 16) + (10 << 8) + 10})
	defer darkgrey.Release()
	fmt.Println("darkgrey:", darkgrey)

	lightgrey := darkgrey.Mult(22.5)
	defer lightgrey.Release()
	fmt.Println("lightgrey:", lightgrey)

	random := Color.NewColor([]int{12, 31, 23})
	defer random.Release()
	fmt.Println("random:", random)
}