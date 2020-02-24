package WesterosHouse

import "fmt"

type HouseOfWesteros interface {
	GetHouseName() string
	GetHouseMotto() string
	GetHouseSeat() string
}

func Introduce(house HouseOfWesteros) {
	name := house.GetHouseName()
	seat := house.GetHouseSeat()
	motto := house.GetHouseMotto()
	fmt.Printf("House %s of %s: \"%s\"\n", name, seat, motto)
}
