package main

import (
	WesterosHouse "d07/ex03"
)

type DrHouse struct {
}

func NewDrHouse() *DrHouse {
	return &DrHouse{}
}

func (d *DrHouse) Diagnose() string {
	return "It's not lupus!"
}

func main() {
	houses := [1]WesterosHouse.HouseOfWesteros{NewDrHouse()}
	for _, house := range houses {
		WesterosHouse.Introduce(house)
	}
}
