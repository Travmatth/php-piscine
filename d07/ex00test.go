package main

import (
	"fmt"

	Tyrion "d07/ex00"
)

func main() {
	tyrion := Tyrion.NewTyrion()
	fmt.Println(tyrion.GetSize())
	fmt.Println(tyrion.HouseMotto())
}
