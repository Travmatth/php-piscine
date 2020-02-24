package main

import (
	Targaryen "d07/ex02"
	"fmt"
)

type Viserys struct {
	*Targaryen.Targaryen
}

func NewViserys() (v *Viserys) {
	v = &Viserys{Targaryen.NewTargaryen()}
	return
}

type Daenerys struct {
	*Targaryen.Targaryen
}

func NewDaenerys() (d *Daenerys) {
	d = &Daenerys{Targaryen.NewTargaryen()}
	return
}

func (d *Daenerys) ResistsFire() bool {
	return true
}

func main() {
	v := NewViserys()
	d := NewDaenerys()
	fmt.Println("Viserys", Targaryen.GetBurned(v))
	fmt.Println("Daenerys", Targaryen.GetBurned(d))
}
