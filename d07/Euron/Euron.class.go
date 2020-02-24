package Euron

import (
	Greyjoy "d07/ex01"
	"fmt"
)

type Euron struct {
	*Greyjoy.Greyjoy
}

func NewEuron() (e *Euron) {
	e = &Euron{Greyjoy.NewGreyjoy()}
	return
}

func (e *Euron) AnnounceMotto() {
	fmt.Println(e.GetFamilyMotto())
}
