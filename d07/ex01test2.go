package main

import (
	Euron "d07/Euron"
	"fmt"
)

func main() {
	euron := Euron.NewEuron()
	fmt.Println(euron.familyMotto)
	fmt.Println(euron.Greyjoy.familyMotto)
}
