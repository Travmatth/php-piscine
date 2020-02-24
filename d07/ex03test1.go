package main

import (
	WesterosHouse "d07/ex03"
)

type HouseStark struct {
}

func NewHouseStark() *HouseStark {
	return &HouseStark{}
}

func (s *HouseStark) GetHouseName() string {
	return "Stark"
}

func (s *HouseStark) GetHouseMotto() string {
	return "Winter is coming"
}

func (s *HouseStark) GetHouseSeat() string {
	return "Winterfell"
}

type HouseMartell struct {
}

func NewHouseMartell() *HouseMartell {
	return &HouseMartell{}
}

func (h *HouseMartell) GetHouseName() string {
	return "Martell"
}

func (h *HouseMartell) GetHouseMotto() string {
	return "Unbowed, Unbent, Unbroken"
}

func (h *HouseMartell) GetHouseSeat() string {
	return "Sunspear"
}

func main() {
	houses := [2]WesterosHouse.HouseOfWesteros{NewHouseStark(), NewHouseMartell()}
	for _, house := range houses {
		WesterosHouse.Introduce(house)
	}
}
